package typescript

import (
	"fmt"
	"strings"
)

func (c *Code) Function(name string) *Code {
	c.buf.WriteString(fmt.Sprintf("function %s ", name))
	return c
}

func (c *Code) Param(params ...Param) *Code {
	c.buf.WriteString("(")

	var paramStrings []string
	for _, p := range params {
		paramStrings = append(paramStrings, p.toString())
	}

	c.buf.WriteString(strings.Join(paramStrings, ", "))

	c.buf.WriteString(")")

	return c
}

func (c *Code) ReturnType(typeName string) *Code {
	c.buf.WriteString(": ")

	c.buf.WriteString(typeName)

	c.buf.WriteString(" ")

	return c
}

func (c *Code) Call(params ...*Code) *Code {
	c.buf.WriteString("(")

	var paramStrings []string
	for _, p := range params {
		paramStrings = append(paramStrings, p.String())
	}

	c.buf.WriteString(strings.Join(paramStrings, ", "))

	c.buf.WriteString(")")

	return c
}
