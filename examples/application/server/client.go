package state

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Client struct {
	handler        *LoginHandler
	room           *Room
	conn           Connector
	messageChannel chan []byte
	id             string
}

func newClient(websocketConnector Connector) (*Client, error) {
	clientID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating client ID: %s", err)
	}

	c := Client{
		handler:        nil,
		conn:           websocketConnector,
		messageChannel: make(chan []byte, 32),
		id:             clientID.String(),
	}

	return &c, nil
}

func (c *Client) SendMessage(msg []byte) {
	c.messageChannel <- msg
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) LeaveRoom() {
	c.room.unregisterClientSync(c)
}

func (c *Client) JoinRoom(room *Room) {
	c.room.unregisterClientSync(c)
	room.registerClientSync(c)
}

func (c *Client) Room() *Room {
	return c.room
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

			errDescription := messageUnmarshallingError(msgBytes, err)

			errMsg, err := Message{MessageKindError, errDescription, nil}.MarshalJSON()
			if err != nil {
				log.Printf("error marshalling error message \"%s\"", errDescription)
				continue
			}

			c.messageChannel <- errMsg

			continue
		}

		msg.client = c

		if msg.Kind == MessageKindGlobal {
			c.handler.processMessageSync(msg)
		} else {
			c.room.processMessageSync(msg)
		}
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
