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
			"barEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"baz": map[interface{}]interface{}{
				"nui": "*fooNonEvent",
				"ban": "fooEvent",
				"buf": "*fooEvent",
				"bal": "[]*fooEvent",
				"oof": "[]fooEvent",
				"bun": "anyOf<barEvent,fooEvent>",
			},
		}

		actualErrors := validateInvalidEventUsage(data)
		expectedErrors := []error{
			newValidationErrorInvalidEventUsage("*fooEvent"),
			newValidationErrorInvalidEventUsage("[]*fooEvent"),
			newValidationErrorInvalidEventUsage("fooEvent"),
			newValidationErrorInvalidEventUsage("anyOf<barEvent,fooEvent>"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail when events and non-events are mixed in valueString", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fooNonEvent": map[interface{}]interface{}{
				"bar": "string",
			},
			"fooEvent": map[interface{}]interface{}{
				"__event__": "true",
				"bar":       "string",
			},
			"baz": map[interface{}]interface{}{
				"ban": "[]anyOf<fooEvent,fooNonEvent>",
			},
		}

		actualErrors := validateInvalidEventUsage(data)
		expectedErrors := []error{
			newValidationErrorUnpureEvent("[]anyOf<fooEvent,fooNonEvent>"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
