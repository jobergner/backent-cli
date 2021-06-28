package state

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Client struct {
	room           *Room
	conn           Connector
	messageChannel chan []byte
	id             uuid.UUID
}

func newClient(websocketConnector Connector) (*Client, error) {
	clientID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating client ID: %s", err)
	}
	c := Client{
		conn:           websocketConnector,
		messageChannel: make(chan []byte, 32),
		id:             clientID,
	}

	return &c, nil
}

func (c *Client) discontinue() {
	c.room.unregisterChannel <- c
	c.conn.Close()
}

func (c *Client) assignToRoom(room *Room) {
	c.room = room
}

func (c *Client) forwardToRoom(msg Message) {
	select {
	case c.room.clientMessageChannel <- msg:
	default:
		log.Println("room's message buffer full -> message dropped:")
		log.Println(printMessage(msg))
	}
}

func (c *Client) runReadMessages() {
	defer c.discontinue()
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("error while reading connection: %s", err)
			continue
		}

		var msg Message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Printf("error parsing message \"%s\" with error %s", string(msgBytes), err)
			c.room.pendingResponsesChannel <- Message{MessageKindError, messageUnmarshallingError(msgBytes, err), c}
			continue
		}

		msg.client = c
		c.forwardToRoom(msg)
	}
}

func (c *Client) runWriteMessages() {
	defer c.discontinue()
	for {
		msg, ok := <-c.messageChannel
		if !ok {
			log.Printf("messageChannel of client %s has been closed", c.id)
			return
		}
		c.conn.WriteMessage(msg)
	}
}
