package typescript

import "fmt"

func (c *Code) Lit(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}
