package typescript

func (c *Code) And() *Code {
	c.buf.WriteString(" && ")
	return c
}

func (c *Code) Equals() *Code {
	c.buf.WriteString(" === ")
	return c
}

func (c *Code) EqualsNot() *Code {
	c.buf.WriteString(" !== ")
	return c
}

func (c *Code) Or() *Code {
	c.buf.WriteString(" || ")
	return c
}
