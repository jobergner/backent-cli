package typescript

import (
	"bytes"
)

const (
	indent = "	"
)

type Code struct {
	buf *bytes.Buffer
}

func NewCode() *Code {
	return &Code{
		buf: new(bytes.Buffer),
	}
}

func Const(name string) *Code {
	return NewCode().Const(name)
}

func Let(name string) *Code {
	return NewCode().Let(name)
}

func Function(name string) *Code {
	return NewCode().Function(name)
}

func Interface(name string, fields ...InterfaceField) *Code {
	return NewCode().Interface(name, fields...)
}
