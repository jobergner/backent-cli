package validator

import (
	"fmt"
)

func validateResponseToUnknownAction(actionsData, responsesData map[interface{}]interface{}) (errs []error) {

	var actionNames []string
	for key := range actionsData {
		actionName := fmt.Sprintf("%v", key)
		actionNames = append(actionNames, actionName)
	}

	var responseNames []string
	for key := range responsesData {
		typeName := fmt.Sprintf("%v", key)
		responseNames = append(responseNames, typeName)
	}

	for _, responseName := range responseNames {
		if !contains(actionNames, responseName) {
			errs = append(errs, newValidationErrorResponseToUnknownAction(responseName))
		}
	}

	return
}
