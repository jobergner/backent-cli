package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writePools() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Var().Id(configType.Name + "CheckPool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Map(Id(Title(configType.Name) + "ID")).Bool()))),
		})
		decls.File.Var().Id(configType.Name + "IDSlicePool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Index().Id(Title(configType.Name)+"ID"), Lit(0)))),
		})
	})

	s.config.RangeRefFields(func(field ast.Field) {
		decls.File.Var().Id(field.ValueTypeName + "CheckPool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Map(Id(Title(field.ValueTypeName) + "ID")).Bool()))),
		})
		decls.File.Var().Id(field.ValueTypeName + "IDSlicePool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Index().Id(Title(field.ValueTypeName)+"ID"), Lit(0)))),
		})
	})

	decls.Render(s.buf)
	return s
}
