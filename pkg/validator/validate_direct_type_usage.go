package validator

import (
	"fmt"
)

// validateDirectTypeUsage is used only for thematical validation of action config data
func validateDirectTypeUsage(data map[interface{}]interface{}, typeIDs []string) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateDirectTypeUsageObject(mapValue, keyName, typeIDs)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateDirectTypeUsageObject(objectData map[interface{}]interface{}, objectName string, typeIDs []string) (errs []error) {
	for _, value := range objectData {
		valueString := fmt.Sprintf("%v", value)
		extractedTypes := extractTypes(valueString)
		// at this point it is known that there has to be exactly 1 type in each valueString
		// if it is not a basic type or any of the known IDs it has to be an already validated user-defined type
		if !isBasicType(extractedTypes[0]) && !contains(typeIDs, extractedTypes[0]) {
			errs = append(errs, newValidationErrorDirectTypeUsage(objectName, valueString))
		}
	}

	return
}
