package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeduplicate() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {
		d := deduplicateWriter{
			typeName: func() string {
				return configType.Name
			},
			idType: func() string {
				return Title(configType.Name) + "ID"
			},
		}

		writeDeduplicate(&decls, d)
	})
	s.config.RangeRefFields(func(field ast.Field) {
		d := deduplicateWriter{
			typeName: func() string {
				return field.ValueTypeName
			},
			idType: func() string {
				return Title(field.ValueTypeName) + "ID"
			},
		}

		writeDeduplicate(&decls, d)
	})

	decls.Render(s.buf)
	return s
}

func writeDeduplicate(decls *DeclSet, d deduplicateWriter) {
	decls.File.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
		d.defineCheck(),
		For(d.clearCheckLoopConditions()).Block(
			d.clearCheckValue(),
		),
		d.defineDeduped(),
		For(d.loopConditions("a")).Block(
			d.checkValue(),
		),
		For(d.loopConditions("b")).Block(
			d.checkValue(),
		),
		d.loopCheck(),
		d.returnCheckToPool(),
		Return(Id("deduped")),
	)
}

func (s *EngineFactory) writeAllIDsMethod() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {
		a := allIDsMehtodWriter{
			typeName: func() string {
				return configType.Name
			},
		}

		writeAllIDsMethod(&decls, a)
	})

	s.config.RangeRefFields(func(field ast.Field) {
		a := allIDsMehtodWriter{
			typeName: func() string {
				return field.ValueTypeName
			},
		}

		writeAllIDsMethod(&decls, a)
	})

	decls.Render(s.buf)
	return s
}

func writeAllIDsMethod(decls *DeclSet, a allIDsMehtodWriter) {
	decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params().Id(a.returns()).Block(
		a.declareStateIDsSlice(),
		For(a.stateIDsLoopConditions()).Block(
			a.appendStateID(),
		),
		a.declarePatchIDsSlice(),
		For(a.patchIDsLoopConditions()).Block(
			a.appendPatchID(),
		),
		a.declareDedupedIDs(),
		a.returnIdSliceToPool("state"),
		a.returnIdSliceToPool("patch"),
		Return(Id("dedupedIDs")),
	)
}
