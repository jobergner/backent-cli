package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

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
	return cw.t.Name
}

func (cw creatorWrapperWriter) createElement() *Statement {
	return Id("engine").Dot("create"+Title(cw.t.Name)).Call(Id("newPath").Call(Id(cw.t.Name+"Identifier")), True())
}

type creatorWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (c creatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c creatorWriter) name() string {
	return "create" + Title(c.t.Name)
}

func (c creatorWriter) returns() string {
	return c.t.Name
}

func (c creatorWriter) params() (*Statement, *Statement) {
	return Id("p").Id("path"), Id("extendWithID").Bool()
}

func (c creatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.t.Name + "Core")
}

func (c creatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c creatorWriter) generateID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(Title(c.t.Name) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c creatorWriter) setHasParent() *Statement {
	return Id("element").Dot("HasParent").Op("=").Len(Id("p")).Op(">").Lit(1)
}

func (c creatorWriter) setPath() *Statement {
	return Id("element").Dot("path").Op("=").Id("p")
}

func (c creatorWriter) extendPathWithID() *Statement {
	return Id("element").Dot("path").Op("=").Id("element").Dot("path").Dot("id").Call(Int().Call(Id("element").Dot("ID")))
}

func (c creatorWriter) setJSONPath() *Statement {
	return Id("element").Dot("Path").Op("=").Id("element").Dot("path").Dot("toJSONPath").Call()
}

func (c creatorWriter) createChildElement() *Statement {
	statement := Id("element" + Title(c.f.Name)).Op(":=").Id("engine").Dot("create" + Title(c.f.ValueTypeName))
	if c.f.HasAnyValue {
		return statement.Call(True(), Id("p").Dot(c.f.Name).Call())
	}
	return statement.Call(Id("p").Dot(c.f.Name).Call(), False())
}
func (c creatorWriter) setChildElement() *Statement {
	return Id("element").Dot(Title(c.f.Name)).Op("=").Id("element" + Title(c.f.Name)).Dot(c.f.ValueTypeName).Dot("ID")
}

func (c creatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c creatorWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(c.t.Name)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c creatorWriter) returnElement() *Statement {
	return Id(c.t.Name).Values(Dict{
		Id(c.t.Name): Id("element"),
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

func (c generatedTypeCreatorWriter) params() *Statement {
	referencedElementType := c.f.ValueType().Name
	if c.f.HasAnyValue {
		referencedElementType = anyNameByField(c.f)
	}
	return List(Id("referencedElementID").Id(Title(referencedElementType)+"ID"), Id("parentID").Id(Title(c.f.Parent.Name)+"ID"))
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

func (c generatedTypeCreatorWriter) setID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(Title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c generatedTypeCreatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c generatedTypeCreatorWriter) assignElementToPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(c.typeName)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c generatedTypeCreatorWriter) createChildElement() *Statement {
	return Id("element"+Title(c.f.ValueType().Name)).Op(":=").Id("engine").Dot("create"+Title(c.f.ValueType().Name)).Call(Id("childElementPath"), False())
}

func (c generatedTypeCreatorWriter) assignChildElement() *Statement {
	return Id("element").Dot(Title(c.f.ValueType().Name)).Op("=").Id("element" + Title(c.f.ValueType().Name)).Dot(c.f.ValueType().Name).Dot("ID")
}

func (c generatedTypeCreatorWriter) assignElementKind() *Statement {
	return Id("element").Dot("ElementKind").Op("=").Id("ElementKind" + Title(c.f.ValueType().Name))
}

func (c generatedTypeCreatorWriter) setChildElementPath() *Statement {
	return Id("element").Dot("ChildElementPath").Op("=").Id("childElementPath")
}

func (c generatedTypeCreatorWriter) returnElement() *Statement {
	return Id(anyNameByField(c.f)).Values(Dict{
		Id(anyNameByField(c.f)): Id("element"),
	})
}
