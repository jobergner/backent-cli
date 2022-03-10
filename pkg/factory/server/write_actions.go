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
			return Id(Title(action.Name)).Id(Title(action.Name) + "Action")
		}),
	)

	s.config.RangeActions(func(action ast.Action) {
		responseName := Id(Title(action.Name) + "Response")
		if action.Response == nil {
			responseName = Empty()
		}

		params := Params(Id("params").Id(Title(action.Name)+"Params"), Id("engine").Id("*Engine"), Id("roomName"), Id("clientID").String())

		s.file.Type().Id(Title(action.Name)+"Action").Struct(
			Id("Broadcast").Func().Add(params).Line(),
			Id("Emit").Func().Add(params).Add(responseName).Line(),
		)
	})

	return s
}
