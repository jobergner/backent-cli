package typescript

import "fmt"

func (c *Code) Const(name string) *Code {
	c.buf.WriteString(fmt.Sprintf("const %s = ", name))
	return c
}

func (c *Code) Let(name string) *Code {
	c.buf.WriteString(fmt.Sprintf("let %s = ", name))
	return c
}
