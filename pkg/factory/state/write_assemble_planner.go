package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeAssemblePlanner() *Factory {

	s.file.Type().Id("assemblePlanner").Struct(
		Id("updatedPaths").Index().Id("path"),
		Id("updatedReferencePaths").Map(Id("ComplexID")).Id("path"),
		Id("updatedElementPaths").Map(Int()).Id("path"),
		Id("includedElements").Map(Int()).Bool(),
	)

	s.file.Func().Id("newAssemblePlanner").Params().Id("*assemblePlanner").Block(
		Return(
			Id("&assemblePlanner").Values(Dict{
				Id("updatedPaths"):          Make(Index().Id("path"), Lit(0)),
				Id("updatedReferencePaths"): Make(Map(Id("ComplexID")).Id("path")),
				Id("updatedElementPaths"):   Make(Map(Int()).Id("path")),
				Id("includedElements"):      Make(Map(Int()).Bool()),
			}),
		),
	)

	return s
}

func (s *Factory) writeAssemblePlannerClear() *Factory {

	s.file.Func().Params(Id("a").Id("*assemblePlanner")).Id("clear").Params().Block(
		Id("a").Dot("updatedPaths").Op("=").Id("a").Dot("updatedPaths").Index(Empty(), Lit(0)),
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

	return s
}

func (s *Factory) writeAssemblePlannerPlan() *Factory {

	ap := assemblePlannerWriter{
		f: nil,
	}

	s.file.Func().Params(Id("ap").Id("*assemblePlanner")).Id("plan").Params(Id("state"), Id("patch").Id("*State")).Block(
		ForEachBasicType(func(b BasicType) *Statement {
			return For(List(Id("_"), Id(b.Value)).Op(":=").Range().Id("patch").Dot(Title(b.Value))).Block(
				Id("ap").Dot("updatedElementPaths").Index(Int().Call(Id(b.Value).Dot("ID"))).Op("=").Id(b.Value).Dot("Path"),
			)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return For(List(Id("_"), Id(configType.Name)).Op(":=").Range().Id("patch").Dot(Title(configType.Name))).Block(
				Id("ap").Dot("updatedElementPaths").Index(Int().Call(Id(configType.Name).Dot("ID"))).Op("=").Id(configType.Name).Dot("Path"),
			)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			return For(List(Id("_"), Id(field.ValueTypeName)).Op(":=").Range().Id("patch").Dot(Title(field.ValueTypeName))).Block(
				Id("ap").Dot("updatedReferencePaths").Index(Id("ComplexID").Call(Id(field.ValueTypeName).Dot("ID"))).Op("=").Id(field.ValueTypeName).Dot("Path"),
			)
		}),
		Id("previousLen").Op(":=").Lit(0),
		For().Block(
			For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedElementPaths")).Block(
				For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
					Id("ap").Dot("includedElements").Index(Id("seg").Dot("ID")).Op("=").True(),
				),
			),
			For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedReferencePaths")).Block(
				For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
					If(Id("seg").Dot("RefID").Op("!=").Params(Id("ComplexID").Values())).Block().Else().Block(
						Id("ap").Dot("includedElements").Index(Id("seg").Dot("ID")).Op("=").True(),
					),
				),
			),
			If(Id("previousLen").Op("==").Id("len").Call(Id("ap").Dot("includedElements"))).Block(
				Break(),
			),
			Id("previousLen").Op("=").Len(Id("ap").Dot("includedElements")),
			ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
				ap.f = &field
				return For(ap.eachRefInState("patch")).Block(
					If(ap.includedElementsContainReferencedElement(), Id("ok")).Block(
						ap.putPathInUpdatedReferencePaths(),
					),
				).Line().
					For(ap.eachRefInState("state")).Block(
					If(ap.pathAlreadyIncluded(), Id("ok")).Block(
						Continue(),
					),
					If(ap.includedElementsContainReferencedElement(), Id("ok")).Block(
						ap.putPathInUpdatedReferencePaths(),
					),
				)
			}),
		),
		For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedElementPaths")).Block(
			Id("ap").Dot("updatedPaths").Op("=").Append(Id("ap").Dot("updatedPaths"), Id("p")),
		),
		For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedReferencePaths")).Block(
			Id("ap").Dot("updatedPaths").Op("=").Append(Id("ap").Dot("updatedPaths"), Id("p")),
		),
	)

	return s
}

func (s *Factory) writeAssemblePlannerFill() *Factory {

	s.file.Func().Params(Id("ap").Id("*assemblePlanner")).Id("fill").Params(Id("state").Id("*State")).Block(
		ForEachBasicType(func(b BasicType) *Statement {
			return For(List(Id("_"), Id(b.Value)).Op(":=").Range().Id("state").Dot(Title(b.Value))).Block(
				Id("ap").Dot("updatedElementPaths").Index(Int().Call(Id(b.Value).Dot("ID"))).Op("=").Id(b.Value).Dot("Path"),
			)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return For(List(Id("_"), Id(configType.Name)).Op(":=").Range().Id("state").Dot(Title(configType.Name))).Block(
				Id("ap").Dot("updatedElementPaths").Index(Int().Call(Id(configType.Name).Dot("ID"))).Op("=").Id(configType.Name).Dot("Path"),
			)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			return For(List(Id("_"), Id(field.ValueTypeName)).Op(":=").Range().Id("state").Dot(Title(field.ValueTypeName))).Block(
				Id("ap").Dot("updatedReferencePaths").Index(Id("ComplexID").Call(Id(field.ValueTypeName).Dot("ID"))).Op("=").Id(field.ValueTypeName).Dot("Path"),
			)
		}),
		For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedElementPaths")).Block(
			Id("ap").Dot("updatedPaths").Op("=").Append(Id("ap").Dot("updatedPaths"), Id("p")),
		),
		For(List(Id("_"), Id("p")).Op(":=").Range().Id("ap").Dot("updatedReferencePaths")).Block(
			Id("ap").Dot("updatedPaths").Op("=").Append(Id("ap").Dot("updatedPaths"), Id("p")),
		),
	)

	return s
}
