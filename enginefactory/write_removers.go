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

			decls.File.Func().Params(r.receiverParams()).Id(r.name()).Params(r.params()).Id(r.returns()).Block(
				r.reassignElement(),
				If(r.isOperationKindDelete()).Block(
					Return(Id(configType.Name)),
				),
				r.declareWereElementsAltered(),
				r.declareNewElements(),
				For(r.existingElementsLoopConditions()).Block(
					r.declareToBeRemovedBool(),
					For(r.elementsToDeteleLoopConditions()).Block(
						If(r.isElementMatching()).Block(
							r.setToBeRemovedTrue(),
							r.setWereElementAlteredTrue(),
							onlyIf(!field.ValueType.IsBasicType, r.deleteElement()),
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

	decls.Render(s.buf)
	return s
}

type remover struct {
	t ast.ConfigType
	f ast.Field
}

func (r remover) receiverParams() *Statement {
	return Id(r.receiverName()).Id(r.t.Name)
}

func (r remover) name() string {
	return "Remove" + title(r.f.Name)
}

func (r remover) toRemoveParamName() string {
	return r.f.Name + "ToRemove"
}

func (r remover) params() *Statement {
	toRemoveParam := Id(r.toRemoveParamName())
	if r.f.ValueType.IsBasicType {
		toRemoveParam = toRemoveParam.Id("..." + r.f.ValueType.Name)
	} else {
		toRemoveParam = toRemoveParam.Id("..." + title(r.f.ValueType.Name) + "ID")
	}
	return List(Id("se").Id("*Engine"), toRemoveParam)
}

func (r remover) returns() string {
	return r.t.Name
}

func (r remover) reassignElement() *Statement {
	return Id(r.t.Name).Op(":=").Id("se").Dot(title(r.t.Name)).Call(Id(r.receiverName()).Dot(r.t.Name).Dot("ID"))
}

func (r remover) isOperationKindDelete() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
}

func (r remover) declareWereElementsAltered() *Statement {
	return Var().Id("wereElementsAltered").Bool()
}

func (r remover) declareNewElements() *Statement {
	newElementsType := "[]"
	if r.f.ValueType.IsBasicType {
		newElementsType = newElementsType + r.f.ValueType.Name
	} else {
		newElementsType = newElementsType + title(r.f.ValueType.Name) + "ID"
	}
	return Var().Id("newElements").Id(newElementsType)
}

func (r remover) existingElementsLoopConditions() *Statement {
	return List(Id("_"), Id("element")).Op(":=").Range().Id(r.t.Name).Dot(r.t.Name).Dot(title(r.f.Name))
}

func (r remover) declareToBeRemovedBool() *Statement {
	return Var().Id("toBeRemoved").Bool()
}

func (r remover) elementsToDeteleLoopConditions() *Statement {
	return List(Id("_"), Id("elementToRemove")).Op(":=").Range().Id(r.toRemoveParamName())
}

func (r remover) isElementMatching() *Statement {
	return Id("element").Op("==").Id("elementToRemove")
}

func (r remover) setToBeRemovedTrue() *Statement {
	return Id("toBeRemoved").Op("=").True()
}

func (r remover) setWereElementAlteredTrue() *Statement {
	return Id("wereElementsAltered").Op("=").True()
}

func (r remover) deleteElement() *Statement {
	return Id("se").Dot("delete" + title(r.f.ValueType.Name)).Call(Id("element"))
}

func (r remover) isElementRemoved() *Statement {
	return Add(Id("!")).Id("toBeRemoved")
}

func (r remover) appendRemainingElement() *Statement {
	return Id("newElements").Op("=").Append(Id("newElements"), Id("element"))
}

func (r remover) isWereElementsAltered() *Statement {
	return Add(Id("!")).Id("wereElementsAltered")
}

func (r remover) setNewElements() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot(title(r.f.Name)).Op("=").Id("newElements")
}

func (r remover) setOperationKind() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
}

func (r remover) updateElementInPatch() *Statement {
	return Id("se").Dot("Patch").Dot(title(r.t.Name)).Index(Id(r.t.Name).Dot(r.t.Name).Dot("ID")).Op("=").Id(r.t.Name).Dot(r.t.Name)
}

func (r remover) receiverName() string {
	return "_" + r.t.Name
}
