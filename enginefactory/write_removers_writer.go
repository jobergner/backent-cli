package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type remover struct {
	t ast.ConfigType
	f ast.Field
	v *ast.ConfigType
}

func (r remover) receiverParams() *Statement {
	return Id(r.receiverName()).Id(r.t.Name)
}

func (r remover) name() string {
	var optionalSuffix string
	if r.f.HasAnyValue {
		optionalSuffix = Title(r.v.Name)
	}
	return "Remove" + Title(r.f.Name) + optionalSuffix
}

func (r remover) toRemoveParamName() string {
	if r.f.HasAnyValue {
		return r.v.Name + "sToRemove"
	}
	return r.f.Name + "ToRemove"
}

func (r remover) params() *Statement {
	toRemoveParam := Id(r.toRemoveParamName())
	if r.v.IsBasicType {
		return toRemoveParam.Id("..." + r.v.Name)
	}
	return toRemoveParam.Id("..." + Title(r.v.Name) + "ID")
}

func (r remover) returns() string {
	return r.t.Name
}

func (r remover) reassignElement() *Statement {
	return Id(r.t.Name).Op(":=").Id(r.receiverName()).Dot(r.t.Name).Dot("engine").Dot(Title(r.t.Name)).Call(Id(r.receiverName()).Dot(r.t.Name).Dot("ID"))
}

func (r remover) isOperationKindDelete() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (r remover) declareWereElementsAltered() *Statement {
	return Var().Id("wereElementsAltered").Bool()
}

func (r remover) declareNewElements() *Statement {
	newElementsType := "[]"
	if r.v.IsBasicType {
		newElementsType = newElementsType + r.f.ValueTypeName
	} else {
		newElementsType = newElementsType + Title(r.f.ValueTypeName) + "ID"
	}
	return Var().Id("newElements").Id(newElementsType)
}

func (r remover) existingElementsLoopConditions() *Statement {
	assignedVariableId := "element"
	if r.f.HasPointerValue {
		assignedVariableId = "refElement"
	} else if r.f.HasAnyValue {
		assignedVariableId = "anyContainerID"
	}
	return List(Id("_"), Id(assignedVariableId)).Op(":=").Range().Id(r.t.Name).Dot(r.t.Name).Dot(Title(r.f.Name))
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
	assignedVariableId := "element"
	if r.f.HasPointerValue {
		assignedVariableId = "refElement"
	}
	deleteCall := "delete" + Title(r.f.ValueTypeName)
	if r.f.HasAnyValue && !r.f.HasPointerValue {
		deleteCall = "delete" + Title(r.v.Name)
	}
	return Id(r.t.Name).Dot(r.t.Name).Dot("engine").Dot(deleteCall).Call(Id(assignedVariableId))
}

func (r remover) isElementRemoved() *Statement {
	return Add(Id("!")).Id("toBeRemoved")
}

func (r remover) appendRemainingElement() *Statement {
	assignedVariableId := Id("element")
	if r.f.HasPointerValue {
		assignedVariableId = Id("refElement")
	} else if r.f.HasAnyValue {
		assignedVariableId = Id("anyContainer").Dot(r.f.ValueTypeName).Dot("ID")
	}
	return Id("newElements").Op("=").Append(Id("newElements"), assignedVariableId)
}

func (r remover) isWereElementsAltered() *Statement {
	return Add(Id("!")).Id("wereElementsAltered")
}

func (r remover) setNewElements() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot(Title(r.f.Name)).Op("=").Id("newElements")
}

func (r remover) setOperationKind() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (r remover) updateElementInPatch() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("engine").Dot("Patch").Dot(Title(r.t.Name)).Index(Id(r.t.Name).Dot(r.t.Name).Dot("ID")).Op("=").Id(r.t.Name).Dot(r.t.Name)
}

func (r remover) receiverName() string {
	return "_" + r.t.Name
}

func (r remover) declareElementFromRef() *Statement {
	if r.f.HasAnyValue {
		return Id("element").Op(":=").Id("anyContainer").Dot(Title(r.v.Name)).Call().Dot("ID").Call()
	}
	return Id("element").Op(":=").Id(r.t.Name).Dot(r.t.Name).Dot("engine").Dot(r.f.ValueTypeName).Call(Id("refElement")).Dot(r.f.ValueTypeName).Dot("ReferencedElementID")
}

func (r remover) declareAnyContainer() *Statement {
	statement := Id("anyContainer").Op(":=").Id(r.t.Name).Dot(r.t.Name).Dot("engine").Dot(r.f.ValueTypeName)
	if r.f.HasAnyValue && !r.f.HasPointerValue {
		return statement.Call(Id("anyContainerID"))
	}
	return statement.Call(Id("refElement")).Dot("Get").Call()
}
