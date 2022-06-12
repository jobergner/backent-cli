package client

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeController() *Factory {

	s.file.Type().Id("Controller").Interface(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return &Statement{
				Id(Title(action.Name)).Add(actionParams(action)).Add(OnlyIf(action.Response != nil, Id("message").Dot(Title(action.Name)+"Response"))).Line(),
			}
		}),
	)
	return s
}

func actionParams(action ast.Action) *Statement {
	return Params(Id("params").Id("message").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state").Dot("Engine"), Id("roomName"), Id("clientID").String())
}
