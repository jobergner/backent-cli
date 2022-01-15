package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeIdentifiers() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("treeFieldIdentifier").String()

	decls.File.Const().Defs(
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

	decls.Render(s.buf)
	return s
}

func writePathSegmentMethod(decls DeclSet, name string) {
	decls.File.Func().Params(Id("p").Id("path")).Id(name).Params().Id("path").Block(
		Id("newPath").Op(":=").Make(Id("[]int"), Len(Id("p")), Len(Id("p")).Op("+").Lit(1)),
		Copy(Id("newPath"), Id("p")),
		Id("newPath").Op("=").Append(Id("newPath"), Id(name+"Identifier")),
		Return(Id("newPath")),
	)
}

func (s *EngineFactory) writePathSegments() *EngineFactory {
	decls := NewDeclSet()

	alreadyWrittenCheck := make(map[string]bool)
	s.config.RangeTypes(func(configType ast.ConfigType) {

		if !alreadyWrittenCheck[configType.Name] {
			writePathSegmentMethod(decls, configType.Name)
			alreadyWrittenCheck[configType.Name] = true
		}

		configType.RangeFields(func(field ast.Field) {
			if alreadyWrittenCheck[field.Name] {
				return
			}
			if field.ValueType().IsBasicType || field.HasPointerValue {
				return
			}
			alreadyWrittenCheck[field.Name] = true
			writePathSegmentMethod(decls, field.Name)
		})

	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writePath() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("path").Index().Id("segment")

	decls.File.Func().Id("newPath").Params().Id("path").Block(
		Return(Make(Id("path"), Lit(0))),
	)

	decls.File.Func().Params(Id("p").Id("path")).Id("toJSONPath").Params().String().Block(
		Id("jsonPath").Op(":=").Lit("$"),
		For(List(Id("_"), Id("seg")).Op(":=").Range().Id("p")).Block(
			Id("jsonPath").Op("+=").Lit(".").Op("+").Id("pathIdentifierToString").Call(Id("seg").Dot("identifier")),
			If(Id("isSliceFieldIdentifier").Call(Id("seg").Dot("identifier"))).Block(
				Id("jsonPath").Op("+=").Lit("[").Op("+").Id("strconv").Dot("Itoa").Call(Id("seg").Dot("id")).Op("+").Lit("]"),
			),
		),
		Return(Id("jsonPath")),
	)

	decls.File.Func().Id("pathIdentifierToString").Params(Id("fieldIdentifier").Id("treeFieldIdentifier")).String().Block(
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

	decls.File.Func().Id("isSliceFieldIdentifier").Params(Id("fieldIdentifier").Id("treeFieldIdentifier")).Bool().Block(
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

	decls.Render(s.buf)
	return s
}
