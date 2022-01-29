package serverfactory

import (
	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeInspectHandler(configJson []byte) *ServerFactory {

	s.file.Func().Id("inspectHandler").Params(Id("w").Id("http").Dot("ResponseWriter"), Id("r").Id("*http").Dot("Request")).Block(
		Id("w").Dot("Header").Call().Dot("Set").Call(Lit("Access-Control-Allow-Origin"), Lit("*")),
		Id("fmt").Dot("Fprintf").Call(Id("w"), Lit(string(configJson))),
	)

	return s
}
