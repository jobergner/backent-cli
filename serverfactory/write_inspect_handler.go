package serverfactory

import (
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeInspectHandler(configJson []byte) *ServerFactory {
	decls := NewDeclSet()

	decls.File.Func().Id("inspectHandler").Params(Id("w").Id("http").Dot("ResponseWriter"), Id("r").Id("*http").Dot("Request")).Block(
		Id("fmt").Dot("Fprintf").Call(Id("w"), Lit(string(configJson))),
	)

	decls.Render(s.buf)
	return s
}
