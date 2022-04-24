package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataInvalidAnyOfDefinition(t *testing.T) {
	t.Run("returns nil on valid anyOf definition", func(t *testing.T) {
		input := "anyOf< bar, baz,foo>"
		assert.Equal(t, nil, validateAnyOfDefinition(input))
		input = "*anyOf< bar, baz,foo>"
		assert.Equal(t, nil, validateAnyOfDefinition(input))
	})
	t.Run("returns error on too few types in definition", func(t *testing.T) {
		input := "anyOf<foo>"
		assert.Equal(t, newValidationErrorInvalidAnyOfDefinition(input), validateAnyOfDefinition(input))
	})
	t.Run("returns error on no types in definition", func(t *testing.T) {
		input := "anyOf<>"
		assert.Equal(t, newValidationErrorInvalidAnyOfDefinition(input), validateAnyOfDefinition(input))
	})
	t.Run("returns error on duplicate type in definition", func(t *testing.T) {
		input := "anyOf<foo , bar, foo >"
		assert.Equal(t, newValidationErrorInvalidAnyOfDefinition(input), validateAnyOfDefinition(input))
	})
	t.Run("returns error on anyOf types not being in alphabetical order", func(t *testing.T) {
		input := "anyOf<foo,bar>"
		assert.Equal(t, newValidationErrorInvalidAnyOfDefinition(input), validateAnyOfDefinition(input))
	})
}
