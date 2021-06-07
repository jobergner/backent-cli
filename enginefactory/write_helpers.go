package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeDeduplicate() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {
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

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAllIDsMethod() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeRefFields(func(field ast.Field) {
		a := allIDsMehtodWriter{
			f: field,
		}

		decls.File.Func().Params(a.receiverParams()).Id(a.name()).Params().Id(a.returns()).Block(
			a.declareStateIDsSlice(),
			For(a.stateIDsLoopConditions()).Block(
				a.appendStateID(),
			),
			a.declarePatchIDsSlice(),
			For(a.patchIDsLoopConditions()).Block(
				a.appendPatchID(),
			),
			Return(a.deduplicatedIDs()),
		)
	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeMergeIDs() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		m := mergeIDsWriter{
			idType: func() string {
				return Title(configType.Name) + "ID"
			},
		}

		writeMergeIDs(&decls, m)
	})

	s.config.RangeRefFields(func(field ast.Field) {
		m := mergeIDsWriter{
			idType: func() string {
				return Title(field.ValueTypeName) + "ID"
			},
		}

		writeMergeIDs(&decls, m)
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		m := mergeIDsWriter{
			idType: func() string {
				return Title(anyNameByField(field)) + "ID"
			},
		}

		writeMergeIDs(&decls, m)
	})

	decls.Render(s.buf)
	return s
}

func writeMergeIDs(decls *DeclSet, m mergeIDsWriter) {

	decls.File.Func().Id(m.name()).Params(m.params()).Id(m.returns()).Block(
		m.declareIDs(),
		m.copyIDs(),
		m.declareCounter(),
		For(m.currentIDsLoopConditions()).Block(
			If(m.idDoesNotMatch()).Block(
				Continue(),
			),
			m.incrementCounter(),
		),
		For(m.nextIDsLoopConditions()).Block(
			m.appendID(),
		),
		Return(Id("ids")),
	)
}
