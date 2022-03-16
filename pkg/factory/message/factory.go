package message

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

// isIDTypeOfType evaluates whether a given type name is the respective ID-Type
// of a user-defined type.
// Background:
// Every user-defined type has a generated ID type.
// E.g. a defined type "person" has its ID-Type "PersonID" generated automatically
func (s Factory) isIDTypeOfType(typeName string) bool {
	for _, configType := range s.config.Types {
		if configType.Name+"ID" == typeName {
			return true
		}
	}
	return false
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

	newFactory(file, config).
		writeMessageKinds().
		writeParameters().
		writeResponses()
}
