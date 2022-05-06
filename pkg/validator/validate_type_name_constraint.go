package validator

import (
	"fmt"
	"strings"
)

var unavailableTypeNamePrefixes = []string{
	"query",
}

var unavailableTypeNameSuffixes = []string{
	"ID",
}

var unavailableTypeNameSubStrings = []string{
	"_",
}

func violatesTypeNameConstraints(typeName string) bool {
	for _, prefix := range unavailableTypeNamePrefixes {
		if strings.HasPrefix(typeName, prefix) {
			return true
		}
	}

	for _, suffix := range unavailableTypeNameSuffixes {
		if strings.HasSuffix(typeName, suffix) {
			return true
		}
	}

	for _, subString := range unavailableTypeNameSubStrings {
		if strings.Contains(typeName, subString) {
			return true
		}
	}

	return false
}

func validateTypeNameConstraintViolation(data map[interface{}]interface{}) (errs []error) {
	for key := range data {
		keyName := fmt.Sprintf("%v", key)

		if violatesTypeNameConstraints(keyName) {
			errs = append(errs, newValidationErrorTypeNameConstraintViolation(keyName))
		}

	}

	return
}
