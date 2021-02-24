package validator

import (
	"fmt"
	"go/parser"
	"go/token"
)

func isValidValueString(value string) bool {
	sourceCodeMock := `
	package main
	type a ` + value

	_, err := parser.ParseFile(token.NewFileSet(), "", sourceCodeMock, 0)
	return err == nil
}

func validateInvalidValueString(data map[interface{}]interface{}) (errs []error) {
	for key, value := range data {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			if !isValidValueString(valueString) {
				errs = append(errs, newValidationErrorInvalidValueString(valueString, keyName, "root"))
			}
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateInvalidValueStringObject(mapValue, keyName)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateInvalidValueStringObject(yamlObjectData map[interface{}]interface{}, objectName string) (errs []error) {
	for key, value := range yamlObjectData {
		keyName := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)

		if !isValidValueString(valueString) {
			errs = append(errs, newValidationErrorInvalidValueString(valueString, keyName, objectName))
		}
	}
	return
}
