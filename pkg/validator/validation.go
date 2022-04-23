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
	if len(valueErrors) != 0 {
		return valueErrors
	}

	return nil
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
	// prevalidate structure so we can make our assumptions about how the data looks
	structuralErrs := structuralValidation(data)
	if len(structuralErrs) != 0 {
		return structuralErrs
	}

	nonObjectTypeErrs := validateNonObjectType(data)
	if len(nonObjectTypeErrs) != 0 {
		return nonObjectTypeErrs
	}

	dataWithoutMetaFields := prepareStateConfig(data)

	dataCombinations, prevalidationErrs := stateConfigCombinationsFrom(dataWithoutMetaFields)
	if len(prevalidationErrs) != 0 {
		return prevalidationErrs
	}

	for _, anyOfTypeCombination := range dataCombinations {
		validationErrs := validateStateConfig(anyOfTypeCombination)
		errs = append(errs, validationErrs...)
	}

	if len(errs) != 0 {
		return errs
	}

	eventErrs := validateInvalidEventUsage(data)
	errs = append(errs, eventErrs...)

	return deduplicateErrs(errs)
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

func ValidateResponsesConfig(stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{}) (errs []error) {
	// prevalidate structure so we can make our assumptions about how the data looks
	structuralErrs := structuralValidation(stateConfigData)
	if len(structuralErrs) != 0 {
		return structuralErrs
	}

	nonObjectTypeErrs := validateNonObjectType(stateConfigData)
	if len(nonObjectTypeErrs) != 0 {
		return nonObjectTypeErrs
	}

	stateConfigDataWithoutEvents := prepareStateConfig(stateConfigData)

	// responses and action share the same restrictions/requirements
	responsesAsActionsValidationErrs := ValidateActionsConfig(stateConfigDataWithoutEvents, responsesConfigData)
	errs = append(errs, responsesAsActionsValidationErrs...)

	responseToUnknownActionErrs := validateResponseToUnknownAction(actionsConfigData, responsesConfigData)
	errs = append(errs, responseToUnknownActionErrs...)

	return
}

func ValidateActionsConfig(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) (errs []error) {
	errs = ValidateStateConfig(stateConfigData)
	if len(errs) != 0 {
		return errs
	}

	stateConfigDataWithoutMetaFields := prepareStateConfig(stateConfigData)

	dataCombinations, prevalidationErrs := stateConfigCombinationsFrom(stateConfigDataWithoutMetaFields)
	if len(prevalidationErrs) != 0 {
		return prevalidationErrs
	}

	// use first combination as it does not matter which types of anyOf<> definitions are taken as value
	stateConfigDataVariant := dataCombinations[0]

	sameNameErrs := validateTypeAndActionWithSameName(stateConfigDataVariant, actionsConfigData)
	errs = append(errs, sameNameErrs...)

	generalErrs := generalActionsConfigValidation(stateConfigDataVariant, actionsConfigData)
	errs = append(errs, generalErrs...)

	actionUsedAsTypeErrs := validateActionUsedAsType(actionsConfigData)
	errs = append(errs, actionUsedAsTypeErrs...)

	var availableTypeIDs []string
	for keyName := range stateConfigData {
		typeName := fmt.Sprintf("%v", keyName)
		availableTypeIDs = append(availableTypeIDs, typeName+"ID")
	}

	thematicalErrs := thematicalValidationActions(actionsConfigData, availableTypeIDs)
	errs = append(errs, thematicalErrs...)

	return
}

func generalActionsConfigValidation(stateConfigData, actionsConfigData map[interface{}]interface{}) []error {

	actionsConfigCopy := copyData(actionsConfigData)

	for keyName := range stateConfigData {
		typeName := fmt.Sprintf("%v", keyName)
		idName := typeName + "ID"
		actionsConfigCopy[typeName] = "int"
		actionsConfigCopy[idName] = "int"
	}

	// general validation with modified actions data so validator is aware of all defined type IDs in
	// the state config data when validating types used within actions
	return generalValidation(actionsConfigCopy)
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
