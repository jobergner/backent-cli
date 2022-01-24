package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

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

	decls.File.Func().Id("newTree").Params().Id("*Tree").Block(
		Return(Id("&Tree").Values(
			ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
				s := treeWriter{configType}
				return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
			}),
		)),
	)

	decls.File.Func().Params(Id("t").Id("*" + "Tree")).Id("clear").Params().Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return For(Id("key").Op(":=").Range().Id("t").Dot(Title(configType.Name))).Block(
				Delete(Id("t").Dot(Title(configType.Name)), Id("key")),
			)
		}),
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
	})

	e := treeElementWriter{}
	decls.File.Type().Id("elementReference").Struct(
		Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
		Id("ElementID").Int().Id(e.metaFieldTag("id")).Line(),
		Id("ElementKind").Id("ElementKind").Id(e.metaFieldTag("elementKind")).Line(),
		Id("ReferencedDataStatus").Id("ReferencedDataStatus").Id(e.metaFieldTag("referencedDataStatus")).Line(),
		Id("ElementPath").String().Id(e.metaFieldTag("elementPath")).Line(),
	)

	decls.Render(s.buf)
	return s
}
