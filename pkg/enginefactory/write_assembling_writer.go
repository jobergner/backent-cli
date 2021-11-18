package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type assembleTreeWriter struct {
	t *ast.ConfigType
}

func (w assembleTreeWriter) clearMap(engineFieldName string, capitalizeSubMapName bool) *Statement {
	subMapName := w.t.Name
	if capitalizeSubMapName {
		subMapName = Title(w.t.Name)
	}
	return For(Id("key").Op(":=").Range().Id("engine").Dot(engineFieldName).Dot(subMapName)).Block(
		Delete(Id("engine").Dot(engineFieldName).Dot(subMapName), Id("key")),
	)
}

func (a assembleTreeWriter) dataElementName() string {
	return a.t.Name + "Data"
}

func (a assembleTreeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (a assembleTreeWriter) params() *Statement {
	return Id("assembleEntireTree").Bool()
}

func (a assembleTreeWriter) patchLoopConditions() *Statement {
	return List(Id("_"), Id(a.dataElementName())).Op(":=").Range().Id("engine").Dot("Patch").Dot(Title(a.t.Name))
}

func (a assembleTreeWriter) elementHasNoParent() *Statement {
	return Id("!" + a.dataElementName()).Dot("HasParent")
}

func (a assembleTreeWriter) elementNonExistentInTree() (*Statement, *Statement) {
	condition := List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("Tree").Dot(Title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID"))
	return condition, Id("!ok")
}

func (a assembleTreeWriter) assembleElement() *Statement {
	variableNames := List(Id(a.t.Name), Id("include"), Id("_"))
	return variableNames.Op(":=").Id("engine").Dot("assemble"+Title(a.t.Name)).Call(Id(a.dataElementName()).Dot("ID"), Nil(), Id("config"))
}

func (a assembleTreeWriter) setElementInTree() *Statement {
	return Id("engine").Dot("Tree").Dot(Title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID")).Op("=").Id(a.t.Name)
}

func (a assembleTreeWriter) stateLoopConditions() *Statement {
	return List(Id("_"), Id(a.dataElementName())).Op(":=").Range().Id("engine").Dot("State").Dot(Title(a.t.Name))
}

func (a assembleTreeWriter) createConfig() *Statement {
	return Id("config").Op(":=").Id("assembleConfig").Values(Dict{Id("forceInclude"): Id("assembleEntireTree")})
}

type assembleElementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (a assembleElementWriter) treeElementName() string {
	return "element"
}

func (a assembleElementWriter) dataElementName() string {
	return a.t.Name + "Data"
}

func (a assembleElementWriter) treeTypeName() string {
	return a.t.Name
}

func (a assembleElementWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (a assembleElementWriter) name() string {
	return "assemble" + Title(a.t.Name)
}

func (a assembleElementWriter) idParam() string {
	return a.t.Name + "ID"
}

func (a assembleElementWriter) params() (*Statement, *Statement, *Statement) {
	return Id(a.idParam()).Id(Title(a.t.Name) + "ID"), Id("check").Id("*recursionCheck"), Id("config").Id("assembleConfig")
}

func (a assembleElementWriter) returns() (*Statement, *Statement, *Statement) {
	return Id(a.treeTypeName()), Bool(), Bool()
}

func (a assembleElementWriter) checkIsDefined() *Statement {
	return Id("check").Op("!=").Nil()
}

func (a assembleElementWriter) elementExistsInCheck() (*Statement, *Statement) {
	return Id("alreadyExists").Op(":=").Id("check").Dot(a.t.Name).Index(Id(a.idParam())), Id("alreadyExists")
}

func (a assembleElementWriter) returnEmpty() (*Statement, *Statement, *Statement) {
	return Id(a.t.Name).Values(), False(), False()
}

func (a assembleElementWriter) checkElement() *Statement {
	return Id("check").Dot(a.t.Name).Index(Id(a.idParam())).Op("=").True()
}

func (a assembleElementWriter) getElementFromPatch() *Statement {
	return List(Id(a.dataElementName()), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(Title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElementWriter) getElementFromState() *Statement {
	return Id(a.dataElementName()).Op("=").Id("engine").Dot("State").Dot(Title(a.t.Name)).Index(Id(a.idParam()))
}

func (a assembleElementWriter) declareTreeElement() *Statement {
	return Var().Id(a.treeElementName()).Id(a.treeTypeName())
}

func (a assembleElementWriter) shouldRetrieveFromCache() (*Statement, *Statement) {
	return List(Id("cached"+Title(a.t.Name)), Id("ok")).Op(":=").Id("engine").Dot("assembleCache").Dot(a.t.Name).Index(Id(a.dataElementName()).Dot("ID")), Id("ok").Op("&&").Id("!config").Dot("forceInclude")
}

func (a assembleElementWriter) returnCachedElement() *Statement {
	return Return(Id("cached"+Title(a.t.Name)).Dot(a.t.Name), Id("cached"+Title(a.t.Name)).Dot("hasUpdated").Op("||").Id("config").Dot("forceInclude"), Id("cached"+Title(a.t.Name)).Dot("hasUpdated"))
}

func (a assembleElementWriter) typeFieldOn(from string) *Statement {
	return Id("engine").Dot(from).Dot(Title(a.t.Name)).Index(Id(a.dataElementName()).Dot("ID")).Dot(Title(a.f.Name))
}

func (a assembleElementWriter) sliceFieldLoopConditions() *Statement {
	loopVars := List(Id("_"), Id(a.f.ValueTypeName+"ID"))
	mergeFuncName := "merge" + Title(a.f.ValueTypeName) + "IDs"
	return loopVars.Op(":=").Range().Id(mergeFuncName).Call(a.typeFieldOn("State"), a.typeFieldOn("Patch"))
}

func (a assembleElementWriter) usedAssembleID(configType ast.ConfigType, field ast.Field, valueType *ast.ConfigType) *Statement {
	if !field.HasPointerValue && !field.HasAnyValue && !field.HasSliceValue {
		return Id(a.dataElementName()).Dot(Title(field.Name))
	} else if field.HasPointerValue && !field.HasAnyValue && !field.HasSliceValue {
		return Id(configType.Name + "ID")
	} else if field.HasPointerValue && field.HasSliceValue {
		return Id(field.ValueTypeName + "ID")
	} else if field.HasPointerValue && field.HasAnyValue && !field.HasSliceValue {
		return Id(configType.Name + "ID")
	}

	return Id(valueType.Name + "ID")
}

func (a assembleElementWriter) elementHasUpdated(valueType *ast.ConfigType, assembleID *Statement) (*Statement, *Statement) {
	handledElementTypeName := a.f.ValueTypeName
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		handledElementTypeName = valueType.Name
	}
	conditionVars := List(Id("tree"+Title(handledElementTypeName)), Id("include"), Id("childHasUpdated"))
	condition := conditionVars.Op(":=").Id("engine").Dot("assemble"+Title(handledElementTypeName)).Call(assembleID, Id("check"), Id("config"))
	return condition, Id("include")
}

func (a assembleElementWriter) setHasUpdatedTrue() *Statement {
	return Id("hasUpdated").Op("=").True()
}

func (a assembleElementWriter) makeMap() *Statement {
	mapValueType := Id(a.f.ValueType().Name + "Reference")
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		mapValueType = Id("interface{}")
	}
	if a.f.HasPointerValue && a.f.HasAnyValue {
		mapValueType = Id(anyNameByField(*a.f) + "Reference")
	}
	if !a.f.HasPointerValue && !a.f.HasAnyValue {
		mapValueType = Id(a.f.ValueType().Name)
	}
	mapKeyType := Id(Title(a.f.ValueType().Name) + "ID")
	if a.f.HasAnyValue {
		mapKeyType = Int()
	}
	return If(Id(a.treeElementName()).Dot(Title(a.f.Name)).Op("==").Nil()).Block(
		Id(a.treeElementName()).Dot(Title(a.f.Name)).Op("=").Make(Map(mapKeyType).Add(mapValueType)),
	)
}

func (a assembleElementWriter) appendToElementsInField(valueType *ast.ConfigType) *Statement {
	appendedType := valueType.Name
	if !a.f.HasAnyValue && a.f.HasPointerValue || (a.f.HasPointerValue && a.f.HasSliceValue && a.f.HasAnyValue) {
		appendedType = a.f.ValueTypeName
	}
	usedID := "ElementID"
	if !a.f.HasPointerValue {
		usedID = "ID"
	}
	mapKeyType := Id("tree" + Title(appendedType)).Dot(usedID)
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		mapKeyType = Int().Call(Id("tree" + Title(appendedType)).Dot(usedID))
	}
	return Id(a.treeElementName()).Dot(Title(a.f.Name)).Index(mapKeyType).Op("=").Id("tree" + Title(appendedType))
}

func (a assembleElementWriter) setFieldElement(valueType *ast.ConfigType) *Statement {
	var optionalShare string
	if !a.f.HasPointerValue {
		optionalShare = "&"
	}
	usedType := a.f.ValueTypeName
	if a.f.HasAnyValue && !a.f.HasPointerValue {
		usedType = valueType.Name
	}
	return Id(a.treeElementName()).Dot(Title(a.f.Name)).Op("=").Id(optionalShare + "tree" + Title(usedType))
}

func (a assembleElementWriter) setField() *Statement {
	return Id(a.treeElementName()).Dot(Title(a.f.Name)).Op("=").Id(a.dataElementName()).Dot(Title(a.f.Name))
}

func (a assembleElementWriter) setID() *Statement {
	return Id(a.treeElementName()).Dot("ID").Op("=").Id(a.dataElementName()).Dot("ID")
}

func (a assembleElementWriter) setOperationKind() *Statement {
	return Id(a.treeElementName()).Dot("OperationKind").Op("=").Id(a.dataElementName()).Dot("OperationKind")
}

func (a assembleElementWriter) putInCache(cacheName string) *Statement {
	return Id("engine").Dot(cacheName).Dot(a.t.Name).Index(Id(a.treeElementName()).Dot("ID")).Op("=").Id(a.t.Name + "CacheElement").Values(List(Id("hasUpdated").Op(":").Id("hasUpdated"), Id(a.t.Name).Op(":").Id(a.treeElementName())))
}

func (a assembleElementWriter) finalReturn() (*Statement, *Statement, *Statement) {
	return Id(a.treeElementName()), Id("hasUpdated").Op("||").Id("config").Dot("forceInclude"), Id("hasUpdated")
}

func (a assembleElementWriter) anyContainerName() string {
	return anyNameByField(*a.f) + "Container"
}

func (a assembleElementWriter) createAnyContainer() *Statement {
	usedID := Id(anyNameByField(*a.f) + "ID")
	if !a.f.HasSliceValue {
		usedID = Id(a.dataElementName()).Dot(Title(a.f.Name))
	}
	return Id(a.anyContainerName()).Op(":=").Id("engine").Dot(anyNameByField(*a.f)).Call(usedID).Dot(anyNameByField(*a.f))
}

func (a assembleElementWriter) assignIDFromAnyContainer(valueType *ast.ConfigType) *Statement {
	return Id(valueType.Name + "ID").Op(":=").Id(a.anyContainerName()).Dot(Title(valueType.Name))
}

type referenceWriterMode int

const (
	referenceWriterModeForceInclude referenceWriterMode = iota
	referenceWriterModeRefCreate
	referenceWriterModeRefDelete
	referenceWriterModeRefReplace
	referenceWriterModeRefUpdate
	referenceWriterModeElementModified
)

type assembleReferenceWriter struct {
	f    ast.Field
	v    *ast.ConfigType
	mode referenceWriterMode
}

func (a *assembleReferenceWriter) setMode(mode referenceWriterMode) *Statement {
	a.mode = mode
	return Empty()
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
	idType := Title(a.f.Parent.Name) + "ID"
	if a.f.HasSliceValue {
		idType = Title(a.f.ValueTypeName) + "ID"
	}
	return Id(a.idParam()).Id(idType), Id("check").Id("*recursionCheck"), Id("config").Id("assembleConfig")
}

func (a assembleReferenceWriter) returns() (*Statement, *Statement, *Statement) {
	optionalPointer := ""
	if !a.f.HasSliceValue {
		optionalPointer = "*"
	}
	return Id(optionalPointer + a.nextValueName() + "Reference"), Bool(), Bool()
}

func (a assembleReferenceWriter) declareStateElement() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Op(":=").Id("engine").Dot("State").Dot(Title(a.f.Parent.Name)).Index(Id(a.idParam()))
}

func (a assembleReferenceWriter) declarepatchElement() *Statement {
	return List(Id("patch"+Title(a.f.Parent.Name)), Id(a.f.Parent.Name+"IsInPatch")).Op(":=").Id("engine").Dot("Patch").Dot(Title(a.f.Parent.Name)).Index(Id(a.idParam()))
}

func (a assembleReferenceWriter) refIsNotSet() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("==").Lit(0).Op("&&").Params(Id("!" + a.f.Parent.Name + "IsInPatch").Op("||").Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("==").Lit(0))
}

// non-slice gen force: ref := engine.playerTargetRef(patchPlayer.Target)
// non-slice gen el upd:ref := engine.playerTargetRef(statePlayer.Target)
// slice * : ref := engine.equipmentSetEquipmentRef(refID).equipmentSetEquipmentRef
func (a assembleReferenceWriter) declareRef() *Statement {
	if a.f.HasSliceValue {
		return Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id("refID")).Dot(a.f.ValueTypeName)
	}
	if a.mode == referenceWriterModeForceInclude && !a.f.HasSliceValue {
		return &Statement{
			Var().Id("ref").Id(Title(a.f.ValueTypeName)).Line(),
			If(Id("patch" + Title(a.f.Parent.Name))).Dot("ID").Op("==").Lit(0).Block(
				Id("ref").Op("=").Id("engine").Dot(a.f.ValueTypeName).Call(Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name))),
			).Else().Block(
				Id("ref").Op("=").Id("engine").Dot(a.f.ValueTypeName).Call(Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name))),
			).Line(),
		}
	}
	usedElement := "patch"
	if !a.f.HasSliceValue && (a.mode == referenceWriterModeRefDelete || a.mode == referenceWriterModeElementModified) {
		usedElement = "state"
	}
	return Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id(usedElement + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)))
}

// non-slice gen: engine.anyOfPlayer_ZoneItem(ref.playerTargetRef.ReferencedElementID)
// slice gen: 		engine.anyOfPlayer_ZoneItem(ref.ReferencedElementID)
// __ ref updated:engine.anyOfPlayer_ZoneItem(patchRef.ReferencedElementID)
func (a assembleReferenceWriter) declareAnyContainer() *Statement {
	usedID := Id("ref").Dot(a.f.ValueTypeName).Dot("ReferencedElementID")
	if a.f.HasSliceValue {
		usedID = Id("ref").Dot("ReferencedElementID")
		if a.mode == referenceWriterModeRefUpdate {
			usedID = Id("patchRef").Dot("ReferencedElementID")
		}
	}
	return Id("anyContainer").Op(":=").Id("engine").Dot(a.nextValueName()).Call(usedID)
}

// slice non gen: referencedElement := engine.Player(ref.ReferencedElementID).player
// slice gen: referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
// non-slice gen:referencedElement := engine.Player(anyContainer.anyOfPlayer_ZoneItem.Player).player
// non-slice non-gen: referencedElement := engine.Player(ref.itemBoundToRef.ReferencedElementID).player
func (a assembleReferenceWriter) declareReferencedElement() *Statement {
	if a.f.HasAnyValue {
		return Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(Id("anyContainer").Dot(a.nextValueName()).Dot(Title(a.v.Name))).Dot(a.v.Name)
	}
	if a.f.HasSliceValue {
		usedRef := "ref"
		if a.mode == referenceWriterModeRefUpdate {
			usedRef = "patchRef"
		}
		return Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(Id(usedRef).Dot("ReferencedElementID")).Dot(a.v.Name)
	}
	return Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(Id("ref").Dot(a.f.ValueTypeName).Dot("ReferencedElementID")).Dot(a.v.Name)
}

// non-slice gen: 									 engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
// __ on referenced element update:  engine.assembleZoneItem(anyContainer.anyOfPlayer_ZoneItem.ZoneItem, check,
// non-slice non-gen:  							 engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
// __ on referenced element update:  engine.assembleZoneItem(ref.ID(), check,
// slice non-gen:		  							 engine.assembleZoneItem(referencedElement.ID, check, config); hasUpdatedDownstream {
// __ on referenced element update:  engine.assembleItem(ref.ReferencedElementID, check

func (a assembleReferenceWriter) assembleReferencedElement() *Statement {
	usedID := Id("referencedElement").Dot("ID")
	if a.mode == referenceWriterModeElementModified {
		if a.f.HasAnyValue {
			usedID = Id("anyContainer").Dot(a.nextValueName()).Dot(Title(a.v.Name))
		} else if !a.f.HasSliceValue {
			usedID = Id("ref").Dot("ID").Call()
		} else {
			usedID = Id("ref").Dot("ReferencedElementID")
		}
	}
	return List(Id("_"), Id("_"), Id("hasUpdatedDownstream")).Op(":=").Id("engine").Dot("assemble"+Title(a.v.Name)).Call(usedID, Id("check"), Id("config"))
}

// non-slice gen force: ref.playerTargetRef.OperationKind
//   "          create: OperationKindUpdate
//							remove: OperationKindDelete
//							replac: OperationKindUpdate
//							   mod: OperationKindUnchanged
// non-slice non-gen force: SAME
// slice gen     force: ref.OperationKind
//   "          update: patchRef.OperationKind
//							   mod: OperationKindUnchanged
// slice non-gen force: SAME

func (a assembleReferenceWriter) defineReference() *Statement {

	operationKindStatement := Id("OperationKindUpdate")
	if a.mode == referenceWriterModeForceInclude {
		if a.f.HasSliceValue {
			operationKindStatement = Id("ref").Dot("OperationKind")
		} else {
			operationKindStatement = Id("ref").Dot(a.f.ValueTypeName).Dot("OperationKind")
		}
	}
	if a.mode == referenceWriterModeRefDelete {
		operationKindStatement = Id("OperationKindDelete")
	}
	if a.mode == referenceWriterModeElementModified {
		operationKindStatement = Id("OperationKindUnchanged")
	}
	if a.mode == referenceWriterModeRefUpdate {
		operationKindStatement = Id("patchRef").Dot("OperationKind")
	}

	usedID := Id("referencedElement").Dot("ID")
	if a.mode == referenceWriterModeElementModified {
		if a.f.HasAnyValue {
			usedID = Id("anyContainer").Dot(a.nextValueName()).Dot(Title(a.v.Name))
		} else if !a.f.HasSliceValue {
			usedID = Id("ref").Dot("ID").Call()
		}
	}
	if !a.f.HasAnyValue && a.f.HasSliceValue {
		usedID = Id("ref").Dot("ReferencedElementID")
	}
	if a.mode == referenceWriterModeRefUpdate && !a.f.HasAnyValue {
		usedID = Id("patchRef").Dot("ReferencedElementID")
	}

	if a.f.HasAnyValue {
		usedID = Int().Call(usedID)
	}
	optionalShare := "&"
	if a.f.HasSliceValue {
		optionalShare = ""
	}

	dataStatus := Id("referencedDataStatus")
	if a.mode == referenceWriterModeElementModified {
		dataStatus = Id("ReferencedDataModified")
	} else if a.mode == referenceWriterModeForceInclude {
		dataStatus = Id("ReferencedDataUnchanged")
	}

	return Id(optionalShare+a.nextValueName()+"Reference").Values(
		operationKindStatement,
		usedID,
		Id("ElementKind"+Title(a.v.Name)),
		dataStatus,
		Id("referencedElement").Dot("Path"),
	)
}

// non-slice gen force: ref.playerTargetRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
//							 create: referencedDataStatus == ReferencedDataModified
// 							 remove: -
//							replace: -
//									mod: true
// non-slice non-gen: SAME
// slice jen  force: ref.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
// 						update: patchRef.OperationKind == OperationKindUpdate || referencedDataStatus == ReferencedDataModified
// slice non-jen: SAME
func (a assembleReferenceWriter) hasUpdated() *Statement {
	dataStatusIsModified := Id("referencedDataStatus").Op("==").Id("ReferencedDataModified")
	if a.mode == referenceWriterModeForceInclude {
		if a.f.HasSliceValue {
			return Id("ref").Dot("OperationKind").Op("==").Id("OperationKindUpdate").Op("||").Add(dataStatusIsModified)
		} else {
			return Id("ref").Dot(a.f.ValueTypeName).Dot("OperationKind").Op("==").Id("OperationKindUpdate").Op("||").Add(dataStatusIsModified)
		}
	}
	if a.mode == referenceWriterModeRefUpdate {
		return Id("patchRef").Dot("OperationKind").Op("==").Id("OperationKindUpdate").Op("||").Add(dataStatusIsModified)
	}
	if a.mode == referenceWriterModeElementModified {
		return True()
	}

	return dataStatusIsModified
}

func (a assembleReferenceWriter) refWasCreated() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("==").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Lit(0))
}

func (a assembleReferenceWriter) refWasRemoved() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("==").Lit(0))
}

func (a assembleReferenceWriter) refHasBeenReplaced() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Lit(0).Op("&&").Params(Id(a.f.Parent.Name + "IsInPatch").Op("&&").Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Lit(0))
}

func (a assembleReferenceWriter) refWasReplaced() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Id("patch" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name))
}

func (a assembleReferenceWriter) referencedElementGotUpdated() *Statement {
	return Id("state" + Title(a.f.Parent.Name)).Dot(Title(a.f.Name)).Op("!=").Lit(0)
}

func (a assembleReferenceWriter) finalReturn() *Statement {
	if !a.f.HasSliceValue {
		return Nil()
	}
	return Id(a.nextValueName() + "Reference").Values()
}

func (a assembleReferenceWriter) writeTreeReferenceForceInclude() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		Return(a.defineReference(), True(), False()).Line(),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefCreated() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.declareReferencedElement().Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		a.assembleReferencedElement().Line(),
		If(Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		Return(a.defineReference(), True(), a.hasUpdated()),
	}
}

func (a assembleReferenceWriter) writeSliceTreeReferenceRefUpdated() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.declareReferencedElement().Line(),
		a.assembleReferencedElement().Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		Return(a.defineReference(), True(), a.hasUpdated()),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefRemoved() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.declareReferencedElement().Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		If(a.assembleReferencedElement(), Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		Return(a.defineReference(), True(), a.hasUpdated()),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefReplaced() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.declareReferencedElement().Line(),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged").Line(),
		a.assembleReferencedElement().Line(),
		If(Id("hasUpdatedDownstream")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		).Line(),
		Return(a.defineReference(), True(), a.hasUpdated()),
	}
}

func (a assembleReferenceWriter) writeNonSliceTreeReferenceRefElementUpdated() *Statement {
	return &Statement{
		If(Id("check").Op("==").Nil()).Block(
			Id("check").Op("=").Id("newRecursionCheck").Call(),
		).Line(),
		a.declareReferencedElement().Line(),
		If(a.assembleReferencedElement(), Id("hasUpdatedDownstream")).Block(
			Return(a.defineReference(), True(), a.hasUpdated()),
		),
	}
}

func (a assembleReferenceWriter) writeSliceTreeReferenceRefElementUpdated() *Statement {
	return &Statement{
		a.declareReferencedElement().Line(),
		If(a.assembleReferencedElement(), Id("hasUpdatedDownstream")).Block(
			Return(a.defineReference(), True(), a.hasUpdated()),
		),
	}
}

func (a assembleReferenceWriter) sliceRefHasUpdated() (*Statement, *Statement) {
	return List(Id("patchRef"), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(Title(a.f.ValueTypeName)).Index(Id("refID")), Id("hasUpdated")
}
