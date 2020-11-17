package statefactory

import (
	"fmt"
	"go/ast"
)

type stateMachineBuilder struct {
	input        *ast.File
	stateMachine *ast.File
}

func newStateMachineBuilder(input *ast.File) *stateMachineBuilder {
	return &stateMachineBuilder{
		input:        input,
		stateMachine: &ast.File{},
	}
}

func main() {
	fmt.Println("vim-go")
}
