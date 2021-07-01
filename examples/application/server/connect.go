package state

import (
	"context"
	"fmt"
	"net/http"

	"nhooyr.io/websocket"
)

type Connector interface {
	Close()
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType []byte) error
}

type Connection struct {
	Conn          *websocket.Conn
	ctx           context.Context
	cancelContext context.CancelFunc
}

func NewConnection(conn *websocket.Conn, r *http.Request) *Connection {
	return &Connection{
		Conn: conn,
		ctx:  context.Background(),
	}
}

func (c *Connection) Close() {
	c.Conn.Close(websocket.StatusNormalClosure, "")
	c.cancelContext()
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
