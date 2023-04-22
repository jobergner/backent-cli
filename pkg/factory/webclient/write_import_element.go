package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeImportElement() *Factory {

	s.config.RangeTypes(func(configType ast.ConfigType) {
		var typeImportBody []*Code

		if configType.IsEvent {
			return
		}

		typeImportBody = append(typeImportBody, If(Id("current").Equals().Null().Or().Id("current").Equals().Undf()).Block(
			Return().Id("update").Sc(),
		))

		configType.RangeFields(func(field ast.Field) {
			if field.ValueType().IsEvent {
				return
			}
			typeImportBody = append(typeImportBody, If(Id("update").Dot(field.Name).EqualsNot().Null().And().Id("update").Dot(field.Name).EqualsNot().Undf()).Block(
				onlyIf(field.ValueType().IsBasicType, Id("current").Dot(field.Name).Assign().Id("update").Dot(field.Name).Sc()),
				onlyIf(field.HasAnyValue && field.HasSliceValue && !field.ValueType().IsBasicType, fieldImportAnySlice(field)),
				onlyIf(!field.HasAnyValue && field.HasSliceValue && !field.ValueType().IsBasicType, fieldImportSlice(field)),
				onlyIf(field.HasAnyValue && !field.HasSliceValue && !field.ValueType().IsBasicType, fieldImportAny(field)),
				onlyIf(!field.HasAnyValue && !field.HasSliceValue && !field.ValueType().IsBasicType, fieldImport(field)),
			))
		})

		typeImportBody = append(typeImportBody,
			Return().Id("current").Sc(),
		)

		s.file.Function("import"+Title(configType.Name)).Param(Param{Id: "current", Type: Id(Title(configType.Name)).OrType("null").OrType("undefined")}, Param{Id: "update", Type: Id(Title(configType.Name))}).ReturnType(Title(configType.Name)).FuncBody(
			typeImportBody...,
		)

	})

	s.file.Function("importElementReference").Param(
		Param{Id: "current", Type: Id("ElementReference").OrType("null").OrType("undefined")},
		Param{Id: "update", Type: Id("ElementReference")},
	).ReturnType("ElementReference").FuncBody(
		Return().Id("update").Sc(),
	)

	return s
}

func fieldImportAny(field ast.Field) *Code {
	if field.HasPointerValue {
		return If(Id("update").Dot(field.Name).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
			Delete().Id("current").Dot(field.Name).Sc(),
		).Id(" else").Block(
			Id("current").Dot(field.Name).Assign().
				Id("importElementReference").Call(Id("current").Dot(field.Name), Id("update").Dot(field.Name)).Sc(),
		)
	}
	return CodeSet(
		rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
			return If(Id("update").Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
				Id("current").Dot(field.Name).Assign().Id("import"+Title(valueType.Name)).Call(
					Id("current").Dot(field.Name).Id(" as "+Title(valueType.Name)),
					Id("update").Dot(field.Name),
				).Sc(),
			)
		})...,
	)
}

func fieldImportAnySlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return CodeSet(
			If(Id("current").Dot(field.Name).Equals().Null().Or().Id("current").Dot(field.Name).Equals().Undf()).Block(
				Id("current").Dot(field.Name).Assign().Id("{}").Sc(),
			),
			ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
				If(Id("update").Dot(field.Name).Index(Id("id")).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
					Delete().Id("current").Dot(field.Name).Index(Id("id")).Sc(),
				).Id(" else").Block(
					Id("current").Dot(field.Name).Index(Id("id")).Assign().
						Id("importElementReference").Call(Id("current").Dot(field.Name).Index(Id("id")), Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
				),
			),
		)
	}

	return CodeSet(
		If(Id("current").Dot(field.Name).Equals().Null().Or().Id("current").Dot(field.Name).Equals().Undf()).Block(
			Id("current").Dot(field.Name).Assign().Id("{}").Sc(),
		),
		ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
				return If(Id("update").Dot(field.Name).Index(Id("id")).Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
					If(Id("update").Dot(field.Name).Index(Id("id")).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
						Delete().Id("current").Dot(field.Name).Index(Id("id")).Sc(),
					).Id(" else").Block(
						Id("current").Dot(field.Name).Index(Id("id")).Assign().
							Id("import"+Title(valueType.Name)).Call(Id("current").Dot(field.Name).Index(Id("id")).Id(" as "+Title(valueType.Name)), Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
					),
				)
			})...,
		),
	)
}

func fieldImportSlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return CodeSet(
			If(Id("current").Dot(field.Name).Equals().Null().Or().Id("current").Dot(field.Name).Equals().Undf()).Block(
				Id("current").Dot(field.Name).Assign().Id("{}").Sc(),
			),
			ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
				If(Id("update").Dot(field.Name).Index(Id("id")).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
					Delete().Id("current").Dot(field.Name).Index(Id("id")).Sc(),
				).Id(" else").Block(
					Id("current").Dot(field.Name).Index(Id("id")).Assign().
						Id("importElementReference").Call(Id("current").Dot(field.Name).Index(Id("id")), Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
				),
			),
		)
	}
	return CodeSet(
		If(Id("current").Dot(field.Name).Equals().Null().Or().Id("current").Dot(field.Name).Equals().Undf()).Block(
			Id("current").Dot(field.Name).Assign().Id("{}").Sc(),
		),
		ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			If(Id("update").Dot(field.Name).Index(Id("id")).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
				Delete().Id("current").Dot(field.Name).Index(Id("id")).Sc(),
			).Id(" else").Block(
				Id("current").Dot(field.Name).Index(Id("id")).Assign().
					Id("import"+Title(field.ValueType().Name)).Call(Id("current").Dot(field.Name).Index(Id("id")), Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
			),
		),
	)
}

func fieldImport(field ast.Field) *Code {
	if field.HasPointerValue {
		return If(Id("update").Dot(field.Name).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
			Delete().Id("current").Dot(field.Name).Sc(),
		).Id(" else").Block(
			Id("current").Dot(field.Name).Assign().
				Id("importElementReference").Call(Id("current").Dot(field.Name), Id("update").Dot(field.Name)).Sc(),
		)
	}
	return Id("current").Dot(field.Name).Assign().Id("import"+Title(field.ValueType().Name)).Call(Id("current").Dot(field.Name), Id("update").Dot(field.Name)).Sc()
}
