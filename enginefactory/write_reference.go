package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeReference() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {

		r := referenceWriter{
			f: field,
		}

		if !field.HasSliceValue {
			decls.File.Func().Params(r.receiverParams()).Id("IsSet").Params().Bool().Block(
				r.reassignRef(),
				r.returnIsSet(),
			)

			decls.File.Func().Params(r.receiverParams()).Id("Unset").Params().Block(
				r.reassignRef(),
				r.deleteSelf(),
				r.declareParent(),
				If(r.parentIsDeleted()).Block(
					Return(),
				),
				r.setRefIDInParent(),
				r.setParentOperationKind(),
				r.updateParentInPatch(),
			)
		}

		decls.File.Func().Params(r.receiverParams()).Id("Get").Params().Id(lower(r.returnTypeOfGet())).Block(
			r.reassignRef(),
			r.returnReferencedElement(),
		)

	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeDereference() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {
		field.RangeValueTypes(func(valueType *ast.ConfigType) {

			d := dereferenceWriter{
				f: field,
				v: *valueType,
			}

			decls.File.Func().Params(d.receiverParams()).Id(d.name()).Params(d.params()).Block(
				For(d.allIDsLoopConditions()).Block(
					d.declareRef(),
					onlyIf(field.HasAnyValue, d.declareAnyContainer()),
					onlyIf(field.HasAnyValue, If(d.anyContainerContainsElemenKind()).Block(
						Continue(),
					)),
					If(d.dereferenceCondition()).Block(
						onlyIf(field.HasSliceValue, &Statement{
							d.declareParent().Line(),
							d.removeChildReferenceFromParent(),
						}),
						onlyIf(!field.HasSliceValue, &Statement{
							d.unsetRef(),
						}),
					),
				),
			)
		})

	})

	decls.Render(s.buf)
	return s
}
