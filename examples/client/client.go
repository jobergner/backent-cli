package client

import (
	"context"
	"sync"
	"time"

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

func NewClient(connectCTX context.Context, controller Controller, fps int) (*Client, error) {

	c, _, err := websocket.Dial(connectCTX, "http://localhost:8080/ws", nil)
	if err != nil {
		log.Err(err).Msg("failed creating client while dialing server")
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	client := Client{
		fps:        fps,
		controller: controller,
		conn:       connect.NewConnection(c, ctx, cancel),
		router: &responseRouter{
			pending: make(map[string]chan []byte),
		},
		receiveID:      make(chan string, 1),
		messageChannel: make(chan []byte),
		patchChannel:   make(chan []byte),
		engine:         state.NewEngine(),
	}

	go client.runReadMessages()
	go client.runWriteMessages()
	go client.emitPatches()

	select {
	case <-connectCTX.Done():
		return nil, ErrResponseTimeout
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

	patchBytes, err := c.engine.Patch.MarshalJSON()
	if err != nil {
		log.Err(err).Msg("failed marshalling patch")
		return
	}

	c.engine.UpdateState()

	c.patchChannel <- patchBytes
}

// TODO maybe return error that signals when anything critical happens
// switch of patchChannel and errorChannel
func (c *Client) ReadUpdate() []byte {
	return <-c.patchChannel
}

func (c *Client) emitPatches() {
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

		log.Debug().Msg("importing patch")
		c.engine.ImportPatch(&patch)

		c.mu.Unlock()
	default:

		c.router.route(msg)
	}

	return nil
}
