package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeWalkElement() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Func().Params(Id("engine").Id("*Engine")).Id("walk"+title(configType.Name)).Params(Id(configType.Name+"ID").Id(title(configType.Name)+"ID"), Id("p").Id("path")).Block(
			List(Id(configType.Name+"Data"), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(title(configType.Name)).Index(Id(configType.Name + "ID")),
		)
	})

	decls.Render(s.buf)
	return s
}
