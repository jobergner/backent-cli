package serverfactory

import (
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeSideEffects() *ServerFactory {
	decls := NewDeclSet()

	decls.File.Comment("easyjson:skip")
	decls.File.Type().Id("SideEffects").Struct(
		Id("OnDeploy").Func().Params(Id("*Engine")),
		Id("OnFrameTick").Func().Params(Id("*Engine")),
	)

	decls.Render(s.buf)
	return s
}
