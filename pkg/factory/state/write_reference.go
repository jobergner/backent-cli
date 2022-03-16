package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeReference() *Factory {
	s.config.RangeRefFields(func(field ast.Field) {

		r := referenceWriter{
			f: field,
		}

		if !field.HasSliceValue {

			s.file.Func().Params(r.receiverParams()).Id("IsSet").Params().Params(r.returns()).Block(
				r.reassignRef(),
				Return(Id("ref"), r.returnIsSet()),
			)

			s.file.Func().Params(r.receiverParams()).Id("Unset").Params().Block(
				r.reassignRef(),
				If(r.isOperationKindDelete()).Block(
					Return(),
				),
				r.deleteSelf(),
				r.declareParent(),
				If(r.parentIsDeleted()).Block(
					Return(),
				),
				r.setRefIDInParent(),
				r.setParentOperationKind(),
				r.signParent(),
				r.updateParentInPatch(),
			)
		}

		if field.HasAnyValue {
			s.file.Func().Params(r.receiverParams()).Id("Get").Params().Id(Lower(r.returnTypeOfGet())+"Ref").Block(
				r.reassignRef(),
				Id("anyContainer").Op(":=").Add(r.getReferencedElement()),
				Return(r.wrapIntoAnyRefWrapper()),
			)
		} else {
			s.file.Func().Params(r.receiverParams()).Id("Get").Params().Id(r.returnTypeOfGet()).Block(
				r.reassignRef(),
				Return(r.getReferencedElement()),
			)
		}

	})

	return s
}

func (s *Factory) writeDereference() *Factory {
	s.config.RangeRefFields(func(field ast.Field) {
		field.RangeValueTypes(func(valueType *ast.ConfigType) {

			d := dereferenceWriter{
				f: field,
				v: *valueType,
			}

			s.file.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params()).Block(
				d.declareAllIDs(),
				For(d.allIDsLoopConditions()).Block(
					d.declareRef(),
					OnlyIf(field.HasAnyValue, d.declareAnyContainer()),
					OnlyIf(field.HasAnyValue, If(d.anyContainerContainsElemenKind()).Block(
						Continue(),
					)),
					If(d.dereferenceCondition()).Block(
						OnlyIf(field.HasSliceValue, &Statement{
							d.declareParent().Line(),
							d.removeChildReferenceFromParent(),
						}),
						OnlyIf(!field.HasSliceValue, &Statement{
							d.unsetRef(),
						}),
					),
				),
				d.returnSliceToPool(),
			)
		})
	})

	return s
}
