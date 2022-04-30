package typescript

import "fmt"

func (c *Code) For(i string, arrayId string) *Code {
	c.buf.WriteString(fmt.Sprintf("for (let %s = 0; %s < %s.length); %s++) ", i, i, arrayId, i))

	return c
}
