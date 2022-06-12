package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeIDs() *Factory {
	RangeBasicTypes(func(b BasicType) {
		s.file.Type().Id(Title(b.Value) + "ID").Int()
	})

	s.config.RangeTypes(func(configType ast.ConfigType) {
		s.file.Type().Id(Title(configType.Name) + "ID").Int()
	})

	s.config.RangeRefFields(func(field ast.Field) {
		s.file.Type().Id(Title(field.ValueTypeName) + "ID").Int()
	})

	s.config.RangeAnyFields(func(field ast.Field) {

		s.file.Type().Id(Title(anyNameByField(field)) + "ID").Int()
	})

	return s
}

func (s *Factory) writeIsEmpty() *Factory {
	s.file.Func().Params(Id("s").Id("State")).Id("IsEmpty").Call().Bool().Block(
		ForEachBasicType(func(b BasicType) *Statement {
			return If(Len(Id("s").Dot(Title(b.Value))).Op("!=").Lit(0)).Block(
				Return(False()),
			)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return If(Len(Id("s").Dot(Title(configType.Name))).Op("!=").Lit(0)).Block(
				Return(False()),
			)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			return If(Len(Id("s").Dot(Title(field.ValueTypeName))).Op("!=").Lit(0)).Block(
				Return(False()),
			)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			return If(Len(Id("s").Dot(Title(anyNameByField(field)))).Op("!=").Lit(0)).Block(
				Return(False()),
			)
		}),
		Return(True()),
	)

	return s
}

func (s *Factory) writeState() *Factory {

	s.file.Type().Id("State").Struct(

		ForEachBasicType(func(basicType BasicType) *Statement {
			return Id(Title(basicType.Value)).Map(Id(Title(basicType.Value) + "ID")).Id(basicType.Value).Id(metaFieldTag(basicType.Value))
		}),

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

	s.file.Func().Id("newState").Params().Id("*State").Block(
		Return(Id("&State").Block(
			ForEachBasicType(func(basicType BasicType) *Statement {
				return Id(Title(basicType.Value)).Id(":").Make(Map(Id(Title(basicType.Value) + "ID")).Id(basicType.Value)).Id(",")
			}),

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

	return s
}

func (s *Factory) writeElements() *Factory {

	RangeBasicTypes(func(b BasicType) {
		s.file.Type().Id(b.Value).Struct(
			Id("ID").Id(Title(b.Value)+"ID").Id(metaFieldTag("id")).Line(),
			Id("Value").Id(b.Name).Id(metaFieldTag("value")).Line(),
			Id("OperationKind").Id("OperationKind").Id(metaFieldTag("operationKind")).Line(),
			Id("JSONPath").String().Id(metaFieldTag("jsonPath")),
			Id("Path").Id("path").Id(metaFieldTag("path")),
			Id("engine").Id("*Engine").Line(),
		)
	})

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := elementWriter{
			t: configType,
		}

		s.file.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(metaFieldTag("id")).Line(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(metaFieldTag("operationKind")).Line(),
			Id("HasParent").Bool().Id(metaFieldTag("hasParent")).Line(),
			Id("JSONPath").String().Id(metaFieldTag("jsonPath")),
			Id("Path").Id("path").Id(metaFieldTag("path")),
			Id("engine").Id("*Engine").Line(),
		)

		s.file.Type().Id(Title(configType.Name)).Struct(Id(configType.Name).Id(e.name()))
	})

	s.config.RangeRefFields(func(field ast.Field) {

		referencedElementName := field.ValueType().Name
		if field.HasAnyValue {
			referencedElementName = anyNameByField(field)
		}

		s.file.Type().Id(field.ValueTypeName+"Core").Struct(
			Id("ID").Id(Title(field.ValueTypeName)+"ID").Id(fieldTag("id")).Line(),
			Id("ParentID").Id(Title(field.Parent.Name)+"ID").Id(fieldTag("parentID")).Line(),
			Id("ChildID").Int().Id(fieldTag("childID")).Line(),
			Id("ReferencedElementID").Id(Title(referencedElementName)+"ID").Id(fieldTag("referencedElementID")).Line(),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("Path").Id("path").Id(metaFieldTag("path")),
			Id("engine").Id("*Engine").Line(),
		)
		s.file.Type().Id(Title(field.ValueTypeName)).Struct(Id(field.ValueTypeName).Id(field.ValueTypeName + "Core"))
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		s.file.Type().Id(anyNameByField(field)+"Core").Struct(
			Id("ID").Id(Title(anyNameByField(field))+"ID").Id(fieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(fieldTag("elementKind")).Line(),
			Id("ParentID").Int().Id(fieldTag("parentID")).Line(),
			Id("ChildID").Int().Id(fieldTag("childID")).Line(),
			Id("ParentElementPath").Id("path").Id(fieldTag("parentElementPath")).Line(),
			Id("FieldIdentifier").Id("treeFieldIdentifier").Id(fieldTag("fieldIdentifier")).Line(),
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("engine").Id("*Engine").Line(),
		)
		s.file.Type().Id(Title(anyNameByField(field))).Struct(Id(anyNameByField(field)).Id(anyNameByField(field) + "Core"))
	})

	return s
}
