package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeReferencedDataStatus() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("ReferencedDataStatus").String()

	decls.File.Const().Defs(
		Id("ReferencedDataModified").Id("ReferencedDataStatus").Op("=").Lit("MODIFIED"),
		Id("ReferencedDataUnchanged").Id("ReferencedDataStatus").Op("=").Lit("UNCHANGED"),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeElementKinds() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("ElementKind").String()

	decls.File.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return Id("ElementKind" + Title(configType.Name)).Id("ElementKind").Op("=").Lit(Title(configType.Name))
		}),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeTree() *EngineFactory {
	decls := NewDeclSet()
	decls.File.Type().Id("Tree").Struct(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			s := treeWriter{configType}
			return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
		}),
	)

	decls.File.Func().Id("newTree").Params().Id("Tree").Block(
		Return(Id("Tree").Values(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				s := treeWriter{configType}
				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
		)),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeTreeElements() *EngineFactory {
	decls := NewDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := treeElementWriter{
			t: configType,
		}

		decls.File.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Add(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
		)

		decls.File.Type().Id(e.name()+"Reference").Struct(
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
			Id("ElementID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(e.metaFieldTag("elementKind")).Line(),
			Id("ReferencedDataStatus").Id("ReferencedDataStatus").Id(e.metaFieldTag("referencedDataStatus")).Line(),
			Id("ElementPath").Id("string").Id(e.metaFieldTag("elementPath")).Line(),
			Id(Title(configType.Name)).Id("*"+Title(configType.Name)).Id(e.metaFieldTag(configType.Name)).Line(),
		)

	})

	s.config.RangeAnyFields(func(field ast.Field) {
		if !field.HasPointerValue {
			return
		}
		decls.File.Type().Id(Title(anyNameByField(field))+"Reference").Struct(
			Id("OperationKind").Id("OperationKind").Id(fieldTag("operationKind")).Line(),
			Id("ElementID").Int().Id(fieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(fieldTag("elementKind")).Line(),
			Id("ReferencedDataStatus").Id("ReferencedDataStatus").Id(fieldTag("referencedDataStatus")).Line(),
			Id("ElementPath").Id("string").Id(fieldTag("elementPath")).Line(),
			Id("Element").Interface().Id(fieldTag("element")).Line(),
		)
	})

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeRecursionCheck() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("recursionCheck").Struct(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			r := elementMapWriter{
				typeName: func() string {
					return configType.Name
				},
			}

			return Id(r.fieldName()).Map(r.mapKey()).Bool().Line()
		}),
	)

	decls.File.Func().Id("newRecursionCheck").Params().Id("*recursionCheck").Block(
		Return(Id("&recursionCheck").Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {

				r := elementMapWriter{
					typeName: func() string {
						return configType.Name
					},
				}

				return Id(r.fieldName()).Id(":").Make(Map(r.mapKey()).Bool()).Id(",")
			}),
		)),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeAssembleCache() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("assembleCache").Struct(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			e := elementMapWriter{
				typeName: func() string {
					return configType.Name
				},
			}

			return Id(e.fieldName()).Map(e.mapKey()).Id(configType.Name + "CacheElement").Line()
		}),
	)

	decls.File.Func().Id("newAssembleCache").Params().Id("assembleCache").Block(
		Return(Id("assembleCache").Block(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {

				e := elementMapWriter{
					typeName: func() string {
						return configType.Name
					},
				}

				return Id(e.fieldName()).Id(":").Make(Map(e.mapKey()).Id(configType.Name + "CacheElement")).Id(",")
			}),
		)),
	)

	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.File.Type().Id(configType.Name+"CacheElement").Struct(
			Id("hasUpdated").Bool(),
			Id(configType.Name).Id(Title(configType.Name)),
		)
	})

	decls.Render(s.buf)
	return s
}
