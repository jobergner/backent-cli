package client

import (
	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/examples/configs"
	"github.com/jobergner/backent-cli/pkg/ast"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, configs.ActionsConfig, configs.ResponsesConfig)
	return simpleAST
}

type Factory struct {
	config *ast.AST
	file   *jen.File
}

func newFactory(file *jen.File, config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   file,
	}
}

// Write writes source code for a given ActionsConfig
func Write(
	file *jen.File,
	stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{},
) {

	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)

	newFactory(file, config).writeActions()
}
