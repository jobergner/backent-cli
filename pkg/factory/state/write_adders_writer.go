package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type adderWriter struct {
	t ast.ConfigType
	f ast.Field
	v *ast.ConfigType
}

func (a adderWriter) receiverParams() *Statement {
	return Id(a.receiverName()).Id(Title(a.t.Name))
}

func (a adderWriter) name() string {
	var optionalSuffix string
	if len(a.f.ValueTypes) > 1 {
		optionalSuffix = Title(a.v.Name)
	}

	return "Add" + Title(Singular(a.f.Name)) + optionalSuffix
}

func (a adderWriter) idParam() string {
	return a.v.Name + "ID"
}

func (a adderWriter) params() *Statement {
	switch {
	case a.v.IsBasicType:
		return Id(Singular(a.f.Name)).Id(a.f.ValueType().Name)
	case a.f.HasPointerValue:
		return Id(a.idParam()).Id(Title(a.v.Name) + "ID")
	default:
		return Empty()
	}
}

func (a adderWriter) returns() string {
	switch {
	case a.f.ValueType().IsBasicType || a.f.HasPointerValue:
		return ""
	default:
		return Title(a.v.Name)
	}
}

func (a adderWriter) reassignElement() *Statement {
	return Id(a.t.Name).Op(":=").Id(a.receiverName()).Dot(a.t.Name).Dot("engine").Dot(Title(a.t.Name)).Call(Id(a.receiverName()).Dot(a.t.Name).Dot("ID"))
}

func (a adderWriter) isOperationKindDelete() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a adderWriter) elementCore() *Statement {
	return Id(a.t.Name).Dot(a.t.Name)
}

func (a adderWriter) engine() *Statement {
	return a.elementCore().Dot("engine")
}

func (a adderWriter) referencedElementDoesntExist() *Statement {
	return a.engine().Dot(Title(a.v.Name)).Call(Id(a.idParam())).Dot(a.v.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a adderWriter) returnIfReferencedElementIsAlreadyReferencedReturnCondition() *Statement {
	switch {
	case a.f.HasAnyValue:
		return Id(Title(a.v.Name) + "ID").Call(Id("anyContainer").Dot(anyNameByField(a.f)).Dot("ChildID")).Op("==").Id(a.idParam())
	default:
		return Id("currentRef").Dot(a.f.ValueTypeName).Dot("ReferencedElementID").Op("==").Id(a.idParam())
	}
}

func (a adderWriter) returnIfReferencedElementIsAlreadyReferenced() *Statement {
	return For(List(Id("_"), Id("currentRefID")).Op(":=").Range().Add(a.elementCore()).Dot(Title(a.f.Name))).Block(
		Id("currentRef").Op(":=").Add(a.engine()).Dot(a.f.ValueTypeName).Call(Id("currentRefID")),
		OnlyIf(a.f.HasAnyValue, Id("anyContainer").Op(":=").Add(a.engine()).Dot(anyNameByField(a.f)).Call(Id("currentRef").Dot(a.f.ValueTypeName).Dot("ReferencedElementID"))),
		If(a.returnIfReferencedElementIsAlreadyReferencedReturnCondition()).Block(
			Return(),
		),
	)
}

func (a adderWriter) returnDeletedElement() *Statement {
	switch {
	case a.v.IsBasicType || a.f.HasPointerValue:
		return Empty()

	default:
		return Id(Title(a.v.Name)).Values(Dict{
			Id(a.v.Name): Id(a.v.Name + "Core").Values(Dict{
				Id("OperationKind"): Id("OperationKindDelete"),
				Id("engine"):        Add(a.engine()),
			})})
	}
}

func (a adderWriter) createNewElement() *Statement {
	switch {
	case a.v.IsBasicType:
		return Id(Singular(a.f.Name)+"Value").Op(":=").Add(a.engine()).Dot("create"+Title(a.v.Name)+"Value").Call(Add(a.elementCore()).Dot("Path"), Id(FieldPathIdentifier(a.f)), Id(Singular(a.f.Name)))
	default:
		return Id(a.v.Name).Op(":=").Add(a.engine()).Dot("create"+Title(a.v.Name)).Call(Add(a.elementCore()).Dot("Path"), Id(FieldPathIdentifier(a.f)))
	}
}

func (a adderWriter) createAnyContainerCallParams() *Statement {
	switch {
	case a.f.HasPointerValue && !a.f.HasAnyValue:
		return Call(False(), Nil(), Lit(""))
	case a.f.HasPointerValue && a.f.HasAnyValue:
		return Call(Int().Call(a.elementID()), Int().Call(Id(a.idParam())), Id("ElementKind"+Title(a.v.Name)), Add(a.elementCore()).Dot("Path"), Id(FieldPathIdentifier(a.f))).Dot(anyNameByField(a.f))
	case a.f.HasAnyValue:
		return Call(Int().Call(a.elementID()), Int().Call(Id(a.v.Name).Dot(a.v.Name).Dot("ID")), Id("ElementKind"+Title(a.v.Name)), Add(a.elementCore()).Dot("Path"), Id(FieldPathIdentifier(a.f))).Dot(anyNameByField(a.f))
	default:
		return Call(False(), Add(a.elementCore()).Dot("path"), Lit(""))
	}
}

func (a adderWriter) createAnyContainer() *Statement {
	return Id("anyContainer").Op(":=").Add(a.engine()).
		Dot("create" + Title(anyNameByField(a.f))).
		Add(a.createAnyContainerCallParams())
}

func (a adderWriter) setAnyContainerCallParams() *Statement {
	switch {
	case a.f.HasPointerValue:
		return Call(Id(a.idParam()), False())
	default:
		return Call(Id(a.v.Name).Dot(a.v.Name).Dot("ID"), False())
	}
}

func (a adderWriter) setAnyContainer() *Statement {
	return Id("anyContainer").Dot("set" + Title(a.v.Name)).Add(a.setAnyContainerCallParams())
}

func (a adderWriter) createRefCallParams() *Statement {
	switch {
	case a.f.HasAnyValue:
		return Call(a.elementCore().Dot("Path"), Id(FieldPathIdentifier(a.f)), Id("anyContainer").Dot("ID"), a.elementCore().Dot("ID"), Id("ElementKind"+Title(a.v.Name)), Int().Call(Id(a.idParam())))
	default:
		return Call(a.elementCore().Dot("Path"), Id(FieldPathIdentifier(a.f)), Id(a.idParam()), Add(a.elementCore()).Dot("ID"))
	}
}

func (a adderWriter) createRef() *Statement {
	return Id("ref").Op(":=").Add(a.engine()).Dot("create" + Title(a.f.ValueTypeName)).Add(a.createRefCallParams())
}

func (a adderWriter) toAppend() *Statement {
	switch {
	case a.f.ValueType().IsBasicType:
		return Id(Singular(a.f.Name) + "Value").Dot("ID")

	case a.f.HasPointerValue:
		return Id("ref").Dot("ID")

	case a.f.HasAnyValue:
		return Id("anyContainer").Dot("ID")

	default:
		return Id(a.f.ValueType().Name).Dot(a.f.ValueType().Name).Dot("ID")
	}
}

func (a adderWriter) appendElement() *Statement {
	return Add(a.elementCore()).Dot(Title(a.f.Name)).Op("=").Append(
		Add(a.elementCore()).Dot(Title(a.f.Name)),
		a.toAppend(),
	)
}

func (a adderWriter) setOperationKindUpdate() *Statement {
	return Add(a.elementCore()).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (a adderWriter) signElement() *Statement {
	return a.elementCore().Dot("Meta").Dot("sign").Call(a.engine().Dot("BroadcastingClientID"))
}

func (a adderWriter) updateElementInPatch() *Statement {
	return a.engine().Dot("Patch").Dot(Title(a.t.Name)).Index(a.elementID()).Op("=").Add(a.elementCore())
}

func (a adderWriter) elementID() *Statement {
	return Add(a.elementCore()).Dot("ID")
}

func (a adderWriter) receiverName() string {
	return "_" + a.t.Name
}
