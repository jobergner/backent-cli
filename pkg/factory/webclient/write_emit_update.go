package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeEmitUpdate() *Factory {
	s.file.Export().Function("emit_Update").Param(Param{Id: "update", Type: Id("Tree")}).Block(s.rangeTypes(func(configType ast.ConfigType) *Code {
		return If(Id("update").Dot(configType.Name).EqualsNot().Null().And().Id("update").Dot(configType.Name).EqualsNot().Undf()).Block(
			ForIn(Const("id"), Id("update").Dot(configType.Name)).Block(
				Id("emit" + Title(configType.Name)).Call(Id("update").Dot(configType.Name).Index(Id("id"))).Sc(),
			),
		)
	})...)

	s.config.RangeTypes(func(configType ast.ConfigType) {

		var typeEmitBody []*Code

		typeEmitBody = append(typeEmitBody,
			If(Id("update").Dot("OperationKind").Equals().Id("OperationKind").Dot("OperationKindDelete").And().Id("elementRegistrar").Index(Id("update").Dot("id")).Equals().Undf()).Block(
				Id("return").Sc(),
			),
			If(Id("update").Dot("OperationKind").Equals().Id("OperationKind").Dot("OperationKindDelete").And().Id("elementRegistrar").Index(Id("update").Dot("id")).EqualsNot().Undf()).Block(
				Delete().Id("elementRegistrar").Index(Id("update").Dot("id")).Sc(),
			),
			If(Id("update").Dot("OperationKind").Equals().Id("OperationKind").Dot("OperationKindUpdate").And().Id("elementRegistrar").Index(Id("update").Dot("id")).Equals().Undf()).Block(
				Id("update").Dot("operationKind").Assign().Id("OperationKind").Dot("OperationKindCreate"),
				Id("elementRegistrar").Index(Id("update").Dot("id")).Assign().Id("true").Sc(),
			),
		)

		configType.RangeFields(func(field ast.Field) {
			if field.HasSliceValue {
				typeEmitBody = append(typeEmitBody, If(Id("update").Dot(field.Name)).EqualsNot().Null().And().Id("update").Dot(field.Name).EqualsNot().Undf().Block(
					ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
						onlyIf(field.HasAnyValue && field.HasSliceValue, fieldEmitAnySlice(field)),
						onlyIf(field.HasSliceValue, fieldEmitSlice(field)),
						onlyIf(field.HasAnyValue, fieldEmitAny(field)),
						onlyIf(field.HasAnyValue, fieldEmit(field)),
					),
				))
			}

		})

		typeEmitBody = append(typeEmitBody,
			Id("eventEmitter").Dot("emit").Call(Id("update").Dot("id"), Id("update")),
		)

		s.file.Function("emit" + Title(configType.Name)).Param(Param{Id: "update", Type: Id(Title(configType.Name))}).Block(
			typeEmitBody...,
		)
	})

	return s
}

func fieldEmitAny(field ast.Field) *Code {
	if field.HasPointerValue {
		return Id("emitElementReference").Call(Id("update").Dot(field.Name))
	}
	return CodeSet(
		rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
			return If(Id("update").Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
				Id("emit" + Title(valueType.Name)).Call(Id("update").Dot(field.Name)),
			)
		})...,
	)
}

func fieldEmitAnySlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			Id("emitElementReference").Call(Id("update").Dot(field.Name).Index(Id("id"))),
		)
	}
	return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
		rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
			return If(Id("update").Dot(field.Name).Index(Id("id")).Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
				Id("emit" + Title(valueType.Name)).Call(Id("update").Dot(field.Name).Index(Id("id")).Dot(field.Name)).Sc(),
			)
		})...,
	)
}

func fieldEmitSlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			Id("emitElementReference").Call(Id("update").Dot(field.Name).Index(Id("id"))),
		)
	}
	return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
		Id("emit" + Title(field.ValueType().Name)).Call(Id("update").Dot(field.Name).Index(Id("id"))),
	)
}

func fieldEmit(field ast.Field) *Code {
	if field.HasPointerValue {
		return Id("emitElementReference").Call(Id("update").Dot(field.Name))
	}
	return Id("emit" + Title(field.Name)).Call(Id("update").Dot(field.Name))
}

func (s *Factory) rangeTypes(fn func(configType ast.ConfigType) *Code) []*Code {
	var code []*Code
	s.config.RangeTypes(func(configType ast.ConfigType) {
		code = append(code, fn(configType))
	})
	return code
}

func rangeValueTypes(field ast.Field, fn func(configType *ast.ConfigType) *Code) []*Code {
	var code []*Code
	field.RangeValueTypes(func(configType *ast.ConfigType) {
		code = append(code, fn(configType))
	})
	return code
}

func onlyIf(condition bool, code *Code) *Code {
	if condition {
		return code
	}
	return Empty()
}
