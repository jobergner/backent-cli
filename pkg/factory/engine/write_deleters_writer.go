package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type deleteTypeWrapperWriter struct {
	t ast.ConfigType
}

func (d deleteTypeWrapperWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteTypeWrapperWriter) name() string {
	return "Delete" + Title(d.t.Name)
}

func (d deleteTypeWrapperWriter) idParam() string {
	return d.t.Name + "ID"
}

func (d deleteTypeWrapperWriter) params() *Statement {
	return Id(d.idParam()).Id(Title(d.t.Name) + "ID")
}

func (d deleteTypeWrapperWriter) getElement() *Statement {
	return Id(d.t.Name).Op(":=").Id("engine").Dot(Title(d.t.Name)).Call(Id(d.idParam())).Dot(d.t.Name)
}

func (d deleteTypeWrapperWriter) hasParent() *Statement {
	return Id(d.t.Name).Dot("HasParent")
}

func (d deleteTypeWrapperWriter) deleteElement() *Statement {
	return Id("engine").Dot("delete" + Title(d.t.Name)).Call(Id(d.idParam()))
}

type deleteTypeWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (d deleteTypeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteTypeWriter) name() string {
	return "delete" + Title(d.t.Name)
}

func (d deleteTypeWriter) idParam() string {
	return d.t.Name + "ID"
}

func (d deleteTypeWriter) params() *Statement {
	return Id(d.idParam()).Id(Title(d.t.Name) + "ID")
}

func (d deleteTypeWriter) getElement() *Statement {
	return Id(d.t.Name).Op(":=").Id("engine").Dot(Title(d.t.Name)).Call(Id(d.idParam())).Dot(d.t.Name)
}

func (d deleteTypeWriter) isOperationKindDelete() *Statement {
	return Id(d.t.Name).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (d deleteTypeWriter) setOperationKind() *Statement {
	return Id(d.t.Name).Dot("OperationKind").Op("=").Id("OperationKindDelete")
}

func (d deleteTypeWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(d.t.Name)).Index(Id(d.t.Name).Dot("ID")).Op("=").Id(d.t.Name)
}

func (d deleteTypeWriter) loopedElementIdentifier() string {
	return Singular(d.f.Name) + "ID"
}

func (d deleteTypeWriter) loopConditions() *Statement {
	return List(Id("_"), Id(d.loopedElementIdentifier())).Op(":=").Range().Id(d.t.Name).Dot(Title(d.f.Name))
}

func (d deleteTypeWriter) deleteElementInLoop() *Statement {
	deleteFunc := Id("engine").Dot("delete" + Title(d.f.ValueTypeName))
	if !d.f.HasPointerValue && d.f.HasAnyValue {
		return deleteFunc.Call(Id(d.loopedElementIdentifier()), True())
	}
	return deleteFunc.Call(Id(d.loopedElementIdentifier()))
}

func (d deleteTypeWriter) deleteElement() *Statement {
	deleteFunc := Id("engine").Dot("delete" + Title(d.f.ValueTypeName))
	if !d.f.HasPointerValue && d.f.HasAnyValue {
		return deleteFunc.Call(Id(d.t.Name).Dot(Title(d.f.Name)), True())
	}
	return deleteFunc.Call(Id(d.t.Name).Dot(Title(d.f.Name)))
}

func (d deleteTypeWriter) existsInState() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(Title(d.t.Name)).Index(Id(d.idParam())).Id(";").Id("ok")
}

func (d deleteTypeWriter) deleteFromPatch() *Statement {
	return Delete(Id("engine").Dot("Patch").Dot(Title(d.t.Name)), Id(d.idParam()))
}

func (d deleteTypeWriter) dereferenceField(field *ast.Field) *Statement {
	var suffix string
	if field.HasAnyValue {
		suffix = Title(d.t.Name)
	}
	return Id("engine").Dot("dereference" + Title(field.Parent.Name) + Title(Singular(field.Name)) + "Refs" + suffix).Call(Id(d.idParam()))
}

type deleteGeneratedTypeWriter struct {
	f             ast.Field
	valueTypeName func() string
}

func (d deleteGeneratedTypeWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d deleteGeneratedTypeWriter) name() string {
	return "delete" + Title(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) idParam() string {
	return d.valueTypeName() + "ID"
}

func (d deleteGeneratedTypeWriter) params() *Statement {
	return Id(d.idParam()).Id(Title(d.valueTypeName()) + "ID")
}

func (d deleteGeneratedTypeWriter) getElement() *Statement {
	return Id(d.valueTypeName()).Op(":=").Id("engine").Dot(d.valueTypeName()).Call(Id(d.idParam())).Dot(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) isOperationKindDelete() *Statement {
	return Id(d.valueTypeName()).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (d deleteGeneratedTypeWriter) deleteChild() *Statement {
	return Id(d.valueTypeName()).Dot("deleteChild").Call()
}

func (d deleteGeneratedTypeWriter) deleteAnyContainer() *Statement {
	return Id("engine").Dot("delete"+Title(anyNameByField(d.f))).Call(Id(d.valueTypeName()).Dot("ReferencedElementID"), False())
}

func (d deleteGeneratedTypeWriter) setOperationKind() *Statement {
	return Id(d.valueTypeName()).Dot("OperationKind").Op("=").Id("OperationKindDelete")
}

func (d deleteGeneratedTypeWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(Title(d.valueTypeName())).Index(Id(d.valueTypeName()).Dot("ID")).Op("=").Id(d.valueTypeName())
}

func (d deleteGeneratedTypeWriter) existsInState() *Statement {
	return List(Id("_"), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(Title(d.valueTypeName())).Index(Id(d.idParam())).Id(";").Id("ok")
}

func (d deleteGeneratedTypeWriter) deleteFromPatch() *Statement {
	return Delete(Id("engine").Dot("Patch").Dot(Title(d.valueTypeName())), Id(d.idParam()))
}
