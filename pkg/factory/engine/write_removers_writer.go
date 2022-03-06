package engine

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
	return "Remove" + Title(r.f.Name) + optionalSuffix
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

func (r remover) zeroValueID() *Statement {
	switch {
	case r.f.HasAnyValue || r.f.HasPointerValue:
		return Id(Title(r.f.ValueTypeName) + "ID").Values()
	default:
		return Lit(0)
	}
}

func (r remover) toRemoveComparator() *Statement {
	switch {
	case r.v.IsBasicType:
		return r.engine().Dot(BasicTypes[r.v.Name]).Call(Id("valID")).Dot("Value")
	case r.f.HasAnyValue || r.f.HasPointerValue:
		return Id(Title(r.v.Name) + "ID").Call(Id("complexID").Dot("ChildID"))
	default:
		return Id(r.v.Name + "ID")
	}
}

func (r remover) eachElementLiteral() *Statement {
	switch {
	case r.v.IsBasicType:
		return Id("valID")
	case !r.f.HasAnyValue && !r.f.HasPointerValue:
		return Id(r.v.Name + "ID")
	default:
		return Id("complexID")
	}
}

func (r remover) eachElement() *Statement {
	return List(Id("i"), r.eachElementLiteral()).Op(":=").Range().Add(r.field())
}

func (r remover) deleteElement() *Statement {
	switch {
	case r.f.HasAnyValue && !r.f.HasPointerValue:
		return r.engine().Dot("delete"+Title(r.f.ValueTypeName)).Call(r.eachElementLiteral(), True())
	default:
		return r.engine().Dot("delete" + Title(r.f.ValueTypeName)).Call(r.eachElementLiteral())
	}
}
