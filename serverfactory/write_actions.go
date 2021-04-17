package serverfactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeActions() *ServerFactory {
	decls := NewDeclSet()

	_iota := 0
	decls.File.Const().Defs(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			_iota += 1
			return Id("messageKindAction_" + action.Name).Id("messageKind").Op("=").Lit(_iota)
		}),
	)

	decls.File.Type().Id("actions").Struct(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id(action.Name).Func().Params(Id(title(action.Name)+"Params"), Id("*Engine"))
		}),
	)

	decls.Render(s.buf)
	return s
}
