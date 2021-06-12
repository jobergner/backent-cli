package serverfactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeMessageKinds() *ServerFactory {
	decls := NewDeclSet()

	_iota := 0
	decls.File.Const().Defs(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			_iota += 1
			return Id("messageKindAction_" + action.Name).Id("messageKind").Op("=").Lit(_iota)
		}),
	)

	decls.Render(s.buf)
	return s
}
