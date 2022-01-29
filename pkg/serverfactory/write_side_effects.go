package serverfactory

import (
	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeSideEffects() *ServerFactory {

	s.file.Comment("easyjson:skip")
	s.file.Type().Id("SideEffects").Struct(
		Id("OnDeploy").Func().Params(Id("*Engine")),
		Id("OnFrameTick").Func().Params(Id("*Engine")),
	)

	return s
}
