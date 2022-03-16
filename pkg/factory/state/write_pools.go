package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writePools() *Factory {

	s.config.RangeTypes(func(configType ast.ConfigType) {

		s.file.Var().Id(configType.Name + "CheckPool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Map(Id(Title(configType.Name) + "ID")).Bool()))),
		})

		s.file.Var().Id(configType.Name + "IDSlicePool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Index().Id(Title(configType.Name)+"ID"), Lit(0)))),
		})

	})

	s.config.RangeRefFields(func(field ast.Field) {

		s.file.Var().Id(field.ValueTypeName + "CheckPool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Map(Id(Title(field.ValueTypeName) + "ID")).Bool()))),
		})

		s.file.Var().Id(field.ValueTypeName + "IDSlicePool").Op("=").Id("sync").Dot("Pool").Values(Dict{
			Id("New"): Func().Params().Interface().Block(Return(Make(Index().Id(Title(field.ValueTypeName)+"ID"), Lit(0)))),
		})

	})

	return s
}
