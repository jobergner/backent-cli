package validator

import (
	"fmt"
	"reflect"
)

var golangBasicTypes = []string{"string", "bool", "int8", "uint8", "byte", "int16", "uint16", "int32", "rune", "uint32", "int64", "uint64", "int", "uint", "uintptr", "float32", "float64", "complex64", "complex128"}

const mockPackageName string = "foobar"

func isString(unknown interface{}) bool {
	v := reflect.ValueOf(unknown)
	if v.Kind() == reflect.String {
		return true
	}
	return false
}

func isSlice(unknown interface{}) bool {
	v := reflect.ValueOf(unknown)
	if v.Kind() == reflect.Slice {
		return true
	}
	return false
}

func isMap(unknown interface{}) bool {
	v := reflect.ValueOf(unknown)
	if v.Kind() == reflect.Map {
		return true
	}
	return false
}

func isNil(unknown interface{}) bool {
	return unknown == nil
}

func isEmptyString(unknown interface{}) bool {
	if !isString(unknown) {
		return false
	}
	valueString := fmt.Sprintf("%v", unknown)
	return valueString == ""
}

func structuralValidation(yamlData map[interface{}]interface{}) (errs []error) {

	valueErrors := validateIllegalValue(yamlData)
	errs = append(errs, valueErrors...)

	return
}

func syntacticalValidation(yamlData map[interface{}]interface{}) (errs []error) {

	illegalTypeNameErrs := validateIllegalTypeName(yamlData)
	errs = append(errs, illegalTypeNameErrs...)

	invalidValueStringErrs := validateInvalidValueString(yamlData)
	errs = append(errs, invalidValueStringErrs...)

	return
}

func logicalValidation(yamlData map[interface{}]interface{}) (errs []error) {

	missingTypeDeclarationErrs := validateTypeNotFound(yamlData)
	errs = append(errs, missingTypeDeclarationErrs...)

	recursiveTypeUsageErrs := validateRecursiveTypeUsage(yamlData)
	errs = append(errs, recursiveTypeUsageErrs...)

	invalidMapKeyErrs := validateIllegalMapKeys(yamlData)
	errs = append(errs, invalidMapKeyErrs...)

	unknownMethodErrs := validateUnknownMethod(yamlData)
	errs = append(errs, unknownMethodErrs...)

	return
}

func thematicalValidation(yamlData map[interface{}]interface{}) (errs []error) {

	nonObjectTypeErrs := validateNonObjectType(yamlData)
	errs = append(errs, nonObjectTypeErrs...)

	if len(errs) != 0 {
		return
	}

	return validateIncompatibleValue(yamlData)
}

func validateYamlData(yamlData map[interface{}]interface{}) (errs []error) {

	structuralErrs := structuralValidation(yamlData)
	errs = append(errs, structuralErrs...)
	if len(errs) != 0 {
		return
	}

	syntacticalErrs := syntacticalValidation(yamlData)
	errs = append(errs, syntacticalErrs...)
	if len(errs) != 0 {
		return
	}

	logicalErrs := logicalValidation(yamlData)
	errs = append(errs, logicalErrs...)

	return
}
