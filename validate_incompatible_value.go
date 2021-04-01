package validator

import (
	"fmt"
	"regexp"
)

func validateIncompatibleValue(data map[interface{}]interface{}) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateIncompatibleValueObject(mapValue, keyName)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateIncompatibleValueObject(
	objectData map[interface{}]interface{},
	objectName string,
) (errs []error) {

	for key, value := range objectData {
		if isString(value) {
			keyName := fmt.Sprintf("%v", key)
			valueString := fmt.Sprintf("%v", value)
			if !isCompatibleValue(valueString) {
				errs = append(
					errs,
					newValidationErrorIncompatibleValue(valueString, keyName, objectName),
				)
			}
		}
	}

	return
}

func isSliceOfType(valueString string) bool {
	re := regexp.MustCompile(`\[\][A-Za-z]+[0-9]*`)
	includesSliceOfType := re.MatchString(valueString)
	if includesSliceOfType && len(re.FindString(valueString)) == len(valueString) {
		return true
	}
	return false
}

// TODO refactor!
func isCompatibleValue(valueString string) bool {
	if isSliceOfType(valueString) {
		return true
	}

	extractedTypes := extractTypes(valueString)
	if len(extractedTypes) != 1 {
		return false
	}

	if len(valueString) != len(extractedTypes[0]) {
		return false
	}

	if isSliceOfSlice(valueString) {
		return false
	}

	return true
}

func isSliceOfSlice(valueString string) bool {
	re := regexp.MustCompile(`\[\]\[\]`)
	return re.MatchString(valueString)
}
