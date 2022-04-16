package server

import (
	"bytes"

	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/configs"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, configs.ActionsConfig, configs.ResponsesConfig)
	return simpleAST
}

type Factory struct {
	config *ast.AST
	file   *jen.File
}

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   jen.NewFile("server"),
	}
}

// Write writes source code for a given ActionsConfig
func (f *Factory) Write(buf *bytes.Buffer) {
	f.writeProcessClientMessage().
		writeController()

	f.file.Render(buf)
}
