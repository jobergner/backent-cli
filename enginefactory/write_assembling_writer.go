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

func (a assembleElement) elementHasUpdated() (*Statement, *Statement) {
	fieldIdentifier := Id(a.f.ValueTypeName + "ID")
	if !a.f.HasPointerValue && !a.f.HasSliceValue {
		fieldIdentifier = Id(a.dataElementName()).Dot(title(a.f.Name))
	} else if a.f.HasPointerValue && !a.f.HasSliceValue {
		fieldIdentifier = Id(a.t.Name + "ID")
	}
	conditionVars := List(Id("tree"+title(a.f.ValueTypeName)), Id("include"), Id("childHasUpdated"))
	condition := conditionVars.Op(":=").Id("engine").Dot("assemble"+title(a.f.ValueTypeName)).Call(fieldIdentifier, Id("check"), Id("config"))
	return condition, Id("include")
}

func (a assembleElement) setHasUpdatedTrue() *Statement {
	return Id("hasUpdated").Op("=").True()
}

func (a assembleElement) appendToElementsInField() *Statement {
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Append(Id(a.treeElementName()).Dot(title(a.f.Name)), Id("tree"+title(a.f.ValueTypeName)))
}

func (a assembleElement) setFieldElement() *Statement {
	var optionalShare string
	if !a.f.HasPointerValue {
		optionalShare = "&"
	}
	return Id(a.treeElementName()).Dot(title(a.f.Name)).Op("=").Id(optionalShare + "tree" + title(a.f.ValueTypeName))
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

func (a assembleElement) createAnyContainer() *Statement {
	return Id("engine").Dot(anyNameByField(*a.f)).Call(Id(anyNameByField(*a.f) + "ID")).Dot(anyNameByField(*a.f))
}
