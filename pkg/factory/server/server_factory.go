package server

import (
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
	configJson []byte,
) {

	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)

	newFactory(file, config).writeProcessClientMessage().writeController()
}
