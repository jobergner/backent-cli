package server 
import (
	"nhooyr.io/websocket"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"fmt"
	"sync"
	"github.com/jobergner/backent-cli/examples/state"
	"errors"
	"time"
	"net/http"
	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
)
type Client struct {
	lobby		*Lobby
	room		*Room
	conn		connect.Connector
	messageChannel	chan [ // easyjson:skip
	]byte
	id	string
}

func newClient(websocketConnector connect.Connector, lobby *Lobby) (*Client, error) {
	clientID, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Msg("failed generating client ID")
		return nil, err
	}
	c := Client{lobby: lobby, conn: websocketConnector, messageChannel: make(chan []byte, 32), id: clientID.String()}
	err = c.sendIdentifyingMessage()
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (c *Client) sendIdentifyingMessage() error {
	msg := Message{Kind: message.MessageKindID, Content: []byte(c.id)}
	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling message")
		return err
	}
	c.messageChannel <- msgBytes
	return nil
}
func (c *Client) SendMessage(msg []byte) {
	c.messageChannel <- msg
}
func (c *Client) ID() string {
	return c.id
}
func (c *Client) RoomName() string {
	return c.room.name
}

func (c *Client) closeConnection(reason string) {
	c.conn.Close(reason)
}
func (c *Client) runReadMessages() {
	defer c.closeConnection("failed reading messages")
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		var msg Message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Err(err).Str(logging.Message, string(msgBytes)).Msg("failed unmarshalling message")
			errMsg, _ := Message{message.MessageIDUnknown, message.MessageKindError, [ // closeConnection closes the client's connection
			// this does not do anything else on its own, but triggers
			// the removal of the client from the system in the
			// http handler
			]byte("invalid message"), nil}.MarshalJSON()
			c.messageChannel <- errMsg
			continue
		}
		msg.client = c
		if msg.Kind == message.MessageKindGlobal {
			c.lobby.processMessageSync(msg)
		} else {
			if c.room != nil {
				c.room.processMessageSync(msg)
			}
		}
	}
}
func (c *Client) runWriteMessages() {
	defer c.closeConnection("failed writing messages")
	for {
		msgBytes, ok := <-c.messageChannel
		if !ok {
			log.Warn().Str(logging.ClientID, c.id).Msg("client message channel was closed")
			break
		}
		c.conn.WriteMessage(msgBytes)
	}
}

type clientRegistrar struct {
	clients	map // easyjson:skip
	[*Client]struct{}
	incomingClients	map[*Client]struct{}
	mu		sync.Mutex
}

func newClientRegistar() *clientRegistrar {
	return &clientRegistrar{clients: make(map[*Client]struct{}), incomingClients: make(map[*Client]struct{})}
}
func (c *clientRegistrar) add(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Debug().Str(logging.ClientID, client.id).Msg("adding client")
	c.incomingClients[client] = struct{}{}
}
func (c *clientRegistrar) remove(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Debug().Str(logging.ClientID, client.id).Msg("removing client")
	delete(c.clients, client)
	delete(c.incomingClients, client)
}

func (c *clientRegistrar) kick(client *Client, reason string) {
	log.Debug().Str(logging.ClientID, client.id).Msg("kicking client")
	client.closeConnection(reason)
}
func (c *clientRegistrar) promote(client *Client) {
	c.mu.Lock()
	defer c.mu.Unlock()
	log.Debug().Str(logging.ClientID, client.id).Msg("promoting client")
	c.clients[client] = struct{}{}
	delete(c.incomingClients, client)
}

var (
	ErrMessageKindUnknown = errors.New("message kind unknown")
)

type Lobby struct {
	mu	sync.Mutex
	Rooms	map // TODO unused
	// easyjson:skip
	[string]*Room
	controller	Controller
	fps		int
}

func newLobby(controller Controller, fps int) *Lobby {
	l := &Lobby{Rooms: make(map[string]*Room), controller: controller, fps: fps}
	l.signalCreation()
	return l
}
func (l *Lobby) CreateRoom(name string) *Room {
	if room, ok := l.Rooms[name]; ok {
		log.Warn().Str(logging.RoomName, name).Msg("attempted to create room which already exists")
		return room
	}
	room := newRoom(l.controller, name)
	l.Rooms[name] = room
	room.Deploy(l.controller, l.fps)
	return room
}
func (l *Lobby) DeleteRoom(name string) {
	room, ok := l.Rooms[name]
	if !ok {
		log.Warn().Str(logging.RoomName, name).Msg("attempted to delete room which does not exists")
		return
	}
	room.mode = RoomModeTerminating
	delete(l.Rooms, name)
}
func (l *Lobby) addClient(client *Client) {
	l.signalClientConnect(client)
}
func (l *Lobby) deleteClient(client *Client) {
	if client.room != nil {
		client.room.clients.remove(client)
	}
	l.signalClientDisconnect(client)
}
func (l *Lobby) processMessageSync(msg Message) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.controller.OnSuperMessage(msg, msg.client.room, msg.client, l)
}
func (l *Lobby) signalClientDisconnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.controller.OnClientDisconnect(client.room, client.id, l)
}
func (l *Lobby) signalClientConnect(client *Client) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.controller.OnClientConnect(client, l)
}
func (l *Lobby) signalCreation() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.controller.OnCreation(l)
}

type Message struct {
	ID	string		`json:"id"`
	Kind	message.Kind	`json:"kind"`
	Content	[]byte		`json:"content"`
	client	*Client
}
type RoomMode int

const (
	RoomModeIdle	RoomMode	= iota
	RoomModeRunning
	RoomModeTerminating
)

type Room struct {
	name		string
	mu		sync.Mutex
	clients		*clientRegistrar
	state		*state.Engine
	controller	Controller
	mode		RoomMode
}

func newRoom(controller Controller, name string) *Room {
	return &Room{name: name, clients: newClientRegistar(), state: state.NewEngine(), controller: controller}
}
func (r *Room) Name() string {
	return r.name
}
func (r *Room) RemoveClient(client *Client) {
	r.clients.remove(client)
}
func (r *Room) AddClient(client *Client) {
	client.room = r
	r.clients.add(client)
}
func (r *Room) AlterState(fn func(*state.Engine)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	fn(r.state)
}
func (r *Room) RangeClients(fn func(client *Client)) {
	for c := // easyjson:skip
	range r.clients.incomingClients {
		fn(c)
	}
	for c := range r.clients.clients {
		fn(c)
	}
}
func (r *Room) processMessageSync(msg Message) {
	r.mu.Lock()
	defer r.mu.Unlock()
	response := r.processClientMessage(msg)
	if response.Kind == message.MessageKindNoResponse {
		return
	}
	responseBytes, err := response.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.Message, string(responseBytes)).Str(logging.MessageKind, string(response.Kind)).Msg("failed marshalling response")
		return
	}
	select {
	case response.client.messageChannel <- responseBytes:
	default:
		log.Warn().Str(logging.ClientID, response.client.id).Msg(logging.ClientBufferFull)
		response.client.closeConnection(logging.ClientBufferFull)
	}
}
func (r *Room) run(controller Controller, fps int) {
	ticker := time.NewTicker(time.Second / time.Duration(fps))
	for {
		<-ticker.C
		r.tickSync(controller)
		if r.mode == RoomModeTerminating {
			break
		}
	}
}
func (r *Room) Deploy(controller Controller, fps int) {
	go r.run(controller, fps)
}

type Server struct {
	HttpServer	*http.Server
	Lobby		*Lobby
}

func NewServer(controller Controller, fps int, configs ...func(*http.Server, *http.ServeMux)) *Server {
	server := Server{HttpServer: new(http.Server), Lobby: newLobby(controller, fps)}
	handler := http.NewServeMux()
	for _, c := // easyjson:skip
	range configs {
		c(server.HttpServer, handler)
	}
	if server.HttpServer.Addr == "" {
		server.HttpServer.Addr = fmt.Sprintf(":%d", 8080)
	}
	handler.HandleFunc("/", homePageHandler)
	handler.HandleFunc("/ws", server.wsEndpoint)
	server.HttpServer.Handler = handler
	return &server
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	websocketConnection, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Err(err).Msg("failed creating websocket connection")
		return
	}
	client, err := newClient(connect.NewConnection(websocketConnection, r.Context()), s.Lobby)
	if err != nil {
		return
	}
	s.Lobby.addClient(client)
	go client.runReadMessages()
	go client.runWriteMessages()
	<-client.conn.Context().Done()
	log.Debug().Msg("client context done")
	s.Lobby.deleteClient(client)
}
func (s *Server) Start() chan error {
	log.Info().Msgf("backent running on port %s\n", s.HttpServer.Addr)
	serverError := make(chan error)
	go func() {
		err := s.HttpServer.ListenAndServe()
		serverError <- err
	}()
	return serverError
}
func (r *Room) tickSync(controller Controller) {
	r.mu.Lock()
	defer r.mu.Unlock()
	controller.OnFrameTick(r.state)
	err := r.publishPatch()
	if err != nil {
		return
	}
	r.state.UpdateState()
	r.handleIncomingClients()
}
func (r *Room) publishPatch() error {
	if r.state.Patch.IsEmpty() {
		return nil
	}
	patchBytes, err := r.state.Patch.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling patch")
		return err
	}
	stateUpdateMsg := Message{Kind: message.MessageKindUpdate, Content: patchBytes}
	stateUpdateBytes, err := stateUpdateMsg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(stateUpdateMsg.Kind)).Msg("failed marshalling message")
		return err
	}
	r.broadcastPatchToClients(stateUpdateBytes)
	return nil
}
func (r *Room) broadcastPatchToClients(stateUpdateBytes []byte) {
	for client := range r.clients.clients {
		select {
		case client.messageChannel <- stateUpdateBytes:
		default:
			log.Warn().Str(logging.ClientID, client.id).Msg(logging.ClientBufferFull)
			client.closeConnection(logging.ClientBufferFull)
		}
	}
}
func (r *Room) handleIncomingClients() {
	if len(r.clients.incomingClients) == 0 {
		return
	}
	stateBytes, err := r.state.State.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling state")
		return
	}
	currentStateMsg := Message{Kind: message.MessageKindCurrentState, Content: stateBytes}
	currentStateMessageBytes, err := currentStateMsg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(currentStateMsg.Kind)).Msg("failed marshalling message")
		return
	}
	for client := range r.clients.incomingClients {
		select {
		case client.messageChannel <- currentStateMessageBytes:
			r.clients.promote(client)
		default:
			log.Warn().Str(logging.ClientID, client.id).Msg(logging.ClientBufferFull)
			client.closeConnection(logging.ClientBufferFull)
		}
	}
}
func (r *Room) processClientMessage(msg Message) Message {
	switch msg.Kind {
	case message.MessageKindAction_addItemToPlayer:
		var params message.AddItemToPlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		r.state.BroadcastingClientID = msg.client.id
		r.controller.AddItemToPlayerBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""
		res := r.controller.AddItemToPlayerEmit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling response content")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}

	case message.MessageKindAction_movePlayer:
		var params message.MovePlayerParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		r.state.BroadcastingClientID = msg.client.id
		r.controller.MovePlayerBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""
		r.controller.MovePlayerEmit(params, r.state, r.name, msg.client.id)

		return Message{
			ID:   msg.ID,
			Kind: message.MessageKindNoResponse,
		}
	case message.MessageKindAction_spawnZoneItems:
		var params message.SpawnZoneItemsParams
		err := params.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Str(logging.MessageContent, string(msg.Content)).Msg("failed unmarshalling params")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		r.state.BroadcastingClientID = msg.client.id
		r.controller.SpawnZoneItemsBroadcast(params, r.state, r.name, msg.client.id)
		r.state.BroadcastingClientID = ""
		res := r.controller.SpawnZoneItemsEmit(params, r.state, r.name, msg.client.id)
		resContent, err := res.MarshalJSON()
		if err != nil {
			log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling response content")
			return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
		}
		return Message{msg.ID, msg.Kind, resContent, msg.client}

	default:
		err := ErrMessageKindUnknown
		log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("unknown message kind")
		return Message{msg.ID, message.MessageKindError, []byte("invalid message"), msg.client}
	}
}

type Controller interface {
	OnSuperMessage(msg Message, room *Room, client *Client, lobby *Lobby)
	OnClientConnect(client *Client, lobby *Lobby)
	OnClientDisconnect(room *Room, clientID string, lobby *Lobby)
	OnCreation(lobby *Lobby)
	OnFrameTick(engine *state.Engine)
	AddItemToPlayerBroadcast(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string)
	AddItemToPlayerEmit(params message.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) message.AddItemToPlayerResponse

	MovePlayerBroadcast(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)
	MovePlayerEmit(params message.MovePlayerParams, engine *state.Engine, roomName, clientID string)

	SpawnZoneItemsBroadcast(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string)
	SpawnZoneItemsEmit(params message.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) message.SpawnZoneItemsResponse
}
