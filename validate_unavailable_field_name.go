package validator

import (
	"fmt"
)

var unavailableFiledNames = []string{
	"id",
	"iD",
	"operationKind_",
	"hasParent_",
}

func isUnavailableFiledName(fieldName string) bool {
	for _, unavailableFieldName := range unavailableFiledNames {
		if unavailableFieldName == fieldName {
			return true
		}
	}
	return false
}

func validateUnavailableFieldName(data map[interface{}]interface{}) (errs []error) {
	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateUnavailableFieldNameObject(mapValue, keyName)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateUnavailableFieldNameObject(objectData map[interface{}]interface{}, objectName string) (errs []error) {
	for key := range objectData {
		keyName := fmt.Sprintf("%v", key)
		if isUnavailableFiledName(keyName) {
			errs = append(errs, newValidationErrorUnavailableFieldName(keyName))
		}
	}
	return
}
