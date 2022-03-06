package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type existsGetterWriter struct {
	t ast.ConfigType
}

func (e existsGetterWriter) receiverName() string {
	return "_" + e.t.Name
}

func (e existsGetterWriter) receiverParams() *Statement {
	return Id(e.receiverName()).Id(Title(e.t.Name))
}

func (e existsGetterWriter) returnTypes() (*Statement, *Statement) {
	return Id(Title(e.t.Name)), Bool()
}

func (e existsGetterWriter) isNotOperationKindDelete() *Statement {
	return Id(e.t.Name).Dot(e.t.Name).Dot("OperationKind").Op("!=").Id("OperationKindDelete")
}

func (e existsGetterWriter) reassignElement() *Statement {
	return Id(e.t.Name).Op(":=").Id(e.receiverName()).Dot(e.t.Name).Dot("engine").Dot(Title(e.t.Name)).Call(Id(e.receiverName()).Dot(e.t.Name).Dot("ID"))
}

type everyTypeGetterWriter struct {
	t ast.ConfigType
}

func (t everyTypeGetterWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t everyTypeGetterWriter) returns() *Statement {
	return Id("[]" + Title(t.t.Name))
}

func (t everyTypeGetterWriter) allIDs() *Statement {
	return Id(t.t.Name + "IDs").Op(":=").Id("engine").Dot("all" + Title(t.t.Name) + "IDs").Call()
}

func (t everyTypeGetterWriter) sortIDs() *Statement {
	return Id("sort").Dot("Slice").Params(Id(t.t.Name+"IDs"), Func().Params(Id("i"), Id("j").Int()).Bool().Block(
		Return(Id(t.t.Name+"IDs").Index(Id("i")).Op("<").Id(t.t.Name+"IDs").Index(Id("j"))),
	))
}

func (t everyTypeGetterWriter) sliceName() string {
	return t.t.Name + "s"
}

func (t everyTypeGetterWriter) declareSlice() *Statement {
	return Var().Id(t.sliceName()).Id("[]" + Title(t.t.Name))
}

func (t everyTypeGetterWriter) loopConditions() *Statement {
	return List(Id("_"), Id(t.t.Name+"ID")).Op(":=").Range().Id(t.t.Name + "IDs")
}

func (t everyTypeGetterWriter) declareElement() *Statement {
	return Id(t.t.Name).Op(":=").Id("engine").Dot(Title(t.t.Name)).Call(Id(t.t.Name + "ID"))
}

func (t everyTypeGetterWriter) elementHasParent() *Statement {
	return Id(t.t.Name).Dot(t.t.Name).Dot("HasParent")
}

func (t everyTypeGetterWriter) elementSatisfiesMatcher() *Statement {
	return Id("matcher").Call(Id(t.t.Name))
}

func (t everyTypeGetterWriter) returnToIdsSliceToPool() *Statement {
	return Id(t.t.Name + "IDSlicePool").Dot("Put").Call(Id(t.t.Name + "IDs"))
}

func (t everyTypeGetterWriter) appendElement() *Statement {
	return Id(t.sliceName()).Op("=").Append(Id(t.sliceName()), Id(t.t.Name))
}

type typeGetterWriter struct {
	name     func() string
	typeName func() string
}

func (t typeGetterWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t typeGetterWriter) idParam() string {
	return t.typeName() + "ID"
}

func (t typeGetterWriter) params() *Statement {
	return Id(t.idParam()).Id(Title(t.typeName()) + "ID")
}

func (t typeGetterWriter) returns() string {
	return Title(t.typeName())
}

func (t typeGetterWriter) definePatchingElement() *Statement {
	return List(Id("patching"+Title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(Title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnPatching() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id("patching" + Title(t.typeName()))})
}

func (t typeGetterWriter) defineCurrentElement() *Statement {
	return List(Id("current"+Title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(Title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetterWriter) earlyReturnCurrent() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id("current" + Title(t.typeName()))})
}

func (t typeGetterWriter) finalReturn() *Statement {
	return Id(Title(t.typeName())).Values(Dict{Id(t.typeName()): Id(t.typeName() + "Core").Values(Dict{Id("OperationKind"): Id("OperationKindDelete"), Id("engine"): Id("engine")})})
}

type idGetterWriter struct {
	typeName        func() string
	returns         func() string
	idFieldToReturn func() string
}

func (i idGetterWriter) receiverName() string {
	return "_" + i.typeName()
}

func (i idGetterWriter) receiverParams() *Statement {
	return Id(i.receiverName()).Id(Title(i.typeName()))
}

func (i idGetterWriter) name() string {
	return "ID"
}

func (i idGetterWriter) returnID() *Statement {
	return Id(i.receiverName()).Dot(i.typeName()).Dot(i.idFieldToReturn())
}

type fieldGetterWriter struct {
	t ast.ConfigType
	f ast.Field
}

func (f fieldGetterWriter) receiverName() string {
	return "_" + f.t.Name
}

func (f fieldGetterWriter) receiverParams() *Statement {
	return Id(f.receiverName()).Id(Title(f.t.Name))
}

func (f fieldGetterWriter) name() string {
	return Title(f.f.Name)
}

func (f fieldGetterWriter) returnedType() string {
	switch {
	case f.f.ValueType().IsBasicType:
		return f.f.ValueType().Name
	case f.f.HasPointerValue:
		return f.f.Parent.Name + Title(Singular(f.f.Name)) + "Ref"
	case f.f.HasAnyValue && !f.f.HasPointerValue:
		return anyNameByField(f.f)
	default:
		return f.f.ValueType().Name
	}
}

func (f fieldGetterWriter) returns() string {
	switch {
	case f.f.ValueType().IsBasicType:
		return f.f.ValueString
	case f.f.HasSliceValue:
		return "[]" + Title(f.returnedType())
	default:
		return Title(f.returnedType())
	}
}

func (f fieldGetterWriter) reassignElement() *Statement {
	return Id(f.t.Name).Op(":=").Id(f.receiverName()).Dot(f.t.Name).Dot("engine").Dot(Title(f.t.Name)).Call(Id(f.receiverName()).Dot(f.t.Name).Dot("ID"))
}

func (f fieldGetterWriter) declareSliceOfElements() *Statement {
	return Var().Id(f.f.Name).Id(f.returns())
}

func (f fieldGetterWriter) loopedElementIdentifier() string {
	switch {
	case f.f.ValueType().IsBasicType:
		return BasicTypes[f.f.ValueTypeName] + "ID"
	case f.f.HasPointerValue:
		return "refID"
	case f.f.HasAnyValue && !f.f.HasPointerValue:
		return anyNameByField(f.f) + "ID"
	default:
		return f.f.ValueType().Name + "ID"
	}
}

func (f fieldGetterWriter) loopConditions() *Statement {
	identifier := f.loopedElementIdentifier()
	return List(Id("_"), Id(identifier)).Op(":=").Range().Id(f.t.Name).Dot(f.t.Name).Dot(Title(f.f.Name))
}

func (f fieldGetterWriter) appendedItem() *Statement {
	switch {
	case f.f.ValueType().IsBasicType:
		return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(BasicTypes[f.f.ValueTypeName]).Call(Id(f.loopedElementIdentifier())).Dot("Value")
	case f.f.HasAnyValue || f.f.HasPointerValue:
		return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(f.returnedType()).Call(Id(f.loopedElementIdentifier()))
	default:
		return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(Title(f.returnedType())).Call(Id(f.loopedElementIdentifier()))
	}
}

func (f fieldGetterWriter) appendElement() *Statement {
	return Id(f.f.Name).Op("=").Append(Id(f.f.Name), f.appendedItem())
}

func (f fieldGetterWriter) returnSliceOfType() *Statement {
	return Id(f.f.Name)
}

func (f fieldGetterWriter) element() *Statement {
	return Id(f.f.Parent.Name).Dot(f.f.Parent.Name)
}

func (f fieldGetterWriter) returnBasicType() *Statement {
	return f.element().Dot("engine").Dot(BasicTypes[f.f.ValueTypeName]).Call(f.element().Dot(Title(f.f.Name))).Dot("Value")
}

func (f fieldGetterWriter) returnNamedType() *Statement {
	engine := Id(f.t.Name).Dot(f.t.Name).Dot("engine")
	returnedType := Lower(f.returnedType())
	if !f.f.HasAnyValue && !f.f.HasPointerValue {
		returnedType = Title(returnedType)
	}
	return engine.Dot(returnedType).Call(Id(f.t.Name).Dot(f.t.Name).Dot(Title(f.f.Name)))
}

func (f fieldGetterWriter) returnSingleType() *Statement {
	switch {
	case f.f.ValueType().IsBasicType:
		return f.returnBasicType()
	default:
		return f.returnNamedType()
	}
}

type pathGetterWriter struct {
	t ast.ConfigType
}

func (p pathGetterWriter) receiverName() string {
	return "_" + p.t.Name
}

func (p pathGetterWriter) receiverParams() *Statement {
	return Id(p.receiverName()).Id(Title(p.t.Name))
}

func (p pathGetterWriter) returnPath() *Statement {
	return Id(p.receiverName()).Dot(p.t.Name).Dot("JSONPath")
}

type parentGetter struct {
	t      ast.ConfigType
	parent ast.ConfigType
}

func (p parentGetter) receiverName() string {
	return "_" + p.t.Name
}

func (p parentGetter) element() *Statement {
	return Id(p.t.Name).Dot(p.t.Name)
}

func (p parentGetter) engine() *Statement {
	return p.element().Dot("engine")
}

func (p parentGetter) receiverParams() *Statement {
	return Id(p.receiverName()).Id(Title(p.t.Name))
}

func (p parentGetter) reassignElement() *Statement {
	return Id(p.t.Name).Op(":=").Id(p.receiverName()).Dot(p.t.Name).Dot("engine").Dot(Title(p.t.Name)).Call(Id(p.receiverName()).Dot(p.t.Name).Dot("ID"))
}

func (p parentGetter) hasNoParent() *Statement {
	return Op("!").Add(p.element()).Dot("HasParent")
}

func (p parentGetter) emptyReturn() *Statement {
	return Id(Title(p.parent.Name)).Values(Dict{Id(p.parent.Name): Id(p.parent.Name + "Core").Values(Dict{Id("OperationKind"): Id("OperationKindDelete"), Id("engine"): p.engine()})})
}

func (p parentGetter) declareParentSeg() *Statement {
	return Id("parentSeg").Op(":=").Add(p.element()).Dot("Path").Index(Len(p.element().Dot("Path")).Op("-").Lit(2))
}

func (p parentGetter) getParent() *Statement {
	return p.engine().Dot(Title(p.parent.Name)).Call(Id(Title(p.parent.Name) + "ID").Call(Id("parentSeg").Dot("ID")))
}

type parentKindGetter struct {
	t ast.ConfigType
}

func (p parentKindGetter) receiverName() string {
	return "_" + p.t.Name
}

func (p parentKindGetter) receiverParams() *Statement {
	return Id(p.receiverName()).Id(Title(p.t.Name))
}

func (p parentKindGetter) element() *Statement {
	return Id(p.receiverName()).Dot(p.t.Name)
}

func (p parentKindGetter) path() *Statement {
	return p.element().Dot("Path")
}

func (p parentKindGetter) hasNoParent() *Statement {
	return Op("!").Add(p.element()).Dot("HasParent")
}

func (p parentKindGetter) parentKind() *Statement {
	return p.path().Index(Len(p.path()).Op("-").Lit(2)).Dot("Kind")
}
