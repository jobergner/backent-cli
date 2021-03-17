package actionsfactory

import (
	"fmt"
	"regexp"
)

var golangBasicTypes = []string{"string", "bool", "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32", "int64", "uint64", "int", "uint", "uintptr", "float32", "float64", "complex64", "complex128"}

type actionsConfigAST struct {
	Actions map[string]Action
}

func newActionsConfigAST() *actionsConfigAST {
	return &actionsConfigAST{
		Actions: make(map[string]Action),
	}
}

type Action struct {
	Name   string
	Params map[string]ActionParameter
}

func newAction() Action {
	return Action{
		Params: make(map[string]ActionParameter),
	}
}

type ActionParameter struct {
	Name         string
	TypeLiteral  string
	IsSliceValue bool
	IsBasicType  bool
}

func buildActionsConfigAST(actionsConfigData map[interface{}]interface{}) *actionsConfigAST {
	ast := newActionsConfigAST()
	for key, value := range actionsConfigData {
		objectValue := value.(map[interface{}]interface{})
		actionName := getSring(key)

		action := buildAction(objectValue, actionName)

		ast.Actions[actionName] = action
	}

	return ast
}

func buildAction(configActionData map[interface{}]interface{}, actionName string) Action {
	action := newAction()
	action.Name = actionName

	for key, value := range configActionData {
		paramName := getSring(key)
		valueString := getSring(value)

		valueType := extractValueType(valueString)
		actionParam := ActionParameter{
			Name:         paramName,
			TypeLiteral:  valueType,
			IsSliceValue: isSliceValue(valueString),
			IsBasicType:  isBasicType(valueType),
		}

		action.Params[paramName] = actionParam
	}

	return action
}

// TODO duplicate: already exists in stateFactory
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

func isBasicType(typeLiteral string) bool {
	for _, basicType := range golangBasicTypes {
		if basicType == typeLiteral {
			return true
		}
	}
	return false
}
