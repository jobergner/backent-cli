package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeIdentifiers() *EngineFactory {

	s.file.Type().Id("treeFieldIdentifier").String()

	s.file.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return Id(configType.Name + "Identifier").Op("=").Lit(configType.Name)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return ForEachFieldInType(configType, func(f ast.Field) *Statement {
				if f.ValueType().IsBasicType {
					return Empty()
				}
				return Id(FieldPathIdentifier(f)).Op("=").Lit(configType.Name + "_" + f.Name)
			})
		}),
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

func (s *EngineFactory) writePath() *EngineFactory {

	s.file.Type().Id("segment").Struct(
		Id("id").Int(),
		Id("identifier").Id("treeFieldIdentifier"),
		Id("kind").Id("ElementKind"),
		Id("refID").Int(),
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
			Id("jsonPath").Op("+=").Lit(".").Op("+").Id("pathIdentifierToString").Call(Id("seg").Dot("identifier")),
			If(Id("isSliceFieldIdentifier").Call(Id("seg").Dot("identifier"))).Block(
				Id("jsonPath").Op("+=").Lit("[").Op("+").Id("strconv").Dot("Itoa").Call(Id("seg").Dot("id")).Op("+").Lit("]"),
			),
		),
		Return(Id("jsonPath")),
	)

	s.file.Func().Id("pathIdentifierToString").Params(Id("fieldIdentifier").Id("treeFieldIdentifier")).String().Block(
		Switch(Id("fieldIdentifier")).Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return Case(Id(configType.Name + "Identifier")).Block(
					Return(Lit(configType.Name)),
				)
			}),
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return ForEachFieldInType(configType, func(f ast.Field) *Statement {
					if f.ValueType().IsBasicType {
						return Empty()
					}
					return Case(Id(FieldPathIdentifier(f))).Block(
						Return(Lit(f.Name)),
					)
				})
			}),
		),
		Return(Lit("")),
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
