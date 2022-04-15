package server

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeController() *Factory {

	s.file.Type().Id("Controller").Interface(
		Id(`OnSuperMessage(msg Message, room *Room, client *Client, lobby *Lobby)
	OnClientConnect(client *Client, lobby *Lobby)
	OnClientDisconnect(room *Room, clientID string, lobby *Lobby)
	OnCreation(lobby *Lobby)
	OnFrameTick(engine *state.Engine)`),
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return &Statement{
				Id(Title(action.Name) + "Broadcast").Add(actionParams(action)).Line(),
				Id(Title(action.Name) + "Emit").Add(actionParams(action)).Add(OnlyIf(action.Response != nil, Id("message").Dot(Title(action.Name)+"Response"))).Line(),
			}
		}),
	)
	return s
}

func actionParams(action ast.Action) *Statement {
	return Params(Id("params").Id("message").Dot(Title(action.Name)+"Params"), Id("engine").Id("*state").Dot("Engine"), Id("roomName"), Id("clientID").String())
}
