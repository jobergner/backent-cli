package validator

import (
	"fmt"
	"regexp"
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
	yamlObjectData map[interface{}]interface{},
) (errs []error) {

	for key := range yamlObjectData {
		keyName := fmt.Sprintf("%v", key)
		if startsWithCapitalLetter(keyName) {
			errs = append(errs, newValidationErrorIllegalCapitalization(keyName, literalKindFieldName))
		}
	}

	return
}

func startsWithCapitalLetter(valueString string) bool {
	re := regexp.MustCompile(`[A-Z][A-Za-z]+[0-9]*`)
	return re.MatchString(valueString)
}
