package typescript

import "fmt"

type EnumField struct {
	Name  string
	Value string
}

func (e EnumField) toString() string {
	return fmt.Sprintf("%s = %s,\n", e.Name, e.Value)
}

func (c *Code) Enum(name string, fields ...EnumField) *Code {
	c.buf.WriteString(fmt.Sprintf("enum %s {\n", name))

	for _, f := range fields {
		c.buf.WriteString(indent)
		c.buf.WriteString(f.toString())
	}

	c.buf.WriteString("}\n")
	return c
}
