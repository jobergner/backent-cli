package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeleters() *EngineFactory {
	s.config.RangeTypes(func(configType ast.ConfigType) {

		w := deleteTypeWrapperWriter{
			t: configType,
		}

		s.file.Func().Params(w.receiverParams()).Id(w.name()).Params(w.params()).Block(
			OnlyIf(!configType.IsRootType, w.getElement()),
			OnlyIf(!configType.IsRootType, If(w.hasParent()).Block(
				Return(),
			)),
			w.deleteElement(),
		)

		d := deleteTypeWriter{
			t: configType,
			f: nil,
		}

		s.file.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params()).Block(
			d.getElement(),
			If(d.isOperationKindDelete()).Block(
				Return(),
			),
			ForEachReferenceOfType(configType, func(field *ast.Field) *Statement {
				return d.dereferenceField(field)
			}),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				d.f = &field
				if field.ValueType().IsBasicType {
					return Empty()
				}
				if field.HasSliceValue {
					return For(d.loopConditions().Block(
						d.deleteElementInLoop(),
					))
				}
				return d.deleteElement()
			}),
			If(d.existsInState()).Block(
				d.setOperationKind(),
				d.updateElementInPatch(),
			).Else().Block(
				d.deleteFromPatch(),
			),
		)
	})

	s.config.RangeRefFields(func(field ast.Field) {
		d := deleteGeneratedTypeWriter{
			f:             field,
			valueTypeName: func() string { return field.ValueTypeName },
		}

		s.file.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params()).Block(
			d.getElement(),
			If(d.isOperationKindDelete()).Block(
				Return(),
			),
			OnlyIf(d.f.HasAnyValue, d.deleteAnyContainer()),
			If(d.existsInState()).Block(
				d.setOperationKind(),
				d.updateElementInPatch(),
			).Else().Block(
				d.deleteFromPatch(),
			),
		)
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		d := deleteGeneratedTypeWriter{
			f:             field,
			valueTypeName: func() string { return anyNameByField(field) },
		}

		s.file.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params(), Id("deleteChild").Bool()).Block(
			d.getElement(),
			If(d.isOperationKindDelete()).Block(
				Return(),
			),
			If(Id("deleteChild")).Block(
				d.deleteChild(),
			),
			If(d.existsInState()).Block(
				d.setOperationKind(),
				d.updateElementInPatch(),
			).Else().Block(
				d.deleteFromPatch(),
			),
		)
	})

	return s
}
