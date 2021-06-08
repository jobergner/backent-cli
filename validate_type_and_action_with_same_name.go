package validator

import (
	"fmt"
)

func validateTypeAndActionWithSameName(stateData, actionsData map[interface{}]interface{}) (errs []error) {

	var typeNames []string
	for key := range stateData {
		typeName := fmt.Sprintf("%v", key)
		typeNames = append(typeNames, typeName)
	}

	var actionNames []string
	for key := range stateData {
		actoinName := fmt.Sprintf("%v", key)
		actionNames = append(actionNames, actoinName)
	}

	for _, typeName := range typeNames {
		if contains(actionNames, typeName) {
			errs = append(errs, newValidationErrorTypeAndActionWithSameName(typeName))
		}
	}

	return
}
