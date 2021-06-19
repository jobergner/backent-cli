package serverfactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeStart() *ServerFactory {
	decls := NewDeclSet()

	decls.File.Func().Id("Start").Params(ForEachActionInAST(s.config, func(action ast.Action) *Statement {
		responseName := Id(Title(action.Name) + "Response")
		if action.Response == nil {
			responseName = Empty()
		}
		return Id(action.Name).Func().Params(Id(Title(action.Name)+"Params"), Id("*Engine")).Add(responseName).Id(",")
	}).Id("onDeploy").Func().Params(Id("*Engine")),
		Id("onFrameTick").Func().Params(Id("*Engine")),
	).Id("error").Block(
		Id("a").Op(":=").Id("actions").Values(ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id(action.Name).Id(",")
		})),
		Id("setupRoutes").Call(Id("a"), Id("onDeploy"), Id("onFrameTick")),
		Id("err").Op(":=").Id("http").Dot("ListenAndServe").Call(Lit(":8080"), Nil()),
		Return(Id("err")),
	)

	decls.Render(s.buf)
	return s
}
