package enginefactory

import (
	. "github.com/dave/jennifer/jen"
)

func (s *stateFactory) writeDeduplicate() *stateFactory {
	decls := newDeclSet()
	s.ast.rangeTypes(func(configType stateConfigType) {

		d := deduplicator{
			t: configType,
		}

		decls.file.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
			d.defineCheck(),
			d.defineDeduped(),
			d.loopIDs("a"),
			d.loopIDs("b"),
			d.loopCheck(),
			Return(Id("deduped")),
		)
	})

	decls.render(s.buf)
	return s
}

type deduplicator struct {
	t stateConfigType
}

func (d deduplicator) name() string {
	return "deduplicate" + title(d.t.Name) + "IDs"
}

func (d deduplicator) idType() string {
	return title(d.t.Name) + "ID"
}

func (d deduplicator) params() *Statement {
	return List(Id("a").Id("[]"+d.idType()), Id("b").Id("[]"+d.idType()))
}

func (d deduplicator) returns() string {
	return "[]" + d.idType()
}

func (d deduplicator) defineCheck() *Statement {
	return Id("check").Op(":=").Make(Map(Id(d.idType())).Bool())
}

func (d deduplicator) defineDeduped() *Statement {
	return Id("deduped").Op(":=").Make(Id(d.returns()), Lit(0))

}

func (d deduplicator) loopIDs(getsLooped string) *Statement {
	loop := For(List(Id("_"), Id("val")).Op(":=").Range().Id(getsLooped)).Block(
		Id("check").Index(Id("val")).Op("=").Id("true"),
	)
	return loop
}

func (d deduplicator) loopCheck() *Statement {
	loop := For(Id("val").Op(":=").Range().Id("check")).Block(
		Id("deduped").Op("=").Append(Id("deduped"), Id("val")),
	)
	return loop
}
