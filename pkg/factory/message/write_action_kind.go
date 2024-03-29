package message

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeMessageKinds() *Factory {

	s.file.Const().Defs(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id("MessageKindAction_" + action.Name).Id("Kind").Op("=").Lit(action.Name)
		}),
	)

	return s
}
