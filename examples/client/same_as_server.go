package state

import (
	"context"
	"fmt"
	"log"

	"nhooyr.io/websocket"
)

type MessageKind string

const (
	// server -> client
	MessageKindError        MessageKind = "error"
	MessageKindCurrentState MessageKind = "currentState"
	MessageKindUpdate       MessageKind = "update"
	MessageKindID           MessageKind = "id"
	// responses to messages which fail to unmarshal
	MessageIDUnknown string = "unknown"
	// client -> server
	MessageKindGlobal MessageKind = "global"
)

type Message struct {
	ID      string      `json:"id"`
	Kind    MessageKind `json:"kind"`
	Content []byte      `json:"content"`
}

func printMessage(msg Message) string {
	b, err := msg.MarshalJSON()
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}

func messageUnmarshallingError(msgContent []byte, err error) []byte {
	return []byte(fmt.Sprintf("error when unmarshalling received message content `%s`: %s", msgContent, err))
}

func responseMarshallingError(msgContent []byte, err error) []byte {
	return []byte(fmt.Sprintf("error when marshalling response to `%s`: %s", msgContent, err))
}

type Connector interface {
	Close()
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType []byte) error
}

type Connection struct {
	Conn *websocket.Conn
	ctx  context.Context
}

func NewConnection(ctx context.Context, conn *websocket.Conn) *Connection {
	return &Connection{
		Conn: conn,
		ctx:  ctx,
	}
}

func (c *Connection) Close() {
	err := c.Conn.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		log.Printf("error closing client connection %s", err)
	}
}

func (c *Connection) ReadMessage() (int, []byte, error) {
	msgType, msg, err := c.Conn.Read(c.ctx)

	if err != nil {
		return 0, nil, fmt.Errorf("error reading message from connection: %s", err)
	}

	return int(msgType), msg, nil
}

func (c *Connection) WriteMessage(msg []byte) error {
	err := c.Conn.Write(c.ctx, websocket.MessageText, msg)

	if err != nil {
		return err
	}

	return nil
}
