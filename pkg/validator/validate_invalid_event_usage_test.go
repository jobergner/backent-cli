package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInvalidEventUsage(t *testing.T) {
	t.Run("should not fail on correct event usage", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fooEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"baz": map[interface{}]interface{}{
				"ban": "[]fooEvent",
			},
		}

		actualErrors := validateInvalidEventUsage(data)

		assert.Empty(t, actualErrors)
	})

	t.Run("should fail on event reference or event non-slice", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fooNonEvent": map[interface{}]interface{}{
				"bar": "string",
			},
			"fooEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"baz": map[interface{}]interface{}{
				"nui": "*fooNonEvent",
				"ban": "fooEvent",
				"buf": "*fooEvent",
				"bal": "[]*fooEvent",
				"oof": "[]fooEvent",
			},
		}

		actualErrors := validateInvalidEventUsage(data)
		expectedErrors := []error{
			newValidationErrorInvalidEventUsage("*fooEvent"),
			newValidationErrorInvalidEventUsage("[]*fooEvent"),
			newValidationErrorInvalidEventUsage("fooEvent"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
