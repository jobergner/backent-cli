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

			for _, vt := range sortValueTypes(field.ValueTypes) {
				if field.HasPointerValue {
					typesDef.Id("ElementReference")
					break
				}

				if vt.IsBasicType {
					typesDef.Id(goTypeToTypescriptType(vt.Name))
					break
				}

				if !first {
					typesDef.OrType(Title(vt.Name))
					continue
				}

				typesDef.Id(Title(vt.Name))

				first = false
			}

			if field.HasSliceValue {
				if field.ValueType().IsBasicType {
					typesDef.Index(Empty())
				} else {
					typesDef = Object(ObjectField{
						Id:   Index(Id("id").Is(Id("number"))),
						Type: typesDef,
					})
				}
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

	s.file.Interface("ElementReference", []InterfaceField{
		{
			Name: "operationKind",
			Type: Id("OperationKind"),
		},
		{
			Name: "elementID",
			Type: Id("number"),
		},
		{
			Name: "elementKind",
			Type: Id("ElementKind"),
		},
		{
			Name: "referencedDataStatus",
			Type: Id("ReferencedDataStatus"),
		},
		{
			Name: "elementPath",
			Type: Id("string"),
		},
	}...)

	return s
}
