package enginefactory

import (
	"github.com/jobergner/backent-cli/ast"
	. "github.com/jobergner/backent-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type stateWriter struct {
	typeName func() string
}

func (s stateWriter) fieldName() string {
	return Title(s.typeName())
}

func (s stateWriter) mapKey() *Statement {
	return Id(Title(s.typeName()) + "ID")
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
		value += Title(e.f.ValueTypeName) + "ID"
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
	return Title(e.f.Name)
}

func (e elementWriter) name() string {
	return e.t.Name + "Core"
}

func (e elementWriter) idType() string {
	return Title(e.t.Name) + "ID"
}
