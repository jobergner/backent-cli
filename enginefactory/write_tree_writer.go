package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

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

type recursionCheckWriter struct {
	typeName func() string
}

func (r recursionCheckWriter) fieldName() string {
	return r.typeName()
}

func (r recursionCheckWriter) mapKey() *Statement {
	return Id(title(r.typeName()) + "ID")
}

func (r recursionCheckWriter) mapValue() string {
	return r.typeName() + "Core"
}

func (r recursionCheckWriter) fieldTag() string {
	return "`json:\"" + r.typeName() + "\"`"
}
