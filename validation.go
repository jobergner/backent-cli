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

func structuralValidation(data map[interface{}]interface{}) (errs []error) {

	valueErrors := validateIllegalValue(data)
	errs = append(errs, valueErrors...)

	return
}

func syntacticalValidation(data map[interface{}]interface{}) (errs []error) {

	illegalTypeNameErrs := validateIllegalTypeName(data)
	errs = append(errs, illegalTypeNameErrs...)

	invalidValueStringErrs := validateInvalidValueString(data)
	errs = append(errs, invalidValueStringErrs...)

	return
}

func logicalValidation(data map[interface{}]interface{}) (errs []error) {

	missingTypeDeclarationErrs := validateTypeNotFound(data)
	errs = append(errs, missingTypeDeclarationErrs...)

	recursiveTypeUsageErrs := validateRecursiveTypeUsage(data)
	errs = append(errs, recursiveTypeUsageErrs...)

	invalidMapKeyErrs := validateIllegalMapKeys(data)
	errs = append(errs, invalidMapKeyErrs...)

	unknownMethodErrs := validateUnknownMethod(data)
	errs = append(errs, unknownMethodErrs...)

	return
}

func thematicalValidation(data map[interface{}]interface{}) (errs []error) {

	nonObjectTypeErrs := validateNonObjectType(data)
	errs = append(errs, nonObjectTypeErrs...)

	capitalizationErrs := validateIllegalCapitalization(data)
	errs = append(errs, capitalizationErrs...)

	incompatibleValueErrs := validateIncompatibleValue(data)
	errs = append(errs, incompatibleValueErrs...)

	conflictingSingularErrs := validateConflictingSingular(data)
	errs = append(errs, conflictingSingularErrs...)

	return
}

func generalValidation(data map[interface{}]interface{}) (errs []error) {

	structuralErrs := structuralValidation(data)
	errs = append(errs, structuralErrs...)
	if len(errs) != 0 {
		return
	}

	syntacticalErrs := syntacticalValidation(data)
	errs = append(errs, syntacticalErrs...)
	if len(errs) != 0 {
		return
	}

	logicalErrs := logicalValidation(data)
	errs = append(errs, logicalErrs...)

	return
}

func Validate(data map[interface{}]interface{}) (errs []error) {
	generalErrs := generalValidation(data)
	errs = append(errs, generalErrs...)

	thematicalErrs := thematicalValidation(data)
	errs = append(errs, thematicalErrs...)

	return
}
