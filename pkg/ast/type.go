package ast

import "sort"

func newConfigType(name string) ConfigType {
	return ConfigType{
		Name:   name,
		Fields: make(map[string]Field),
	}
}

type ConfigType struct {
	Name          string
	Fields        map[string]Field
	ImplementedBy []*ConfigType // types which implement this type directly
	ReferencedBy  []*Field      // fields which have a reference of this type via pointer
	IsBasicType   bool          // is of one of Go's basic types (string, rune, int etc.)
	IsRootType    bool          // is not implemented into any other types and thus can not have a parent
	IsLeafType    bool          // does not implement any other user-defined types in any of its fields
	IsEvent       bool          // contains an `"__event__": true,` fields
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
	referencedBy := make([]*Field, len(t.ReferencedBy))
	copy(referencedBy, t.ReferencedBy)
	sort.Slice(referencedBy, valueTypeNameSort(referencedBy))
	for _, field := range referencedBy {
		fn(field)
	}
}

func (t *ConfigType) RangeImplementedBy(fn func(configType *ConfigType)) {
	implementedBy := make([]*ConfigType, len(t.ImplementedBy))
	copy(implementedBy, t.ImplementedBy)
	sort.Slice(implementedBy, typeNameSort(implementedBy))
	for _, t := range implementedBy {
		fn(t)
	}
}
