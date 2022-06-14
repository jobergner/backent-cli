package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeDeduplicate() *Factory {

	s.config.RangeTypes(func(configType ast.ConfigType) {

		d := deduplicateWriter{
			typeName: func() string {
				return configType.Name
			},
			idType: func() string {
				return Title(configType.Name) + "ID"
			},
		}

		writeDeduplicate(s.file, d)
	})

	s.config.RangeRefFields(func(field ast.Field) {

		d := deduplicateWriter{
			typeName: func() string {
				return ValueTypeName(&field)
			},
			idType: func() string {
				return Title(ValueTypeName(&field)) + "ID"
			},
		}

		writeDeduplicate(s.file, d)
	})

	return s
}

func writeDeduplicate(file *File, d deduplicateWriter) {
	file.Func().Id(d.name()).Params(d.params()).Id(d.returns()).Block(
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

func (s *Factory) writeAllIDsMethod() *Factory {

	s.config.RangeTypes(func(configType ast.ConfigType) {
		a := allIDsMehtodWriter{
			typeName: func() string {
				return configType.Name
			},
		}

		writeAllIDsMethod(s.file, a)
	})

	s.config.RangeRefFields(func(field ast.Field) {
		a := allIDsMehtodWriter{
			typeName: func() string {
				return ValueTypeName(&field)
			},
		}

		writeAllIDsMethod(s.file, a)
	})

	return s
}

func writeAllIDsMethod(file *File, a allIDsMehtodWriter) {
	file.Func().Params(a.receiverParams()).Id(a.name()).Params().Id(a.returns()).Block(
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
