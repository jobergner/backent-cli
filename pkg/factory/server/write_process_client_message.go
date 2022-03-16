package server

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeProcessClientMessage() *Factory {

	p := processClientMessageWriter{}

	s.file.Func().Params(p.receiverParams()).Id("processClientMessage").Params(p.params()).Params(Id("Message"), Id("error")).Block(
		Switch().Id("MessageKind").Call(Id("msg").Dot("Kind")).Block(
			ForEachActionInAST(s.config, func(action ast.Action) *Statement {
				p.a = &action
				return Case(Id(p.actionMessageKind())).Block(
					p.declareParams(),
					p.unmarshalMessageContent(),
					If(Id("err").Op("!=").Nil()).Block(
						Return(p.returnErrorMessage()),
					),
					If(p.actionBroadcastIsDefined()).Block(
						p.callActionBroadcast(),
					),
					If(p.actionEmitIsUndefined()).Block(
						Break(),
					),
					p.callActionEmit(),
					OnlyIf(action.Response != nil, p.marshalResponseContent()),
					OnlyIf(action.Response != nil, p.returnMarshallingError()),
					p.returnResponse(),
				)
			}),
			Default().Block(
				Return(p.unknownMessageKindResponse(), Id("fmt").Dot("Errorf").Call(Lit("unknown message kind in: %s"), Id("printMessage").Call(Id("msg")))),
			),
		),
		Return(Id("Message").Values(Id("ID").Op(":").Id("msg").Dot("ID")), Nil()),
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

func (p processClientMessageWriter) actionMessageKind() string {
	return "MessageKindAction_" + p.a.Name
}

func (p processClientMessageWriter) callActionBroadcast() *Statement {
	return Id("r").Dot("actions").Dot(Title(p.a.Name)).Dot("Broadcast").Add(p.actionCall())
}

func (p processClientMessageWriter) actionBroadcastIsDefined() *Statement {
	return Id("r").Dot("actions").Dot(Title(p.a.Name)).Dot("Broadcast").Op("!=").Nil()
}

func (p processClientMessageWriter) actionEmitIsUndefined() *Statement {
	return Id("r").Dot("actions").Dot(Title(p.a.Name)).Dot("Emit").Op("==").Nil()
}

func (p processClientMessageWriter) declareParams() *Statement {
	return Var().Id("params").Id(Title(p.a.Name) + "Params")
}

func (p processClientMessageWriter) unmarshalMessageContent() *Statement {
	return Id("err").Op(":=").Id("params").Dot("UnmarshalJSON").Call(Id("msg").Dot("Content"))
}

func (p processClientMessageWriter) returnErrorMessage() (*Statement, *Statement) {
	return Id("Message").Values(List(Id("msg").Dot("ID"), Id("MessageKindError"), Id("messageUnmarshallingError").Call(Id("msg").Dot("Content"), Id("err")), Id("msg").Dot("client"))), Id("err")
}

func (p processClientMessageWriter) actionCall() *Statement {
	return Call(Id("params"), Id("r").Dot("state"), Id("r").Dot("name"), Id("msg").Dot("client").Dot("id"))
}

func (p processClientMessageWriter) callActionEmit() *Statement {
	call := Id("r").Dot("actions").Dot(Title(p.a.Name)).Dot("Emit").Add(p.actionCall())
	if p.a.Response != nil {
		return Id("res").Op(":=").Add(call)
	}
	return call
}

func (p processClientMessageWriter) marshalResponseContent() *Statement {
	return List(Id("resContent"), Id("err")).Op(":=").Id("res").Dot("MarshalJSON").Call()
}

func (p processClientMessageWriter) returnMarshallingError() *Statement {
	return If(Id("err").Op("!=").Nil()).Block(
		Return(Id("Message").Values(List(Id("msg").Dot("ID"), Id("MessageKindError"), Id("responseMarshallingError").Call(Id("msg").Dot("Content"), Id("err")), Id("msg").Dot("client"))), Id("err")),
	)
}

func (p processClientMessageWriter) returnResponse() *Statement {
	if p.a.Response == nil {
		return Return(Id("Message").Values(Id("ID").Op(":").Id("msg").Dot("ID")), Nil())
	}
	return Return(Id("Message").Values(Id("msg").Dot("ID"), Id("msg").Dot("Kind"), Id("resContent"), Id("msg").Dot("client")), Nil())
}

func (p processClientMessageWriter) unknownMessageKindResponse() *Statement {
	return Id("Message").Values(List(Id("msg").Dot("ID"), Id("MessageKindError"), Index().Byte().Call(Lit("unknown message kind ").Op("+").Id("msg").Dot("Kind")), Id("msg").Dot("client")))
}
