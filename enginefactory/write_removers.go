package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeRemovers() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {

			if !field.HasSliceValue {
				return
			}

			r := remover{
				t: configType,
				f: field,
			}

			field.RangeValueTypes(func(valueType *ast.ConfigType) {

				r.v = valueType

				decls.File.Func().Params(r.receiverParams()).Id(r.name()).Params(r.params()).Id(r.returns()).Block(
					r.reassignElement(),
					If(r.isOperationKindDelete()).Block(
						Return(Id(configType.Name)),
					),
					r.declareWereElementsAltered(),
					r.declareNewElements(),
					For(r.existingElementsLoopConditions()).Block(
						OnlyIf(field.HasAnyValue, r.declareAnyContainer()),
						OnlyIf(field.HasPointerValue || field.HasAnyValue, r.declareElementFromRef()),
						OnlyIf(field.HasAnyValue, If(r.elementIsNil()).Block(
							Continue(),
						)),
						r.declareToBeRemovedBool(),
						For(r.elementsToDeteleLoopConditions()).Block(
							If(r.isElementMatching()).Block(
								r.setToBeRemovedTrue(),
								r.setWereElementAlteredTrue(),
								OnlyIf(!field.ValueType().IsBasicType, r.deleteElement()),
								Break(),
							),
						),
						If(r.isElementRemoved()).Block(
							r.appendRemainingElement(),
						),
					),
					If(r.isWereElementsAltered()).Block(
						Return(Id(configType.Name)),
					),
					r.setNewElements(),
					r.setOperationKind(),
					r.updateElementInPatch(),
					Return(Id(configType.Name)),
				)
			})

		})
	})

	decls.Render(s.buf)
	return s
}
