package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writePathTrack() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("pathTrack").Struct(
		Id("_iterations").Int(),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return Id(configType.Name).Map(Id(title(configType.Name) + "ID")).Id("path")
		}),
	)

	decls.File.Func().Id("newPathTrack").Params().Id("pathTrack").Block(
		Return(Id("pathTrack").Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				return Id(configType.Name).Id(":").Make(Map(Id(title(configType.Name) + "ID")).Id("path")).Id(",")
			}),
		)),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeIdentifiers() *EngineFactory {
	decls := NewDeclSet()

	alreadyWrittenCheck := make(map[string]bool)
	identifierValue := 0
	decls.File.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			typeIdentifierShouldBeWritten := alreadyWrittenCheck[configType.Name]
			alreadyWrittenCheck[configType.Name] = true
			if !typeIdentifierShouldBeWritten {
				identifierValue -= 1
			}
			return &Statement{
				onlyIf(!typeIdentifierShouldBeWritten, Id(configType.Name+"Identifier").Int().Op("=").Lit(identifierValue).Line()),
				ForEachFieldInType(configType, func(field ast.Field) *Statement {
					if alreadyWrittenCheck[field.Name] {
						return Empty()
					}
					if field.ValueType().IsBasicType || field.HasPointerValue {
						return Empty()
					}
					alreadyWrittenCheck[field.Name] = true
					identifierValue -= 1
					return Id(field.Name + "Identifier").Int().Op("=").Lit(identifierValue)
				}),
			}
		}),
	)

	decls.Render(s.buf)
	return s
}
