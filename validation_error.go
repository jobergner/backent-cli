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
