package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type assembleBranchWriter struct {
	t ast.ConfigType
	f *ast.Field
	v *ast.ConfigType
}

func (a assembleBranchWriter) assembleNextSeg() *Statement {
	if a.f.ValueType().IsBasicType {
		switch {
		case !a.f.HasSliceValue:
			return a.assembleBasicNonSliceValue()
		default:
			return a.assembleBasicSliceValue()
		}
	}

	if a.f.HasPointerValue {
		switch {
		case !a.f.HasSliceValue && a.f.HasAnyValue:
			return a.assemblePointerNonSliceAnyValue()
		case a.f.HasSliceValue && !a.f.HasAnyValue:
			return a.assemblePointerSliceNonAnyValue()
		case !a.f.HasSliceValue && !a.f.HasAnyValue:
			return a.assemblePointerNonSliceNonAnyValue()
		default: // a.f.HasSliceValue && a.f.HasAnyValue:
			return a.assemblePointerSliceAnyValue()
		}
	}

	switch {
	case !a.f.HasSliceValue && a.f.HasAnyValue:
		return a.assembleNonPointerNonSliceAnyValue()
	case a.f.HasSliceValue && !a.f.HasAnyValue:
		return a.assembleNonPointerSliceNonAnyValue()
	case !a.f.HasSliceValue && !a.f.HasAnyValue:
		return a.assembleNonPointerNonSliceNonAnyValue()
	default: // a.f.HasSliceValue && a.f.HasAnyValue:
		return a.assembleNonPointerSliceAnyValue()
	}
}

func (a assembleBranchWriter) assembleBasicSliceValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		If(Id("element").Dot(Title(a.f.Name)).Op("==").Nil()).Block(
			If(List(Id("state"+Title(a.t.Name)), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(Title(a.t.Name)).Index(Id("element").Dot("ID")), Id("ok")).Block(
				Id("element").Dot(Title(a.f.Name)).Op("=").Make(Index().Id(a.f.ValueType().Name), Lit(0), Len(Id("state"+Title(a.t.Name)).Dot(Title(a.f.Name)))),
				For(List(Id("_"), Id("valID")).Op(":=").Range().Id("state"+Title(a.t.Name)).Dot(Title(a.f.Name))).Block(
					Id("val").Op(":=").Id("engine").Dot(a.f.ValueType().Name+"Value").Call(Id("valID")),
					Id("element").Dot(Title(a.f.Name)).Op("=").Append(Id("element").Dot(Title(a.f.Name)), Id("val").Dot("Value")),
				),
			).Else().Block(
				Id("element").Dot(Title(a.f.Name)).Op("=").Make(Index().Id(a.f.ValueType().Name), Lit(0), Len(Id(a.t.Name+"Data").Dot(Title(a.f.Name)))),
			),
		),
		Id("child").Op(":=").Id("engine").Dot(a.f.ValueType().Name+"Value").Call(Id(Title(a.f.ValueType().Name)+"ValueID").Call(Id("nextSeg").Dot("ID"))),
		Switch(Id("child").Dot("OperationKind")).Block(
			Case(Id("OperationKindUnchanged")).Block(
				Break(),
			),
			Case(Id("OperationKindDelete")).Block(
				Var().Id("newValues").Index().Id(a.f.ValueType().Name),
				For(List(Id("_"), Id("val")).Op(":=").Range().Id("element").Dot(Title(a.f.Name))).Block(
					If(Id("val").Op("!=").Id("child").Dot("Value")).Block(
						Id("newValues").Op("=").Append(Id("newValues"), Id("val")),
					),
				),
				Id("element").Dot(Title(a.f.Name)).Op("=").Id("newValues"),
			),
			Case(Id("OperationKindUpdate")).Block(
				Id("element").Dot(Title(a.f.Name)).Op("=").Append(Id("element").Dot(Title(a.f.Name)), Id("child").Dot("Value")),
			),
		),
	)
}

func (a assembleBranchWriter) assembleBasicNonSliceValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("child").Op(":=").Id("engine").Dot(BasicTypes[ValueTypeName(a.f)]).Call(Id(a.t.Name+"Data").Dot(Title(a.f.Name))),
		Id("element").Dot("OperationKind").Op("=").Id("child").Dot("OperationKind"),
		Id("element").Dot(Title(a.f.Name)).Op("=").Id("&child").Dot("Value"),
	)
}

func (a assembleBranchWriter) assemblePointerNonSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(ValueTypeName(a.f)).Call(Id(Title(ValueTypeName(a.f))+"ID").Call(Id("nextSeg").Dot("RefID"))).Dot(ValueTypeName(a.f)),
		If(a.field().Op("!=").Nil().Op("&&").Id("ref").Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
			Break(),
		),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Id("ref").Dot("ChildID")), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Switch(Id("nextSeg").Dot("Kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(a.valueTypeID().Call(Id("ref").Dot("ChildID"))).Dot(a.v.Name),
					Id("treeRef").Op(":=").Id("elementReference").Values(
						Id("ID").Op(":").Int().Call(Id("ref").Dot("ID")),
						Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
						Id("ElementID").Op(":").Id("ref").Dot("ChildID"),
						Id("ElementKind").Op(":").Id("ElementKind"+Title(a.v.Name)),
						Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
						Id("ElementPath").Op(":").Id("referencedElement").Dot("JSONPath"),
					),
					a.field().Op("=").Id("&treeRef"),
				)
			}),
		),
	)
}
func (a assembleBranchWriter) assemblePointerSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(ValueTypeName(a.f)).Call(Id(Title(ValueTypeName(a.f))+"ID").Call(Id("nextSeg").Dot("RefID"))).Dot(ValueTypeName(a.f)),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Int().Call(Id("ref").Dot("ReferencedElementID"))), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.f.ValueType().Name)).Call(Id("ref").Dot("ReferencedElementID")).Dot(a.f.ValueType().Name),
		Id("treeRef").Op(":=").Id("elementReference").Values(
			Id("ID").Op(":").Int().Call(Id("ref").Dot("ID")),
			Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
			Id("ElementID").Op(":").Int().Call(Id("ref").Dot("ReferencedElementID")),
			Id("ElementKind").Op(":").Id("ElementKind"+Title(a.f.ValueType().Name)),
			Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
			Id("ElementPath").Op(":").Id("referencedElement").Dot("JSONPath"),
		),
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(a.fieldTypeID()).Id("elementReference")),
		),
		a.field().Index(Id("referencedElement").Dot("ID")).Op("=").Id("treeRef"),
	)
}
func (a assembleBranchWriter) assemblePointerNonSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(ValueTypeName(a.f)).Call(Id(Title(ValueTypeName(a.f))+"ID").Call(Id("nextSeg").Dot("RefID"))).Dot(ValueTypeName(a.f)),
		If(a.field().Op("!=").Nil().Op("&&").Id("ref").Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
			Break(),
		),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Int().Call(Id("ref").Dot("ReferencedElementID"))), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.f.ValueType().Name)).Call(Id("ref").Dot("ReferencedElementID")).Dot(a.f.ValueType().Name),
		Id("treeRef").Op(":=").Id("elementReference").Values(
			Id("ID").Op(":").Int().Call(Id("ref").Dot("ID")),
			Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
			Id("ElementID").Op(":").Int().Call(Id("ref").Dot("ReferencedElementID")),
			Id("ElementKind").Op(":").Id("ElementKind"+Title(a.f.ValueType().Name)),
			Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
			Id("ElementPath").Op(":").Id("referencedElement").Dot("JSONPath"),
		),
		a.field().Op("=").Id("&treeRef"),
	)
}

func (a assembleBranchWriter) assemblePointerSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(Int()).Id("elementReference")),
		),
		Id("ref").Op(":=").Id("engine").Dot(ValueTypeName(a.f)).Call(Id(Title(ValueTypeName(a.f))+"ID").Call(Id("nextSeg").Dot("RefID"))).Dot(ValueTypeName(a.f)),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Id("ref").Dot("ChildID")), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Switch(Id("nextSeg").Dot("Kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(a.valueTypeID().Call(Id("ref").Dot("ChildID"))).Dot(a.v.Name),
					Id("treeRef").Op(":=").Id("elementReference").Values(
						Id("ID").Op(":").Int().Call(Id("ref").Dot("ID")),
						Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
						Id("ElementID").Op(":").Id("ref").Dot("ChildID"),
						Id("ElementKind").Op(":").Id("ElementKind"+Title(a.v.Name)),
						Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"), Id("ElementPath").Op(":").Id("referencedElement").Dot("JSONPath"),
					),
					a.field().Index(Id("ref").Dot("ChildID")).Op("=").Id("treeRef"),
				)
			}),
		),
	)
}

func (a assembleBranchWriter) assembleNonPointerNonSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Switch(Id("nextSeg").Dot("Kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Assert(Id("*"+a.v.Name)),
					If(Id("!ok").Op("||").Id("child").Op("==").Nil()).Block(
						Id("child").Op("=").Id("&"+a.v.Name).Values(Dict{Id("ID"): a.valueTypeID().Call(Id("nextSeg").Dot("ID"))}),
					),
					Id("engine").Dot("assemble"+Title(a.v.Name)+"Path").Call(Id("child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
					If(Id("child").Dot("OperationKind").Op("==").Id("OperationKindDelete").Op("&&").Id("element").Dot(Title(a.f.Name)).Op("!=").Nil()).Block(
						Break(),
					),
					a.field().Op("=").Id("child"),
				)
			}),
		),
	)
}
func (a assembleBranchWriter) assembleNonPointerSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(a.fieldTypeID()).Id(a.f.ValueType().Name)),
		),
		List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Index(a.fieldTypeID().Call(Id("nextSeg").Dot("ID"))),
		If(Id("!ok")).Block(
			Id("child").Op("=").Id(a.f.ValueType().Name).Values(Dict{Id("ID"): a.fieldTypeID().Call(Id("nextSeg").Dot("ID"))}),
		),
		Id("engine").Dot("assemble"+Title(a.f.ValueType().Name)+"Path").Call(Id("&child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
		a.field().Index(Id("child").Dot("ID")).Op("=").Id("child"),
	)
}
func (a assembleBranchWriter) assembleNonPointerNonSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("child").Op(":=").Add(a.field()),
		If(Id("child").Op("==").Nil()).Block(
			Id("child").Op("=").Id("&"+a.f.ValueType().Name).Values(Dict{Id("ID"): a.fieldTypeID().Call(Id("nextSeg").Dot("ID"))}),
		),
		Id("engine").Dot("assemble"+Title(a.f.ValueType().Name)+"Path").Call(Id("child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
		a.field().Op("=").Id("child"),
	)
}
func (a assembleBranchWriter) assembleNonPointerSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(Int()).Interface()),
		),
		Switch(Id("nextSeg").Dot("Kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Index(Id("nextSeg").Dot("ID")).Assert(Id(a.v.Name)),
					If(Id("!ok")).Block(
						Id("child").Op("=").Id(a.v.Name).Values(Dict{Id("ID"): a.valueTypeID().Call(Id("nextSeg").Dot("ID"))}),
					),
					Id("engine").Dot("assemble"+Title(a.v.Name)+"Path").Call(Id("&child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
					a.field().Index(Id("nextSeg").Dot("ID")).Op("=").Id("child"),
				)
			}),
		),
	)
}

func (a assembleBranchWriter) field() *Statement {
	return Id("element").Dot(Title(a.f.Name))
}

func (a assembleBranchWriter) fieldTypeID() *Statement {
	return Id(Title(a.f.ValueType().Name) + "ID")
}

func (a assembleBranchWriter) valueTypeID() *Statement {
	return Id(Title(a.v.Name) + "ID")
}
