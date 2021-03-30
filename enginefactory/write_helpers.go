package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeduplicate() *EngineFactory {
	decls := newDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		d := deduplicator{
			t: configType,
		}

		decls.file.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
			d.defineCheck(),
			d.defineDeduped(),
			For(d.loopConditions("a")).Block(
				d.checkValue(),
			),
			For(d.loopConditions("b")).Block(
				d.checkValue(),
			),
			d.loopCheck(),
			Return(Id("deduped")),
		)
	})

	decls.render(s.buf)
	return s
}

type deduplicator struct {
	t ast.ConfigType
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

func (d deduplicator) loopConditions(getsLooped string) *Statement {
	return List(Id("_"), Id("val")).Op(":=").Range().Id(getsLooped)
}

func (d deduplicator) checkValue() *Statement {
	return Id("check").Index(Id("val")).Op("=").Id("true")
}

func (d deduplicator) loopCheck() *Statement {
	loop := For(Id("val").Op(":=").Range().Id("check")).Block(
		Id("deduped").Op("=").Append(Id("deduped"), Id("val")),
	)
	return loop
}
