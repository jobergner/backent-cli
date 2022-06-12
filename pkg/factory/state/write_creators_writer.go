package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type creatorWrapperWriter struct {
	t ast.ConfigType
}

func (cw creatorWrapperWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (cw creatorWrapperWriter) name() string {
	return "Create" + Title(cw.t.Name)
}

func (cw creatorWrapperWriter) returns() string {
	return Title(cw.t.Name)
}

func (cw creatorWrapperWriter) createElement() *Statement {
	return Id("engine").Dot("create"+Title(cw.t.Name)).Call(Id("newPath").Call(), Id(cw.t.Name+"Identifier"))
}

type creatorWriter struct {
	typeName string
	f        *ast.Field
}

func (c creatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c creatorWriter) name() string {
	return "create" + Title(c.typeName)
}

func (c creatorWriter) returns() string {
	return Title(c.typeName)
}

func (c creatorWriter) params() (*Statement, *Statement) {
	return Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier")
}

func (c creatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.typeName + "Core")
}

func (c creatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c creatorWriter) generateID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(Title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c creatorWriter) assignExtendedPath() *Statement {
	return Id("element").Dot("Path").Op("=").Id("p").Dot("extendAndCopy").Call(Id("fieldIdentifier"), Int().Call(Id("element").Dot("ID")), Id("ElementKind"+Title(c.typeName)), Lit(0))
}

func (c creatorWriter) assignJsonPath() *Statement {
	return Id("element").Dot("JSONPath").Op("=").Id("element").Dot("Path").Dot("toJSONPath").Call()
}

func (c creatorWriter) setHasParent() *Statement {
	return Id("element").Dot("HasParent").Op("=").Len(Id("element").Dot("Path")).Op(">").Lit(1)
}

func (c creatorWriter) createChildElementCall() *Statement {
	switch {
	case c.f.ValueType().IsBasicType:
		return Call(Id("element").Dot("Path"), Id(FieldPathIdentifier(*c.f)), Lit(defaultValueForBasicType(c.f.ValueTypeName)))
	case c.f.HasAnyValue && c.f.HasPointerValue:
		return Call(True(), Id("p").Dot(c.f.Name).Call())
	case c.f.HasAnyValue && !c.f.HasPointerValue:
		// (int(element.ID), int(originElement.player.ID), ElementKindPlayer, element.Path, item_originIdentifier)
		return Call(Int().Call(Id("element").Dot("ID")), Int().Call(Id(c.f.Name+"Element").Dot(c.f.ValueType().Name).Dot("ID")), Id("ElementKind"+Title(c.f.ValueType().Name)), Id("element").Dot("Path"), Id(FieldPathIdentifier(*c.f)))
	default:
		return Call(Id("element").Dot("Path"), Id(FieldPathIdentifier(*c.f)))
	}
}

func (c creatorWriter) createChildSubElement() *Statement {
	return Id(c.f.Name+"Element").Op(":=").Id("engine").Dot("create"+Title(c.f.ValueType().Name)).Call(Id("element").Dot("Path"), Id(FieldPathIdentifier(*c.f)))
}

func (c creatorWriter) createChildElement() *Statement {
	switch {
	case c.f.ValueType().IsBasicType:
		return Id("element" + Title(c.f.Name)).Op(":=").Id("engine").Dot("create" + Title(BasicTypes[c.f.ValueTypeName])).Add(c.createChildElementCall())
	default:
		return Id("element" + Title(c.f.Name)).Op(":=").Id("engine").Dot("create" + Title(c.f.ValueTypeName)).Add(c.createChildElementCall())
	}
}
func (c creatorWriter) setChildElement() *Statement {
	switch {
	case c.f.ValueType().IsBasicType:
		return Id("element").Dot(Title(c.f.Name)).Op("=").Id("element" + Title(c.f.Name)).Dot("ID")
	default:
		return Id("element").Dot(Title(c.f.Name)).Op("=").Id("element" + Title(c.f.Name)).Dot(c.f.ValueTypeName).Dot("ID")
	}
}

func (c creatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c creatorWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(c.typeName)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c creatorWriter) returnElement() *Statement {
	return Id(Title(c.typeName)).Values(Dict{
		Id(c.typeName): Id("element"),
	})
}

type generatedTypeCreatorWriter struct {
	f        ast.Field
	typeName string
}

func (c generatedTypeCreatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c generatedTypeCreatorWriter) name() string {
	return "create" + Title(c.typeName)
}

func (c generatedTypeCreatorWriter) referencedElementIDParam() string {
	switch {
	case c.f.HasAnyValue:
		return Title(anyNameByField(c.f)) + "ID"
	default:
		return Title(c.f.ValueType().Name) + "ID"
	}
}

func (c generatedTypeCreatorWriter) params() *Statement {
	switch {
	case c.f.HasAnyValue:
		return List(Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier"), Id("referencedElementID").Id(c.referencedElementIDParam()), Id("parentID").Id(Title(c.f.Parent.Name)+"ID"), Id("childKind").Id("ElementKind"), Id("childID").Int())
	default:
		return List(Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier"), Id("referencedElementID").Id(c.referencedElementIDParam()), Id("parentID").Id(Title(c.f.Parent.Name)+"ID"), Id("childID").Int())
	}
}

func (c generatedTypeCreatorWriter) returns() string {
	return c.typeName + "Core"
}

func (c generatedTypeCreatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.typeName + "Core")
}

func (c generatedTypeCreatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c generatedTypeCreatorWriter) setReferencedElementID() *Statement {
	return Id("element").Dot("ReferencedElementID").Op("=").Id("referencedElementID")
}

func (c generatedTypeCreatorWriter) setParentID() *Statement {
	return Id("element").Dot("ParentID").Op("=").Id("parentID")
}

func (c generatedTypeCreatorWriter) setIDRef() *Statement {
	return Id("element").Dot("ID").Op("=").Id(Title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
	// switch {
	// case c.f.HasAnyValue:
	// 	return Id("element").Dot("ID").Op("=").Id(Title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
	// default:
	// 	return Id("element").Dot("ID").Op("=").Id(Title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
	// }
}

func (c generatedTypeCreatorWriter) setIDAny() *Statement {
	return Id("element").Dot("ID").Op("=").Id(Title(anyNameByField(c.f)) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c generatedTypeCreatorWriter) setChildID() *Statement {
	return Id("element").Dot("ChildID").Op("=").Id("childID")
}

func (c generatedTypeCreatorWriter) assignPathCall() *Statement {
	switch {
	case c.f.HasAnyValue:
		return Call(Id("fieldIdentifier"), Lit(0), Id("childKind"), Int().Call(Id("element").Dot("ID")))
	default:
		return Call(Id("fieldIdentifier"), Lit(0), Id("ElementKind"+Title(c.f.ValueType().Name)), Int().Call(Id("element").Dot("ID")))
	}
}

func (c generatedTypeCreatorWriter) assignPath() *Statement {
	return Id("element").Dot("Path").Op("=").Id("p").Dot("extendAndCopy").Add(c.assignPathCall())
}

func (c generatedTypeCreatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c generatedTypeCreatorWriter) assignElementToPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(c.typeName)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c generatedTypeCreatorWriter) createChildElement() *Statement {
	return Id("element"+Title(c.f.ValueType().Name)).Op(":=").Id("engine").Dot("create"+Title(c.f.ValueType().Name)).Call(Id("p"), Id("fieldIdentifier"))
}

// func (c generatedTypeCreatorWriter) assignChildElement() *Statement {
// 	return Id("element").Dot(Title(c.f.ValueType().Name)).Op("=").Id("element" + Title(c.f.ValueType().Name)).Dot(c.f.ValueType().Name).Dot("ID")
// }

func (c generatedTypeCreatorWriter) assignElementKind() *Statement {
	return Id("element").Dot("ElementKind").Op("=").Id("childKind")
}

func (c generatedTypeCreatorWriter) setChildElementPath() *Statement {
	return Id("element").Dot("ParentElementPath").Op("=").Id("p")
}

func (c generatedTypeCreatorWriter) setFieldIdentifier() *Statement {
	return Id("element").Dot("FieldIdentifier").Op("=").Id("fieldIdentifier")
}

func (c generatedTypeCreatorWriter) returnElement() *Statement {
	return Id(Title(anyNameByField(c.f))).Values(Dict{
		Id(anyNameByField(c.f)): Id("element"),
	})
}
