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
				Id(field.Parent.Name).Op(":=").Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot(title(field.Parent.Name)).Call(Id("ref").Dot(field.ValueTypeName).Dot("ParentID")).Dot(field.Parent.Name),
				If(Id(field.Parent.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
					Return(),
				),
				Id(field.Parent.Name).Dot(title(field.Name)).Op("=").Lit(0),
				Id(field.Parent.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate"),
				Id("ref").Dot(field.ValueTypeName).Dot("engine").Dot("Patch").Dot(title(field.Parent.Name)).Index(Id(field.Parent.Name).Dot("ID")).Op("=").Id(field.Parent.Name),
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
	})

	decls.Render(s.buf)
	return s
}
