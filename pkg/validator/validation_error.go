package validator

import (
	"errors"
	"fmt"
	"strings"
)

type literalKind string

const (
	literalKindType      literalKind = "type"
	literalKindFieldName             = "field name"
)

func newValidationErrorTypeNotFound(missingTypeLiteral, parentItemName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrTypeNotFound: type with name \"%s\" in \"%s\" was not found",
			missingTypeLiteral,
			parentItemName,
		),
	)
}
func newValidationErrorIllegalValue(keyName, parentItemName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIllegalValue: value assigned to key \"%s\" in \"%s\" is invalid",
			keyName,
			parentItemName,
		),
	)
}
func newValidationErrorInvalidValueString(valueString, keyName, parentItemName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrInvalidValueString: value \"%s\" assigned to \"%s\" in \"%s\" is invalid",
			valueString,
			keyName,
			parentItemName,
		),
	)
}
func newValidationErrorIllegalTypeName(keyName, parentItemName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIllegalTypeName: illegal type name \"%s\" in \"%s\"",
			keyName,
			parentItemName,
		),
	)
}
func newValidationErrorRecursiveTypeUsage(keysResultingInRecursiveness []string) error {
	keys := strings.Join(keysResultingInRecursiveness, "->")
	return errors.New(
		fmt.Sprintf(
			"ErrRecursiveTypeUsage: illegal recursive type detected for \"%s\"",
			keys,
		),
	)
}
func newValidationErrorInvalidMapKey(mapKey, valueString string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrInvalidMapKey: \"%s\" in \"%s\" is not a valid map key",
			mapKey,
			valueString,
		),
	)
}
func newValidationErrorUnknownMethod(typeName, unknownMethod string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIllegalValue: type \"%s\" has no method \"%s\"",
			typeName,
			unknownMethod,
		),
	)
}
func newValidationErrorNonObjectType(keyName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrNonObjectType: type \"%s\" is not an object type",
			keyName,
		),
	)
}
func newValidationErrorIncompatibleValue(valueString, keyName, parentItemName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIncompatibleValue: value \"%s\" assigned to \"%s\" in \"%s\" is incompatible",
			valueString,
			keyName,
			parentItemName,
		),
	)
}
func newValidationErrorIllegalCapitalization(literal string, literalKind literalKind) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIllegalCapitalization: %s \"%s\" starts with a capital letter",
			literalKind,
			literal,
		),
	)
}
func newValidationErrorConflictingSingular(keyName1, keyName2, singularForm string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrConflictingSingular: \"%s\" and \"%s\" share the same singular form \"%s\"",
			keyName1,
			keyName2,
			singularForm,
		),
	)
}
func newValidationErrorUnavailableFieldName(keyName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrUnavailableFieldName: \"%s\" not an available name",
			keyName,
		),
	)
}
func newValidationErrorDirectTypeUsage(actionName, typeName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrDirectTypeUsage: the type \"%s\" was used directly in \"%s\" instead of it's ID (\"%sID\")",
			typeName,
			actionName,
			typeName,
		),
	)
}
func newValidationErrorIllegalPointerParameter(typeName, fieldName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrIllegalPointerParameter: the parameter \"%s\" in \"%s\" contains a pointer value",
			fieldName,
			typeName,
		),
	)
}
func newValidationErrorTypeAndActionWithSameName(name string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrTypeAndActionWithSameName: type and action \"%s\" have the same name",
			name,
		),
	)
}
func newValidationErrorInvalidAnyOfDefinition(valueString string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrInvalidAnyOfDefinition: \"%s\" is not a valid `anyOf` definition",
			valueString,
		),
	)
}
func newValidationErrorResponseToUnknownAction(responseName string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrResponeToUnknownAction: there is no action defined for response \"%s\"",
			responseName,
		),
	)
}
func newValidationErrorInvalidEventUsage(valueString string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrInvalidEventUsage: \"%s\" is not a valid usage of an event (events can only be slices of non-pointers)",
			valueString,
		),
	)
}
func newValidationErrorUnpureEvent(valueString string) error {
	return errors.New(
		fmt.Sprintf(
			"ErrUnpureEvent: \"%s\" contains event types and non-event types",
			valueString,
		),
	)
}
