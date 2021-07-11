package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type adderWriter struct {
	t ast.ConfigType
	f ast.Field
	v *ast.ConfigType
}

func (a adderWriter) receiverParams() *Statement {
	return Id(a.receiverName()).Id(a.t.Name)
}

func (a adderWriter) name() string {
	var optionalSuffix string
	if len(a.f.ValueTypes) > 1 {
		optionalSuffix = Title(a.v.Name)
	}
	if a.f.ValueType().IsBasicType {
		return "Add" + Title(a.f.Name)
	}
	return "Add" + Title(Singular(a.f.Name)) + optionalSuffix
}

func (a adderWriter) idParam() string {
	return a.v.Name + "ID"
}

func (a adderWriter) params() *Statement {
	if a.v.IsBasicType {
		return Id(a.f.Name).Id("..." + a.f.ValueType().Name)
	}
	if a.f.HasPointerValue {
		return Id(a.idParam()).Id(Title(a.v.Name) + "ID")
	}
	return Empty()
}

func (a adderWriter) returns() string {
	if a.f.ValueType().IsBasicType || a.f.HasPointerValue {
		return ""
	}
	return a.v.Name
}

func (a adderWriter) reassignElement() *Statement {
	return Id(a.t.Name).Op(":=").Id(a.receiverName()).Dot(a.t.Name).Dot("engine").Dot(Title(a.t.Name)).Call(Id(a.receiverName()).Dot(a.t.Name).Dot("ID"))
}

func (a adderWriter) isOperationKindDelete() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a adderWriter) referencedElementDoesntExist() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("engine").Dot(Title(a.v.Name)).Call(Id(a.idParam())).Dot(a.v.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (a adderWriter) returnDeletedElement() *Statement {
	if a.v.IsBasicType || a.f.HasPointerValue {
		return Empty()
	}
	return Id(a.v.Name).Values(Dict{
		Id(a.v.Name): Id(a.v.Name + "Core").Values(Dict{
			Id("OperationKind"): Id("OperationKindDelete"),
			Id("engine"):        Id(a.t.Name).Dot(a.t.Name).Dot("engine"),
		})})
}

func (a adderWriter) createNewElement() *Statement {
	return Id(a.v.Name).Op(":=").Id(a.t.Name).Dot(a.t.Name).Dot("engine").Dot("create"+Title(a.v.Name)).Call(Id(a.t.Name).Dot(a.t.Name).Dot("path").Dot(a.f.Name).Call(), True())
}

func (a adderWriter) createAnyContainer() *Statement {
	secondParam := Id(a.t.Name).Dot(a.t.Name).Dot("path").Dot(a.f.Name).Call()
	if a.f.HasPointerValue {
		secondParam = Nil()
	}
	return Id("anyContainer").Op(":=").Id(a.t.Name).Dot(a.t.Name).Dot("engine").Dot("create"+Title(anyNameByField(a.f))).Call(False(), secondParam).Dot(anyNameByField(a.f))
}

func (a adderWriter) setAnyContainer() *Statement {
	statement := Id("anyContainer").Dot("set" + Title(a.v.Name))
	if a.f.HasPointerValue {
		return statement.Call(Id(a.idParam()), False())
	}
	return statement.Call(Id(a.v.Name).Dot(a.v.Name).Dot("ID"), False())
}

func (a adderWriter) createRef() *Statement {
	statement := Id("ref").Op(":=").Id(a.t.Name).Dot(a.t.Name).Dot("engine").Dot("create" + Title(a.f.ValueTypeName))

	if a.f.HasAnyValue {
		return statement.Call(Id("anyContainer").Dot("ID"), Id(a.t.Name).Dot(a.t.Name).Dot("ID"))
	}

	return statement.Call(Id(a.idParam()), Id(a.t.Name).Dot(a.t.Name).Dot("ID"))
}

func (a adderWriter) appendElement() *Statement {

	var toAppend *Statement
	if a.f.ValueType().IsBasicType {
		toAppend = Id(a.f.Name + "...")
	} else {
		if a.f.HasPointerValue {
			toAppend = Id("ref").Dot("ID")
		} else if a.f.HasAnyValue {
			toAppend = Id("anyContainer").Dot("ID")
		} else {
			toAppend = Id(a.f.ValueType().Name).Dot(a.f.ValueType().Name).Dot("ID")
		}
	}

	appendStatement := Id(a.t.Name).Dot(a.t.Name).Dot(Title(a.f.Name)).Op("=").Append(
		Id(a.t.Name).Dot(a.t.Name).Dot(Title(a.f.Name)),
		toAppend,
	)

	return appendStatement
}

func (a adderWriter) setOperationKindUpdate() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (a adderWriter) updateElementInPatch() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("engine").Dot("Patch").Dot(Title(a.t.Name)).Index(a.elementID()).Op("=").Id(a.t.Name).Dot(a.t.Name)
}

func (a adderWriter) elementID() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("ID")
}

func (a adderWriter) receiverName() string {
	return "_" + a.t.Name
}
