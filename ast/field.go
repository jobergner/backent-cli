package ast

import "sort"

type Field struct {
	Name            string
	ValueTypes      map[string]*ConfigType // references the field's value's Type
	Parent          *ConfigType            // references the field's parent (not use when field is action param)
	ValueTypeName   string                 // the name of the to-be-generated type
	ValueString     string                 // the original value represented as string (eg. "[]Person")
	HasSliceValue   bool                   // if the value is a slice value (eg. []string)
	HasPointerValue bool                   // if the value is a pointer value (eg. *foo, []*foo)
	HasAnyValue     bool
}

func (f *Field) RangeValueTypes(fn func(configType *ConfigType)) {
	var keys []string
	for key := range f.ValueTypes {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(f.ValueTypes[key])
	}
}

func (f Field) ValueType() *ConfigType {
	var valueType *ConfigType
	f.RangeValueTypes(func(configType *ConfigType) {
		if valueType == nil {
			valueType = configType
		}
	})
	return valueType
}
