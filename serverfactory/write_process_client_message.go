package serverfactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeProcessClientMessage() *ServerFactory {
	decls := NewDeclSet()

	p := processClientMessageWriter{}

	decls.File.Func().Params(p.receiverParams()).Id("processClientMessage").Params(p.params()).Id("error").Block(
		Switch().Id("messageKind").Call(Id("msg").Dot("Kind")).Block(
			ForEachActionInAST(s.config, func(action ast.Action) *Statement {
				p.a = &action
				return Case(Id(p.actionMessageKind())).Block(
					p.declareParams(),
					p.unmarshalMessageContent(),
					If(Id("err").Op("!=").Nil()).Block(
						Return(Id("err")),
					),
					p.callAction(),
				)
			}),
			Default().Block(
				Return(Id("errors").Dot("New").Call(Lit("unknown message kind"))),
			),
		),
		Return(Nil()),
	)

	decls.Render(s.buf)
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
	return Id("msg").Id("message")
}

func (p processClientMessageWriter) actionMessageKind() string {
	return "messageKindAction_" + p.a.Name
}

func (p processClientMessageWriter) declareParams() *Statement {
	return Var().Id("params").Id(Title(p.a.Name) + "Params")
}

func (p processClientMessageWriter) unmarshalMessageContent() *Statement {
	return Id("err").Op(":=").Id("params").Dot("UnmarshalJSON").Call(Id("msg").Dot("Content"))
}

func (p processClientMessageWriter) callAction() *Statement {
	return Id("r").Dot("actions").Dot(p.a.Name).Call(Id("params"), Id("r").Dot("state"))
}
