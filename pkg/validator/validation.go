package validator

import (
	"fmt"
)

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

// generalValidation validates the given data as if it is actual pseudo go-code
// which means that maps, arrays and non-object types are considered valid as well
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

func thematicalActionsValidation(stateConfigDataVariant, actionsConfigData map[interface{}]interface{}) (errs []error) {
	// prevalidate structure so we can make our assumptions about how the data looks
	nonObjectTypeErrs := validateNonObjectType(actionsConfigData)
	if len(nonObjectTypeErrs) != 0 {
		return nonObjectTypeErrs
	}

	sameNameErrs := validateTypeAndActionWithSameName(stateConfigDataVariant, actionsConfigData)
	errs = append(errs, sameNameErrs...)

	actionUsedAsTypeErrs := validateActionUsedAsType(actionsConfigData)
	errs = append(errs, actionUsedAsTypeErrs...)

	var availableTypeIDs []string
	for keyName := range stateConfigDataVariant {
		typeName := fmt.Sprintf("%v", keyName)
		availableTypeIDs = append(availableTypeIDs, typeName+"ID")
	}

	capitalizationErrs := validateIllegalCapitalization(actionsConfigData)
	errs = append(errs, capitalizationErrs...)

	incompatibleValueErrs := validateIncompatibleValue(actionsConfigData)
	errs = append(errs, incompatibleValueErrs...)

	directTypeUsageErrs := validateDirectTypeUsage(actionsConfigData, availableTypeIDs)
	errs = append(errs, directTypeUsageErrs...)

	pointerParameterErrs := validateIllegalPointerParameter(actionsConfigData)
	errs = append(errs, pointerParameterErrs...)

	return
}

func thematicalStateValidation(data map[interface{}]interface{}) (errs []error) {

	capitalizationErrs := validateIllegalCapitalization(data)
	errs = append(errs, capitalizationErrs...)

	incompatibleValueErrs := validateIncompatibleValue(data)
	errs = append(errs, incompatibleValueErrs...)

	typeNameConstraintViolationErrs := validateTypeNameConstraintViolation(data)
	errs = append(errs, typeNameConstraintViolationErrs...)

	conflictingSingularErrs := validateConflictingSingular(data)
	errs = append(errs, conflictingSingularErrs...)

	unavailableFieldNameErrs := validateUnavailableFieldName(data)
	errs = append(errs, unavailableFieldNameErrs...)

	return
}

func ValidateStateConfig(data map[interface{}]interface{}) []error {
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

	dataCombinations, combinatorErrs := stateConfigCombinationsFrom(dataWithoutMetaFields)
	if len(combinatorErrs) != 0 {
		return combinatorErrs
	}

	var configVariantErrs []error
	for _, configVariant := range dataCombinations {

		generalErrs := generalValidation(configVariant)
		configVariantErrs = append(configVariantErrs, generalErrs...)

		thematicalErrs := thematicalStateValidation(configVariant)
		configVariantErrs = append(configVariantErrs, thematicalErrs...)

	}
	if len(configVariantErrs) != 0 {
		return configVariantErrs
	}

	eventErrs := validateInvalidEventUsage(data)
	if len(eventErrs) != 0 {
		return eventErrs
	}

	return nil
}

func ValidateResponsesConfig(stateConfigData, actionsConfigData, responsesConfigData map[interface{}]interface{}) []error {
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
	if len(responsesAsActionsValidationErrs) != 0 {
		return responsesAsActionsValidationErrs
	}

	responseToUnknownActionErrs := validateResponseToUnknownAction(actionsConfigData, responsesConfigData)
	if len(responseToUnknownActionErrs) != 0 {
		return responseToUnknownActionErrs
	}

	return nil
}

func ValidateActionsConfig(stateConfigData map[interface{}]interface{}, actionsConfigData map[interface{}]interface{}) []error {
	stateConfigErrs := ValidateStateConfig(stateConfigData)
	if len(stateConfigErrs) != 0 {
		return stateConfigErrs
	}

	stateConfigDataWithoutMetaFields := prepareStateConfig(stateConfigData)

	dataCombinations, combinatorErrs := stateConfigCombinationsFrom(stateConfigDataWithoutMetaFields)
	if len(combinatorErrs) != 0 {
		return combinatorErrs
	}

	// use first combination as it does not matter which types of anyOf<> definitions are taken as value
	stateConfigDataVariant := dataCombinations[0]

	generalErrs := generalActionsConfigValidation(stateConfigDataVariant, actionsConfigData)
	if len(generalErrs) != 0 {
		return generalErrs
	}

	thematicalErrs := thematicalActionsValidation(stateConfigDataVariant, actionsConfigData)
	if len(thematicalErrs) != 0 {
		return thematicalErrs
	}

	return nil
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
