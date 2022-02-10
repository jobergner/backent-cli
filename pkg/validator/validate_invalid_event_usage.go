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
		// only one type is expected at this point
		if isBasicType(extractedTypes[0]) {
			continue
		}

		typeFields := data[extractedTypes[0]]

		if !isEvent(typeFields.(map[interface{}]interface{})) {
			continue
		}

		// at this point we know this field's valueString contains an event
		if !hasSliceValue(valueString) {
			errs = append(errs, newValidationErrorInvalidEventUsage(valueString))
			continue
		}

		if hasPointerValue(valueString) {
			errs = append(errs, newValidationErrorInvalidEventUsage(valueString))
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
