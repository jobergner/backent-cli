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
					if field.HasAnyValue && !field.HasPointerValue {
						return For(a.sliceFieldLoopConditions()).Block(
							onlyIf(!field.HasPointerValue, a.createAnyContainer()),
							forEachFieldValueComparison(field, *Id(a.anyContainerName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
								return &Statement{
									a.assignIDFromAnyContainer(valueType).Line(),
									If(a.elementHasUpdated(valueType, a.usedAssembleID(configType, field, valueType))).Block(
										If(Id("childHasUpdated")).Block(
											a.setHasUpdatedTrue(),
										),
										a.appendToElementsInField(valueType),
									),
								}
							}),
						)
					} else {
						return For(a.sliceFieldLoopConditions()).Block(
							If(a.elementHasUpdated(field.ValueType(), a.usedAssembleID(configType, field, field.ValueType()))).Block(
								If(Id("childHasUpdated")).Block(
									a.setHasUpdatedTrue(),
								),
								a.appendToElementsInField(field.ValueType()),
							),
						)
					}
				}

				if field.HasAnyValue && !field.HasPointerValue {
					return &Statement{
						onlyIf(!field.HasPointerValue, a.createAnyContainer().Line()),
						forEachFieldValueComparison(field, *Id(a.anyContainerName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
							return &Statement{
								a.assignIDFromAnyContainer(valueType).Line(),
								If(a.elementHasUpdated(valueType, a.usedAssembleID(configType, field, valueType))).Block(
									If(Id("childHasUpdated")).Block(
										a.setHasUpdatedTrue(),
									),
									a.setFieldElement(valueType),
								),
							}

						}),
					}
				}
				return If(a.elementHasUpdated(field.ValueType(), a.usedAssembleID(configType, field, field.ValueType()))).Block(
					If(Id("childHasUpdated")).Block(
						a.setHasUpdatedTrue(),
					),
					a.setFieldElement(field.ValueType()),
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

func (s *EngineFactory) writeAssembleTreeReference() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeRefFields(func(field ast.Field) {
		if field.HasAnyValue && !field.HasSliceValue {
			a := assembleReferenceWriter{
				f: field,
			}
			decls.File.Func().Params(a.receiverParams()).Id("assemble"+title(field.ValueTypeName)).Params(a.params()).Params(a.returns()).Block(
				a.declareStateElement(),
				a.declarepatchElement(),
				If(a.refIsNotSet()).Block(
					Return(Nil(), False(), False()),
				),
				If(Id("config").Dot("forceInclude")).Block(
					a.declareRef(false),
					a.declareAnyContainer(false),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeTreeReferenceForceInclude()
					}),
				),
				If(a.refWasCreated()).Block(
					Id("config").Dot("forceInclude").Op("=").True(),
					a.declareRef(false),
					a.declareAnyContainer(false),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeNonSliceTreeReferenceRefCreated()
					}),
				),
				If(a.refWasRemoved()).Block(
					a.declareRef(true),
					a.declareAnyContainer(false),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeNonSliceTreeReferenceRefRemoved()
					}),
				),
				If(a.refHasBeenReplaced()).Block(
					If(a.refWasReplaced()).Block(
						a.declareRef(false),
						a.declareAnyContainer(false),
						forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
							a.v = valueType
							return a.writeNonSliceTreeReferenceRefReplaced()
						}),
					),
				),
				If(a.referencedElementGotUpdated()).Block(
					a.declareRef(true),
					a.declareAnyContainer(false),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeNonSliceTreeReferenceRefElementUpdated()
					}),
				),
				Return(Nil(), False(), False()),
			)
		}
		if !field.HasAnyValue && !field.HasSliceValue {
			a := assembleReferenceWriter{
				f: field,
				v: field.ValueType(),
			}
			decls.File.Func().Params(a.receiverParams()).Id("assemble"+title(field.ValueTypeName)).Params(a.params()).Params(a.returns()).Block(
				a.declareStateElement(),
				a.declarepatchElement(),
				If(a.refIsNotSet()).Block(
					Return(Nil(), False(), False()),
				),
				If(Id("config").Dot("forceInclude")).Block(
					a.declareRef(false),
					a.writeTreeReferenceForceInclude(),
				),
				If(a.refWasCreated()).Block(
					Id("config").Dot("forceInclude").Op("=").True(),
					a.declareRef(false),
					a.writeNonSliceTreeReferenceRefCreated(),
				),
				If(a.refWasRemoved()).Block(
					a.declareRef(true),
					a.writeNonSliceTreeReferenceRefRemoved(),
				),
				If(a.refHasBeenReplaced()).Block(
					If(a.refWasReplaced()).Block(
						a.declareRef(false),
						a.writeNonSliceTreeReferenceRefReplaced(),
					),
				),
				If(a.referencedElementGotUpdated()).Block(
					a.declareRef(true),
					a.writeNonSliceTreeReferenceRefElementUpdated(),
				),
				Return(Nil(), False(), False()),
			)
		}
		if field.HasAnyValue && field.HasSliceValue {
			a := assembleReferenceWriter{
				f: field,
			}
			decls.File.Func().Params(a.receiverParams()).Id("assemble"+title(field.ValueTypeName)).Params(a.params()).Params(a.returns()).Block(
				If(Id("config").Dot("forceInclude")).Block(
					a.declareRef(false),
					a.declareAnyContainer(false),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeTreeReferenceForceInclude()
					}),
				),
				If(a.sliceRefHasUpdated()).Block(
					If(Id("patchRef").Dot("OperationKind").Op("==").Id("OperationKindUpdate")).Block(
						Id("config").Dot("forceInclude").Op("=").True(),
					),
					a.declareAnyContainer(true),
					forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
						a.v = valueType
						return a.writeSliceTreeReferenceRefUpdated()
					}),
				),
				a.declareRef(false),
				a.declareAnyContainer(false),
				If(Id("check").Op("==").Nil()).Block(
					Id("check").Op("=").Id("newRecursionCheck").Call(),
				),
				a.declareAnyContainer(false),
				forEachFieldValueComparison(field, *Id("anyContainer").Dot(a.nextValueName()).Dot("ElementKind"), func(valueType *ast.ConfigType) *Statement {
					a.v = valueType
					return a.writeSliceTreeReferenceRefElementUpdated()
				}),
				Return(Nil(), False(), False()),
			)
		}
		if !field.HasAnyValue && field.HasSliceValue {
			a := assembleReferenceWriter{
				f: field,
				v: field.ValueType(),
			}
			decls.File.Func().Params(a.receiverParams()).Id("assemble"+title(field.ValueTypeName)).Params(a.params()).Params(a.returns()).Block(
				If(Id("config").Dot("forceInclude")).Block(
					a.declareRef(false),
					a.writeTreeReferenceForceInclude(),
				),
				If(a.sliceRefHasUpdated()).Block(
					If(Id("patchRef").Dot("OperationKind").Op("==").Id("OperationKindUpdate")).Block(
						Id("config").Dot("forceInclude").Op("=").True(),
					),
					a.writeSliceTreeReferenceRefUpdated(),
				),
				a.declareRef(false),
				If(Id("check").Op("==").Nil()).Block(
					Id("check").Op("=").Id("newRecursionCheck").Call(),
				),
				a.writeSliceTreeReferenceRefElementUpdated(),
				Return(Nil(), False(), False()),
			)
		}
	})

	decls.Render(s.buf)
	return s
}
