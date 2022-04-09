package connect

import (
	"context"

	"github.com/jobergner/backent-cli/examples/logging"
	"github.com/rs/zerolog/log"
	"nhooyr.io/websocket"
)

type Connector interface {
	Close(reason string)
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType []byte) error
}

type Connection struct {
	Conn     *websocket.Conn
	ctx      context.Context
	cancelFn func()
}

func NewConnection(conn *websocket.Conn, ctx context.Context, cancel func()) *Connection {
	return &Connection{
		Conn:     conn,
		ctx:      ctx,
		cancelFn: cancel,
	}
}

func (c *Connection) Close(reason string) {
	err := c.Conn.Close(websocket.StatusNormalClosure, reason)
	if err != nil {
		log.Warn().Msg("failed closing connection")
	}
	c.cancelFn()
}

func (c *Connection) ReadMessage() (int, []byte, error) {
	msgType, msg, err := c.Conn.Read(c.ctx)

	if err != nil {
		log.Err(err).Msg("failed reading from connection")
		return 0, nil, err
	}

	return int(msgType), msg, nil
}

func (c *Connection) WriteMessage(msg []byte) error {
	err := c.Conn.Write(c.ctx, websocket.MessageText, msg)

	if err != nil {
		log.Err(err).Str(logging.Message, string(msg)).Msg("failed writing to connection")
		return err
	}

	return nil
}
