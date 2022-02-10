package validator

import (
	"fmt"
)

func validateInvalidEventUsage(data map[interface{}]interface{}) (errs []error) {

	for _, value := range data {

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateInvalidEventUsageObject(data, mapValue)
			errs = append(errs, objectValidationErrs...)
		}

	}

	return
}

func validateInvalidEventUsageObject(data, objectData map[interface{}]interface{}) (errs []error) {

	for key, value := range objectData {
		keyName := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)

		// field iself is an event meta field
		if keyName == eventMetaField.name && valueString == eventMetaField.value {
			continue
		}

		extractedTypes := extractTypes(valueString)
		if isAnyOfTypes(valueString) {
			extractedTypes = extractedTypes[1:] // extractTypes considers anyOf identifier as type, so we cut it
		}

		var hasEventValue bool
		var hasNonEventValue bool

		for _, extractedType := range extractedTypes {
			if isBasicType(extractedType) {
				continue
			}

			typeFields := data[extractedType]

			if !isEvent(typeFields.(map[interface{}]interface{})) {
				hasNonEventValue = true
				continue
			}
			hasEventValue = true

		}

		if hasEventValue && hasNonEventValue {
			errs = append(errs, newValidationErrorUnpureEvent(valueString))
		}

		if hasEventValue {
			if !hasSliceValue(valueString) {
				errs = append(errs, newValidationErrorInvalidEventUsage(valueString))
				continue
			}

			if hasPointerValue(valueString) {
				errs = append(errs, newValidationErrorInvalidEventUsage(valueString))
			}
		}

	}

	return
}

func isEvent(typeFields map[interface{}]interface{}) bool {

	for key, value := range typeFields {
		keyName := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)

		if keyName == eventMetaField.name && valueString == eventMetaField.value {
			return true
		}
	}

	return false
}
