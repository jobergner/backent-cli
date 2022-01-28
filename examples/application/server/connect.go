package state

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

type Connector interface {
	Close()
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType []byte) error
}

type Connection struct {
	Conn *websocket.Conn
	ctx  context.Context
}

func NewConnection(conn *websocket.Conn, r *http.Request) *Connection {
	return &Connection{
		Conn: conn,
		ctx:  r.Context(),
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
