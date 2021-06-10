package serverfactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeActions() *ServerFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("actions").Struct(
		ForEachActionInAST(s.config, func(action ast.Action) *Statement {
			return Id(action.Name).Func().Params(Id(Title(action.Name)+"Params"), Id("*Engine"))
		}),
	)

	decls.Render(s.buf)
	return s
}
