package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateYamlDataIllegalTypeName(t *testing.T) {
	t.Run("should not fail on valid key inputs", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"baz": map[interface{}]interface{}{
				"ban": "int",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on spaces in key literal", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fo o": "int",
			"baz": map[interface{}]interface{}{
				"oof":  "int",
				"ba n": "int",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{
			newValidationErrorIllegalTypeName("fo o", "root"),
			newValidationErrorIllegalTypeName("ba n", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not fail on usage of allowed type names", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of keywords", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"break": "int",
			"bar":   "string",
			"baz": map[interface{}]interface{}{
				"const": "int32",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{
			newValidationErrorIllegalTypeName("break", "root"),
			newValidationErrorIllegalTypeName("const", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage special characters", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"*":    "int",
			"<":    "string",
			"fo$o": "int",
			"baz": map[interface{}]interface{}{
				">-":    "int32",
				"bent{": "int32",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{
			newValidationErrorIllegalTypeName("*", "root"),
			newValidationErrorIllegalTypeName("<", "root"),
			newValidationErrorIllegalTypeName("fo$o", "root"),
			newValidationErrorIllegalTypeName(">-", "baz"),
			newValidationErrorIllegalTypeName("bent{", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestIsIllegalTypeName(t *testing.T) {
	t.Run("should return false if the type names are valid", func(t *testing.T) {
		assert.Equal(t, false, isIllegalTypeName("foo_"), isIllegalTypeName("b_ar"), isIllegalTypeName("BA2Z"))
	})
	t.Run("should return true if the type names are illegal", func(t *testing.T) {
		assert.Equal(t, true, isIllegalTypeName("fo o"), isIllegalTypeName("b*ar"), isIllegalTypeName("B+2Z"))
	})
}
