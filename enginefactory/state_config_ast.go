package enginefactory

import (
	"fmt"
	"regexp"
	"sort"
)

// stateConfigAST is an abstract syntax tree of a state configuration.
// I could have used Go's own AST, since the way state is configured leans very heavily onto
// Go's structs, but that would have made things more complicated than they needed to be.
// This way I was also able to add functionality I needed and will be more flexible in the future.
type stateConfigAST struct {
	Types map[string]stateConfigType
}

func (a *stateConfigAST) rangeTypes(fn func(configType stateConfigType)) {
	var keys []string
	for key := range a.Types {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fn(a.Types[key])
	}
}

func newStateConfigAST() *stateConfigAST {
	return &stateConfigAST{
		Types: make(map[string]stateConfigType),
	}
}

func buildStateConfigASTFrom(stateConfigData map[interface{}]interface{}) *stateConfigAST {
	return buildRudimentaryStateConfigAST(stateConfigData).
		fillInReferences().
		fillInParentalInfo()
}

type stateConfigType struct {
	Name        string
	Fields      map[string]stateConfigField
	IsBasicType bool // is of one of Go's basic types (string, rune, int etc.)
	IsRootType  bool // is not implemented into any other types and thus can not have a parent
	IsLeafType  bool // does not implement any other user-defined types in any of its fields
}

func (a *stateConfigType) rangeFields(fn func(field stateConfigField)) {
	var keys []string
	for key := range a.Fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fn(a.Fields[key])
	}
}

func newConfigType(name string) stateConfigType {
	return stateConfigType{
		Name:   name,
		Fields: make(map[string]stateConfigField),
	}
}

type stateConfigField struct {
	Name          string
	Parent        *stateConfigType // references the field's Parent's Type
	ValueType     *stateConfigType // references the field's value's Type
	ValueString   string           // the original value represented as string (eg. "[]Person")
	HasSliceValue bool             // if the value is a slice value (eg. []string)
}

// buildRudimentaryStateConfigAST builds the ast structure including all types and fields
// this needs to happen first so the types in "Parent" and "ValueType" can be referenced
// in fillInReferences
func buildRudimentaryStateConfigAST(stateConfigData map[interface{}]interface{}) *stateConfigAST {
	ast := newStateConfigAST()
	for key, value := range stateConfigData {
		objectValue := value.(map[interface{}]interface{})
		typeName := getSring(key)

		configType := buildRudimentaryStateConfigType(objectValue, typeName)

		ast.Types[typeName] = configType
	}

	return ast
}

func buildRudimentaryStateConfigType(configTypeData map[interface{}]interface{}, typeName string) stateConfigType {
	configType := newConfigType(typeName)

	for key, value := range configTypeData {
		fieldName := getSring(key)
		valueString := getSring(value)

		configField := stateConfigField{
			Name:          fieldName,
			HasSliceValue: isSliceValue(valueString),
			ValueString:   valueString,
		}

		configType.Fields[fieldName] = configField
	}

	return configType
}

// fillInReferences fills in the references of "Parent" and "ValueType"
// in each stateConfigType
func (s *stateConfigAST) fillInReferences() *stateConfigAST {
	for configTypeName, configType := range s.Types {
		ss := configType
		for stateConfigFieldName, stateConfigField := range ss.Fields {
			sf := stateConfigField
			sf.Parent = &ss
			referencedType, ok := s.Types[extractValueType(sf.ValueString)]
			if ok {
				sf.ValueType = &referencedType
			} else {
				sf.ValueType = &stateConfigType{Name: extractValueType(sf.ValueString), IsBasicType: true}
			}
			ss.Fields[stateConfigFieldName] = sf
		}
		s.Types[configTypeName] = ss
	}
	return s
}

// fills in "IsLeafType" and "isRootType" in each stateConfigField
func (s *stateConfigAST) fillInParentalInfo() *stateConfigAST {
	s.evalRootTypes()
	s.evalLeafTypes()
	return s
}

func (s *stateConfigAST) evalLeafTypes() {
	for stateConfigTypeName, stateConfigType := range s.Types {
		isLeafType := true
		for _, stateConfigField := range stateConfigType.Fields {
			if !stateConfigField.ValueType.IsBasicType {
				isLeafType = false
			}
		}
		if isLeafType {
			stateConfigType.IsLeafType = true
			s.Types[stateConfigTypeName] = stateConfigType
		}
	}
}

func (s *stateConfigAST) evalRootTypes() {
	for stateConfigTypeName, stateConfigType := range s.Types {
		isRootType := true
		for _, _stateConfigType := range s.Types {
			for _, stateConfigField := range _stateConfigType.Fields {
				if stateConfigField.ValueType.Name == stateConfigTypeName {
					isRootType = false
				}
			}
		}
		if isRootType {
			stateConfigType.IsRootType = true
			s.Types[stateConfigTypeName] = stateConfigType
		}
	}
}

// "[]string" -> true
// "string" -> false
func isSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
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
