package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeReference() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {

		if !field.HasSliceValue {
			decls.File.Func().Params(Id("_ref").Id(field.ValueTypeName)).Id("IsSet").Params().Bool().Block(
				Id("ref").Op(":=").Id("_ref").Dot(field.ValueTypeName).Dot("engine").Dot(field.ValueTypeName).Call(Id("_ref").Dot(field.ValueTypeName).Dot("ID")),
				Return(Id("ref").Dot(field.ValueTypeName).Dot("ID")).Op("!=").Lit(0),
			)

			decls.File.Func().Params(Id("_ref").Id(field.ValueTypeName)).Id("Unset").Params().Block(
				Id("ref").Op(":=").Id("_ref").Dot(field.ValueTypeName).Dot("engine").Dot(field.ValueTypeName).Call(Id("_ref").Dot(field.ValueTypeName).Dot("ID")),
				Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot("delete"+title(field.ValueTypeName)).Call(Id("ref").Dot(field.ValueTypeName).Dot("ID")),
				Id("parent").Op(":=").Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot(title(field.Parent.Name)).Call(Id("ref").Dot(field.ValueTypeName).Dot("ParentID")).Dot(field.Parent.Name),
				If(Id("parent").Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
					Return(),
				),
				Id("parent").Dot(title(field.Name)).Op("=").Lit(0),
				Id("parent").Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
				Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot("Patch").Dot(title(field.Parent.Name)).Index(Id("parent").Dot("ID")).Op("=").Id("parent"),
			)
		}

		getReturnType := title(field.ValueType().Name)
		if field.HasAnyValue {
			getReturnType = anyNameByField(field)
		}

		decls.File.Func().Params(Id("_ref").Id(field.ValueTypeName)).Id("Get").Params().Id(lower(getReturnType)).Block(
			Id("ref").Op(":=").Id("_ref").Dot(field.ValueTypeName).Dot("engine").Dot(field.ValueTypeName).Call(Id("_ref").Dot(field.ValueTypeName).Dot("ID")),
			Return(Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot(getReturnType).Call(Id("ref").Dot(field.ValueTypeName).Dot("ReferencedElementID"))),
		)

		field.RangeValueTypes(func(valueType *ast.ConfigType) {

			dereferenceCondition := Id("ref").Dot(field.ValueTypeName).Dot("ReferencedElementID").Op("==").Id(valueType.Name + "ID")
			if field.HasAnyValue {
				dereferenceCondition = Id("anyContainer").Dot(anyNameByField(field)).Dot(title(valueType.Name)).Op("==").Id(valueType.Name + "ID")
			}

			var optionalSuffix string
			if len(field.ValueTypes) > 1 {
				optionalSuffix = title(valueType.Name)
			}

			decls.File.Func().Params(Id("engine").Id("*Engine")).Id("dereference" + title(field.ValueTypeName) + "s" + optionalSuffix).Params(Id(valueType.Name + "ID").Id(title(valueType.Name) + "ID")).Block(
				For(List(Id("_"), Id("refID")).Op(":=").Range().Id("engine").Dot("all"+title(field.ValueTypeName)+"IDs").Call()).Block(
					Id("ref").Op(":=").Id("engine").Dot(field.ValueTypeName).Call(Id("refID")),
					onlyIf(field.HasAnyValue, Id("anyContainer").Op(":=").Id("ref").Dot("Get").Call()),
					onlyIf(field.HasAnyValue, If(Id("anyContainer").Dot(anyNameByField(field)).Dot("ElementKind").Op("!=").Id("ElementKind"+title(valueType.Name))).Block(
						Continue(),
					)),
					If(dereferenceCondition).Block(
						onlyIf(field.HasSliceValue, &Statement{
							Id("parent").Op(":=").Id("engine").Dot(title(field.Parent.Name)).Call(Id("ref").Dot(field.ValueTypeName).Dot("ParentID")).Line(),
							Id("parent").Dot("Remove" + title(field.Name) + optionalSuffix).Call(Id(valueType.Name + "ID")),
						}),
						onlyIf(!field.HasSliceValue, &Statement{
							Id("ref").Dot("Unset").Call(),
						}),
					),
				),
			)
		})

	})

	decls.Render(s.buf)
	return s
}
