package server

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeProcessClientMessage() *Factory {

	p := processClientMessageWriter{}

	s.file.Func().Params(p.receiverParams()).Id("processClientMessage").Params(p.params()).Id("Message").Block(
		Switch(Id("msg").Dot("Kind")).Block(
			ForEachActionInAST(s.config, func(action ast.Action) *Statement {
				p.a = &action
				return Case(Id("message").Dot(p.actionMessageKind())).Block(
					p.declareParams(),
					p.unmarshalMessageContent(),
					If(Id("err").Op("!=").Nil()).Block(
						p.logWithMessageKind().Dot("Str").Call(Id("logging").Dot("MessageContent"), String().Call(Id("msg").Dot("Content"))).Dot("Msg").Call(Lit("failed unmarshalling params")),
						p.returnErrorMessage(),
					),
					Id("r").Dot("state").Dot("BroadcastingClientID").Op("=").Id("msg").Dot("client").Dot("id"),
					Id("r").Dot("controller").Dot(Title(action.Name)+"Broadcast").Call(Id("params"), Id("r").Dot("state"), Id("r").Dot("name"), Id("msg").Dot("client").Dot("id")),
					Id("r").Dot("state").Dot("BroadcastingClientID").Op("=").Lit(""),
					OnlyIf(action.Response != nil, Id("res").Op(":=")).Id("r").Dot("controller").Dot(Title(action.Name)+"Emit").Call(Id("params"), Id("r").Dot("state"), Id("r").Dot("name"), Id("msg").Dot("client").Dot("id")),
					OnlyIf(action.Response != nil, &Statement{
						List(Id("resContent"), Id("err")).Op(":=").Id("res").Dot("MarshalJSON").Call().Line(),
						If(Id("err").Op("!=").Nil()).Block(
							p.logWithMessageKind().Dot("Msg").Call(Lit("failed marshalling response content")),
							p.returnErrorMessage(),
						),
					}),
					OnlyIf(action.Response != nil, Return(Id("Message").Values(Id("msg").Dot("ID"), Id("msg").Dot("Kind"), Id("resContent"), Id("msg").Dot("client")))),
					OnlyIf(action.Response == nil, Return(Id("Message").Values(Dict{
						Id("ID"):   Id("msg").Dot("ID"),
						Id("Kind"): Id("message").Dot("MessageKindNoResponse"),
					}))),
				)
			}),
			Default().Block(
				Id("err").Op(":=").Id("ErrMessageKindUnknown"),
				p.logWithMessageKind().Dot("Msg").Call(Lit("unknown message kind")),
				p.returnErrorMessage(),
			),
		),
	)

	return s
}

type processClientMessageWriter struct {
	a *ast.Action
	p *ast.Field
}

func (p processClientMessageWriter) receiverParams() *Statement {
	return Id("r").Id("*Room")
}

func (p processClientMessageWriter) params() *Statement {
	return Id("msg").Id("Message")
}

func (p processClientMessageWriter) logWithMessageKind() *Statement {
	return Id("log").Dot("Err").Call(Id("err")).Dot("Str").Call(Id("logging").Dot("MessageKind"), String().Call(Id("msg").Dot("Kind")))
}

func (p processClientMessageWriter) returnErrorMessage() *Statement {
	return Return(Id("Message").Values(Id("msg").Dot("ID"), Id("message").Dot("MessageKindError"), Index().Byte().Call(Lit("invalid message")), Id("msg").Dot("client")))
}

func (p processClientMessageWriter) actionMessageKind() string {
	return "MessageKindAction_" + p.a.Name
}

func (p processClientMessageWriter) declareParams() *Statement {
	return Var().Id("params").Id("message").Dot(Title(p.a.Name) + "Params")
}

func (p processClientMessageWriter) unmarshalMessageContent() *Statement {
	return Id("err").Op(":=").Id("params").Dot("UnmarshalJSON").Call(Id("msg").Dot("Content"))
}

func (p processClientMessageWriter) actionCall() *Statement {
	return Call(Id("params"), Id("r").Dot("state"), Id("r").Dot("name"), Id("msg").Dot("client").Dot("id"))
}
