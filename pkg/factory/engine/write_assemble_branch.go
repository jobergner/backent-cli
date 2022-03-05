package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAssembleBranch() *EngineFactory {
	s.config.RangeTypes(func(configType ast.ConfigType) {
		a := assembleBranchWriter{
			t: configType,
			f: nil,
			v: nil,
		}

		s.file.Func().Params(Id("engine").Id("*Engine")).Id("assemble"+Title(a.t.Name)+"Path").Params(Id("element").Id("*"+a.t.Name), Id("p").Id("path"), Id("pIndex").Int(), Id("includedElements").Map(Int()).Bool()).Block(
			List(Id(a.t.Name+"Data"), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(Title(a.t.Name)).Index(Id("element").Dot("ID")),
			If(Id("!ok")).Block(
				Id(a.t.Name+"Data").Op("=").Id("engine").Dot("State").Dot(Title(a.t.Name)).Index(Id("element").Dot("ID")),
			),
			If(Id("element").Dot("OperationKind").Op("!=").Id("OperationKindUpdate").Op("&&").Id("element").Dot("OperationKind").Op("!=").Id("OperationKindDelete")).Block(
				Id("element").Dot("OperationKind").Op("=").Id(a.t.Name+"Data").Dot("OperationKind"),
			),
			If(Id("pIndex").Op("+").Lit(1).Op("==").Len(Id("p"))).Block(
				Return(),
			).Line(),
			Id("nextSeg").Op(":=").Id("p").Index(Id("pIndex").Op("+").Lit(1)).Line(),
			Switch(Id("nextSeg").Dot("Identifier")).Block(
				ForEachFieldInType(configType, func(field ast.Field) *Statement {
					a.f = &field
					return a.assembleNextSeg()
				}),
			),
			Id("_").Op("=").Id(a.t.Name+"Data"),
		)
	})

	return s
}
