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
	id             uuid.UUID
}

func newClient(websocketConnector Connector, handler *LoginHandler) (*Client, error) {
	clientID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating client ID: %s", err)
	}

	c := Client{
		handler:        handler,
		conn:           websocketConnector,
		messageChannel: make(chan []byte, 32),
		id:             clientID,
	}

	return &c, nil
}

func (c *Client) handleRoomKick() {
	c.conn.Close()
}

func (c *Client) handleInernalError() {
	c.room.unregisterClientSync(c)
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
