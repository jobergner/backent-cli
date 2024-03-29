package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeAssembleTree() *Factory {

	s.file.Func().Params(Id("engine").Id("*Engine")).Id("AssembleUpdateTree").Params().Block(
		Id("engine").Dot("planner").Dot("clear").Call(),
		Id("engine").Dot("Tree").Dot("clear").Call(),
		Id("engine").Dot("planner").Dot("plan").Call(Id("engine").Dot("State"), Id("engine").Dot("Patch")),
		Id("engine").Dot("assembleTree").Call(),
	)

	s.file.Func().Params(Id("engine").Id("*Engine")).Id("AssembleFullTree").Params().Block(
		Id("engine").Dot("planner").Dot("clear").Call(),
		Id("engine").Dot("Tree").Dot("clear").Call(),
		Id("engine").Dot("planner").Dot("fill").Call(Id("engine").Dot("State")),
		Id("engine").Dot("assembleTree").Call(),
	)

	s.file.Func().Params(Id("engine").Id("*Engine")).Id("assembleTree").Params().Block(
		For(List(Id("_"), Id("p")).Op(":=").Range().Id("engine").Dot("planner").Dot("updatedPaths")).Block(
			Switch(Id("p").Index(Lit(0)).Dot("Identifier")).Block(
				ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
					return Case(Id(configType.Name+"Identifier")).Block(
						List(Id("child"), Id("ok")).Op(":=").Id("engine").Dot("Tree").Dot(Title(configType.Name)).Index(Id(Title(configType.Name)+"ID").Call(Id("p").Index(Lit(0)).Dot("ID"))),
						If(Id("!ok")).Block(
							Id("child").Op("=").Id(configType.Name).Values(Dict{Id("ID"): Id(Title(configType.Name) + "ID").Call(Id("p").Index(Lit(0)).Dot("ID"))}),
						),
						Id("engine").Dot("assemble"+Title(configType.Name)+"Path").Params(Id("&child"), Id("p"), Lit(0), Id("engine").Dot("planner").Dot("includedElements")),
						Id("engine").Dot("Tree").Dot(Title(configType.Name)).Index(Id(Title(configType.Name)+"ID").Call(Id("p").Index(Lit(0)).Dot("ID"))).Op("=").Id("child"),
					)
				}),
			),
		),
	)

	return s
}
