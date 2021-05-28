package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

type setterWriter struct {
	t ast.ConfigType
	f ast.Field
}

func (s setterWriter) receiverParams() *Statement {
	return Id(s.receiverName()).Id(s.t.Name)
}

func (s setterWriter) name() string {
	if s.f.ValueType().IsBasicType {
		return "Set" + title(s.f.Name)
	}
	return "Add" + title(pluralizeClient.Singular(s.f.Name))
}

func (s setterWriter) newValueParam() string {
	return "new" + title(s.f.Name)
}

func (s setterWriter) params() *Statement {
	return Id(s.newValueParam()).Id(s.f.ValueTypeName)
}

func (s setterWriter) returns() string {
	return s.t.Name
}

func (s setterWriter) reassignElement() *Statement {
	return Id(s.t.Name).Op(":=").Id(s.receiverName()).Dot(s.t.Name).Dot("engine").Dot(title(s.t.Name)).Call(Id(s.receiverName()).Dot(s.t.Name).Dot("ID"))
}

func (s setterWriter) isOperationKindDelete() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setterWriter) setAttribute() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot(title(s.f.Name)).Op("=").Id(s.newValueParam())
}

func (s setterWriter) setOperationKind() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setterWriter) setOperationKindUpdate() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setterWriter) updateElementInPatch() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("engine").Dot("Patch").Dot(title(s.t.Name)).Index(s.elementID()).Op("=").Id(s.t.Name).Dot(s.t.Name)
}

func (s setterWriter) elementID() *Statement {
	return Id(s.t.Name).Dot(s.t.Name).Dot("ID")
}

func (s setterWriter) receiverName() string {
	return "_" + s.t.Name
}

type setRefFieldWeiter struct {
	f ast.Field
	v *ast.ConfigType
}

func (s setRefFieldWeiter) receiverParams() *Statement {
	return Id("_" + s.f.Parent.Name).Id(s.f.Parent.Name)
}

func (s setRefFieldWeiter) name() string {
	var optionalSuffix string
	if s.f.HasAnyValue {
		optionalSuffix = title(s.v.Name)
	}
	return "Set" + title(s.f.Name) + optionalSuffix
}

func (s setRefFieldWeiter) idParam() string {
	return s.v.Name + "ID"
}

func (s setRefFieldWeiter) params() *Statement {
	return Id(s.idParam()).Id(title(s.v.Name) + "ID")
}

func (s setRefFieldWeiter) returns() string {
	return s.f.Parent.Name
}

func (s setRefFieldWeiter) reassignElement() *Statement {
	return Id(s.f.Parent.Name).Op(":=").Id("_" + s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot(title(s.f.Parent.Name)).Call(Id("_" + s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID"))
}

func (s setRefFieldWeiter) isOperationKindDelete() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setRefFieldWeiter) isReferencedElementDeleted() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot(title(s.v.Name)).Call(Id(s.idParam())).Dot(s.v.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (s setRefFieldWeiter) isRefAlreadyAssigned() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(title(s.f.Name)).Op("!=").Lit(0)
}

func (s setRefFieldWeiter) deleteExistingRef() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("delete" + title(s.f.ValueTypeName)).Call(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(title(s.f.Name)))
}

func (s setRefFieldWeiter) createAnyContainer() *Statement {
	return Id("anyContainer").Op(":=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("create" + title(anyNameByField(s.f))).Call(False())
}

func (s setRefFieldWeiter) setAnyContainer() *Statement {
	return Id("anyContainer").Dot(anyNameByField(s.f)).Dot("set" + title(s.v.Name)).Call(Id(s.idParam()))
}

func (s setRefFieldWeiter) createNewRef() *Statement {
	referencedElementStatement := Id(s.v.Name + "ID")
	if s.f.HasAnyValue {
		referencedElementStatement = Id("anyContainer").Dot(anyNameByField(s.f)).Dot("ID")
	}
	return Id("ref").Op(":=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("create"+title(s.f.ValueTypeName)).Call(referencedElementStatement, Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID"))
}

func (s setRefFieldWeiter) setNewRef() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot(title(s.f.Name)).Op("=").Id("ref").Dot("ID")
}

func (s setRefFieldWeiter) setOperationKind() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (s setRefFieldWeiter) setItemInPatch() *Statement {
	return Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("engine").Dot("Patch").Dot(title(s.f.Parent.Name)).Index(Id(s.f.Parent.Name).Dot(s.f.Parent.Name).Dot("ID")).Op("=").Id(s.f.Parent.Name).Dot(s.f.Parent.Name)
}
