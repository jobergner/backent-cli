package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type anyKindWriter struct {
	f ast.Field
}

func (a anyKindWriter) receiverParams() *Statement {
	return Id("_any").Id(Title(AnyValueTypeName(&a.f)))
}

func (a anyKindWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot(AnyValueTypeName(&a.f)).Dot("engine").Dot(AnyValueTypeName(&a.f)).Call(Id("_any").Dot(AnyValueTypeName(&a.f)).Dot("ID"))
}

func (a anyKindWriter) containedElementKind() *Statement {
	return Id("any").Dot(AnyValueTypeName(&a.f)).Dot("ElementKind")
}

type anySetterWriter struct {
	f  ast.Field
	v  ast.ConfigType
	_v *ast.ConfigType
}

func (a anySetterWriter) wrapperReceiverParams() *Statement {
	return Id("_any").Id(Title(AnyValueTypeName(&a.f)))
}

func (a anySetterWriter) reassignAnyContainerWrapper() *Statement {
	return Id("any").Op(":=").Id("_any").Dot(AnyValueTypeName(&a.f)).Dot("engine").Dot(AnyValueTypeName(&a.f)).Call(Id("_any").Dot(AnyValueTypeName(&a.f)).Dot("ID"))
}

func (a anySetterWriter) isAlreadyRequestedElement() *Statement {
	return Id("any").Dot(AnyValueTypeName(&a.f)).Dot("ElementKind").Op("==").Id("ElementKind" + Title(a.v.Name))
}

func (a anySetterWriter) isOperationKindDelete() *Statement {
	return Id("any").Dot(AnyValueTypeName(&a.f)).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a anySetterWriter) currentElement() *Statement {
	return Id("any").Dot(Title(a.v.Name)).Call()
}

func (a anySetterWriter) createChild() *Statement {
	return Id(a.v.Name).Op(":=").Id("any").Dot(AnyValueTypeName(&a.f)).Dot("engine").Dot("create"+Title(a.v.Name)).Call(Id("any").Dot(AnyValueTypeName(&a.f)).Dot("ParentElementPath"), Id("any").Dot(AnyValueTypeName(&a.f)).Dot("FieldIdentifier"))
}

func (a anySetterWriter) callSetter() *Statement {
	return Id("any").Dot(AnyValueTypeName(&a.f)).Dot("be"+Title(a.v.Name)).Call(Id(a.v.Name).Dot("ID").Call(), True())
}

func (a anySetterWriter) receiverParams() *Statement {
	return Id("_any").Id(AnyValueTypeName(&a.f) + "Core")
}

func (a anySetterWriter) params() (*Statement, *Statement) {
	return Id(a.v.Name + "ID").Id(Title(a.v.Name + "ID")), Id("deleteCurrentChild").Bool()
}

func (a anySetterWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot("engine").Dot(AnyValueTypeName(&a.f)).Call(Id("_any").Dot("ID")).Dot(AnyValueTypeName(&a.f))
}

func (a anySetterWriter) setElementKind() *Statement {
	return Id("any").Dot("ElementKind").Op("=").Id("ElementKind" + Title(a.v.Name))
}

func (a anySetterWriter) setChildID() *Statement {
	return Id("any").Dot(Title(a.v.Name)).Op("=").Id(a.v.Name + "ID")
}

func (a anySetterWriter) updateContainerInPatch() *Statement {
	return Id("any").Dot("engine").Dot("Patch").Dot(Title(AnyValueTypeName(&a.f))).Index(Id("any").Dot("ID")).Op("=").Id("any")
}

type anyDeleteChildWriter struct {
	f ast.Field
	v *ast.ConfigType
}

func (d anyDeleteChildWriter) receiverParams() *Statement {
	return Id("_any").Id(AnyValueTypeName(&d.f) + "Core")
}

func (d anyDeleteChildWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot("engine").Dot(AnyValueTypeName(&d.f)).Call(Id("_any").Dot("ID")).Dot(AnyValueTypeName(&d.f))
}

func (d anyDeleteChildWriter) deleteChild() *Statement {
	return Id("any").Dot("engine").Dot("delete" + Title(d.v.Name)).Call(Id(Title(d.v.Name) + "ID").Call(Id("any").Dot("ChildID")))
}

type anyRefWriter struct {
	f ast.Field
	v *ast.ConfigType
}

func (a anyRefWriter) typeName() string {
	return AnyValueTypeName(&a.f)
}

func (a anyRefWriter) wrapperName() string {
	return a.typeName() + "Wrapper"
}

func (a anyRefWriter) typeRefName() string {
	return Title(a.typeName()) + "Ref"
}

func (a anyRefWriter) receiverParams() *Statement {
	return Id("_any").Id(a.typeRefName())
}

func (a anyRefWriter) elementKind() *Statement {
	return Id("_any").Dot(a.wrapperName()).Dot("Kind").Call()
}

func (a anyRefWriter) child() *Statement {
	return Id("_any").Dot(a.wrapperName()).Dot(Title(a.v.Name)).Call()
}
