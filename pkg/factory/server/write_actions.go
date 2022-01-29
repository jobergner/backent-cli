package server

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeActions() *ServerFactory {

	s.file.Comment("easyjson:skip")
	s.file.Type().Id("Actions").Struct(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			responseName := Id(Title(action.Name) + "Response")
			if action.Response == nil {
				responseName = Empty()
			}
			return Id(Title(action.Name)).Func().Params(Id(Title(action.Name)+"Params"), Id("*Engine")).Add(responseName)
		}),
	)

	return s
}
