package ast

import "sort"

func newConfigType(name string) ConfigType {
	return ConfigType{
		Name:   name,
		Fields: make(map[string]Field),
	}
}

type ConfigType struct {
	Name         string
	Fields       map[string]Field
	ReferencedBy []*Field
	IsBasicType  bool // is of one of Go's basic types (string, rune, int etc.)
	IsRootType   bool // is not implemented into any other types and thus can not have a parent
	IsLeafType   bool // does not implement any other user-defined types in any of its fields
}

func (t *ConfigType) RangeFields(fn func(field Field)) {
	var keys []string
	for key := range t.Fields {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(t.Fields[key])
	}
}

func (t *ConfigType) RangeReferencedBy(fn func(field *Field)) {
	for _, field := range t.ReferencedBy {
		fn(field)
	}
}
