package enginefactory

import (
	. "github.com/dave/jennifer/jen"

	"bar-cli/ast"
)

func (s *EngineFactory) writeAssembleTree() *EngineFactory {
	decls := newDeclSet()

	a := assembleTreeWriter{}

	decls.file.Func().Params(a.receiverParams()).Id("assembleTree").Params().Id("Tree").Block(
		Id("tree").Op(":=").Id("newTree").Call(),
		forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.patchLoopConditions()).Block(
						a.assembleItem(),
						If(Id("hasUpdated")).Block(
							a.setElementInTree(),
						),
					),
				}
			}

			return &Statement{
				For(a.patchLoopConditions()).Block(
					If(a.elementHasNoParent()).Block(
						a.assembleItem(),
						If(Id("hasUpdated")).Block(
							a.setElementInTree(),
						),
					),
				),
			}

		}),
		forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			a.t = &configType

			if a.t.IsRootType {
				return &Statement{
					For(a.stateLoopConditions()).Block(
						If(a.elementNonExistentInTree()).Block(
							a.assembleItem(),
							If(Id("hasUpdated")).Block(
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
							a.assembleItem(),
							If(Id("hasUpdated")).Block(
								a.setElementInTree(),
							),
						),
					),
				),
			}

		}),
		Return(Id("tree")),
	)

	decls.render(s.buf)
	return s
}

type assembleTreeWriter struct {
	t *ast.ConfigType
}

func (a assembleTreeWriter) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (a assembleTreeWriter) patchLoopConditions() *Statement {
	return List(Id("_"), Id(a.t.Name)).Op(":=").Range().Id("se").Dot("Patch").Dot(title(a.t.Name))
}

func (a assembleTreeWriter) elementHasNoParent() *Statement {
	return Id("!" + a.t.Name).Dot("HasParent_")
}

func (a assembleTreeWriter) elementNonExistentInTree() (*Statement, *Statement) {
	condition := List(Id("_"), Id("ok")).Op(":=").Id("tree").Dot(title(a.t.Name)).Index(Id(a.t.Name).Dot("ID"))
	return condition, Id("!ok")
}

func (a assembleTreeWriter) assembleItem() *Statement {
	variableNames := List(Id("tree"+title(a.t.Name)), Id("hasUpdated"))
	return variableNames.Op(":=").Id("se").Dot("assemble" + title(a.t.Name)).Call(Id(a.t.Name).Dot("ID"))
}

func (a assembleTreeWriter) setElementInTree() *Statement {
	return Id("tree").Dot(title(a.t.Name)).Index(Id(a.t.Name).Dot("ID")).Op("=").Id("tree" + title(a.t.Name))
}

func (a assembleTreeWriter) stateLoopConditions() *Statement {
	return List(Id("_"), Id(a.t.Name)).Op(":=").Range().Id("se").Dot("State").Dot(title(a.t.Name))
}

func (s *EngineFactory) writeAssembleTreeElement() *EngineFactory {
	decls := newDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		a := assembleElement{
			t: configType,
			f: nil,
		}

		decls.file.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Params(a.returns()).Block(
			a.getElementFromPatch(),
			If(Id("!hasUpdated")).Block(
				a.earlyReturn(),
			),
			a.declareTreeElement(),
			forEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if a.f.ValueType.IsBasicType {
					return Empty()
				}

				if field.HasSliceValue {
					return For(a.sliceFieldLoopConditions()).Block(
						If(a.elementHasUpdated(Id(field.ValueType.Name+"ID"))).Block(
							a.setHasUpdatedTrue(),
							a.appendToElementsInField(),
						),
					)
				}
				return If(a.elementHasUpdated(Id(configType.Name).Dot(title(field.Name)))).Block(
					a.setHasUpdatedTrue(),
					a.setFieldElement(),
				)
			}),
			a.setID(),
			a.setOperationKind(),
			forEachFieldInType(configType, func(field ast.Field) *Statement {
				a.f = &field

				if !a.f.ValueType.IsBasicType {
					return Empty()
				}

				return a.setField()
			}),
			Return(a.finalReturn()),
		)
	})

	decls.render(s.buf)
	return s

}

type assembleElement struct {
	t ast.ConfigType
	f *ast.Field
}

func (a assembleElement) treeElementName() string {
	return "tree" + title(a.t.Name)
}

func (a assembleElement) treeTypeName() string {
	return "t" + title(a.t.Name)
}

func (a assembleElement) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (a assembleElement) name() string {
	return "assemble" + title(a.t.Name)
}

func (a assembleElement) idParam() string {
	return a.t.Name + "ID"
}

func (a assembleElement) params() *Statement {
	return Id(a.idParam()).Id(title(a.t.Name) + "ID")
}

func (a assembleElement) returns() (*Statement, *Statement) {
	return Id(a.treeTypeName()), Bool()
}

func (a assembleElement) getElementFromPatch() *Statement {
	return List(Id(a.t.Name), Id("hasUpdated")).Op(":=").Id("se").Dot("Patch").Dot(title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElement) earlyReturn() *Statement {
	if a.t.IsLeafType {
		return Return(List(Id(a.treeTypeName()).Values(), Lit(false)))
	}
	return Id(a.t.Name).Op("=").Id("se").Dot("State").Dot(title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElement) declareTreeElement() *Statement {
	return Var().Id(a.treeElementName()).Id(a.treeTypeName())
}

func (a assembleElement) typeFieldOn(from string) *Statement {
	return Id("se").Dot(from).Dot(title(a.t.Name)).Index(Id(a.t.Name).Dot("ID")).Dot(title(a.f.Name))
}

func (a assembleElement) sliceFieldLoopConditions() *Statement {
	loopVars := List(Id("_"), Id(a.f.ValueType.Name+"ID"))
	deduplicateFuncName := "deduplicate" + title(a.f.ValueType.Name) + "IDs"
	return loopVars.Op(":=").Range().Id(deduplicateFuncName).Call(a.typeFieldOn("State"), a.typeFieldOn("Patch"))
}

func (a assembleElement) elementHasUpdated(fieldIdentifier *Statement) (*Statement, *Statement) {
	elementHasUpdatedId := Id(a.f.ValueType.Name + "HasUpdated")
	conditionVars := List(Id("tree"+title(a.f.ValueType.Name)), elementHasUpdatedId)
	condition := conditionVars.Op(":=").Id("se").Dot("assemble" + title(a.f.ValueType.Name)).Call(fieldIdentifier)
	return condition, elementHasUpdatedId
}

func (a assembleElement) setHasUpdatedTrue() *Statement {
	return Id("hasUpdated").Op("=").True()
}

func (a assembleElement) appendToElementsInField() *Statement {
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Append(Id(a.treeElementName()).Dot(title(a.f.Name)), Id("tree"+title(a.f.ValueType.Name)))
}

func (a assembleElement) setFieldElement() *Statement {
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Id("&" + "tree" + title(a.f.ValueType.Name))
}

func (a assembleElement) setField() *Statement {
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Id(a.t.Name).Dot(title(a.f.Name))
}

func (a assembleElement) setID() *Statement {
	return Id(a.treeElementName()).Dot("ID").Op("=").Id(a.t.Name).Dot("ID")
}

func (a assembleElement) setOperationKind() *Statement {
	return Id(a.treeElementName()).Dot("OperationKind_").Op("=").Id(a.t.Name).Dot("OperationKind_")
}

func (a assembleElement) finalReturn() (*Statement, *Statement) {
	elementName := Id(a.treeElementName())
	if a.t.IsLeafType {
		return elementName, True()
	}
	return elementName, Id("hasUpdated")
}
