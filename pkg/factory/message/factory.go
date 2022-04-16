package message

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

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   jen.NewFile("message"),
	}
}

// Write writes source code for a given ActionsConfig
func (f *Factory) Write(buf *bytes.Buffer) {
	f.writeMessageKinds().
		writeParameters().
		writeResponses()

	f.file.Render(buf)
}
