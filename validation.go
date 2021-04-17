package validator

import (
	"errors"
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

func thematicalValidationActions(data map[interface{}]interface{}) (errs []error) {
	capitalizationErrs := validateIllegalCapitalization(data)
	errs = append(errs, capitalizationErrs...)

	nonObjectTypeErrs := validateNonObjectType(data)
	errs = append(errs, nonObjectTypeErrs...)

	incompatibleValueErrs := validateIncompatibleValue(data)
	errs = append(errs, incompatibleValueErrs...)

	return
}

func thematicalValidationState(data map[interface{}]interface{}) (errs []error) {

	nonObjectTypeErrs := validateNonObjectType(data)
	errs = append(errs, nonObjectTypeErrs...)

	capitalizationErrs := validateIllegalCapitalization(data)
	errs = append(errs, capitalizationErrs...)

	incompatibleValueErrs := validateIncompatibleValue(data)
	errs = append(errs, incompatibleValueErrs...)

	conflictingSingularErrs := validateConflictingSingular(data)
	errs = append(errs, conflictingSingularErrs...)

	unavailableFieldNameErrs := validateUnavailableFieldName(data)
	errs = append(errs, unavailableFieldNameErrs...)

	return
}

func ValidateStateConfig(data map[interface{}]interface{}) (errs []error) {
	generalErrs := generalValidation(data)
	errs = append(errs, generalErrs...)

	thematicalErrs := thematicalValidationState(data)
	errs = append(errs, thematicalErrs...)

	return
}

func ValidateActionsConfig(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) (errs []error) {

	stateConfigErrs := ValidateStateConfig(stateConfigData)
	errs = append(errs, stateConfigErrs...)

	// join configs and treat them as one. Actions are treated as types, params are treated as fields
	jointConfigData := joinConfigs(stateConfigData, actionsConfigData)

	// general validation with joint config so validator is aware of all defined types in
	// the state config data when validating types used within actions
	generalErrs := generalValidation(jointConfigData)
	errs = append(errs, generalErrs...)

	// collect all names of actions
	var actionsNames []string
	for _definedActionName := range actionsConfigData {
		definedActionName := fmt.Sprintf("%v", _definedActionName)
		actionsNames = append(actionsNames, definedActionName)
	}

	// pass names of actions as rejectedTypeNames since an action could falsely
	// be used as a type of a parameter
	actionUsedAsTypeErr := validateTypeNotFound(jointConfigData, actionsNames...)
	errs = append(errs, actionUsedAsTypeErr...)

	thematicalErrs := thematicalValidationActions(jointConfigData)
	errs = append(errs, thematicalErrs...)

	// deduplicate errors as stateConfigData is being validated twice (ValidateStateConfig, generalValidation)
	return deduplicateErrs(errs)
}

func deduplicateErrs(errs []error) []error {

	check := make(map[string]bool)
	deduped := make([]error, 0)
	for _, val := range errs {
		check[val.Error()] = true
	}

	for val := range check {
		deduped = append(deduped, errors.New(val))
	}

	return deduped
}

func joinConfigs(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) map[interface{}]interface{} {
	jointConfigData := make(map[interface{}]interface{})
	for k, v := range stateConfigData {
		jointConfigData[k] = v
	}
	for k, v := range actionsConfigData {
		jointConfigData[k] = v
	}
	return jointConfigData
}
