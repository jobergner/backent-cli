package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type remover struct {
	t ast.ConfigType
	f ast.Field
	v *ast.ConfigType
}

func (r remover) receiverParams() *Statement {
	return Id(r.receiverName()).Id(Title(r.t.Name))
}

func (r remover) receiverName() string {
	return "_" + r.t.Name
}

func (r remover) name() string {
	var optionalSuffix string
	if r.f.HasAnyValue {
		optionalSuffix = Title(r.v.Name)
	}
	return "Remove" + Title(Singular(r.f.Name)) + optionalSuffix
}

func (r remover) toRemoveParamName() string {
	switch {
	case r.f.HasAnyValue:
		return r.v.Name + "ToRemove"
	default:
		return Singular(r.f.Name) + "ToRemove"
	}
}

func (r remover) params() *Statement {
	switch {
	case r.v.IsBasicType:
		return Id(r.toRemoveParamName()).Id(r.v.Name)
	default:
		return Id(r.toRemoveParamName()).Id(Title(r.v.Name) + "ID")
	}
}

func (r remover) returns() string {
	return Title(r.t.Name)
}

func (r remover) elementCore() *Statement {
	return Id(r.t.Name).Dot(r.t.Name)
}

func (r remover) engine() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("engine")
}

func (r remover) reassignElement() *Statement {
	return Id(r.t.Name).Op(":=").Id(r.receiverName()).Dot(r.t.Name).Dot("engine").Dot(Title(r.t.Name)).Call(Id(r.receiverName()).Dot(r.t.Name).Dot("ID"))
}

func (r remover) isOperationKindDelete() *Statement {
	return Id(r.t.Name).Dot(r.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (r remover) field() *Statement {
	return r.elementCore().Dot(Title(r.f.Name))
}

func (r remover) eachElementLiteral() *Statement {
	switch {
	case r.v.IsBasicType:
		return Id("valID")
	case !r.f.HasAnyValue && !r.f.HasPointerValue:
		return Id(r.v.Name + "ID")
	default:
		return Id("id")
	}
}

func (r remover) eachElement() *Statement {
	return List(Id("i"), r.eachElementLiteral()).Op(":=").Range().Add(r.field())
}

func (r remover) valueTypeID() string {
	switch {
	case r.f.ValueType().IsBasicType:
		return Title(r.f.ValueType().Name) + "ValueID"
	default:
		return Title(ValueTypeName(&r.f)) + "ID"
	}
}

func (r remover) deleteElement() *Statement {
	switch {
	case r.f.HasAnyValue && !r.f.HasPointerValue:
		return r.engine().Dot("delete"+Title(ValueTypeName(&r.f))).Call(r.eachElementLiteral(), True())
	case r.f.ValueType().IsBasicType:
		return r.engine().Dot("delete" + Title(r.f.ValueType().Name) + "Value").Call(r.eachElementLiteral())
	default:
		return r.engine().Dot("delete" + Title(ValueTypeName(&r.f))).Call(r.eachElementLiteral())
	}
}

func (r remover) idsDontMatch() []Code {
	switch {
	case r.f.HasPointerValue || r.f.HasAnyValue:
		return []Code{Id("childID").Op(":=").Add(r.engine().Dot(ValueTypeName(&r.f)).Call(Id("id"))).Dot(ValueTypeName(&r.f)).Dot("ChildID"), Id(Title(r.v.Name) + "ID").Call(Id("childID")).Op("!=").Id(r.toRemoveParamName())}
	case r.v.IsBasicType:
		return []Code{r.engine().Dot(r.v.Name + "Value").Call(r.eachElementLiteral()).Dot("Value").Op("!=").Id(r.toRemoveParamName())}
	default:
		return []Code{r.eachElementLiteral().Add(Op("!=").Id(r.toRemoveParamName()))}
	}
}
