package validator

import (
	"fmt"
)

func validateNonObjectType(data map[interface{}]interface{}) (errs []error) {

	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			errs = append(errs, newValidationErrorNonObjectType(keyName))
		}

	}

	return
}
