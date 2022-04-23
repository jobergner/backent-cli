package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataIllegalValue(t *testing.T) {
	t.Run("should not fail on usage of allowed values", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
			},
		}

		actualErrors := validateIllegalValue(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of numbers", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": 1,
			"bar": 1.2,
			"baz": map[interface{}]interface{}{
				"ban": 3,
			},
		}

		actualErrors := validateIllegalValue(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("foo", "root"),
			newValidationErrorIllegalValue("bar", "root"),
			newValidationErrorIllegalValue("ban", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of empty and nil values", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": nil,
			"bar": "",
			"baz": map[interface{}]interface{}{
				"ban": nil,
				"baf": "",
			},
		}

		actualErrors := validateIllegalValue(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("foo", "root"),
			newValidationErrorIllegalValue("bar", "root"),
			newValidationErrorIllegalValue("ban", "baz"),
			newValidationErrorIllegalValue("baf", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of invalid list values", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"mant": []interface{}{},
			},
			"rant": []interface{}{},
		}

		actualErrors := validateIllegalValue(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("mant", "baz"),
			newValidationErrorIllegalValue("rant", "root"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of invalid nested object values", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bant": map[interface{}]interface{}{},
			},
		}

		actualErrors := validateIllegalValue(data)
		expectedErrors := []error{
			newValidationErrorIllegalValue("bant", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
