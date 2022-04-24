package validator

import (
	"fmt"
)

var supportedTypes = []string{
	"bool",
	"string",
	"int64",
	"float64",
}

func isSupportedBasicType(valueString string) bool {
	for _, supportedType := range supportedTypes {
		if valueString == supportedType {
			return true
		}
	}

	return false
}

func validateUsupportedType(data map[interface{}]interface{}) (errs []error) {
	for _, value := range data {
		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateUsupportedTypeObject(mapValue)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateUsupportedTypeObject(
	objectData map[interface{}]interface{},
) (errs []error) {

	for _, value := range objectData {
		valueString := fmt.Sprintf("%v", value)
		valueTypes := extractTypes(valueString)
		if isBasicType(valueTypes[0]) && !isSupportedBasicType(valueTypes[0]) {
			errs = append(errs, newValidationErrorUnsupportedType(valueString))
		}
	}

	return
}
