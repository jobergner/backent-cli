package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

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
	return s.t.Name
}

func (s treeWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
}

type treeElementWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (e treeElementWriter) fieldValueMapDefinition() *Statement {
	mapValueType := Id(e.f.ValueType().Name + "Reference")
	if e.f.HasPointerValue {
		mapValueType = Id("elementReference")
	}
	mapKeyType := Id(Title(e.f.ValueType().Name) + "ID")
	if e.f.HasAnyValue {
		mapKeyType = Int()
	}
	return Map(mapKeyType).Add(mapValueType)
}

func (e treeElementWriter) fieldValue() *Statement {

	if e.f.ValueType().IsBasicType {
		switch {
		case e.f.HasSliceValue:
			return Index().Id(ValueTypeName(e.f))
		default:
			return Id("*" + ValueTypeName(e.f))
		}
	}

	if e.f.HasAnyValue {
		switch {
		case e.f.HasPointerValue && !e.f.HasSliceValue:
			return Id("*elementReference")
		case !e.f.HasPointerValue && e.f.HasSliceValue:
			return Map(Int()).Interface()
		case !e.f.HasPointerValue && !e.f.HasSliceValue:
			return Interface()
		default: // e.f.HasPointerValue && e.f.HasSliceValue:
			return Map(Int()).Id("elementReference")
		}
	}

	switch {
	case e.f.HasPointerValue && e.f.HasSliceValue:
		return Map(Id(Title(e.f.ValueType().Name) + "ID")).Id("elementReference")
	case e.f.HasPointerValue && !e.f.HasSliceValue:
		return Id("*elementReference")
	case !e.f.HasPointerValue && e.f.HasSliceValue:
		return Map(Id(Title(e.f.ValueType().Name) + "ID")).Id(e.f.ValueType().Name)
	default: // !e.f.HasPointerValue && !e.f.HasSliceValue:
		return Id("*" + e.f.ValueType().Name)
	}
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
	return Lower(e.t.Name)
}

func (e treeElementWriter) idType() string {
	return Title(e.t.Name) + "ID"
}

type elementMapWriter struct {
	typeName func() string
}

func (e elementMapWriter) fieldName() string {
	return e.typeName()
}

func (e elementMapWriter) mapKey() *Statement {
	return Id(Title(e.typeName()) + "ID")
}
