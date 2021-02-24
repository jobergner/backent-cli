package validator

import (
	"fmt"
)

func validateNonObjectType(yamlData map[interface{}]interface{}) (errs []error) {

	for key, value := range yamlData {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			errs = append(errs, newValidationErrorNonObjectType(keyName))
		}

	}

	return
}
