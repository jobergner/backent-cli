package enginefactory

import (
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"

	"bar-cli/ast"
)

func (s *EngineFactory) writeAssembleTree() *EngineFactory {
	decls := NewDeclSet()

	a := assembleTreeWriter{}

	decls.File.Func().Params(a.receiverParams()).Id("assembleTree").Params().Id("Tree").Block(
		a.createConfig(),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.patchLoopConditions()).Block(
						a.assembleElement(),
						If(Id("include")).Block(
							a.setElementInTree(),
						),
					),
				}
			}

			return &Statement{
				For(a.patchLoopConditions()).Block(
					If(a.elementHasNoParent()).Block(
						a.assembleElement(),
						If(Id("include")).Block(
							a.setElementInTree(),
						),
					),
				),
			}

		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.stateLoopConditions()).Block(
						If(a.elementNonExistentInTree()).Block(
							a.assembleElement(),
							If(Id("include")).Block(
								a.setElementInTree(),
							),
						),
					),
				}
			}

			return &Statement{
				For(a.stateLoopConditions()).Block(
					If(a.elementHasNoParent()).Block(
						If(a.elementNonExistentInTree()).Block(
							a.assembleElement(),
							If(Id("include")).Block(
								a.setElementInTree(),
							),
						),
					),
				),
			}

		}),
		Return(Id("engine").Dot("Tree")),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAssembleTreeElement() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		a := assembleElement{
			t: configType,
			f: nil,
		}

		decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Params(a.returns()).Block(
			If(a.checkIsDefined()).Block(
				If(a.elementExistsInCheck()).Block(
					Return(a.returnEmpty()),
				).Else().Block(
					a.checkElement(),
				),
			),
			a.getElementFromPatch(),
			If(Id("!hasUpdated")).Block(
				a.getElementFromState(),
			),
			a.declareTreeElement(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if a.f.ValueType().IsBasicType {
					return Empty()
				}

				if field.HasSliceValue {
					if field.HasAnyValue {
						return For(a.sliceFieldLoopConditions()).Block(
							If(a.elementHasUpdated()).Block(
								If(Id("childHasUpdated")).Block(
									a.setHasUpdatedTrue(),
								),
								a.appendToElementsInField(),
							),
						)
					} else {
						return For(a.sliceFieldLoopConditions()).Block(
							If(a.elementHasUpdated()).Block(
								If(Id("childHasUpdated")).Block(
									a.setHasUpdatedTrue(),
								),
								a.appendToElementsInField(),
							),
						)
					}
				}
				return If(a.elementHasUpdated()).Block(
					If(Id("childHasUpdated")).Block(
						a.setHasUpdatedTrue(),
					),
					a.setFieldElement(),
				)
			}),
			a.setID(),
			a.setOperationKind(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if !a.f.ValueType().IsBasicType {
					return Empty()
				}

				return a.setField()
			}),
			Return(a.finalReturn()),
		)
	})

	decls.Render(s.buf)
	return s

}
