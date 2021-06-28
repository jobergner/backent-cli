package validator

import (
	"fmt"
	"strings"
)

func validateIllegalCapitalization(data map[interface{}]interface{}) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if startsWithCapitalLetter(keyName) {
			errs = append(errs, newValidationErrorIllegalCapitalization(keyName, literalKindType))
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateIllegalCapitalizationObject(mapValue)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateIllegalCapitalizationObject(
	objectData map[interface{}]interface{},
) (errs []error) {

	for key := range objectData {
		keyName := fmt.Sprintf("%v", key)
		if startsWithCapitalLetter(keyName) {
			errs = append(errs, newValidationErrorIllegalCapitalization(keyName, literalKindFieldName))
		}
	}

	return
}

func startsWithCapitalLetter(keyName string) bool {
	if keyName == strings.Title(keyName) {
		return true
	}
	return false
}
