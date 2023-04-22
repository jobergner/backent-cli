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
	f.
		writeClient().
		writeElementKind().
		writeEmitElement().
		writeEmitUpdate().
		writeImportElement().
		writeImportUpdate().
		writeMessageKind().
		writeResponseDefinitions().
		writeStaticCode().
		writeTree().
		writeTypeDefinitions()

	return f.file.String()
}

func NewFactory(config *ast.AST) *Factory {
	return &Factory{
		config: config,
		file:   typescript.NewCode(),
	}
}

func (f *Factory) goTypeToTypescriptType(s string) string {
	for typeName := range f.config.Types {
		if typeName+"ID" == s {
			return "number"
		}
	}

	switch s {
	case "float64", "int64":
		return "number"
	default:
		return s
	}
}

func (s *Factory) rangeTypes(fn func(configType ast.ConfigType) *typescript.Code) []*typescript.Code {
	var code []*typescript.Code
	s.config.RangeTypes(func(configType ast.ConfigType) {
		code = append(code, fn(configType))
	})
	return code
}

func (s *Factory) rangeActions(fn func(action ast.Action) *typescript.Code) []*typescript.Code {
	var code []*typescript.Code
	s.config.RangeActions(func(action ast.Action) {
		code = append(code, fn(action))
	})
	return code
}

func rangeValueTypes(field ast.Field, fn func(configType *ast.ConfigType) *typescript.Code) []*typescript.Code {
	var code []*typescript.Code
	field.RangeValueTypes(func(configType *ast.ConfigType) {
		code = append(code, fn(configType))
	})
	return code
}

func rangeParams(action ast.Action, fn func(param ast.Field) typescript.Param) []typescript.Param {
	var params []typescript.Param
	action.RangeParams(func(param ast.Field) {
		params = append(params, fn(param))
	})
	return params
}

func onlyIf(condition bool, code *typescript.Code) *typescript.Code {
	if condition {
		return code
	}
	return typescript.Empty()
}
