package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/typescript"
)

type Factory struct {
	config *ast.AST
	file   *typescript.Code
}

// Write writes source code for a given StateConfig
func (f *Factory) Write() string {
	f.writeTypeDefinitions()

	return f.file.String()
}

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   typescript.NewCode(),
	}
}

func goTypeToTypescriptType(t string) string {
	switch t {
	case "float64", "int64":
		return "number"
	default:
		return t
	}
}
