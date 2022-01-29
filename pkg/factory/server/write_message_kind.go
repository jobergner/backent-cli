package server

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeMessageKinds() *ServerFactory {

	s.file.Const().Defs(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id("MessageKindAction_" + action.Name).Id("MessageKind").Op("=").Lit(action.Name)
		}),
	)

	return s
}
