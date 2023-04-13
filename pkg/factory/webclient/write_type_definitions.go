package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeTypeDefinitions() *Factory {
	s.config.RangeTypes(func(configType ast.ConfigType) {

		var fields []InterfaceField
		fields = append(fields, InterfaceField{
			Name: "id",
			Type: Id("number"),
		})

		configType.RangeFields(func(field ast.Field) {

			typesDef := NewCode()
			first := true

			for _, vt := range field.ValueTypes {
				if vt.IsBasicType {
					typesDef.Id(goTypeToTypescriptType(vt.Name))
					continue
				}

				if !first {
					typesDef.OrType(Title(vt.Name))
					continue
				}

				typesDef.Id(Title(vt.Name))

				first = false
			}

			fields = append(fields, InterfaceField{
				Name:     field.Name,
				Type:     typesDef,
				Optional: true,
			})
		})

		fields = append(fields, InterfaceField{
			Name: "operationKind",
			Type: Id("OperationKind"),
		}, InterfaceField{
			Name: "elementKind",
			Type: Id("ElementKind"),
		})

		s.file.Interface(Title(configType.Name), fields...)
	})

	return s
}
