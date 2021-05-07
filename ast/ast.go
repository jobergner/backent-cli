package ast

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// AST is an abstract syntax tree of a state and actions configuration.
// I could have used Go's own AST, since the way state is configured leans very heavily onto
// Go's structs, but that would have made things more complicated than they needed to be.
// This way I was also able to add functionality I needed and will be more flexible in the future.
type AST struct {
	Types   map[string]ConfigType
	Actions map[string]Action
}

func (a *AST) RangeTypes(fn func(configType ConfigType)) {
	var keys []string
	for key := range a.Types {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(a.Types[key])
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

type Action struct {
	Name   string
	Params map[string]Field
}

func (a *AST) RangeActions(fn func(action Action)) {
	var keys []string
	for key := range a.Actions {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(a.Actions[key])
	}
}

func (a *Action) RangeParams(fn func(field Field)) {
	var keys []string
	for key := range a.Params {
		keys = append(keys, key)
	}
	sort.Slice(keys, caseInsensitiveSort(keys))
	for _, key := range keys {
		fn(a.Params[key])
	}
}

func caseInsensitiveSort(keys []string) func(i, j int) bool {
	return func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	}
}

type Field struct {
	Name            string
	ValueTypes      map[string]*ConfigType // references the field's value's Type
	ValueString     string                 // the original value represented as string (eg. "[]Person")
	HasSliceValue   bool                   // if the value is a slice value (eg. []string)
	HasPointerValue bool                   // if the value is a pointer value (eg. *foo, []*foo)
	HasAnyValue     bool
}

func (f Field) ValueType() *ConfigType {
	for _, valueType := range f.ValueTypes {
		return valueType
	}
	return nil
}

func newAST() *AST {
	return &AST{
		Types:   make(map[string]ConfigType),
		Actions: make(map[string]Action),
	}
}

func Parse(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) *AST {
	return buildASTStructure(stateConfigData, actionsConfigData).
		fillInReferences().
		fillInParentalInfo()
}

func newConfigType(name string) ConfigType {
	return ConfigType{
		Name:   name,
		Fields: make(map[string]Field),
	}
}

func newAction(name string) Action {
	return Action{
		Name:   name,
		Params: make(map[string]Field),
	}
}

// buildASTStructure builds the ast structure including all types and fields
// this needs to happen first so the types in "Parent" and "ValueType" can be referenced
// in fillInReferences
func buildASTStructure(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) *AST {
	ast := newAST()
	for key, value := range stateConfigData {
		objectValue := value.(map[interface{}]interface{})
		typeName := getSring(key)

		configType := buildTypeStructure(objectValue, typeName)

		ast.Types[typeName] = configType
	}

	for key, value := range actionsConfigData {
		objectValue := value.(map[interface{}]interface{})
		actionName := getSring(key)

		action := builActionStructure(objectValue, actionName)

		ast.Actions[actionName] = action
	}

	return ast
}

func buildTypeStructure(configTypeData map[interface{}]interface{}, typeName string) ConfigType {
	configType := newConfigType(typeName)

	for key, value := range configTypeData {
		fieldName := getSring(key)
		valueString := getSring(value)

		field := Field{
			ValueTypes:      make(map[string]*ConfigType),
			Name:            fieldName,
			HasSliceValue:   isSliceValue(valueString),
			HasPointerValue: isPointerValue(valueString),
			ValueString:     valueString,
			HasAnyValue:     isAnyValue(valueString),
		}

		configType.Fields[fieldName] = field
	}

	return configType
}

func builActionStructure(configActionData map[interface{}]interface{}, actionName string) Action {
	action := newAction(actionName)

	for key, value := range configActionData {
		paramName := getSring(key)
		valueString := getSring(value)

		param := Field{
			ValueTypes:    make(map[string]*ConfigType),
			Name:          paramName,
			HasSliceValue: isSliceValue(valueString),
			ValueString:   valueString,
		}

		action.Params[paramName] = param
	}

	return action
}

// fillInReferences fills in the references of "Parent" and "ValueType"
// in Fields
func (a *AST) fillInReferences() *AST {
	for configTypeName, _configType := range a.Types {
		configType := _configType
		for fieldName, field := range configType.Fields {
			a.assignFieldTypeReference(&field)
			if field.HasPointerValue {
				for _, fieldValueType := range field.ValueTypes {
					if configTypeName == fieldValueType.Name {
						f := field
						configType.ReferencedBy = append(configType.ReferencedBy, &f)
					} else {
						f := field
						fieldValueType.ReferencedBy = append(fieldValueType.ReferencedBy, &f)
						a.Types[fieldValueType.Name] = *fieldValueType
					}
				}
			}
			configType.Fields[fieldName] = field
		}
		a.Types[configTypeName] = configType
	}

	for actionName, action := range a.Actions {
		for paramName, param := range action.Params {
			a.assignFieldTypeReference(&param)
			action.Params[paramName] = param
		}
		a.Actions[actionName] = action
	}

	return a
}

func (a *AST) assignFieldTypeReference(field *Field) {
	if field.HasAnyValue {
		for _, typeName := range extractAnyTypes(field.ValueString) {
			referencedType, _ := a.Types[typeName]
			field.ValueTypes[referencedType.Name] = &referencedType
		}
	} else {
		referencedType, ok := a.Types[extractValueType(field.ValueString)]
		if ok {
			field.ValueTypes[referencedType.Name] = &referencedType
		} else {
			field.ValueTypes[extractValueType(field.ValueString)] = &ConfigType{Name: extractValueType(field.ValueString), IsBasicType: true}
		}
	}
}

// fills in "IsLeafType" and "IsRootType" in each stateConfigField
func (a *AST) fillInParentalInfo() *AST {
	a.evalRootTypes()
	a.evalLeafTypes()
	return a
}

func (s *AST) evalLeafTypes() {
	for stateConfigTypeName, stateConfigType := range s.Types {
		isLeafType := true
		for _, stateConfigField := range stateConfigType.Fields {
			if !stateConfigField.HasPointerValue {
				for _, fieldValueType := range stateConfigField.ValueTypes {
					if !fieldValueType.IsBasicType {
						isLeafType = false
					}
				}
			}
		}
		if isLeafType {
			stateConfigType.IsLeafType = true
			s.Types[stateConfigTypeName] = stateConfigType
		}
	}
}

func (a *AST) evalRootTypes() {
	for stateConfigTypeName, stateConfigType := range a.Types {
		isRootType := true
		for _, _stateConfigType := range a.Types {
			for _, stateConfigField := range _stateConfigType.Fields {
				for _, fieldValueType := range stateConfigField.ValueTypes {
					if fieldValueType.Name == stateConfigTypeName {
						if !stateConfigField.HasPointerValue {
							isRootType = false
						}
					}
				}
			}
		}
		if isRootType {
			stateConfigType.IsRootType = true
			a.Types[stateConfigTypeName] = stateConfigType
		}
	}
}

// TODO: all this needs explanation

// "[]string" -> true
// "string" -> false
func isSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
}

func isPointerValue(valueString string) bool {
	re := regexp.MustCompile(`\*`)
	return re.MatchString(valueString)
}

func isAnyValue(valueString string) bool {
	re := regexp.MustCompile(`anyOf<`)
	return re.MatchString(valueString)
}

func extractAnyTypes(valueString string) []string {
	re := regexp.MustCompile(`<.*>`)
	s := re.FindString(valueString)
	typesRe := regexp.MustCompile(`[A-Za-z]*`)
	return typesRe.FindAllString(s, -1)
}

// "[]float64" -> float64
// "float64" -> float64
func extractValueType(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*`)
	return re.FindString(valueString)
}

func getSring(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
