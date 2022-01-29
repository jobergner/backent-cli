package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type referenceWriter struct {
	f ast.Field
}

func (r referenceWriter) receiverParams() *Statement {
	return Id("_ref").Id(Title(r.f.ValueTypeName))
}

func (r referenceWriter) returns() (*Statement, *Statement) {
	return Id(Title(r.f.ValueTypeName)), Bool()
}

func (r referenceWriter) reassignRef() *Statement {
	return Id("ref").Op(":=").Id("_ref").Dot(r.f.ValueTypeName).Dot("engine").Dot(r.f.ValueTypeName).Call(Id("_ref").Dot(r.f.ValueTypeName).Dot("ID"))
}

func (r referenceWriter) isOperationKindDelete() *Statement {
	return Id("ref").Dot(r.f.ValueTypeName).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (r referenceWriter) returnIsSet() *Statement {
	return Id("ref").Dot(r.f.ValueTypeName).Dot("ID").Op("!=").Lit(0)
}

func (r referenceWriter) deleteSelf() *Statement {
	return Id("ref").Dot(r.f.ValueTypeName).Dot("engine").Dot("delete" + Title(r.f.ValueTypeName)).Call(Id("ref").Dot(r.f.ValueTypeName).Dot("ID"))
}

func (r referenceWriter) declareParent() *Statement {
	return Id("parent").Op(":=").Id("ref").Dot(r.f.ValueTypeName).Dot("engine").Dot(Title(r.f.Parent.Name)).Call(Id("ref").Dot(r.f.ValueTypeName).Dot("ParentID")).Dot(r.f.Parent.Name)
}

func (r referenceWriter) parentIsDeleted() *Statement {
	return Id("parent").Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (r referenceWriter) setRefIDInParent() *Statement {
	return Id("parent").Dot(Title(r.f.Name)).Op("=").Lit(0)
}

func (r referenceWriter) setParentOperationKind() *Statement {
	return Id("parent").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (r referenceWriter) updateParentInPatch() *Statement {
	return Id("ref").Dot(r.f.ValueTypeName).Dot("engine").Dot("Patch").Dot(Title(r.f.Parent.Name)).Index(Id("parent").Dot("ID")).Op("=").Id("parent")
}

func (r referenceWriter) returnTypeOfGet() string {
	if r.f.HasAnyValue {
		return anyNameByField(r.f)
	}
	return Title(r.f.ValueType().Name)
}

func (r referenceWriter) getReferencedElement() *Statement {
	return Id("ref").Dot(r.f.ValueTypeName).Dot("engine").Dot(r.returnTypeOfGet()).Call(Id("ref").Dot(r.f.ValueTypeName).Dot("ReferencedElementID"))
}

func (r referenceWriter) wrapIntoAnyRefWrapper() *Statement {
	return Id(anyNameByField(r.f) + "Ref").Values(Dict{
		Id(anyNameByField(r.f) + "Wrapper"): Id("anyContainer"),
		Id(anyNameByField(r.f)):             Id("anyContainer").Dot(anyNameByField(r.f)),
	})
}

func (r referenceWriter) dereferenceCondition() *Statement {
	return Return(Id("ref").Dot(r.f.ValueTypeName).Dot("engine").Dot(r.returnTypeOfGet()).Call(Id("ref").Dot(r.f.ValueTypeName).Dot("ReferencedElementID")))
}

type dereferenceWriter struct {
	f ast.Field
	v ast.ConfigType
}

func (d dereferenceWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (d dereferenceWriter) name() string {
	return "dereference" + Title(d.f.ValueTypeName) + "s" + d.optionalSuffix()
}

func (d dereferenceWriter) optionalSuffix() string {
	if len(d.f.ValueTypes) > 1 {
		return Title(d.v.Name)
	}
	return ""
}

func (d dereferenceWriter) idParam() string {
	return d.v.Name + "ID"
}

func (d dereferenceWriter) params() *Statement {
	return Id(d.idParam()).Id(Title(d.v.Name) + "ID")
}

func (d dereferenceWriter) dereferenceCondition() *Statement {
	if d.f.HasAnyValue {
		return Id("anyContainer").Dot(anyNameByField(d.f)).Dot(Title(d.v.Name)).Op("==").Id(d.v.Name + "ID")
	}
	return Id("ref").Dot(d.f.ValueTypeName).Dot("ReferencedElementID").Op("==").Id(d.v.Name + "ID")
}

func (d dereferenceWriter) allIdsSliceName() string {
	return "all" + Title(d.f.ValueTypeName) + "IDs"
}

func (d dereferenceWriter) declareAllIDs() *Statement {
	return Id(d.allIdsSliceName()).Op(":=").Id("engine").Dot("all" + Title(d.f.ValueTypeName) + "IDs").Call()
}

func (d dereferenceWriter) allIDsLoopConditions() *Statement {
	return List(Id("_"), Id("refID")).Op(":=").Range().Id(d.allIdsSliceName())
}

func (d dereferenceWriter) declareRef() *Statement {
	return Id("ref").Op(":=").Id("engine").Dot(d.f.ValueTypeName).Call(Id("refID"))
}

func (d dereferenceWriter) declareAnyContainer() *Statement {
	return Id("anyContainer").Op(":=").Id("ref").Dot("Get").Call()
}

func (d dereferenceWriter) anyContainerContainsElemenKind() *Statement {
	return Id("anyContainer").Dot(anyNameByField(d.f)).Dot("ElementKind").Op("!=").Id("ElementKind" + Title(d.v.Name))
}

func (d dereferenceWriter) declareParent() *Statement {
	return Id("parent").Op(":=").Id("engine").Dot(Title(d.f.Parent.Name)).Call(Id("ref").Dot(d.f.ValueTypeName).Dot("ParentID"))
}

func (d dereferenceWriter) removeChildReferenceFromParent() *Statement {
	return Id("parent").Dot("Remove" + Title(d.f.Name) + d.optionalSuffix()).Call(Id(d.v.Name + "ID"))
}

func (d dereferenceWriter) unsetRef() *Statement {
	return Id("ref").Dot("Unset").Call()
}

func (d dereferenceWriter) returnSliceToPool() *Statement {
	return Id(d.f.ValueTypeName + "IDSlicePool").Dot("Put").Call(Id(d.allIdsSliceName()))
}
