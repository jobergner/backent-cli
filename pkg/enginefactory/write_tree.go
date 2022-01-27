package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeReferencedDataStatus() *EngineFactory {

	s.file.Type().Id("ReferencedDataStatus").String()

	s.file.Const().Defs(
		Id("ReferencedDataModified").Id("ReferencedDataStatus").Op("=").Lit("MODIFIED"),
		Id("ReferencedDataUnchanged").Id("ReferencedDataStatus").Op("=").Lit("UNCHANGED"),
	)

	return s
}

func (s *EngineFactory) writeElementKinds() *EngineFactory {

	s.file.Type().Id("ElementKind").String()

	s.file.Const().Defs(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return Id("ElementKind" + Title(configType.Name)).Id("ElementKind").Op("=").Lit(Title(configType.Name))
		}),
	)

	return s
}

func (s *EngineFactory) writeTree() *EngineFactory {
	s.file.Type().Id("Tree").Struct(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			s := treeWriter{configType}
			return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
		}),
	)

	s.file.Func().Id("newTree").Params().Id("*Tree").Block(
		Return(Id("&Tree").Values(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				s := treeWriter{configType}
				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
		)),
	)

	s.file.Func().Params(Id("t").Id("*" + "Tree")).Id("clear").Params().Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return For(Id("key").Op(":=").Range().Id("t").Dot(Title(configType.Name))).Block(
				Delete(Id("t").Dot(Title(configType.Name)), Id("key")),
			)
		}),
	)

	return s
}

func (s *EngineFactory) writeTreeElements() *EngineFactory {

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := treeElementWriter{
			t: configType,
		}

		s.file.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Add(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
		)
	})

	e := treeElementWriter{}
	s.file.Type().Id("elementReference").Struct(
		Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
		Id("ElementID").Int().Id(e.metaFieldTag("id")).Line(),
		Id("ElementKind").Id("ElementKind").Id(e.metaFieldTag("elementKind")).Line(),
		Id("ReferencedDataStatus").Id("ReferencedDataStatus").Id(e.metaFieldTag("referencedDataStatus")).Line(),
		Id("ElementPath").String().Id(e.metaFieldTag("elementPath")).Line(),
	)

	return s
}
