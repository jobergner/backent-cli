package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeAny() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeAnyFields(func(field ast.Field) {

		decls.File.Func().Params(Id("_any").Id(anyNameByField(field))).Id("Kind").Params().Id("ElementKind").Block(
			Id("any").Op(":=").Id("_any").Dot(anyNameByField(field)).Dot("engine").Dot(anyNameByField(field)).Call(Id("_any").Dot(anyNameByField(field)).Dot("ID")),
			Return(Id("any").Dot(anyNameByField(field)).Dot("ElementKind")),
		)

		field.RangeValueTypes(func(configType *ast.ConfigType) {
			decls.File.Func().Params(Id("_any").Id(anyNameByField(field))).Id("Set"+title(configType.Name)).Params().Id(configType.Name).Block(
				Id(configType.Name).Op(":=").Id("_any").Dot(anyNameByField(field)).Dot("engine").Dot("create"+title(configType.Name)).Call(True()),
				Id("_any").Dot(anyNameByField(field)).Dot("set"+title(configType.Name)).Call(Id(configType.Name).Dot("ID").Call()),
				Return(Id(configType.Name)),
			)
			decls.File.Func().Params(Id("_any").Id(anyNameByField(field)+"Core")).Id("set"+title(configType.Name)).Params(Id(configType.Name+"ID").Id(title(configType.Name+"ID"))).Block(
				Id("any").Op(":=").Id("_any").Dot("engine").Dot(anyNameByField(field)).Call(Id("_any").Dot("ID")).Dot(anyNameByField(field)),
				ForEachValueOfField(field, func(valueType *ast.ConfigType) *Statement {
					if valueType.Name == configType.Name {
						return Empty()
					}
					return If(Id("any").Dot(title(valueType.Name)).Op("!=").Lit(0)).Block(
						Id("any").Dot("engine").Dot("delete"+title(valueType.Name)).Call(Id("any").Dot(title(valueType.Name))),
						Id("any").Dot(title(valueType.Name)).Op("=").Lit(0),
					)
				}),
				Id("any").Dot("ElementKind").Op("=").Id("ElementKind"+title(configType.Name)),
				Id("any").Dot(title(configType.Name)).Op("=").Id(configType.Name+"ID"),
				Id("any").Dot("engine").Dot("Patch").Dot(title(anyNameByField(field))).Index(Id("any").Dot("ID")).Op("=").Id("any"),
			)
		})

		decls.File.Func().Params(Id("_any").Id(anyNameByField(field)+"Core")).Id("deleteChild").Params().Block(
			Id("any").Op(":=").Id("_any").Dot("engine").Dot(anyNameByField(field)).Call(Id("_any").Dot("ID")).Dot(anyNameByField(field)),
			Switch(Id("any").Dot("ElementKind")).Block(
				ForEachValueOfField(field, func(valueType *ast.ConfigType) *Statement {
					return Case(Id("ElementKind" + title(valueType.Name))).Block(
						Id("any").Dot("engine").Dot("delete" + title(valueType.Name)).Call(Id("any").Dot(title(valueType.Name))),
					)
				}),
			),
		)

	})

	decls.Render(s.buf)
	return s
}
