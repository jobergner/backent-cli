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

func thematicalValidationActions(data map[interface{}]interface{}, availableTypeIDs []string) (errs []error) {
	capitalizationErrs := validateIllegalCapitalization(data)
	errs = append(errs, capitalizationErrs...)

	nonObjectTypeErrs := validateNonObjectType(data)
	errs = append(errs, nonObjectTypeErrs...)

	incompatibleValueErrs := validateIncompatibleValue(data)
	errs = append(errs, incompatibleValueErrs...)

	directTypeUsageErrs := validateDirectTypeUsage(data, availableTypeIDs)
	errs = append(errs, directTypeUsageErrs...)

	pointerParameterErrs := validateIllegalPointerParameter(data)
	errs = append(errs, pointerParameterErrs...)

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
	cmb := newAnyOfTypeCombinator(data)
	structuralErrs := cmb.build()
	if len(structuralErrs) != 0 {
		return structuralErrs
	}

	cmb.generateCombinations()

	for _, anyOfTypeCombination := range cmb.dataCombinations {
		validationErrs := validateStateConfig(anyOfTypeCombination)
		errs = append(errs, validationErrs...)
	}

	return
}

func validateStateConfig(data map[interface{}]interface{}) (errs []error) {
	generalErrs := generalValidation(data)
	errs = append(errs, generalErrs...)
	if len(errs) != 0 {
		return
	}

	thematicalErrs := thematicalValidationState(data)
	errs = append(errs, thematicalErrs...)

	return
}

func ValidateActionsConfig(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) (errs []error) {

	sameNameErrs := validateTypeAndActionWithSameName(stateConfigData, actionsConfigData)
	errs = append(errs, sameNameErrs...)

	// join configs and treat them as one. Actions are treated as types, params are treated as fields
	jointConfigData := joinConfigs(stateConfigData, actionsConfigData)

	// collect all availableIDs and insert them into the jointConfigData so the validator cosinders them to be valid types
	var availableTypeIDs []string
	for keyName := range stateConfigData {
		typeName := fmt.Sprintf("%v", keyName)
		idName := typeName + "ID"
		availableTypeIDs = append(availableTypeIDs, idName)
		jointConfigData[idName] = "int"
	}

	// general validation with joint config so validator is aware of all defined types in
	// the state config data when validating types used within actions
	generalErrs := generalValidation(jointConfigData)
	errs = append(errs, generalErrs...)

	// collect all names of actions
	var actionsNames []string
	for keyName := range actionsConfigData {
		definedActionName := fmt.Sprintf("%v", keyName)
		actionsNames = append(actionsNames, definedActionName)
	}

	// pass names of actions as rejectedTypeNames since an action could falsely
	// be used as a type of a parameter
	actionUsedAsTypeErr := validateTypeNotFound(jointConfigData, actionsNames...)
	errs = append(errs, actionUsedAsTypeErr...)

	thematicalErrs := thematicalValidationActions(actionsConfigData, availableTypeIDs)
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
