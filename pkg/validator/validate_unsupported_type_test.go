package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUnsupportedType(t *testing.T) {
	t.Run("should fail on usage of two field names with same singular", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"foot":      "string",
				"feet":      "bar",
				"equipment": "rune",
			},
		}

		actualErrors := validateUsupportedType(data)
		expectedErrors := []error{
			newValidationErrorUnsupportedType("rune"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
