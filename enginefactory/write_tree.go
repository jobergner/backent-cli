package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeTree() *EngineFactory {
	decls := newDeclSet()
	decls.file.Type().Id("Tree").Struct(forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
		s := treeWriter{configType}
		return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
	}))

	decls.file.Func().Id("newTree").Params().Id("Tree").Block(
		Return(Id("Tree").Values(forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			s := treeWriter{configType}
			return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
		}))),
	)

	decls.render(s.buf)
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
	return "t" + title(s.t.Name)
}

func (s treeWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
}

func (s *EngineFactory) writeTreeElements() *EngineFactory {
	decls := newDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := treeElementWriter{
			t: configType,
		}

		decls.file.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			forEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind_").Id("OperationKind").Id(e.metaFieldTag("operationKind_")).Line(),
		)
	})

	decls.render(s.buf)
	return s
}

type treeElementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (e treeElementWriter) fieldValue() string {
	if e.f.ValueType.IsBasicType {
		if e.f.HasSliceValue {
			return "[]" + e.f.ValueType.Name
		}
		return e.f.ValueType.Name
	}

	if e.f.HasSliceValue {
		return "[]" + "t" + title(e.f.ValueType.Name)
	}

	return "*" + "t" + title(e.f.ValueType.Name)
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
	return "t" + title(e.t.Name)
}

func (e treeElementWriter) idType() string {
	return title(e.t.Name) + "ID"
}
