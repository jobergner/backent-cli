package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type assembleBranchWriter struct {
	t ast.ConfigType
	f *ast.Field
	v *ast.ConfigType
}

func (a assembleBranchWriter) assembleNextSeg() *Statement {
	if a.f.ValueType().IsBasicType {
		return Empty()
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

func (a assembleBranchWriter) assemblePointerNonSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id(Title(a.f.ValueTypeName)+"ID").Call(Id("nextSeg").Dot("refID"))).Dot(a.f.ValueTypeName),
		If(a.field().Op("!=").Nil().Op("&&").Id("ref").Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
			Break(),
		),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Id("nextSeg").Dot("id")), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Switch(Id("nextSeg").Dot("kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(a.valueTypeID().Call(Id("nextSeg").Dot("id"))).Dot(a.v.Name),
					Id("treeRef").Op(":=").Id("elementReference").Values(
						Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
						Id("ElementID").Op(":").Id("nextSeg").Dot("id"),
						Id("ElementKind").Op(":").Id("ElementKind"+Title(a.v.Name)),
						Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
						Id("ElementPath").Op(":").Id("referencedElement").Dot("Path"),
					),
					a.field().Op("=").Id("&treeRef"),
				)
			}),
		),
	)
}
func (a assembleBranchWriter) assemblePointerSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id(Title(a.f.ValueTypeName)+"ID").Call(Id("nextSeg").Dot("refID"))).Dot(a.f.ValueTypeName),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Int().Call(Id("ref").Dot("ReferencedElementID"))), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.f.ValueType().Name)).Call(Id("ref").Dot("ReferencedElementID")).Dot(a.f.ValueType().Name),
		Id("treeRef").Op(":=").Id("elementReference").Values(
			Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
			Id("ElementID").Op(":").Int().Call(Id("ref").Dot("ReferencedElementID")),
			Id("ElementKind").Op(":").Id("ElementKind"+Title(a.f.ValueType().Name)),
			Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
			Id("ElementPath").Op(":").Id("referencedElement").Dot("Path"),
		),
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(a.fieldTypeID()).Id("elementReference")),
		),
		a.field().Index(Id("referencedElement").Dot("ID")).Op("=").Id("treeRef"),
	)
}
func (a assembleBranchWriter) assemblePointerNonSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id(Title(a.f.ValueTypeName)+"ID").Call(Id("nextSeg").Dot("id"))).Dot(a.f.ValueTypeName),
		If(a.field().Op("!=").Nil().Op("&&").Id("ref").Dot("OperationKind").Op("==").Id("OperationKindDelete")).Block(
			Break(),
		),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Int().Call(Id("ref").Dot("ReferencedElementID"))), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.f.ValueType().Name)).Call(Id("ref").Dot("ReferencedElementID")).Dot(a.f.ValueType().Name),
		Id("treeRef").Op(":=").Id("elementReference").Values(
			Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
			Id("ElementID").Op(":").Int().Call(Id("ref").Dot("ReferencedElementID")),
			Id("ElementKind").Op(":").Id("ElementKind"+Title(a.f.ValueType().Name)),
			Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"),
			Id("ElementPath").Op(":").Id("referencedElement").Dot("Path"),
		),
		a.field().Op("=").Id("&treeRef"),
	)
}

func (a assembleBranchWriter) assemblePointerSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		If(a.field().Op("==").Nil()).Block(
			a.field().Op("=").Make(Map(Int()).Id("elementReference")),
		),
		Id("ref").Op(":=").Id("engine").Dot(a.f.ValueTypeName).Call(Id(Title(a.f.ValueTypeName)+"ID").Call(Id("nextSeg").Dot("refID"))).Dot(a.f.ValueTypeName),
		Id("referencedDataStatus").Op(":=").Id("ReferencedDataUnchanged"),
		If(List(Id("_"), Id("ok")).Op(":=").Id("includedElements").Index(Id("nextSeg").Dot("id")), Id("ok")).Block(
			Id("referencedDataStatus").Op("=").Id("ReferencedDataModified"),
		),
		Switch(Id("nextSeg").Dot("kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					Id("referencedElement").Op(":=").Id("engine").Dot(Title(a.v.Name)).Call(a.valueTypeID().Call(Id("nextSeg").Dot("id"))).Dot(a.v.Name),
					Id("treeRef").Op(":=").Id("elementReference").Values(
						Id("OperationKind").Op(":").Id("ref").Dot("OperationKind"),
						Id("ElementID").Op(":").Id("nextSeg").Dot("id"),
						Id("ElementKind").Op(":").Id("ElementKind"+Title(a.v.Name)),
						Id("ReferencedDataStatus").Op(":").Id("referencedDataStatus"), Id("ElementPath").Op(":").Id("referencedElement").Dot("Path"),
					),
					a.field().Index(Id("nextSeg").Dot("id")).Op("=").Id("treeRef"),
				)
			}),
		),
	)
}

func (a assembleBranchWriter) assembleNonPointerNonSliceAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Switch(Id("nextSeg").Dot("kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Assert(Id("*"+a.v.Name)),
					If(Id("!ok").Op("||").Id("child").Op("==").Nil()).Block(
						Id("child").Op("=").Id("&"+a.v.Name).Values(Dict{Id("ID"): a.valueTypeID().Call(Id("nextSeg").Dot("id"))}),
					),
					Id("engine").Dot("assemble"+Title(a.v.Name)+"Path").Call(Id("child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
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
		List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Index(a.fieldTypeID().Call(Id("nextSeg").Dot("id"))),
		If(Id("!ok")).Block(
			Id("child").Op("=").Id(a.f.ValueType().Name).Values(Dict{Id("ID"): a.fieldTypeID().Call(Id("nextSeg").Dot("id"))}),
		),
		Id("engine").Dot("assemble"+Title(a.f.ValueType().Name)+"Path").Call(Id("&child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
		a.field().Index(Id("child").Dot("ID")).Op("=").Id("child"),
	)
}
func (a assembleBranchWriter) assembleNonPointerNonSliceNonAnyValue() *Statement {
	return Case(Id(FieldPathIdentifier(*a.f))).Block(
		Id("child").Op(":=").Add(a.field()),
		If(Id("child").Op("==").Nil()).Block(
			Id("child").Op("=").Id("&"+a.f.ValueType().Name).Values(Dict{Id("ID"): a.fieldTypeID().Call(Id("nextSeg").Dot("id"))}),
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
		Switch(Id("nextSeg").Dot("kind")).Block(
			ForEachValueOfField(*a.f, func(valueType *ast.ConfigType) *Statement {
				a.v = valueType
				return Case(Id("ElementKind"+Title(a.v.Name))).Block(
					List(Id("child"), Id("ok")).Op(":=").Add(a.field()).Index(Id("nextSeg").Dot("id")).Assert(Id(a.v.Name)),
					If(Id("!ok")).Block(
						Id("child").Op("=").Id(a.v.Name).Values(Dict{Id("ID"): a.valueTypeID().Call(Id("nextSeg").Dot("id"))}),
					),
					Id("engine").Dot("assemble"+Title(a.v.Name)+"Path").Call(Id("&child"), Id("p"), Id("pIndex").Op("+").Lit(1), Id("includedElements")),
					a.field().Index(Id("nextSeg").Dot("id")).Op("=").Id("child"),
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
