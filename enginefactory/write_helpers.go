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

func (s *EngineFactory) writeAllIDsMethod() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if !field.HasPointerValue {
				return
			}

			d := allIDsMehtodWriter{
				f: field,
			}

			decls.File.Func().Params(d.receiverParams()).Id(d.name()).Params().Id(d.returns()).Block(
				d.declareStateIDsSlice(),
				For(d.stateIDsLoopConditions()).Block(
					d.appendStateID(),
				),
				d.declarePatchIDsSlice(),
				For(d.patchIDsLoopConditions()).Block(
					d.appendPatchID(),
				),
				Return(d.deduplicatedIDs()),
			)
		})
	})

	decls.Render(s.buf)
	return s
}

type allIDsMehtodWriter struct {
	f ast.Field
}

func (d allIDsMehtodWriter) typeName() string {
	return title(d.f.Parent.Name) + title(pluralizeClient.Singular(d.f.Name)) + "Ref"
}

func (d allIDsMehtodWriter) idType() string {
	return d.typeName() + "ID"
}

func (d allIDsMehtodWriter) name() string {
	return "all" + d.idType() + "s"
}

func (d allIDsMehtodWriter) receiverParams() *Statement {
	return Id("engine").Id("Engine")
}

func (d allIDsMehtodWriter) returns() string {
	return "[]" + d.idType()
}

func (d allIDsMehtodWriter) idSliceName(prefix string) string {
	return prefix + d.idType() + "s"
}

func (d allIDsMehtodWriter) declareStateIDsSlice() *Statement {
	return Var().Id(d.idSliceName("state")).Id("[]" + d.idType())
}

func (d allIDsMehtodWriter) stateIDsLoopConditions() *Statement {
	return Id(lower(d.idType())).Op(":=").Range().Id("engine").Dot("State").Dot(d.typeName())
}

func (d allIDsMehtodWriter) appendStateID() *Statement {
	return Id(d.idSliceName("state")).Op("=").Append(Id(d.idSliceName("state")), Id(lower(d.idType())))
}

func (d allIDsMehtodWriter) declarePatchIDsSlice() *Statement {
	return Var().Id(d.idSliceName("patch")).Id("[]" + d.idType())
}

func (d allIDsMehtodWriter) patchIDsLoopConditions() *Statement {
	return Id(lower(d.idType())).Op(":=").Range().Id("engine").Dot("Patch").Dot(d.typeName())
}

func (d allIDsMehtodWriter) appendPatchID() *Statement {
	return Id(d.idSliceName("patch")).Op("=").Append(Id(d.idSliceName("patch")), Id(lower(d.idType())))
}

func (d allIDsMehtodWriter) deduplicatedIDs() *Statement {
	dedu := deduplicateWriter{d.f}
	return Id(dedu.name()).Call(Id(d.idSliceName("state")), Id(d.idSliceName("patch")))
}
