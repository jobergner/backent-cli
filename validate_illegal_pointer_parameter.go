package validator

import (
	"fmt"
	"regexp"
)

// validateDirectTypeUsage is used only for thematical validation of action config data
func validateIllegalPointerParameter(data map[interface{}]interface{}) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateIllegalPointerParameterObject(mapValue, keyName)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateIllegalPointerParameterObject(objectData map[interface{}]interface{}, objectName string) (errs []error) {
	for key, value := range objectData {
		keyName := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)
		if hasPointerValue(valueString) {
			errs = append(errs, newValidationErrorIllegalPointerParameter(objectName, keyName))
		}
	}

	return
}

func hasPointerValue(valueString string) bool {
	re := regexp.MustCompile(`\*[A-Za-z]+[0-9]*`)
	return re.MatchString(valueString)
}
