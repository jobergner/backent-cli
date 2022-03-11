package state

import (
	"context"
	"log"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

type Client struct {
	id             string
	mu             sync.Mutex
	actions        Actions
	engine         *Engine
	conn           Connector
	router         *responseRouter
	idSignal       chan string
	messageChannel chan []byte
}

func NewClient(actions Actions) (*Client, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	c, _, err := websocket.Dial(ctx, "ws://localhost:8080", nil)
	if err != nil {
		return nil, cancel, err
	}

	client := Client{
		actions: actions,
		engine:  newEngine(),
		conn:    NewConnection(ctx, c),
		router: &responseRouter{
			pending: make(map[string]chan []byte),
		},
		idSignal:       make(chan string, 1),
		messageChannel: make(chan []byte),
	}

	go client.runReadMessages()
	go client.runWriteMessages()

	client.id = <-client.idSignal

	return &client, cancel, nil
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) handleInernalError() {
	c.conn.Close()
}

func (c *Client) runReadMessages() {
	defer c.handleInernalError()

	for {
		_, msgBytes, err := c.conn.ReadMessage()

		if err != nil {
			log.Printf("unregistering client due to error while reading connection: %s", err)
			break
		}

		var msg Message
		err = msg.UnmarshalJSON(msgBytes)

		if err != nil {
			log.Printf("error parsing message \"%s\" with error %s", string(msgBytes), err)
			continue
		}

		c.processMessageSync(msg)
	}
}

func (c *Client) runWriteMessages() {
	defer c.handleInernalError()

	for {
		msg, ok := <-c.messageChannel

		if !ok {
			log.Printf("messageChannel of client %s has been closed", c.id)
			break
		}

		c.conn.WriteMessage(msg)
	}
}

func (c *Client) processMessageSync(msg Message) error {
	switch msg.Kind {
	case MessageKindID:
		c.idSignal <- msg.ID
	case MessageKindUpdate, MessageKindCurrentState:
		var state State

		err := state.UnmarshalJSON(msg.Content)
		if err != nil {
			return err
		}

		c.engine.importPatch(&state)
	default:
		panic("DA")
	}

	return nil
}
