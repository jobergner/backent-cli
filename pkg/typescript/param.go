package typescript

import (
	"fmt"
)

type Param struct {
	Id   string
	Type string
}

func (p Param) toString() string {
	return fmt.Sprintf("%s : %s", p.Id, p.Type)
}
