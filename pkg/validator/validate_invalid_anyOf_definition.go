package validator

import (
	"sort"
	"strings"
)

func validateAnyOfDefinition(valueString string) error {

	containedTypes := extractTypes(valueString)[1:] // extractTypes considers anyOf identifier as type, so we cut it

	if len(containedTypes) < 2 {
		return newValidationErrorInvalidAnyOfDefinition(valueString)
	}

	duplicateCheck := make(map[string]bool)
	for _, containedType := range containedTypes {
		if duplicateCheck[containedType] {
			return newValidationErrorInvalidAnyOfDefinition(valueString)
		}
		duplicateCheck[containedType] = true
	}

	containedTypesCopy := make([]string, len(containedTypes))
	copy(containedTypesCopy, containedTypes)
	sort.Slice(containedTypesCopy, caseInsensitiveSort(containedTypesCopy))
	for i, containedType := range containedTypes {
		containedTypeCopy := containedTypesCopy[i]
		if containedTypeCopy != containedType {
			return newValidationErrorInvalidAnyOfDefinition(valueString)
		}
	}

	return nil
}

func caseInsensitiveSort(keys []string) func(i, j int) bool {
	return func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	}
}
