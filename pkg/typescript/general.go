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

func (c *Code) Id(s string) *Code {
	c.buf.WriteString(s)
	return c
}

func (c *Code) toString() string {
	return c.buf.String()
}

func (c *Code) P(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}
