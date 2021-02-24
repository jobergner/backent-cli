package validator

import (
	"fmt"
	"regexp"
)

func validateUnknownMethod(yamlData map[interface{}]interface{}) (errs []error) {

	for key, value := range yamlData {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			if hasDotAccessedMethod(valueString) {
				errs = append(
					errs,
					newValidationErrorUnknownMethod(ExtractFirstLiteralBeforeDot(valueString), ExtractFirstLiteralAfterDot(valueString)),
				)
			}
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateUnknownMethodObject(mapValue, keyName)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateUnknownMethodObject(
	yamlObjectData map[interface{}]interface{},
	objectName string,
) (errs []error) {

	for _, value := range yamlObjectData {
		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			if hasDotAccessedMethod(valueString) {
				errs = append(
					errs,
					newValidationErrorUnknownMethod(ExtractFirstLiteralBeforeDot(valueString), ExtractFirstLiteralAfterDot(valueString)),
				)
			}
		}
	}

	return
}

func hasDotAccessedMethod(valueString string) bool {
	re := regexp.MustCompile(`\.[A-Za-z]+[0-9]*`)
	return re.MatchString(valueString)
}

func ExtractFirstLiteralAfterDot(valueString string) string {
	re := regexp.MustCompile(`\.[A-Za-z]+[0-9]*`)
	return re.FindAllString(valueString, 1)[0][1:]
}

func ExtractFirstLiteralBeforeDot(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*\.`)
	match := re.FindAllString(valueString, 1)[0]
	return match[:len(match)-1]
}
