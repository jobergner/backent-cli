package typescript

import (
	"fmt"
	"strings"
)

func (c *Code) Null() *Code {
	c.buf.WriteString("null")
	return c
}

func (c *Code) Undf() *Code {
	c.buf.WriteString("undefined")
	return c
}

func (c *Code) Index(code *Code) *Code {
	c.buf.WriteString(fmt.Sprintf("[%s]", code.String()))
	return c
}

func (c *Code) Id(s string) *Code {
	c.buf.WriteString(fmt.Sprintf("%s", s))
	return c
}

func (c *Code) Dot(s string) *Code {
	c.buf.WriteString(fmt.Sprintf(".%s", s))
	return c
}

func (c *Code) Is(code *Code) *Code {
	c.buf.WriteString(fmt.Sprintf(": %s", code.String()))
	return c
}

func (c *Code) If(code *Code) *Code {
	c.buf.WriteString("if (")
	c.buf.WriteString(code.String())
	c.buf.WriteString(")")
	return c
}

func (c *Code) ForIn(decl, iterable *Code) *Code {
	c.buf.WriteString(fmt.Sprintf("for (%s in %s)", decl.String(), iterable.String()))
	return c
}

func (c *Code) CodeSet(code ...*Code) *Code {
	for i, segment := range code {
		s := segment.String()
		if len(s) == 0 {
			continue
		}
		c.buf.WriteString(s)
		if i == len(code)-1 {
			continue
		}
		c.buf.WriteString("\n")
	}

	return c
}

func (c *Code) FuncBody(code ...*Code) *Code {
	c.Block(code...)
	c.buf.WriteString("\n")
	return c
}

func (c *Code) Block(code ...*Code) *Code {
	c.buf.WriteString(" {\n")

	for _, segment := range code {
		s := segment.String()
		if len(s) == 0 {
			continue
		}
		lines := strings.Split(s, "\n")
		for i := range lines {
			lines[i] = indent + lines[i]
		}
		s = strings.Join(lines, "\n")
		c.buf.WriteString(s)
		c.buf.WriteString("\n")
	}

	c.buf.WriteString("}")

	return c
}

type ObjectField struct {
	Id   *Code
	Type *Code
}

func (o ObjectField) toString() string {
	return fmt.Sprintf("%s: %s", o.Id, o.Type.String())
}

func (c *Code) Object(fields ...ObjectField) *Code {

	var fieldStrings []string
	for _, f := range fields {
		fieldStrings = append(fieldStrings, f.toString())
	}

	c.buf.WriteString(fmt.Sprintf("{%s}", strings.Join(fieldStrings, ", ")))
	return c
}

func (c *Code) ObjectSpaced(fields ...ObjectField) *Code {

	var fieldStrings []string
	for _, f := range fields {
		fieldStrings = append(fieldStrings, f.toString())
	}

	c.buf.WriteString(fmt.Sprintf("{ %s }", strings.Join(fieldStrings, ", ")))
	return c
}

func (c *Code) Return() *Code {
	c.buf.WriteString("return ")
	return c
}

func (c *Code) Assign() *Code {
	c.buf.WriteString(" = ")
	return c
}

func (c *Code) Sc() *Code {
	c.buf.WriteString(";")
	return c
}

func (c *Code) Export() *Code {
	c.buf.WriteString("export ")
	return c
}

func (c *Code) Delete() *Code {
	c.buf.WriteString("delete ")
	return c
}

func (c *Code) Class() *Code {
	c.buf.WriteString("class ")
	return c
}

func (c *Code) Public() *Code {
	c.buf.WriteString("public ")
	return c
}

func (c *Code) Private() *Code {
	c.buf.WriteString("private ")
	return c
}

func (c *Code) New() *Code {
	c.buf.WriteString("new ")
	return c
}

func (c *Code) Promise(typeName string) *Code {
	c.buf.WriteString(fmt.Sprintf("Promise<%s> ", typeName))
	return c
}

func (c *Code) ArrowFunc(code ...*Code) *Code {
	c.Call(code...)
	c.buf.WriteString("=> ")
	return c
}

func (c *Code) Switch() *Code {
	c.buf.WriteString("switch ")
	return c
}

func (c *Code) Case(matcher *Code, code ...*Code) *Code {
	c.buf.WriteString(fmt.Sprintf("case %s:", matcher.String()))

	for _, segment := range code {
		s := segment.String()
		if len(s) == 0 {
			continue
		}
		lines := strings.Split(s, "\n")
		for i := range lines {
			lines[i] = indent + lines[i]
		}
		s = strings.Join(lines, "\n")
		c.buf.WriteString(s)
		c.buf.WriteString("\n")
	}

	return c
}

func (c *Code) Default(code ...*Code) *Code {
	c.buf.WriteString("default:")

	for _, segment := range code {
		s := segment.String()
		if len(s) == 0 {
			continue
		}
		lines := strings.Split(s, "\n")
		for i := range lines {
			lines[i] = indent + lines[i]
		}
		s = strings.Join(lines, "\n")
		c.buf.WriteString(s)
		c.buf.WriteString("\n")
	}

	return c
}
