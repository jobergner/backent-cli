package typescript

import (
	"fmt"
	"strings"
)

func (c *Code) Index(s string) *Code {
	c.buf.WriteString(fmt.Sprintf("[%s] ", s))
	return c
}

func (c *Code) Id(s string) *Code {
	c.buf.WriteString(fmt.Sprintf("%s ", s))
	return c
}

func (c *Code) Dot(s string) *Code {
	c.buf.WriteString(fmt.Sprintf(".%s ", s))
	return c
}

func (c *Code) Is(s string) *Code {
	c.buf.WriteString(fmt.Sprintf(": %s ", s))
	return c
}

func (c *Code) If(code *Code) *Code {
	c.buf.WriteString("if (")
	c.buf.WriteString(code.toString())
	c.buf.WriteString(") ")
	return c
}

func (c *Code) ForIn(decl, iterable *Code) *Code {
	c.buf.WriteString(fmt.Sprintf("for (%s in %s) ", decl.toString(), iterable.toString()))
	return c
}

func (c *Code) Block(code ...*Code) *Code {
	c.buf.WriteString("{\n")

	for _, line := range code {
		c.buf.WriteString(indent)
		c.buf.WriteString(line.toString())
		c.buf.WriteString("\n")
	}

	c.buf.WriteString("}")

	return c
}

type ObjectField struct {
	Id   string
	Type *Code
}

func (o ObjectField) toString() string {
	return fmt.Sprintf("%s : %s", o.Id, o.Type.toString())
}

func (c *Code) Object(fields ...ObjectField) *Code {

	var fieldStrings []string
	for _, f := range fields {
		fieldStrings = append(fieldStrings, f.toString())
	}

	c.buf.WriteString(fmt.Sprintf("{%s} ", strings.Join(fieldStrings, ", ")))
	return c
}

func (c *Code) Return(s string) *Code {
	c.buf.WriteString(fmt.Sprintf("return %s", s))
	return c
}
