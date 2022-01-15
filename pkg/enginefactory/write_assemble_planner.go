package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAssemblePlanner() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("assemblePlanner").Struct(
		Id("updatedPaths").Map(Int()).Id("path"),
		Id("updatedReferencePaths").Map(Int()).Id("path"),
		Id("updatedElementPaths").Map(Int()).Id("path"),
		Id("includedElements").Map(Int()).Bool(),
	)

	decls.File.Func().Id("newAssemblePlanner").Params().Id("*assemblePlanner").Block(
		Return(
			Id("&assemblePlanner").Values(Dict{
				Id("updatedPaths"):          Make(Map(Int()).Id("path")),
				Id("updatedReferencePaths"): Make(Map(Int()).Id("path")),
				Id("updatedElementPaths"):   Make(Map(Int()).Id("path")),
				Id("includedElements"):      Make(Map(Int()).Bool()),
			}),
		),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAssemblePlannerClear() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Func().Params(Id("a").Id("*assemblePlanner")).Id("clear").Params().Block(
		For(Id("key").Op(":=").Range().Id("a").Dot("updatedPaths")).Block(
			Delete(Id("a").Dot("updatedPaths"), Id("key")),
		),
		For(Id("key").Op(":=").Range().Id("a").Dot("updatedElementPaths")).Block(
			Delete(Id("a").Dot("updatedElementPaths"), Id("key")),
		),
		For(Id("key").Op(":=").Range().Id("a").Dot("updatedReferencePaths")).Block(
			Delete(Id("a").Dot("updatedReferencePaths"), Id("key")),
		),
		For(Id("key").Op(":=").Range().Id("a").Dot("includedElements")).Block(
			Delete(Id("a").Dot("includedElements"), Id("key")),
		),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAssemblePlannerPlan() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Func().Params(Id("ap").Id("*assemblePlanner")).Id("plan").Params(Id("state"), Id("patch").Id("*State")).Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return For(List(Id("_"), Id(configType.Name)).Op(":=").Range().Id("patch").Dot(Title(configType.Name))).Block(
				Id("ap").Dot("updatedElementPaths").Index(Int().Call(Id(configType.Name).Dot("ID"))).Op("=").Id(configType.Name).Dot("path"),
			)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			return For(List(Id("_"), Id(field.ValueTypeName)).Op(":=").Range().Id("patch").Dot(Title(field.ValueTypeName))).Block(
				Id("ap").Dot("updatedReferencePaths").Index(Int().Call(Id(field.ValueTypeName).Dot("ID"))).Op("=").Id(field.ValueTypeName).Dot("path"),
			)
		}),
		Id("previousLen").Op(":=").Lit(0),
		For().Block(
			For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedElementPaths")).Block(
				For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
					Id("ap").Dot("includedElements").Index(Id("seg").Dot("id")).Op("=").True(),
				),
			),
			For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedReferencePaths")).Block(
				For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
					If(Id("seg").Dot("refID").Op("!=").Lit(0)).Block(
						Id("ap").Dot("includedElements").Index(Id("seg").Dot("refID")).Op("=").True(),
					).Else().Block(
						Id("ap").Dot("includedElements").Index(Id("seg").Dot("id")).Op("=").True(),
					),
				),
			),
		),
		If(Id("previousLen").Op("==").Id("len").Call(Id("ap").Dot("includedElements"))).Block(
			Break(),
		),
		Id("previousLen").Op("=").Len(Id("ap").Dot("includedElements")),
	)

	decls.Render(s.buf)
	return s
}
