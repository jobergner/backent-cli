package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

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
