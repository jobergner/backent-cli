package enginefactory

import (
	. "github.com/dave/jennifer/jen"

	"bar-cli/ast"
)

type assembleTreeWriter struct {
	t *ast.ConfigType
}

func (a assembleTreeWriter) dataElementName() string {
	return a.t.Name + "Data"
}

func (a assembleTreeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (a assembleTreeWriter) patchLoopConditions() *Statement {
	return List(Id("_"), Id(a.dataElementName())).Op(":=").Range().Id("engine").Dot("Patch").Dot(title(a.t.Name))
}

func (a assembleTreeWriter) elementHasNoParent() *Statement {
	return Id("!" + a.dataElementName()).Dot("HasParent")
}

func (a assembleTreeWriter) elementNonExistentInTree() (*Statement, *Statement) {
	condition := List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("Tree").Dot(title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID"))
	return condition, Id("!ok")
}

func (a assembleTreeWriter) assembleElement() *Statement {
	variableNames := List(Id(a.t.Name), Id("include"), Id("_"))
	return variableNames.Op(":=").Id("engine").Dot("assemble"+title(a.t.Name)).Call(Id(a.dataElementName()).Dot("ID"), Nil(), Id("config"))
}

func (a assembleTreeWriter) setElementInTree() *Statement {
	return Id("engine").Dot("Tree").Dot(title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID")).Op("=").Id(a.t.Name)
}

func (a assembleTreeWriter) stateLoopConditions() *Statement {
	return List(Id("_"), Id(a.dataElementName())).Op(":=").Range().Id("engine").Dot("State").Dot(title(a.t.Name))
}

func (a assembleTreeWriter) createConfig() *Statement {
	return Id("config").Op(":=").Id("assembleConfig").Values(Dict{Id("forceInclude"): False()})
}

type assembleElement struct {
	t ast.ConfigType
	f *ast.Field
}

func (a assembleElement) treeElementName() string {
	return a.t.Name
}

func (a assembleElement) dataElementName() string {
	return a.t.Name + "Data"
}

func (a assembleElement) treeTypeName() string {
	return title(a.t.Name)
}

func (a assembleElement) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (a assembleElement) name() string {
	return "assemble" + title(a.t.Name)
}

func (a assembleElement) idParam() string {
	return a.t.Name + "ID"
}

func (a assembleElement) params() (*Statement, *Statement, *Statement) {
	return Id(a.idParam()).Id(title(a.t.Name) + "ID"), Id("check").Id("*recursionCheck"), Id("config").Id("assembleConfig")
}

func (a assembleElement) returns() (*Statement, *Statement, *Statement) {
	return Id(a.treeTypeName()), Bool(), Bool()
}

func (a assembleElement) checkIsDefined() *Statement {
	return Id("check").Op("!=").Nil()
}

func (a assembleElement) elementExistsInCheck() (*Statement, *Statement) {
	return Id("alreadyExists").Op(":=").Id("check").Dot(a.t.Name).Index(Id(a.idParam())), Id("alreadyExists")
}

func (a assembleElement) returnEmpty() (*Statement, *Statement, *Statement) {
	return Id(title(a.t.Name)).Values(), False(), False()
}

func (a assembleElement) checkElement() *Statement {
	return Id("check").Dot(a.t.Name).Index(Id(a.idParam())).Op("=").True()
}

func (a assembleElement) getElementFromPatch() *Statement {
	return List(Id(a.dataElementName()), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElement) getElementFromState() *Statement {
	return Id(a.dataElementName()).Op("=").Id("engine").Dot("State").Dot(title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElement) declareTreeElement() *Statement {
	return Var().Id(a.treeElementName()).Id(a.treeTypeName())
}

func (a assembleElement) typeFieldOn(from string) *Statement {
	return Id("engine").Dot(from).Dot(title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID")).Dot(title(a.f.Name))
}

func (a assembleElement) sliceFieldLoopConditions() *Statement {
	loopVars := List(Id("_"), Id(a.f.ValueTypeName+"ID"))
	mergeFuncName := "merge" + title(a.f.ValueTypeName) + "IDs"
	return loopVars.Op(":=").Range().Id(mergeFuncName).Call(a.typeFieldOn("State"), a.typeFieldOn("Patch"))
}

func (a assembleElement) usedAssembleID(configType ast.ConfigType, field ast.Field, valueType *ast.ConfigType) *Statement {
	if !field.HasPointerValue && !field.HasAnyValue && !field.HasSliceValue {
		return Id(a.dataElementName()).Dot(title(field.Name))
	} else if field.HasPointerValue && !field.HasAnyValue && !field.HasSliceValue {
		return Id(configType.Name + "ID")
	} else if field.HasPointerValue && field.HasSliceValue {
		return Id(field.ValueTypeName + "ID")
	}

	return Id(valueType.Name + "ID")
}

func (a assembleElement) elementHasUpdated(valueType *ast.ConfigType, assembleID *Statement) (*Statement, *Statement) {
	handledElementTypeName := a.f.ValueTypeName
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		handledElementTypeName = valueType.Name
	}
	conditionVars := List(Id("tree"+title(handledElementTypeName)), Id("include"), Id("childHasUpdated"))
	condition := conditionVars.Op(":=").Id("engine").Dot("assemble"+title(handledElementTypeName)).Call(assembleID, Id("check"), Id("config"))
	return condition, Id("include")
}

func (a assembleElement) setHasUpdatedTrue() *Statement {
	return Id("hasUpdated").Op("=").True()
}

func (a assembleElement) appendToElementsInField(valueType *ast.ConfigType) *Statement {
	appendedType := valueType.Name
	if !a.f.HasAnyValue && a.f.HasPointerValue || (a.f.HasPointerValue && a.f.HasSliceValue && a.f.HasAnyValue) {
		appendedType = a.f.ValueTypeName
	}
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Append(Id(a.treeElementName()).Dot(title(a.f.Name)), Id("tree"+title(appendedType)))
}

func (a assembleElement) setFieldElement(valueType *ast.ConfigType) *Statement {
	var optionalShare string
	if !a.f.HasPointerValue {
		optionalShare = "&"
	}
	usedType := a.f.ValueTypeName
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		usedType = valueType.Name
	}
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Id(optionalShare + "tree" + title(usedType))
}

func (a assembleElement) setField() *Statement {
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Id(a.dataElementName()).Dot(title(a.f.Name))
}

func (a assembleElement) setID() *Statement {
	return Id(a.treeElementName()).Dot("ID").Op("=").Id(a.dataElementName()).Dot("ID")
}

func (a assembleElement) setOperationKind() *Statement {
	return Id(a.treeElementName()).Dot("OperationKind").Op("=").Id(a.dataElementName()).Dot("OperationKind")
}

func (a assembleElement) finalReturn() (*Statement, *Statement, *Statement) {
	return Id(a.t.Name), Id("hasUpdated").Op("||").Id("config").Dot("forceInclude"), Id("hasUpdated")
}

func (a assembleElement) anyContainerName() string {
	return anyNameByField(*a.f) + "Container"
}

func (a assembleElement) createAnyContainer() *Statement {
	usedID := Id(anyNameByField(*a.f) + "ID")
	if !a.f.HasSliceValue {
		usedID = Id(a.dataElementName()).Dot(title(a.f.Name))
	}
	return Id(a.anyContainerName()).Op(":=").Id("engine").Dot(anyNameByField(*a.f)).Call(usedID).Dot(anyNameByField(*a.f))
}

func (a assembleElement) assignIDFromAnyContainer(valueType *ast.ConfigType) *Statement {
	return Id(valueType.Name + "ID").Op(":=").Id(a.anyContainerName()).Dot(title(valueType.Name))
}

type assembleReferenceWriter struct {
	f ast.Field
	v *ast.ConfigType
}

func (a assembleReferenceWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (a assembleReferenceWriter) idParam() string {
	if a.f.HasSliceValue {
		return "refID"
	}
	return a.f.Parent.Name + "ID"
}

func (a assembleReferenceWriter) nextValueName() string {
	if a.f.HasAnyValue {
		return anyNameByField(a.f)
	}
	return a.f.ValueType().Name
}

func (a assembleReferenceWriter) params() (*Statement, *Statement, *Statement) {
	idType := title(a.f.Parent.Name) + "ID"
	if a.f.HasSliceValue {
		idType = title(a.f.ValueTypeName) + "ID"
	}
	return Id(a.idParam()).Id(idType), Id("check").Id("*recursionCheck"), Id("config").Id("assembleConfig")
}

func (a assembleReferenceWriter) returns() (*Statement, *Statement, *Statement) {
	return Id(title(a.nextValueName()) + "Reference"), Bool(), Bool()
}

func (a assembleReferenceWriter) declareStateElement() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Op(":=").Id("engine").Dot("State").Dot(title(a.f.Parent.Name)).Index(Id(a.idParam()))
}

func (a assembleReferenceWriter) declarepatchElement() *Statement {
	return List(Id("patch"+title(a.f.Parent.Name)), Id(a.f.Parent.Name+"IsInPatch")).Op(":=").Id("engine").Dot("Patch").Dot(title(a.f.Parent.Name)).Index(Id(a.idParam()))
}

func (a assembleReferenceWriter) refIsNotSet() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("==").Lit(0).Op("&&").Params(Id("!" + a.f.Parent.Name + "IsInPatch").Op("||").Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("==").Lit(0))
}

func (a assembleReferenceWriter) declareRef() *Statement {
	if a.f.HasAnyValue {
		return Id("ref").Op(":=").Id("engine").Dot(a.nextValueName()).Call(Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name)))
	}
	return Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id("refID")).Dot(a.f.ValueTypeName)
}

func (a assembleReferenceWriter) declareAnyContainer() *Statement {
	return Id("anyContainer").Dot(a.nextValueName()).Call(Id("ref").Dot(a.f.Name).Dot("ReferencedElementID"))
}

func (a assembleReferenceWriter) declareReferencedElement() *Statement {
	if a.f.HasAnyValue {
		return Id("referencedElement").Op(":=").Id("engine").Dot(title(a.v.Name)).Call(Id("anyContainer").Dot(a.nextValueName()).Dot(title(a.v.Name))).Dot(a.v.Name)
	}
	return Id("referencedElement").Op(":=").Id("engine").Dot(title(a.v.Name)).Call(Id("ref").Dot("ReferencedElementID")).Dot(a.v.Name)
}

func (a assembleReferenceWriter) assembleReferencedElement(declareElement bool, useContainer bool) *Statement {
	firstReturn := "_"
	if declareElement {
		firstReturn = "element"
	}
	usedID := Id("referencedElement").Dot("ID")
	if a.f.HasSliceValue {
		usedID = Id("ref").Dot("ReferencedElementID")
	}
	if useContainer {
		usedID = Id("anyContainer").Dot(a.nextValueName()).Dot(title(a.v.Name))
	}
	return List(Id(firstReturn), Id("_"), Id("hasUpdatedDownstream")).Op(":=").Id("engine").Dot("assemble"+title(a.v.Name)).Call(usedID, Id("check"), Id("config"))
}

func (a assembleReferenceWriter) retrievePath(useContainer bool) *Statement {
	usedID := Id("referencedElement").Dot("ID")
	if useContainer {
		usedID = Id("anyContainer").Dot(a.nextValueName()).Dot(title(a.v.Name))
	}
	return List(Id("path"), Id("_")).Op(":=").Id("engine").Dot("PathTrack").Dot(a.v.Name).Index(usedID)
}

func (a assembleReferenceWriter) defineReference(useElement bool, operationKind string, useContainer bool) *Statement {
	element := Nil()
	if useElement {
		element = Id("&Element")
	}
	operationKindStatement := Id("ref").Dot(a.f.Name).Dot("OperationKind")
	if operationKind == "delete" {
		operationKindStatement = Id("OperationKindDelete")
	} else if operationKind == "update" {
		operationKindStatement = Id("OperationKindUpdate")
	}
	usedID := Id("referencedElement").Dot("ID")
	if useContainer {
		usedID = Id("anyContainer").Dot(a.nextValueName()).Dot(title(a.v.Name))
	}
	optionalShare := "&"
	if a.f.HasSliceValue {
		optionalShare = ""
	}
	return Id(optionalShare+title(a.nextValueName())+"Reference").Values(
		operationKindStatement,
		Int().Call(usedID),
		Id("ElementKind"+title(a.v.Name)),
		Id("referencedDataStatus"),
		Id("path").Dot("toJSONPath").Call(),
		element,
	)
}

func (a assembleReferenceWriter) dataHasUpdated() *Statement {
	return Id("ref").Dot(a.f.ValueTypeName).Dot("OperationKind").Op("==").Id("OperationKindUpdate").Op("||").Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")
}

func (a assembleReferenceWriter) refWasCreated() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("==").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Lit(0))
}

func (a assembleReferenceWriter) refWasRemoved() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("==").Lit(0))
}

func (a assembleReferenceWriter) refHasBeenReplaced() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Lit(0))
}

func (a assembleReferenceWriter) refWasReplaced() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Id("patch" + title(a.f.Parent.Name)).Dot(title(a.f.Name))
}

func (a assembleReferenceWriter) referencedElementGotUpdated() *Statement {
	return Id("state" + title(a.f.Parent.Name)).Dot(title(a.f.Name)).Op("!=").Lit(0)
}

func (a assembleReferenceWriter) writeTreeReferenceForceInclude() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(a.assembleReferencedElement(false, false), Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		a.retrievePath(false).Line(),
		Return(a.defineReference(false, "", false), True(), a.dataHasUpdated()).Line(),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefCreated() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		a.assembleReferencedElement(true, false).Line(),
		If(Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		a.retrievePath(false).Line(),
		Return(a.defineReference(true, "update", false), True(), Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")),
	}
}

func (a assembleReferenceWriter) writeSliceTreeReferenceRefUpdated() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.assembleReferencedElement(true, false).Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		Var().Id("el").Id("*" + title(a.v.Name)).Line(),
		If(Id("patchRef").Dot("OperationKind").Op("==").Id("OperationKindUpdate")).Block(
			Id("el").Op("=").Id("&element"),
		).Line(),
		a.retrievePath(false).Line(),
		Return(a.defineReference(true, "update", false), True(), Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefRemoved() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(a.assembleReferencedElement(false, false), Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		a.retrievePath(false).Line(),
		Return(a.defineReference(false, "delete", false), True(), Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefReplaced() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(a.assembleReferencedElement(false, false), Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		a.retrievePath(false).Line(),
		Return(a.defineReference(false, "update", false), True(), Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefElementUpdated() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		If(a.assembleReferencedElement(false, true), Id("hasUpdatedDownstream")).Block(
			a.retrievePath(true),
			Return(a.defineReference(false, "update", true), True(), True()),
		),
	}
}

func (a assembleReferenceWriter) writeSliceTreeReferenceRefElementUpdated() *Statement {
	return &Statement{
		If(a.assembleReferencedElement(false, true), Id("hasUpdatedDownstream")).Block(
			a.retrievePath(true),
			Return(a.defineReference(false, "update", true), True(), True()),
		),
	}
}

func (a assembleReferenceWriter) sliceRefHasUpdated() (*Statement, *Statement) {
	return List(Id("patchRef"), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(title(a.f.ValueTypeName)).Index(Id("refID")), Id("hasUpdated")
}
