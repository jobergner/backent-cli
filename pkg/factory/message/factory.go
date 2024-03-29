package message

import (
	"bytes"

	"github.com/dave/jennifer/jen"

	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/utils"
)

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
		file:   jen.NewFile(utils.PackageName),
	}
}

// Write writes source code for a given ActionsConfig
func (f *Factory) Write() string {
	f.writeMessageKinds().
		writeParameters().
		writeResponses()

	buf := bytes.NewBuffer(nil)
	f.file.Render(buf)

	return utils.TrimPackageClause(buf.String())
}
