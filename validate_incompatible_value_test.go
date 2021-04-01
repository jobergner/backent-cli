package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataIncompatibleMethod(t *testing.T) {
	t.Run("should generally fail when values are not compatible", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{},
			"foo": map[interface{}]interface{}{
				"ban": "map[int]string",
				"bal": "[2]int",
				"buf": "*int",
				"fan": "[]float64",
				"lan": "bar",
				"boe": "**[]int",
			},
		}

		actualErrors := validateIncompatibleValue(data)
		expectedErrors := []error{
			newValidationErrorIncompatibleValue("map[int]string", "ban", "foo"),
			newValidationErrorIncompatibleValue("[2]int", "bal", "foo"),
			newValidationErrorIncompatibleValue("*int", "buf", "foo"),
			newValidationErrorIncompatibleValue("**[]int", "boe", "foo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
