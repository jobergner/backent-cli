package typescript

import "fmt"

type InterfaceField struct {
	Name     string
	Type     string
	Optional bool
}

func (i InterfaceField) toString() string {
	if i.Optional {
		return fmt.Sprintf("%s?: %s;\n", i.Name, i.Type)
	}
	return fmt.Sprintf("%s: %s;\n", i.Name, i.Type)
}

func (c *Code) Interface(name string, fields ...InterfaceField) *Code {
	c.buf.WriteString(fmt.Sprintf("interface %s {\n", name))

	for _, f := range fields {
		c.buf.WriteString(indent)
		c.buf.WriteString(f.toString())
	}

	c.buf.WriteString("}\n")

	return c
}
