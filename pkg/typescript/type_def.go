package typescript

import "fmt"

type InterfaceField struct {
	Name     string
	Type     *Code
	Optional bool
}

func (i InterfaceField) toString() string {
	if i.Optional {
		return fmt.Sprintf("%s?: %s;\n", i.Name, i.Type.String())
	}
	return fmt.Sprintf("%s: %s;\n", i.Name, i.Type.String())
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

type EnumField struct {
	Name  string
	Value *Code
}

func (e EnumField) String() string {
	return fmt.Sprintf("%s = %s,\n", e.Name, e.Value.String())
}

func (c *Code) Enum(name string, fields ...EnumField) *Code {
	c.buf.WriteString(fmt.Sprintf("enum %s {\n", name))

	for _, f := range fields {
		c.buf.WriteString(indent)
		c.buf.WriteString(f.String())
	}

	c.buf.WriteString("}\n")
	return c
}

func (c *Code) OrType(s string) *Code {
	c.buf.WriteString(fmt.Sprintf(" | %s", s))
	return c
}
