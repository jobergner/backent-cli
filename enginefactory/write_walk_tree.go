package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeWalkElement() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Func().Params(Id("engine").Id("*Engine")).Id("walk"+title(configType.Name)).Params(Id(configType.Name+"ID").Id(title(configType.Name)+"ID"), Id("p").Id("path")).Block(
			List(Id(configType.Name+"Data"), Id("hasUpdated")).Op(":=").Id("engine").Dot("Patch").Dot(title(configType.Name)).Index(Id(configType.Name+"ID")),
			If(Id("!hasUpdated")).Block(
				Id(configType.Name+"Data").Op("=").Id("engine").Dot("State").Dot(title(configType.Name)).Index(Id(configType.Name+"ID")),
			),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				if !field.HasSliceValue {
					return &Statement{
						Var().Id(field.Name + "Path").Id("path").Line(),
						If(List(Id("existingPath"), Id("pathExists")).Op(":=").Id("engine").Dot("PathTrack").Dot(field.ValueTypeName).Index(Id(configType.Name+"Data").Dot(title(field.Name))), Id("!pathExists")).Block(
							Id(field.Name + "Path").Op("=").Id("p").Dot(field.ValueTypeName).Call(),
						).Else().Block(
							Id(field.Name + "Path").Op("=").Id("existingPath"),
						).Line(),
						Id("engine").Dot("walk"+title(field.ValueTypeName)).Call(Id(configType.Name+"Data").Dot(title(field.Name)), Id(field.Name+"Path")),
					}
				}
				return &Statement{
					For(List(Id("i"), Id(field.ValueTypeName)).Op(":=").Range().Id("merge"+title(field.ValueTypeName)+"IDs").Call(
						Id("engine").Dot("State").Dot(title(configType.Name)).Index(Id(configType.Name+"Data").Dot("ID")).Dot(title(field.Name)),
						Id("engine").Dot("Patch").Dot(title(configType.Name)).Index(Id(configType.Name+"Data").Dot("ID")).Dot(title(field.Name)),
					)).Block(
						Var().Id(field.Name+"Path").Id("path").Line(),
						If(List(Id("existingPath"), Id("pathExists")).Op(":=").Id("engine").Dot("PathTrack").Dot(field.ValueTypeName).Index(Id(configType.Name+"Data").Dot(title(field.Name))), Id("!pathExists")).Block(
							Id(field.Name+"Path").Op("=").Id("p").Dot(field.ValueTypeName).Call(),
						).Else().Block(
							Id(field.Name+"Path").Op("=").Id("existingPath"),
						).Line(),
						Id("engine").Dot("walk"+title(field.ValueTypeName)).Call(Id(configType.Name+"Data").Dot(title(field.Name)), Id(field.Name+"Path")),
					),
				}
			}),
		)
	})

	decls.Render(s.buf)
	return s
}
