package serverfactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeActions() *ServerFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("actions").Struct(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			responseName := Id(Title(action.Name) + "Response")
			if action.Response == nil {
				responseName = Empty()
			}
			return Id(action.Name).Func().Params(Id(Title(action.Name)+"Params"), Id("*Engine")).Add(responseName)
		}),
	)

	decls.Render(s.buf)
	return s
}
