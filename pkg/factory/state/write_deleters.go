package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeDeleters() *Factory {

	RangeBasicTypes(func(b BasicType) {
		d := deleteTypeWriter{
			typeName: b.Value,
			f:        nil,
		}

		s.file.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params()).Block(
			Id(b.Value).Op(":=").Id("engine").Dot(b.Value).Call(Id(b.Value+"ID")),
			If(d.isOperationKindDelete()).Block(
				Return(),
			),
			If(d.existsInState()).Block(
				d.setOperationKind(),
				d.updateElementInPatch(),
			).Else().Block(
				d.deleteFromPatch(),
			),
		)
	})

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
			typeName: configType.Name,
			f:        nil,
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
				if field.HasSliceValue {
					return For(d.loopConditions().Block(
						d.deleteElementInLoop(),
					))
				}
				return d.deleteChild()
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
			valueTypeName: func() string { return ValueTypeName(&field) },
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
			valueTypeName: func() string { return AnyValueTypeName(&field) },
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
