package webclient

import (
	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeResponseDefinitions() *Factory {

	s.config.RangeActions(func(action ast.Action) {

		if action.Response == nil {
			return
		}

		var fields []InterfaceField

		ForEachResponseValueInAction(action, func(param ast.Field) *jen.Statement {
			typesDef := NewCode()

			param.RangeValueTypes(func(vt *ast.ConfigType) {
				typesDef.Id(s.goTypeToTypescriptType(vt.Name))
			})

			if param.HasSliceValue {
				typesDef.Index(Empty())
			}

			fields = append(fields, InterfaceField{
				Name: param.Name,
				Type: typesDef,
			})
			return nil
		})

		s.file.Export().Interface(Title(action.Name)+"Response", fields...)
	})

	return s
}
