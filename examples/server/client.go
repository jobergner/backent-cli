package server

import (
	"github.com/google/uuid"
	"github.com/jobergner/backent-cli/examples/connect"
	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/jobergner/backent-cli/examples/message"
	"github.com/rs/zerolog/log"
)

type Client struct {
	lobby          *Lobby
	room           *Room
	conn           connect.Connector
	messageChannel chan []byte
	id             string
}

func newClient(websocketConnector connect.Connector, lobby *Lobby) (*Client, error) {
	clientID, err := uuid.NewRandom()
	if err != nil {
		log.Err(err).Msg("failed generating client ID")
		return nil, err
	}

	c := Client{
		lobby:          lobby,
		conn:           websocketConnector,
		messageChannel: make(chan []byte, 32),
		id:             clientID.String(),
	}

	msg := Message{
		Kind:    message.MessageKindID,
		Content: []byte(c.id),
	}

	msgBytes, err := msg.MarshalJSON()
	if err != nil {
		log.Err(err).Str(logging.MessageKind, string(msg.Kind)).Msg("failed marshalling message")
		return nil, err
	}

	c.messageChannel <- msgBytes

	return &c, nil
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

// closeConnection closes the client's connection
// this does not do anything else at its own, but triggers
// the removal of the client from the system in the
// http handler
func (c *Client) closeConnection() {
	log.Debug().Str(logging.ClientID, c.id).Msg("closing client connection")
	c.conn.Close()
}

func (c *Client) runReadMessages() {
	defer c.closeConnection()

	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg Message
		err = msg.UnmarshalJSON(msgBytes)
		if err != nil {
			log.Err(err).Str(logging.Message, string(msgBytes)).Msg("failed unmarshalling message")

			errMsg, _ := Message{message.MessageIDUnknown, message.MessageKindError, []byte("invalid message"), nil}.MarshalJSON()

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
	defer c.closeConnection()

	for {
		msg, ok := <-c.messageChannel

		if !ok {
			log.Warn().Str(logging.ClientID, c.id).Msg("client message channel was closed")
			break
		}

		c.conn.WriteMessage(msg)
	}
}
