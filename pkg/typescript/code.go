package typescript

import (
	"bytes"
)

const (
	indent = "  "
)

type Code struct {
	buf *bytes.Buffer
}

func NewCode() *Code {
	return &Code{
		buf: new(bytes.Buffer),
	}
}

func (c *Code) String() string {
	return c.buf.String()
}
