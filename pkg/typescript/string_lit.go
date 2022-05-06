package typescript

import "fmt"

func (c *Code) S(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}
