package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateYamlInvalidValueString(t *testing.T) {
	t.Run("should not fail on usage of allowed values strings", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "map[int]string",
			"baz": map[interface{}]interface{}{
				"ban": "[]int32",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of '['/']' in the wrong places", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"a":   "[]string",
			"b":   "map[int]string]",
			"foo": "int[]",
			"bar": "[]map[int]string",
			"bu":  "bar",
			"baz": map[interface{}]interface{}{
				"ban": "[]in[t32",
				"fan": "[]",
				"c":   "map[int][]string",
				"d":   "map[int]string*",
				"e":   "[*]int32",
				"f":   "int*",
				"g":   "[]in t32",
				"h":   " ",
				"i":   "[]in@t32",
				"j":   "foo",
			},
		}

		actualErrors := syntacticalValidation(data)
		expectedErrors := []error{
			newValidationErrorInvalidValueString("map[int]string]", "b", "root"),
			newValidationErrorInvalidValueString("int[]", "foo", "root"),
			newValidationErrorInvalidValueString("[]in[t32", "ban", "baz"),
			newValidationErrorInvalidValueString("[]", "fan", "baz"),
			newValidationErrorInvalidValueString("map[int]string*", "d", "baz"),
			newValidationErrorInvalidValueString("[*]int32", "e", "baz"),
			newValidationErrorInvalidValueString("int*", "f", "baz"),
			newValidationErrorInvalidValueString("[]in t32", "g", "baz"),
			newValidationErrorInvalidValueString(" ", "h", "baz"),
			newValidationErrorInvalidValueString("[]in@t32", "i", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestIsValidValueString(t *testing.T) {
	t.Run("is valid value", func(t *testing.T) {
		assert.Equal(t, isValidValueString("map[int]string"), true)
		assert.Equal(t, isValidValueString("[]int32"), true)
		assert.Equal(t, isValidValueString("int"), true)
		assert.Equal(t, isValidValueString("int*"), false)
		assert.Equal(t, isValidValueString("in+t"), false)
		assert.Equal(t, isValidValueString("map[int]st&ring"), false)
		assert.Equal(t, isValidValueString("[]in@t32"), false)
		assert.Equal(t, isValidValueString("@"), false)
		assert.Equal(t, isValidValueString("in t"), false)
		assert.Equal(t, isValidValueString(" "), false)
		assert.Equal(t, isValidValueString("[]in t32"), false)
		assert.Equal(t, isValidValueString("map[int]*string"), true)
		assert.Equal(t, isValidValueString("*string"), true)
		assert.Equal(t, isValidValueString("map[int*]string"), false)
		assert.Equal(t, isValidValueString("[*]int32"), false)
		assert.Equal(t, isValidValueString("*"), false)
		assert.Equal(t, isValidValueString("map[int]string*"), false)
		assert.Equal(t, isValidValueString("["), false)
		assert.Equal(t, isValidValueString("]"), false)
		assert.Equal(t, isValidValueString("]["), false)
		assert.Equal(t, isValidValueString("[]string"), true)
		assert.Equal(t, isValidValueString("map[int]string]"), false)
		assert.Equal(t, isValidValueString("[]map[int]string"), true)
		assert.Equal(t, isValidValueString("**[]map[**int]****string"), true)
		assert.Equal(t, isValidValueString("int[]"), false)
		assert.Equal(t, isValidValueString("[]in[t32"), false)
		assert.Equal(t, isValidValueString("[]"), false)
		assert.Equal(t, isValidValueString("foo"), true)
		assert.Equal(t, isValidValueString("bar"), true)
	})
}
