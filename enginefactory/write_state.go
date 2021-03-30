package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

func (s *stateFactory) writeIDs() *stateFactory {
	decls := newDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		decls.file.Type().Id(title(configType.Name) + "ID").Int()
	})

	decls.render(s.buf)
	return s
}

func (s *stateFactory) writeState() *stateFactory {
	decls := newDeclSet()
	decls.file.Type().Id("State").Struct(forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
		s := stateWriter{configType}
		return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
	}))

	decls.file.Func().Id("newState").Params().Id("State").Block(
		Return(Id("State").Values(forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			s := stateWriter{configType}
			return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
		}))),
	)

	decls.render(s.buf)
	return s
}

type stateWriter struct {
	t ast.ConfigType
}

func (s stateWriter) fieldName() string {
	return title(s.t.Name)
}

func (s stateWriter) mapKey() *Statement {
	return Id(title(s.t.Name) + "ID")
}

func (s stateWriter) mapValue() string {
	return s.t.Name + "Core"
}

func (s stateWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
}

func (s *stateFactory) writeElements() *stateFactory {
	decls := newDeclSet()

	s.config.RangeTypes(func(configType ast.ConfigType) {

		e := elementWriter{
			t: configType,
		}

		decls.file.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			forEachFieldInType(configType, func(field ast.Field) *Statement {
				e.f = &field
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind_").Id("OperationKind").Id(e.metaFieldTag("operationKind_")).Line(),
			onlyIf(!configType.IsRootType, Id("HasParent_").Bool().Id(e.metaFieldTag("hasParent_")).Line()),
		)

		decls.file.Type().Id(title(configType.Name)).Struct(Id(configType.Name).Id(e.name()))
	})

	decls.render(s.buf)
	return s
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

	if e.f.ValueType.IsBasicType {
		value += e.f.ValueType.Name
	} else {
		value += title(e.f.ValueType.Name) + "ID"
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
