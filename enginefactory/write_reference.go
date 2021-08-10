package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeReference() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {

		r := referenceWriter{
			f: field,
		}

		if !field.HasSliceValue {
			decls.File.Func().Params(r.receiverParams()).Id("IsSet").Params().Params(r.returns()).Block(
				r.reassignRef(),
				Return(Id("ref"), r.returnIsSet()),
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

		if field.HasAnyValue {
			decls.File.Func().Params(r.receiverParams()).Id("Get").Params().Id(Lower(r.returnTypeOfGet())+"Ref").Block(
				r.reassignRef(),
				Id("anyContainer").Op(":=").Add(r.getReferencedElement()),
				Return(r.wrapIntoAnyRefWrapper()),
			)
		} else {
			decls.File.Func().Params(r.receiverParams()).Id("Get").Params().Id(Lower(r.returnTypeOfGet())).Block(
				r.reassignRef(),
				Return(r.getReferencedElement()),
			)
		}

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

	decls.Render(s.buf)
	return s
}
