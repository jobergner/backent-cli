package validator

import (
	"fmt"
)

func validateActionUsedAsType(actionsData map[interface{}]interface{}) (errs []error) {

	var actionsNames []string
	for key := range actionsData {
		keyName := fmt.Sprintf("%v", key)
		actionsNames = append(actionsNames, keyName)
	}

	for key, value := range actionsData {
		if isMap(value) {
			keyName := fmt.Sprintf("%v", key)
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateActionUsedAsTypeObject(keyName, mapValue, actionsNames)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateActionUsedAsTypeObject(
	parentName string,
	objectData map[interface{}]interface{},
	actionNames []string,
) (errs []error) {

	for _, value := range objectData {
		valueString := fmt.Sprintf("%v", value)
		extractedTypes := extractTypes(valueString)

		for _, typeName := range extractedTypes {
			for _, actionName := range actionNames {
				if typeName == actionName {
					errs = append(errs, newValidationErrorTypeNotFound(typeName, parentName))
				}
			}
		}
	}

	return
}
