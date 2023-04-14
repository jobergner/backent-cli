package typescript

import (
	"fmt"
)

type Param struct {
	Id   string
	Type *Code
}

func (p Param) String() string {
	return fmt.Sprintf("%s: %s", p.Id, p.Type.String())
}
