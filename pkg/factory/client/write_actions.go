package client

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeActions() *Factory {

	s.config.RangeActions(func(action ast.Action) {
		s.file.Func().Params(Id("c").Id("*Client")).Id(Title(action.Name)).Params(Id("params").Id("message").Dot(Title(action.Name)+"Params")).Add(returnParams(action)).Block(
			List(Id("msgContent"), Id("err")).Op(":=").Id("params").Dot("MarshalJSON").Call(),
			If(Id("err").Op("!=").Nil()).Block(
				logWithMessageKind(action).Dot("Msg").Call(Lit("failed marshalling parameters")),
				returnError(action, "err"),
			),
			List(Id("id"), Id("err")).Op(":=").Id("newMessageID").Call(),
			If(Id("err").Op("!=").Nil()).Block(
				returnError(action, "err"),
			),
			Id("msg").Op(":=").Id("Message").Values(Id("id"), Id("message").Dot(messageKind(action)), Id("msgContent")),
			List(Id("msgBytes"), Id("err")).Op(":=").Id("msg").Dot("MarshalJSON").Call(),
			If(Id("err").Op("!=").Nil()).Block(
				logWithMessageKind(action).Dot("Int").Call(Id("logging").Dot("MessageID"), Id("msg").Dot("ID")).Dot("Str").Call(Id("logging").Dot("Message"), String().Call(Id("msgBytes"))).Dot("Msg").Call(Lit("failed marshalling message")),
				returnError(action, "err"),
			),
			OnlyIf(action.Response != nil, &Statement{
				Id("responseChan").Op(":=").Make(Chan().Index().Byte()).Line(),
				Id("c").Dot("router").Dot("add").Call(Id("id"), Id("responseChan")).Line(),
				Defer().Id("c").Dot("router").Dot("remove").Call(Id("id")).Line(),
			}),
			Id("c").Dot("messageChannel").Op("<-").Id("msgBytes"),
			OnlyIf(action.Response != nil, &Statement{
				Select().Block(
					Case(Op("<-").Id("time").Dot("After").Call(Lit(2).Op("*").Id("time").Dot("Second"))).Block(
						Id("log").Dot("Err").Call(Id("ErrResponseTimeout")).Dot("Int").Call(Id("logging").Dot("MessageID"), Id("msg").Dot("ID")).Dot("Msg").Call(Lit("timed out waiting for response")),
						returnError(action, "ErrResponseTimeout"),
					),
					Case(Id("responseBytes").Op(":=").Op("<-").Id("responseChan")).Block(
						Var().Id("res").Id("message").Dot(Title(action.Name)+"Response"),
						Id("err").Op(":=").Id("res").Dot("UnmarshalJSON").Call(Id("responseBytes")),
						If(Id("err").Op("!=").Nil()).Block(
							logWithMessageKind(action).Dot("Int").Call(Id("logging").Dot("MessageID"), Id("msg").Dot("ID")).Dot("Msg").Call(Lit("failed unmarshalling response")),
							returnError(action, "err"),
						),

						Return(Id("res"), Nil()),
					),
				),
			}),
			OnlyIf(action.Response == nil, Return(Nil())),
		)
	})

	return s
}

func messageKind(action ast.Action) string {
	return "MessageKindAction_" + action.Name
}

func logWithMessageKind(action ast.Action) *Statement {
	return Id("log").Dot("Err").Call(Id("err")).Dot("Str").Call(Id("logging").Dot("MessageKind"), String().Call(Id("message").Dot(messageKind(action))))
}

func returnError(action ast.Action, errId string) *Statement {
	switch {
	case action.Response == nil:
		return Return(Id(errId))
	default:
		return Return(Id("message").Dot(Title(action.Name)+"Response").Values(), Id(errId))
	}
}

func returnParams(action ast.Action) *Statement {
	switch {
	case action.Response == nil:
		return Error()
	default:
		return Params(Id("message").Dot(Title(action.Name)+"Response"), Error())
	}
}
