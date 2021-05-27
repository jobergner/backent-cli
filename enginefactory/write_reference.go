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

			decls.File.Func().Params(Id("engine").Id("*Engine")).Id("dereference"+title(field.ValueTypeName)+title(valueType.Name)+"Refs").Params(Id(valueType.Name+"ID").Id(title(valueType.Name)+"ID")).Block(
				Id("ref").Op(":=").Id("_ref").Dot(field.ValueTypeName).Dot("engine").Dot(field.ValueTypeName).Call(Id("_ref").Dot(field.ValueTypeName).Dot("ID")),
				Return(Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot(getReturnType).Call(Id("ref").Dot(field.ValueTypeName).Dot("ReferencedElementID"))),
				For(List(Id("_"), Id("refID")).Op(":=").Range().Id("engine").Dot("all"+title(field.ValueTypeName)+"IDs")).Block(
					Id("ref").Op(":=").Id("engine").Dot(field.ValueTypeName).Call(Id("refID")),
					onlyIf(field.HasSliceValue, &Statement{}),
					onlyIf(!field.HasSliceValue, &Statement{}),
				),
			)
		})

	})

	decls.Render(s.buf)
	return s
}
