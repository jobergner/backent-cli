package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeIDs() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Type().Id(Title(configType.Name) + "ID").Int()
	})

	s.config.RangeRefFields(func(field ast.Field) {
		decls.File.Type().Id(Title(field.ValueTypeName) + "ID").Int()
	})

	s.config.RangeAnyFields(func(field ast.Field) {

		decls.File.Type().Id(Title(anyNameByField(field)) + "ID").Int()
	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeState() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("State").Struct(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			s := stateWriter{
				typeName: func() string {
					return configType.Name
				},
			}

			return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			s := stateWriter{
				typeName: func() string {
					return field.ValueTypeName
				},
			}

			return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			s := stateWriter{
				typeName: func() string {
					return anyNameByField(field)
				},
			}

			return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
		}),
	)

	decls.File.Func().Id("newState").Params().Id("State").Block(
		Return(Id("State").Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {

				s := stateWriter{
					typeName: func() string {
						return configType.Name
					},
				}

				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
			ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
				s := stateWriter{
					typeName: func() string {
						return field.ValueTypeName
					},
				}

				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
			ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
				s := stateWriter{
					typeName: func() string {
						return anyNameByField(field)
					},
				}

				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
		)),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeElements() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := elementWriter{
			t: configType,
		}

		decls.File.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
			Id("HasParent").Bool().Id(e.metaFieldTag("hasParent")).Line(),
			Id("Path").String().Id(e.metaFieldTag("path")),
			Id("path").Id("path"),
			Id("engine").Id("*Engine").Line(),
		)

		decls.File.Type().Id(configType.Name).Struct(Id(configType.Name).Id(e.name()))
	})

	s.config.RangeRefFields(func(field ast.Field) {

		referencedElementName := field.ValueType().Name
		if field.HasAnyValue {
			referencedElementName = anyNameByField(field)
		}

		decls.File.Type().Id(field.ValueTypeName+"Core").Struct(
			Id("ID").Id(Title(field.ValueTypeName)+"ID").Id(fieldTag("id")).Line(),
			Id("ParentID").Id(Title(field.Parent.Name)+"ID").Id(fieldTag("parentID")).Line(),
			Id("ReferencedElementID").Id(Title(referencedElementName)+"ID").Id(fieldTag("referencedElementID")).Line(),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("engine").Id("*Engine").Line(),
		)
		decls.File.Type().Id(field.ValueTypeName).Struct(Id(field.ValueTypeName).Id(field.ValueTypeName + "Core"))
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		decls.File.Type().Id(anyNameByField(field)+"Core").Struct(
			Id("ID").Id(Title(anyNameByField(field))+"ID").Id(fieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(fieldTag("elementKind")).Line(),
			Id("ChildElementPath").Id("path").Id(fieldTag("childElementPath")).Line(),
			ForEachValueOfField(field, func(configType *ast.ConfigType) *Statement {
				return Id(Title(configType.Name)).Id(Title(configType.Name) + "ID").Id(fieldTag(configType.Name)).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("engine").Id("*Engine").Line(),
		)
		decls.File.Type().Id(anyNameByField(field)).Struct(Id(anyNameByField(field)).Id(anyNameByField(field) + "Core"))
	})

	decls.Render(s.buf)
	return s
}
