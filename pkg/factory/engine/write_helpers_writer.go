package engine

import (
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type deduplicateWriter struct {
	typeName func() string
	idType   func() string
}

func (d deduplicateWriter) name() string {
	return "deduplicate" + d.idType() + "s"
}

func (d deduplicateWriter) params() *Statement {
	return List(Id("a").Id("[]"+d.idType()), Id("b").Id("[]"+d.idType()))
}

func (d deduplicateWriter) returns() string {
	return "[]" + d.idType()
}

func (d deduplicateWriter) defineCheck() *Statement {
	return Id("check").Op(":=").Id(d.typeName() + "CheckPool").Dot("Get").Call().Assert(Map(Id(d.idType())).Bool())
}

func (d deduplicateWriter) clearCheckLoopConditions() *Statement {
	return Id("k").Op(":=").Range().Id("check")
}

func (d deduplicateWriter) clearCheckValue() *Statement {
	return Delete(Id("check"), Id("k"))
}

func (d deduplicateWriter) defineDeduped() *Statement {
	return Id("deduped").Op(":=").Id(d.typeName() + "IDSlicePool").Dot("Get").Call().Assert(Id(d.returns())).Index(Op(":").Lit(0))

}

func (d deduplicateWriter) loopConditions(getsLooped string) *Statement {
	return List(Id("_"), Id("val")).Op(":=").Range().Id(getsLooped)
}

func (d deduplicateWriter) checkValue() *Statement {
	return Id("check").Index(Id("val")).Op("=").Id("true")
}

func (d deduplicateWriter) loopCheck() *Statement {
	loop := For(Id("val").Op(":=").Range().Id("check")).Block(
		Id("deduped").Op("=").Append(Id("deduped"), Id("val")),
	)
	return loop
}

func (d deduplicateWriter) returnCheckToPool() *Statement {
	return Id(d.typeName() + "CheckPool").Dot("Put").Call(Id("check"))
}

type allIDsMehtodWriter struct {
	typeName func() string
}

func (a allIDsMehtodWriter) idType() string {
	return a.typeName() + "ID"
}

func (a allIDsMehtodWriter) name() string {
	return "all" + Title(a.idType()) + "s"
}

func (a allIDsMehtodWriter) receiverParams() *Statement {
	return Id("engine").Id("Engine")
}

func (a allIDsMehtodWriter) returns() string {
	return "[]" + Title(a.idType())
}

func (a allIDsMehtodWriter) idSliceName(prefix string) string {
	return prefix + Title(a.idType()) + "s"
}

func (a allIDsMehtodWriter) declareStateIDsSlice() *Statement {
	return Id(a.idSliceName("state")).Op(":=").Id(a.idType() + "SlicePool").Dot("Get").Call().Assert(Index().Id(Title(a.idType()))).Index(Op(":").Lit(0))
}

func (a allIDsMehtodWriter) stateIDsLoopConditions() *Statement {
	return Id(Lower(a.idType())).Op(":=").Range().Id("engine").Dot("State").Dot(Title(a.typeName()))
}

func (a allIDsMehtodWriter) appendStateID() *Statement {
	return Id(a.idSliceName("state")).Op("=").Append(Id(a.idSliceName("state")), Id(Lower(a.idType())))
}

func (a allIDsMehtodWriter) declarePatchIDsSlice() *Statement {
	return Id(a.idSliceName("patch")).Op(":=").Id(a.idType() + "SlicePool").Dot("Get").Call().Assert(Index().Id(Title(a.idType()))).Index(Op(":").Lit(0))
}

func (a allIDsMehtodWriter) patchIDsLoopConditions() *Statement {
	return Id(Lower(a.idType())).Op(":=").Range().Id("engine").Dot("Patch").Dot(Title(a.typeName()))
}

func (a allIDsMehtodWriter) appendPatchID() *Statement {
	return Id(a.idSliceName("patch")).Op("=").Append(Id(a.idSliceName("patch")), Id(Lower(a.idType())))
}

func (a allIDsMehtodWriter) declareDedupedIDs() *Statement {
	return Id("dedupedIDs").Op(":=").Id("deduplicate"+Title(a.typeName())+"IDs").Call(Id(a.idSliceName("state")), Id(a.idSliceName("patch")))
}

func (a allIDsMehtodWriter) returnIdSliceToPool(prefix string) *Statement {
	return Id(a.idType() + "SlicePool").Dot("Put").Call(Id(prefix + Title(a.idType()) + "s"))
}
