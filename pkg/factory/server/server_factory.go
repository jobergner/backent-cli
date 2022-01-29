package server

import (
	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/pkg/ast"
)

type ServerFactory struct {
	config *ast.AST
	file   *jen.File
}

// isIDTypeOfType evaluates whether a given type name is the respective ID-Type
// of a user-defined type.
// Background:
// Every user-defined type has a generated ID type.
// E.g. a defined type "person" has its ID-Type "PersonID" generated automatically
func (s ServerFactory) isIDTypeOfType(typeName string) bool {
	for _, configType := range s.config.Types {
		if configType.Name+"ID" == typeName {
			return true
		}
	}
	return false
}

func newServerFactory(file *jen.File, config *ast.AST) *ServerFactory {
	return &ServerFactory{
		config: config,
		file:   file,
	}
}

// WriteServerFrom writes source code for a given ActionsConfig
func WriteServer(
	file *jen.File,
	stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{},
	configJson []byte,
) {

	config := ast.Parse(stateConfigData, actionsConfigData, responsesConfigData)

	newServerFactory(file, config).
		writeMessageKinds().
		writeParameters().
		writeResponses().
		writeProcessClientMessage().
		writeInspectHandler(configJson).
		writeActions().
		writeSideEffects()
}
