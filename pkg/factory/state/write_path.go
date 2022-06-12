package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeIdentifiers() *Factory {

	s.file.Type().Id("treeFieldIdentifier").Int()

	var identifierValue = 0

	s.file.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			identifierValue = identifierValue + 1
			return Id(configType.Name + "Identifier").Id("treeFieldIdentifier").Op("=").Lit(identifierValue)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return ForEachFieldInType(configType, func(f ast.Field) *Statement {
				identifierValue = identifierValue + 1
				return Id(FieldPathIdentifier(f)).Id("treeFieldIdentifier").Op("=").Lit(identifierValue)
			})
		}),
	)

	s.file.Func().Params(Id("t").Id("treeFieldIdentifier")).Id("toString").Params().String().Block(
		Switch(Id("t")).Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return Case(Id(configType.Name + "Identifier")).Block(
					Return(Lit(configType.Name)),
				)
			}),
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return ForEachFieldInType(configType, func(f ast.Field) *Statement {
					return Case(Id(FieldPathIdentifier(f))).Block(
						Return(Lit(f.Name)),
					)
				})
			}),
			Default().Block(
				Panic(Id("fmt.Sprintf(\"no string found for identifier <%d>\", t)")),
			),
		),
	)

	return s
}

func writePathSegmentMethod(file *File, name string) {
	file.Func().Params(Id("p").Id("path")).Id(name).Params().Id("path").Block(
		Id("newPath").Op(":=").Make(Id("[]int"), Len(Id("p")), Len(Id("p")).Op("+").Lit(1)),
		Copy(Id("newPath"), Id("p")),
		Id("newPath").Op("=").Append(Id("newPath"), Id(name+"Identifier")),
		Return(Id("newPath")),
	)
}

func (s *Factory) writePath() *Factory {

	s.file.Type().Id("segment").Struct(
		Id("ID").Int().Add(Id(metaFieldTag("id"))),
		Id("Identifier").Id("treeFieldIdentifier").Add(Id(metaFieldTag("identifier"))),
		Id("Kind").Id("ElementKind").Add(Id(metaFieldTag("kind"))),
		Id("RefID").Int().Add(Id(metaFieldTag("refID"))),
	)

	s.file.Type().Id("path").Index().Id("segment")

	s.file.Func().Id("newPath").Params().Id("path").Block(
		Return(Make(Id("path"), Lit(0))),
	)

	s.file.Func().Params(Id("p").Id("path")).Id("extendAndCopy").Params(Id("fieldIdentifier").Id("treeFieldIdentifier"), Id("id").Int(), Id("kind").Id("ElementKind"), Id("refID").Int()).Id("path").Block(
		Id("newPath").Op(":=").Make(Id("path"), Len(Id("p")), Len(Id("p")).Op("+").Lit(1)),
		Copy(Id("newPath"), Id("p")),
		Id("newPath").Op("=").Append(Id("newPath"), Id("segment").Values(Id("id"), Id("fieldIdentifier"), Id("kind"), Id("refID"))),
		Return(Id("newPath")),
	)

	s.file.Func().Params(Id("p").Id("path")).Id("toJSONPath").Params().String().Block(
		Id("jsonPath").Op(":=").Lit("$"),
		For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
			Id("jsonPath").Op("+=").Lit(".").Op("+").Id("seg").Dot("Identifier").Dot("toString").Call(),
			If(Id("isSliceFieldIdentifier").Call(Id("seg").Dot("Identifier"))).Block(
				Id("jsonPath").Op("+=").Lit("[").Op("+").Id("strconv").Dot("Itoa").Call(Id("seg").Dot("ID")).Op("+").Lit("]"),
			),
		),
		Return(Id("jsonPath")),
	)

	s.file.Func().Id("isSliceFieldIdentifier").Params(Id("fieldIdentifier").Id("treeFieldIdentifier")).Bool().Block(
		Switch(Id("fieldIdentifier")).Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return Case(Id(configType.Name + "Identifier")).Block(
					Return(True()),
				)
			}),
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return ForEachFieldInType(configType, func(f ast.Field) *Statement {
					if f.ValueType().IsBasicType || !f.HasSliceValue {
						return Empty()
					}
					return Case(Id(FieldPathIdentifier(f))).Block(
						Return(True()),
					)
				})
			}),
		),
		Return(False()),
	)

	return s
}
