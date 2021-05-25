package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

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
			return Id("ElementKind" + title(configType.Name)).Id("ElementKind").Op("=").Lit(title(configType.Name))
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

type treeWriter struct {
	t ast.ConfigType
}

func (s treeWriter) fieldName() string {
	return title(s.t.Name)
}

func (s treeWriter) mapKey() *Statement {
	return Id(title(s.t.Name) + "ID")
}

func (s treeWriter) mapValue() string {
	return title(s.t.Name)
}

func (s treeWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
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
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
		)

		decls.File.Type().Id(e.name()+"Reference").Struct(
			Id("OperationKind").Id("OperationKind").Id(e.metaFieldTag("operationKind")).Line(),
			Id("ElementID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			Id("ElementKind").Id("ElementKind").Id(e.metaFieldTag("elementKind")).Line(),
			Id("ReferencedDataStatus").Id("ReferencedDataStatus").Id(e.metaFieldTag("referencedDataStatus")).Line(),
			Id("ElementPath").Id("string").Id(e.metaFieldTag("elementPath")).Line(),
			Id(title(configType.Name)).Id("*"+title(configType.Name)).Id(e.metaFieldTag(configType.Name)).Line(),
		)

	})

	decls.Render(s.buf)
	return s
}

type treeElementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (e treeElementWriter) fieldValue() string {
	var typeName string

	if e.f.HasAnyValue && !e.f.HasPointerValue {
		typeName = "interface{}"
		if e.f.HasSliceValue {
			return "[]" + typeName
		}
		return typeName
	}

	if e.f.ValueType().IsBasicType {
		typeName = e.f.ValueTypeName
	} else if e.f.HasPointerValue {
		if e.f.HasAnyValue {
			typeName = title(anyNameByField(*e.f))
		} else {
			typeName = title(e.f.ValueType().Name)
		}
		typeName = typeName + "Reference"
	} else {
		typeName = title(e.f.ValueTypeName)
	}

	if e.f.HasSliceValue {
		typeName = "[]" + typeName
	} else if !e.f.ValueType().IsBasicType {
		typeName = "*" + typeName
	}

	return typeName
}

func (e treeElementWriter) fieldTag() string {
	return "`json:\"" + e.f.Name + "\"`"
}

func (e treeElementWriter) metaFieldTag(name string) string {
	return "`json:\"" + name + "\"`"
}

func (e treeElementWriter) fieldName() string {
	return title(e.f.Name)
}

func (e treeElementWriter) name() string {
	return title(e.t.Name)
}

func (e treeElementWriter) idType() string {
	return title(e.t.Name) + "ID"
}
