package validator

import (
	"fmt"
	"regexp"
)

// returns errors if types are used which are not declared in the YAML file
// order of declaration is irrelevant
func validateTypeNotFound(yamlData map[interface{}]interface{}) (errs []error) {

	var definedTypes []string

	for key := range yamlData {
		keyName := fmt.Sprintf("%v", key)
		definedTypes = append(definedTypes, keyName)
	}

	for key, value := range yamlData {
		keyName := fmt.Sprintf("%v", key)

		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			extractedTypes := extractTypes(valueString)
			undefinedTypes := findUndefinedTypesIn(extractedTypes, definedTypes)
			for _, undefinedType := range undefinedTypes {
				errs = append(errs, newValidationErrorTypeNotFound(undefinedType, "root"))
			}
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateTypeNotFoundObject(mapValue, keyName, definedTypes)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateTypeNotFoundObject(
	yamlObjectData map[interface{}]interface{},
	objectName string,
	definedTypes []string,
) (errs []error) {

	for _, value := range yamlObjectData {
		if !isString(value) || isEmptyString(value) {
			continue
		}
		valueString := fmt.Sprintf("%v", value)
		extractedTypes := extractTypes(valueString)
		undefinedTypes := findUndefinedTypesIn(extractedTypes, definedTypes)
		for _, undefinedType := range undefinedTypes {
			errs = append(errs, newValidationErrorTypeNotFound(undefinedType, objectName))
		}
	}

	return
}

// extracts all types which are defined in a type definition
// map[string]int => []string{"string", "int"}
func extractTypes(typeDefinitionString string) (extractedTypes []string) {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*`)
	matches := re.FindAllString(typeDefinitionString, -1)
	for _, match := range matches {
		if match == "map" || match == "" {
			continue
		}
		extractedTypes = append(extractedTypes, match)
	}
	return
}

func findUndefinedTypesIn(usedTypes, definedTypes []string) (undefinedTypes []string) {
	allKnownTypes := append(definedTypes, golangBasicTypes...)
	for _, usedType := range usedTypes {
		var isDefined bool
		for _, knownType := range allKnownTypes {
			if knownType == usedType {
				isDefined = true
				break
			}
		}
		if !isDefined {
			undefinedTypes = append(undefinedTypes, usedType)
		}
	}
	return
}
