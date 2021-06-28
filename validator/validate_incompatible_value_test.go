package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataIncompatibleValue(t *testing.T) {
	t.Run("should generally fail when values are not compatible", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{},
			"foo": map[interface{}]interface{}{
				"ban": "map[int]string",
				"bal": "[2]int",
				"buf": "*int",
				"luf": "*bar",
				"fan": "[]float64",
				"lan": "bar",
				"boe": "**[]int",
				"bor": "[][]int",
				"yan": "*[]int",
				"wan": "[]*int",
				"qan": "[]*bar",
				"pan": "***int",
			},
		}

		actualErrors := validateIncompatibleValue(data)
		expectedErrors := []error{
			newValidationErrorIncompatibleValue("map[int]string", "ban", "foo"),
			newValidationErrorIncompatibleValue("[2]int", "bal", "foo"),
			newValidationErrorIncompatibleValue("**[]int", "boe", "foo"),
			newValidationErrorIncompatibleValue("[][]int", "bor", "foo"),
			newValidationErrorIncompatibleValue("*[]int", "yan", "foo"),
			newValidationErrorIncompatibleValue("***int", "pan", "foo"),
			newValidationErrorIncompatibleValue("*int", "buf", "foo"),
			newValidationErrorIncompatibleValue("[]*int", "wan", "foo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
