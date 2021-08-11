package validator

import (
	"fmt"
	"regexp"
)

func validateUnknownMethod(data map[interface{}]interface{}) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			if hasDotAccessedMethod(valueString) {
				errs = append(
					errs,
					newValidationErrorUnknownMethod(extractFirstLiteralBeforeDot(valueString), extractFirstLiteralAfterDot(valueString)),
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
	objectData map[interface{}]interface{},
	objectName string,
) (errs []error) {

	for _, value := range objectData {
		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			if hasDotAccessedMethod(valueString) {
				errs = append(
					errs,
					newValidationErrorUnknownMethod(extractFirstLiteralBeforeDot(valueString), extractFirstLiteralAfterDot(valueString)),
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

func extractFirstLiteralAfterDot(valueString string) string {
	re := regexp.MustCompile(`\.[A-Za-z]+[0-9]*`)
	return re.FindAllString(valueString, 1)[0][1:]
}

func extractFirstLiteralBeforeDot(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*\.`)
	match := re.FindAllString(valueString, 1)[0]
	return match[:len(match)-1]
}
