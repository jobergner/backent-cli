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
			done := false

			field.RangeValueTypes(func(vt *ast.ConfigType) {
				if done {
					return
				}

				if field.HasPointerValue {
					typesDef.Id("ElementReference")
					done = true
					return
				}

				if vt.IsBasicType {
					typesDef.Id(s.goTypeToTypescriptType(vt.Name))
					done = true
					return
				}

				if !first {
					typesDef.OrType(Title(vt.Name))
					return
				}

				typesDef.Id(Title(vt.Name))

				first = false
			})

			if field.HasSliceValue {
				if field.ValueType().IsBasicType {
					typesDef.Index(Empty())
				} else {
					typesDef = ObjectSpaced(ObjectField{
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
			Name: "id",
			Type: Id("number"),
		},
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
