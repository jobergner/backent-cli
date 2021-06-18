package enginefactory

import (
	"github.com/Java-Jonas/bar-cli/ast"
	. "github.com/Java-Jonas/bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

type treeWriter struct {
	t ast.ConfigType
}

func (s treeWriter) fieldName() string {
	return Title(s.t.Name)
}

func (s treeWriter) mapKey() *Statement {
	return Id(Title(s.t.Name) + "ID")
}

func (s treeWriter) mapValue() string {
	return Title(s.t.Name)
}

func (s treeWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
}

type treeElementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (e treeElementWriter) fieldValueMapDefinition() *Statement {
	mapValueType := Id(Title(e.f.ValueType().Name) + "Reference")
	if e.f.HasAnyValue && !e.f.HasPointerValue {
		mapValueType = Id("interface{}")
	}
	if e.f.HasPointerValue && e.f.HasAnyValue {
		mapValueType = Id(Title(anyNameByField(*e.f) + "Reference"))
	}
	if !e.f.HasPointerValue && !e.f.HasAnyValue {
		mapValueType = Id(Title(e.f.ValueType().Name))
	}
	mapKeyType := Id(Title(e.f.ValueType().Name) + "ID")
	if e.f.HasAnyValue {
		mapKeyType = Int()
	}
	return Map(mapKeyType).Add(mapValueType)
}

func (e treeElementWriter) fieldValue() *Statement {
	var typeName string

	if e.f.HasAnyValue && !e.f.HasPointerValue {
		if !e.f.HasSliceValue {
			return Id("interface{}")
		}
		return e.fieldValueMapDefinition()
	}

	if e.f.ValueType().IsBasicType {
		typeName = e.f.ValueTypeName
	} else if e.f.HasPointerValue {
		if e.f.HasAnyValue {
			typeName = Title(anyNameByField(*e.f))
		} else {
			typeName = Title(e.f.ValueType().Name)
		}
		typeName = typeName + "Reference"
	} else {
		typeName = Title(e.f.ValueTypeName)
	}

	if e.f.HasSliceValue {
		if e.f.ValueType().IsBasicType {
			return Id("[]" + typeName)
		}
		return e.fieldValueMapDefinition()
	} else if !e.f.ValueType().IsBasicType {
		return Id("*" + typeName)
	}

	return Id(typeName)
}

func (e treeElementWriter) fieldTag() string {
	return "`json:\"" + e.f.Name + "\"`"
}

func (e treeElementWriter) metaFieldTag(name string) string {
	return "`json:\"" + name + "\"`"
}

func (e treeElementWriter) fieldName() string {
	return Title(e.f.Name)
}

func (e treeElementWriter) name() string {
	return Title(e.t.Name)
}

func (e treeElementWriter) idType() string {
	return Title(e.t.Name) + "ID"
}

type recursionCheckWriter struct {
	typeName func() string
}

func (r recursionCheckWriter) fieldName() string {
	return r.typeName()
}

func (r recursionCheckWriter) mapKey() *Statement {
	return Id(Title(r.typeName()) + "ID")
}

func (r recursionCheckWriter) mapValue() string {
	return r.typeName() + "Core"
}

func (r recursionCheckWriter) fieldTag() string {
	return "`json:\"" + r.typeName() + "\"`"
}
