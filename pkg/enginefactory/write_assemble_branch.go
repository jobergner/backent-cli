package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAssembleBranch() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		a := assembleBranchWriter{
			t: configType,
			f: nil,
			v: nil,
		}

		decls.File.Func().Params(Id("engine"), Id("*Engine")).Id("assemble"+Title(a.t.Name)+"Path").Params(Id("element").Id("*"+a.t.Name), Id("p").Id("path"), Id("pIndex").Int(), Id("includedElements").Map(Int()).Bool()).Block(
			List(Id(a.t.Name+"Data"), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(Title(a.t.Name)).Index(Id("element").Dot("ID")),
			If(Id("!ok")).Block(
				Id(a.t.Name+"Data").Op("=").Id("engine").Dot("State").Dot(Title(a.t.Name)).Index(Id("element").Dot("ID")),
			),
			Id("element").Dot("OperationKind").Op("=").Id(a.t.Name+"Data").Dot("OperationKind"),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				if !field.ValueType().IsBasicType {
					return Empty()
				}
				a.f = &field
				return Id("element").Dot(Title(a.f.Name)).Op("=").Id(a.t.Name + "Data").Dot(Title(a.f.Name))
			}),
			If(Id("pIndex").Op("+").Lit(1).Op("==").Len(Id("p"))).Block(
				Return(),
			),
			Id("nextSeg").Op(":=").Id("p").Index(Id("pIndex").Op("+").Lit(1)),
			OnlyIf(!a.t.IsLeafType,
				Switch(Id("nextSeg").Dot("identifier")).Block(
					ForEachFieldInType(configType, func(field ast.Field) *Statement {
						if field.ValueType().IsBasicType {
							return Empty()
						}
						a.f = &field
						return a.assembleNextSeg()
					})),
			),
			Id("_").Op("=").Id(a.t.Name+"Data"),
		)
	})

	decls.Render(s.buf)
	return s
}
