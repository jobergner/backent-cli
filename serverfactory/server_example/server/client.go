package state

import (
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
		return nil, err
	}
	client := Client{
		conn:           websocketConnector,
		messageChannel: make(chan []byte, 256),
		id:             clientID,
	}

	return &client, nil
}

func (c *Client) discontinue() {
	c.room.unregisterChannel <- c
	c.conn.Close()
}

func (c *Client) assignToRoom(room *Room) {
	c.room = room
}

func (c *Client) forwardToRoom(msg message) {
	c.room.clientMessageChannel <- msg
}

func (c *Client) runReadMessages() {
	defer c.discontinue()
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var msg message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Println(err)
		}

		if msg.Kind == messageKindInit {
			c.room.requestInit(c)
		} else {
			c.forwardToRoom(msg)
		}
	}
}

func (c *Client) runWriteMessages() {
	defer c.discontinue()
	for {
		msg, ok := <-c.messageChannel
		if !ok {
			return
		}
		c.conn.WriteMessage(msg)
	}
}
