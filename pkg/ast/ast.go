package ast

import (
	"sort"
)

// TODO clean up!!!

func newAST() *AST {
	return &AST{
		Types:   make(map[string]ConfigType),
		Actions: make(map[string]Action),
	}
}

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

func (a *AST) RangeAnyFields(fn func(field Field)) {
	alreadyWrittenCheck := make(map[string]bool)
	a.RangeTypes(func(configType ConfigType) {
		configType.RangeFields(func(field Field) {
			if alreadyWrittenCheck[anyNameByField(field)] {
				return
			}

			if !field.HasAnyValue {
				return
			}

			alreadyWrittenCheck[anyNameByField(field)] = true

			fn(field)
		})
	})
}

func (a *AST) RangeRefFields(fn func(field Field)) {
	a.RangeTypes(func(configType ConfigType) {
		configType.RangeFields(func(field Field) {
			if !field.HasPointerValue {
				return
			}

			fn(field)
		})
	})
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

func Parse(
	stateConfigData map[interface{}]interface{},
	actionsConfigData map[interface{}]interface{},
	responsesConfigData map[interface{}]interface{},
) *AST {
	return buildASTStructure(stateConfigData, actionsConfigData, responsesConfigData).
		fillInReferences().
		fillInParentalInfo()
}

// buildASTStructure builds the ast structure including all types and fields
// this needs to happen first so the types in "Parent" and "ValueType" can be referenced
// in fillInReferences
func buildASTStructure(
	stateConfigData map[interface{}]interface{},
	actionsConfigData map[interface{}]interface{},
	responsesConfigData map[interface{}]interface{},
) *AST {
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

		if responseData, ok := responsesConfigData[actionName]; ok {
			responseDataObject := responseData.(map[interface{}]interface{})
			action := builActionStructure(objectValue, actionName, responseDataObject)
			ast.Actions[actionName] = action
		} else {
			action := builActionStructure(objectValue, actionName, nil)
			ast.Actions[actionName] = action
		}
	}

	return ast
}

func buildTypeStructure(configTypeData map[interface{}]interface{}, typeName string) ConfigType {
	configType := newConfigType(typeName)

	for key, value := range configTypeData {
		fieldName := getSring(key)
		valueString := getSring(value)

		if fieldName == "__event__" && valueString == "true" {
			configType.IsEvent = true
			continue
		}

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

func builActionStructure(
	configActionData map[interface{}]interface{},
	actionName string,
	responseConfigData map[interface{}]interface{},
) Action {
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

	if responseConfigData != nil {
		action.Response = buildResponeStructure(responseConfigData)
	}

	return action
}

func buildResponeStructure(responseData map[interface{}]interface{}) map[string]Field {
	response := make(map[string]Field)

	for key, value := range responseData {
		valueName := getSring(key)
		valueString := getSring(value)

		responseValue := Field{
			ValueTypes:    make(map[string]*ConfigType),
			Name:          valueName,
			HasSliceValue: isSliceValue(valueString),
			ValueString:   valueString,
		}

		response[valueName] = responseValue
	}

	return response
}

// fillInReferences fills in the references of "Parent" and "ValueType"
// in Fields, and "ReferencedBy" in types
func (a *AST) fillInReferences() *AST {
	for configTypeName, _configType := range a.Types {
		configType := _configType
		configType.RangeFields(func(field Field) {
			field.Parent = &configType
			a.assignFieldTypeReference(&field)
			field.ValueTypeName = fieldValueTypeName(field)
			if field.HasPointerValue {
				field.RangeValueTypes(func(fieldValueType *ConfigType) {
					if configTypeName == fieldValueType.Name {
						f := field
						configType.ReferencedBy = append(configType.ReferencedBy, &f)
					} else {
						f := field
						fieldValueType.ReferencedBy = append(fieldValueType.ReferencedBy, &f)
						a.Types[fieldValueType.Name] = *fieldValueType
					}
				})
			}
			configType.Fields[field.Name] = field
		})
		a.Types[configTypeName] = configType
	}

	for actionName, action := range a.Actions {
		for _, param := range action.Params {
			a.assignFieldTypeReference(&param)
		}
		for _, responseValue := range action.Response {
			a.assignFieldTypeReference(&responseValue)
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
		referencedType, isUserDefinedType := a.Types[extractValueType(field.ValueString)]
		if isUserDefinedType {
			field.ValueTypes[referencedType.Name] = &referencedType
		} else {
			// TODO: maybe be more explicit
			// IDs of types (eg. playerID) are treated this way as well
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
