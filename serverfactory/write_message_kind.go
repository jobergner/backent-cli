package serverfactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeMessageKinds() *ServerFactory {
	decls := NewDeclSet()

	decls.File.Const().Defs(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id("MessageKindAction_" + action.Name).Id("MessageKind").Op("=").Lit(action.Name)
		}),
	)

	decls.Render(s.buf)
	return s
}
