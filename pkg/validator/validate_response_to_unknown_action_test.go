package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateResponseToUnknownAction(t *testing.T) {
	t.Run("should fail on usage of named type", func(t *testing.T) {
		actionsData := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
		}
		responsesData := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": "string",
		}

		actualErrors := validateResponseToUnknownAction(actionsData, responsesData)
		expectedErrors := []error{
			newValidationErrorResponseToUnknownAction("baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
