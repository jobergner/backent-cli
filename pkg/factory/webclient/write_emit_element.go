package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeEmitElement() *Factory {

	registrarCheck := []*Code{
		If(Id("update").Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete").And().Id("elementRegistrar").Index(Id("update").Dot("id")).Equals().Undf()).Block(
			Id("return").Sc(),
		),
		If(Id("update").Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete").And().Id("elementRegistrar").Index(Id("update").Dot("id")).EqualsNot().Undf()).Block(
			Delete().Id("elementRegistrar").Index(Id("update").Dot("id")).Sc(),
		),
		If(Id("update").Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindUpdate").And().Id("elementRegistrar").Index(Id("update").Dot("id")).Equals().Undf()).Block(
			Id("update").Dot("operationKind").Assign().Id("OperationKind").Dot("OperationKindCreate").Sc(),
			Id("elementRegistrar").Index(Id("update").Dot("id")).Assign().Id("true").Sc(),
		),
	}

	s.config.RangeTypes(func(configType ast.ConfigType) {
		var typeEmitBody []*Code

		typeEmitBody = append(typeEmitBody, registrarCheck...)

		configType.RangeFields(func(field ast.Field) {
			if field.ValueType().IsBasicType {
				return
			}
			typeEmitBody = append(typeEmitBody, If(Id("update").Dot(field.Name).EqualsNot().Null().And().Id("update").Dot(field.Name).EqualsNot().Undf()).Block(
				onlyIf(field.HasAnyValue && field.HasSliceValue, fieldEmitAnySlice(field)),
				onlyIf(!field.HasAnyValue && field.HasSliceValue, fieldEmitSlice(field)),
				onlyIf(field.HasAnyValue && !field.HasSliceValue, fieldEmitAny(field)),
				onlyIf(!field.HasAnyValue && !field.HasSliceValue, fieldEmit(field)),
			))
		})

		typeEmitBody = append(typeEmitBody,
			Id("eventEmitter").Dot("emit").Call(Id("update").Dot("id"), Id("update")).Sc(),
		)

		s.file.Function("emit" + Title(configType.Name)).Param(Param{Id: "update", Type: Id(Title(configType.Name))}).FuncBody(
			typeEmitBody...,
		)
	})

	var emitElementReferenceBody []*Code
	emitElementReferenceBody = append(emitElementReferenceBody, registrarCheck...)
	emitElementReferenceBody = append(
		emitElementReferenceBody,
		Id("eventEmitter").Dot("emit").Call(Id("update").Dot("id"), Id("update")).Sc(),
	)
	s.file.Function("emitElementReference").Param(Param{Id: "update", Type: Id("ElementReference")}).FuncBody(
		emitElementReferenceBody...,
	)

	return s
}

func fieldEmitAny(field ast.Field) *Code {
	if field.HasPointerValue {
		return Id("emitElementReference").Call(Id("update").Dot(field.Name)).Sc()
	}
	return CodeSet(
		rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
			return If(Id("update").Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
				Id("emit" + Title(valueType.Name)).Call(Id("update").Dot(field.Name)).Sc(),
			)
		})...,
	)
}

func fieldEmitAnySlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			Id("emitElementReference").Call(Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
		)
	}
	return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
		rangeValueTypes(field, func(valueType *ast.ConfigType) *Code {
			return If(Id("update").Dot(field.Name).Index(Id("id")).Dot("elementKind").Equals().Id("ElementKind").Dot("ElementKind" + Title(valueType.Name))).Block(
				Id("emit" + Title(valueType.Name)).Call(Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
			)
		})...,
	)
}

func fieldEmitSlice(field ast.Field) *Code {
	if field.HasPointerValue {
		return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
			Id("emitElementReference").Call(Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
		)
	}
	return ForIn(Const("id"), Id("update").Dot(field.Name)).Block(
		Id("emit" + Title(field.ValueType().Name)).Call(Id("update").Dot(field.Name).Index(Id("id"))).Sc(),
	)
}

func fieldEmit(field ast.Field) *Code {
	if field.HasPointerValue {
		return Id("emitElementReference").Call(Id("update").Dot(field.Name)).Sc()
	}
	return Id("emit" + Title(field.Name)).Call(Id("update").Dot(field.Name)).Sc()
}
