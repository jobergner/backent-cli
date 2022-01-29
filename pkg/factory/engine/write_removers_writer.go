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

func (r remover) eachWrapperRangee() *Statement {
	switch {
	case !r.f.HasPointerValue || !r.f.HasAnyValue:
		return r.field()
	default:
		return Id("refs")
	}
}

func (r remover) getWrapper() *Statement {
	switch {
	case r.f.HasAnyValue:
		return r.engine().Dot(anyNameByField(r.f)).Call(Id("wrapperID"))
	default:
		return r.engine().Dot(r.f.ValueTypeName).Call(Id("wrapperID"))
	}
}

func (r remover) eachWrapperIndexLit() *Statement {
	switch {
	case !r.f.HasPointerValue || !r.f.HasAnyValue:
		return Id("_")
	default:
		return Id("refID")
	}
}

func (r remover) eachWrapper() *Statement {
	return List(r.eachWrapperIndexLit(), Id("wrapperID")).Op(":=").Range().Add(r.eachWrapperRangee())
}

func (r remover) usedWrapperID() *Statement {
	switch {
	case r.f.HasPointerValue && r.f.HasAnyValue:
		return Id("refID")
	default:
		return Id("wrapperID")
	}
}

func (r remover) getElementID() *Statement {
	switch {
	case r.f.HasAnyValue:
		return Id("wrapper").Dot(Title(r.v.Name)).Call().Dot("ID").Call()
	default:
		return Id("wrapper").Dot(r.f.ValueTypeName).Dot("ReferencedElementID")
	}
}

func (r remover) toRemoveComparator() *Statement {
	switch {
	case r.v.IsBasicType:
		return Id("val")
	default:
		return Id(r.v.Name + "ID")
	}
}

func (r remover) eachElementLiteral() *Statement {
	switch {
	case r.v.IsBasicType:
		return Id("val")
	case !r.f.HasAnyValue && !r.f.HasPointerValue:
		return Id(r.v.Name + "ID")
	default:
		return Id("wrapperID")
	}
}

func (r remover) eachElement() *Statement {
	return List(Id("i"), r.eachElementLiteral()).Op(":=").Range().Add(r.field())
}

func (r remover) defaultValueForBasicType(typeLiteral string) interface{} {
	switch typeLiteral {
	case "bool":
		return bool(false)
	case "string":
		return string("")
	case "int8":
		return int8(0)
	case "byte":
		return byte(0)
	case "int16":
		return int16(0)
	case "uint16":
		return uint16(0)
	case "rune":
		return rune(0)
	case "uint32":
		return uint32(0)
	case "int64":
		return int64(0)
	case "uint64":
		return uint64(0)
	case "int":
		return int(0)
	case "uint":
		return uint(0)
	case "uintptr":
		return uintptr(0)
	case "float32":
		return float32(0)
	case "float64":
		return float64(0)
	case "complex64":
		return complex64(0)
	case "complex128":
		return complex128(0)
	}

	return 0
}

func (r remover) deleteElement() *Statement {
	switch {
	case r.f.HasAnyValue && !r.f.HasPointerValue:
		return r.engine().Dot("delete"+Title(r.f.ValueTypeName)).Call(r.eachElementLiteral(), True())
	default:
		return r.engine().Dot("delete" + Title(r.f.ValueTypeName)).Call(r.eachElementLiteral())
	}
}
