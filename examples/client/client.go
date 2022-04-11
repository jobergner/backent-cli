package client

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/jobergner/backent-cli/examples/state"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

// easyjson:skip
type Client struct {
	fps            int
	id             string
	mu             sync.Mutex
	controller     Controller
	engine         *state.Engine
	conn           connect.Connector
	router         *responseRouter
	receiveID      chan string
	messageChannel chan []byte
	patchChannel   chan []byte
}

func NewClient(ctx context.Context, controller Controller, fps int) (*Client, error) {

	// TODO i dont know about all this
	dialCTX, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, _, err := websocket.Dial(dialCTX, "http://localhost:8080/ws", nil)
	if err != nil {
		log.Err(err).Msg("failed creating client while dialing server")
		return nil, err
	}

	client := Client{
		fps:            fps,
		controller:     controller,
		conn:           connect.NewConnection(c, ctx),
		router:         newReponseRouter(),
		receiveID:      make(chan string),
		messageChannel: make(chan []byte),
		patchChannel:   make(chan []byte),
		engine:         state.NewEngine(),
	}

	go client.runReadMessages()
	go client.runWriteMessages()
	go client.runEmitPatches()

	select {
	case <-time.After(2 * time.Second):
		cancel()
		return nil, dialCTX.Err()

	case clientID := <-client.receiveID:
		client.id = clientID
		client.engine.ThisClientID = clientID
		break

	}

	return &client, nil
}

func (c *Client) tick() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.engine.Patch.IsEmpty() {
		return
	}

	c.engine.AssembleUpdateTree()
	updateTreeBytes, err := c.engine.Tree.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling patch")
		return
	}

	c.engine.UpdateState()

	c.patchChannel <- updateTreeBytes
}

// TODO maybe return error that signals when anything critical happens
// switch of patchChannel and errorChannel
func (c *Client) ReadUpdate() []byte {
	return <-c.patchChannel
}

func (c *Client) runEmitPatches() {
	ticker := time.NewTicker(time.Second / time.Duration(c.fps))

	for {
		<-ticker.C
		c.tick()
	}
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) closeConnection(reason string) {
	c.conn.Close(reason)
}

func (c *Client) runReadMessages() {
	defer c.closeConnection("failed reading messages")

	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Err(err).Msg("failed reading message")
			break
		}

		var msg Message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Err(err).Str(logging.Message, string(msgBytes)).Msg("failed unmarshalling message")
			continue
		}

		c.processMessage(msg)
	}
}

func (c *Client) runWriteMessages() {
	defer c.closeConnection("failed writing messages")

	for {
		msg, ok := <-c.messageChannel

		if !ok {
			log.Warn().Msg("failed while attempted sending to closed client message channel")
			break
		}

		c.conn.WriteMessage(msg)
	}
}

func (c *Client) processMessage(msg Message) error {
	switch msg.Kind {
	case message.MessageKindID:
		c.receiveID <- string(msg.Content)

	case message.MessageKindUpdate, message.MessageKindCurrentState:
		var patch state.State

		err := patch.UnmarshalJSON(msg.Content)
		if err != nil {
			log.Err(err).Msg("failed unmarshalling patch")
			return err
		}

		c.mu.Lock()
		c.engine.ImportPatch(&patch)
		c.mu.Unlock()

	default:
		c.router.route(msg)
	}

	return nil
}

func (c *Client) SuperMessage(b []byte) error {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(message.MessageKindGlobal)).Msg("failed generating message ID")
		return err
	}

	idString := id.String()

	msg := Message{idString, message.MessageKindGlobal, b}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageID, msg.ID).Str(logging.Message, string(msgBytes)).Str(logging.MessageKind, string(message.MessageKindGlobal)).Msg("failed marshalling message")
		return err
	}

	c.messageChannel <- msgBytes

	return nil
}
