package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeIDs() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Type().Id(title(configType.Name) + "ID").Int()
	})

	s.config.RangeRefFields(func(field ast.Field) {
		decls.File.Type().Id(title(field.ValueTypeName) + "ID").Int()
	})

	s.config.RangeAnyFields(func(field ast.Field) {

		decls.File.Type().Id(title(anyNameByField(field)) + "ID").Int()
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

type stateWriter struct {
	typeName func() string
}

func (s stateWriter) fieldName() string {
	return title(s.typeName())
}

func (s stateWriter) mapKey() *Statement {
	return Id(title(s.typeName()) + "ID")
}

func (s stateWriter) mapValue() string {
	return s.typeName() + "Core"
}

func (s stateWriter) fieldTag() string {
	return "`json:\"" + s.typeName() + "\"`"
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
			onlyIf(!configType.IsRootType, Id("HasParent").Bool().Id(e.metaFieldTag("hasParent")).Line()),
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
			Id("ID").Id(title(field.ValueTypeName)+"ID").Id(fieldTag("id")).Line(),
			Id("ParentID").Id(title(field.Parent.Name)+"ID").Id(fieldTag("parentID")).Line(),
			Id("ReferencedElementID").Id(title(referencedElementName)+"ID").Id(fieldTag("referencedElementID")).Line(),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("engine").Id("*Engine").Line(),
		)
		decls.File.Type().Id(field.ValueTypeName).Struct(Id(field.ValueTypeName).Id(field.ValueTypeName + "Core"))
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		decls.File.Type().Id(anyNameByField(field)+"Core").Struct(
			Id("ID").Id(title(anyNameByField(field))+"ID").Id(fieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(fieldTag("elementKind")).Line(),
			ForEachValueOfField(field, func(configType *ast.ConfigType) *Statement {
				return Id(title(configType.Name)).Id(title(configType.Name) + "ID").Id(fieldTag(configType.Name)).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("engine").Id("*Engine").Line(),
		)
		decls.File.Type().Id(anyNameByField(field)).Struct(Id(anyNameByField(field)).Id(anyNameByField(field) + "Core"))
	})

	decls.Render(s.buf)
	return s
}

func fieldTag(name string) string {
	return "`json:\"" + name + "\"`"
}

type elementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (e elementWriter) fieldValue() string {
	var value string

	if e.f.HasSliceValue {
		value = "[]"
	}

	if e.f.ValueType().IsBasicType {
		value += e.f.ValueTypeName
	} else {
		value += title(e.f.ValueTypeName) + "ID"
	}

	return value
}

func (e elementWriter) fieldTag() string {
	return "`json:\"" + e.f.Name + "\"`"
}

func (e elementWriter) metaFieldTag(name string) string {
	return "`json:\"" + name + "\"`"
}

func (e elementWriter) fieldName() string {
	return title(e.f.Name)
}

func (e elementWriter) name() string {
	return e.t.Name + "Core"
}

func (e elementWriter) idType() string {
	return title(e.t.Name) + "ID"
}
