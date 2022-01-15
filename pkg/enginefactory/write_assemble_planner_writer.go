package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type assemblePlannerWriter struct {
	t ast.ConfigType
	f ast.Field
	v *ast.ConfigType
}

func (a assemblePlannerWriter) receiverParams() *Statement {
	return Id(a.receiverName()).Id(Title(a.t.Name))
}

func (a assemblePlannerWriter) name() string {
	var optionalSuffix string
	if len(a.f.ValueTypes) > 1 {
		optionalSuffix = Title(a.v.Name)
	}

	return "Add" + Title(Singular(a.f.Name)) + optionalSuffix
}

func (a assemblePlannerWriter) idParam() string {
	return a.v.Name + "ID"
}

func (a assemblePlannerWriter) params() *Statement {
	switch {
	case a.v.IsBasicType:
		return Id(Singular(a.f.Name)).Id(a.f.ValueType().Name)
	case a.f.HasPointerValue:
		return Id(a.idParam()).Id(Title(a.v.Name) + "ID")
	default:
		return Empty()
	}
}

func (a assemblePlannerWriter) returns() string {
	switch {
	case a.f.ValueType().IsBasicType || a.f.HasPointerValue:
		return ""
	default:
		return Title(a.v.Name)
	}
}

func (a assemblePlannerWriter) reassignElement() *Statement {
	return Id(a.t.Name).Op(":=").Id(a.receiverName()).Dot(a.t.Name).Dot("engine").Dot(Title(a.t.Name)).Call(Id(a.receiverName()).Dot(a.t.Name).Dot("ID"))
}

func (a assemblePlannerWriter) isOperationKindDelete() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a assemblePlannerWriter) elementCore() *Statement {
	return Id(a.t.Name).Dot(a.t.Name)
}

func (a assemblePlannerWriter) engine() *Statement {
	return a.elementCore().Dot("engine")
}

func (a assemblePlannerWriter) referencedElementDoesntExist() *Statement {
	return a.engine().Dot(Title(a.v.Name)).Call(Id(a.idParam())).Dot(a.v.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a assemblePlannerWriter) returnIfReferencedElementIsAlreadyReferencedReturnCondition() *Statement {
	switch {
	case a.f.HasAnyValue:
		return Id("anyContainer").Dot(anyNameByField(a.f)).Dot(Title(a.v.Name)).Op("==").Id(a.idParam())
	default:
		return Id("currentRef").Dot(a.f.ValueTypeName).Dot("ReferencedElementID").Op("==").Id(a.idParam())
	}
}

func (a assemblePlannerWriter) returnIfReferencedElementIsAlreadyReferenced() *Statement {
	return For(List(Id("_"), Id("currentRefID")).Op(":=").Range().Add(a.elementCore()).Dot(Title(a.f.Name))).Block(
		Id("currentRef").Op(":=").Add(a.engine()).Dot(a.f.ValueTypeName).Call(Id("currentRefID")),
		OnlyIf(a.f.HasAnyValue, Id("anyContainer").Op(":=").Add(a.engine()).Dot(anyNameByField(a.f)).Call(Id("currentRef").Dot(a.f.ValueTypeName).Dot("ReferencedElementID"))),
		If(a.returnIfReferencedElementIsAlreadyReferencedReturnCondition()).Block(
			Return(),
		),
	)
}

func (a assemblePlannerWriter) returnDeletedElement() *Statement {
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

func (a assemblePlannerWriter) createNewElement() *Statement {
	return Id(a.v.Name).Op(":=").Add(a.engine()).Dot("create"+Title(a.v.Name)).
		Call(Add(a.elementCore()).Dot("path"), Id(FieldPathIdentifier(a.f)))
}

func (a assemblePlannerWriter) createAnyContainerCallParams() *Statement {
	switch {
	case a.f.HasPointerValue && !a.f.HasAnyValue:
		return Call(False(), Nil(), Lit(""))
	case a.f.HasPointerValue && a.f.HasAnyValue:
		return Call(False(), Add(a.elementCore()).Dot("path"), Lit("")).Dot(anyNameByField(a.f))
	case a.f.HasAnyValue:
		return Call(False(), Add(a.elementCore()).Dot("path"), Id(FieldPathIdentifier(a.f))).Dot(anyNameByField(a.f))
	default:
		return Call(False(), Add(a.elementCore()).Dot("path"), Lit(""))
	}
}

func (a assemblePlannerWriter) createAnyContainer() *Statement {
	return Id("anyContainer").Op(":=").Add(a.engine()).
		Dot("create" + Title(anyNameByField(a.f))).
		Add(a.createAnyContainerCallParams())
}

func (a assemblePlannerWriter) setAnyContainerCallParams() *Statement {
	switch {
	case a.f.HasPointerValue:
		return Call(Id(a.idParam()), False())
	default:
		return Call(Id(a.v.Name).Dot(a.v.Name).Dot("ID"), False())
	}
}

func (a assemblePlannerWriter) setAnyContainer() *Statement {
	return Id("anyContainer").Dot("set" + Title(a.v.Name)).Add(a.setAnyContainerCallParams())
}

func (a assemblePlannerWriter) createRefCallParams() *Statement {
	switch {
	case a.f.HasAnyValue:
		return Call(a.elementCore().Dot("path"), Id(FieldPathIdentifier(a.f)), Id("anyContainer").Dot("ID"), a.elementCore().Dot("ID"), Id("ElementKind"+Title(a.v.Name)), Int().Call(Id(a.idParam())))
	default:
		return Call(a.elementCore().Dot("path"), Id(FieldPathIdentifier(a.f)), Id(a.idParam()), Add(a.elementCore()).Dot("ID"))
	}
}

func (a assemblePlannerWriter) createRef() *Statement {
	return Id("ref").Op(":=").Add(a.engine()).Dot("create" + Title(a.f.ValueTypeName)).Add(a.createRefCallParams())
}

func (a assemblePlannerWriter) toAppend() *Statement {
	switch {
	case a.f.ValueType().IsBasicType:
		return Id(Singular(a.f.Name))

	case a.f.HasPointerValue:
		return Id("ref").Dot("ID")

	case a.f.HasAnyValue:
		return Id("anyContainer").Dot("ID")

	default:
		return Id(a.f.ValueType().Name).Dot(a.f.ValueType().Name).Dot("ID")
	}
}

func (a assemblePlannerWriter) appendElement() *Statement {
	return Add(a.elementCore()).Dot(Title(a.f.Name)).Op("=").Append(
		Add(a.elementCore()).Dot(Title(a.f.Name)),
		a.toAppend(),
	)
}

func (a assemblePlannerWriter) setOperationKindUpdate() *Statement {
	return Add(a.elementCore()).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (a assemblePlannerWriter) updateElementInPatch() *Statement {
	return a.engine().Dot("Patch").Dot(Title(a.t.Name)).Index(a.elementID()).Op("=").Add(a.elementCore())
}

func (a assemblePlannerWriter) elementID() *Statement {
	return Add(a.elementCore()).Dot("ID")
}

func (a assemblePlannerWriter) receiverName() string {
	return "_" + a.t.Name
}
