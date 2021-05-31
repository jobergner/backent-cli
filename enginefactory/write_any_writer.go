package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

type anyKindWriter struct {
	f ast.Field
}

func (a anyKindWriter) receiverParams() *Statement {
	return Id("_any").Id(anyNameByField(a.f))
}

func (a anyKindWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot(anyNameByField(a.f)).Dot("engine").Dot(anyNameByField(a.f)).Call(Id("_any").Dot(anyNameByField(a.f)).Dot("ID"))
}

func (a anyKindWriter) containedElementKind() *Statement {
	return Id("any").Dot(anyNameByField(a.f)).Dot("ElementKind")
}

type anySetterWriter struct {
	f  ast.Field
	v  ast.ConfigType
	_v *ast.ConfigType
}

func (a anySetterWriter) wrapperReceiverParams() *Statement {
	return Id("_any").Id(anyNameByField(a.f))
}

func (a anySetterWriter) createChild() *Statement {
	return Id(a.v.Name).Op(":=").Id("_any").Dot(anyNameByField(a.f)).Dot("engine").Dot("create" + title(a.v.Name)).Call(True())
}

func (a anySetterWriter) callSetter() *Statement {
	return Id("_any").Dot(anyNameByField(a.f)).Dot("set" + title(a.v.Name)).Call(Id(a.v.Name).Dot("ID").Call())
}

func (a anySetterWriter) receiverParams() *Statement {
	return Id("_any").Id(anyNameByField(a.f) + "Core")
}

func (a anySetterWriter) params() *Statement {
	return Id(a.v.Name + "ID").Id(title(a.v.Name + "ID"))
}

func (a anySetterWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot("engine").Dot(anyNameByField(a.f)).Call(Id("_any").Dot("ID")).Dot(anyNameByField(a.f))
}

func (a anySetterWriter) otherValueIsSet() *Statement {
	return Id("any").Dot(title(a._v.Name)).Op("!=").Lit(0)
}

func (a anySetterWriter) deleteOtherValue() *Statement {
	return Id("any").Dot("engine").Dot("delete" + title(a._v.Name)).Call(Id("any").Dot(title(a._v.Name)))
}

func (a anySetterWriter) unsetIDInContainer() *Statement {
	return Id("any").Dot(title(a._v.Name)).Op("=").Lit(0)
}

func (a anySetterWriter) setElementKind() *Statement {
	return Id("any").Dot("ElementKind").Op("=").Id("ElementKind" + title(a.v.Name))
}

func (a anySetterWriter) setChildID() *Statement {
	return Id("any").Dot(title(a.v.Name)).Op("=").Id(a.v.Name + "ID")
}

func (a anySetterWriter) updateContainerInPatch() *Statement {
	return Id("any").Dot("engine").Dot("Patch").Dot(title(anyNameByField(a.f))).Index(Id("any").Dot("ID")).Op("=").Id("any")
}

type anyDeleteChildWriter struct {
	f ast.Field
	v *ast.ConfigType
}

func (d anyDeleteChildWriter) receiverParams() *Statement {
	return Id("_any").Id(anyNameByField(d.f) + "Core")
}

func (d anyDeleteChildWriter) reassignAnyContainer() *Statement {
	return Id("any").Op(":=").Id("_any").Dot("engine").Dot(anyNameByField(d.f)).Call(Id("_any").Dot("ID")).Dot(anyNameByField(d.f))
}

func (d anyDeleteChildWriter) deleteChild() *Statement {
	return Id("any").Dot("engine").Dot("delete" + title(d.v.Name)).Call(Id("any").Dot(title(d.v.Name)))
}
