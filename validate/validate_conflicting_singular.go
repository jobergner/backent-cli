package validator

import (
	"fmt"

	"github.com/gertd/go-pluralize"
)

var pluralizeClient *pluralize.Client = pluralize.NewClient()

func validateConflictingSingular(data map[interface{}]interface{}) (errs []error) {

	for _, value := range data {
		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateConflictingSingularObject(mapValue)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateConflictingSingularObject(
	objectData map[interface{}]interface{},
) (errs []error) {

	for key := range objectData {
		keyName := fmt.Sprintf("%v", key)
		for _key := range objectData {
			_keyName := fmt.Sprintf("%v", _key)
			if keyName == _keyName {
				continue
			}
			if left, right := pluralizeClient.Singular(keyName), pluralizeClient.Singular(_keyName); left == right {
				errs = append(errs, newValidationErrorConflictingSingular(_keyName, keyName, left))
			}
		}
	}

	return
}
