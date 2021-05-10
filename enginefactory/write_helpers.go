package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeduplicate() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if !field.HasPointerValue {
				return
			}

			d := deduplicateWriter{
				f: field,
			}

			decls.File.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
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
	})

	decls.Render(s.buf)
	return s
}

type deduplicateWriter struct {
	f ast.Field
}

func (d deduplicateWriter) idType() string {
	return title(d.f.Parent.Name) + title(pluralizeClient.Singular(d.f.Name)) + "RefID"
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
	return Id("check").Op(":=").Make(Map(Id(d.idType())).Bool())
}

func (d deduplicateWriter) defineDeduped() *Statement {
	return Id("deduped").Op(":=").Make(Id(d.returns()), Lit(0))

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
