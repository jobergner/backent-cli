package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTypeAndActionWithSameName(t *testing.T) {
	t.Run("should fail on usage of named type", func(t *testing.T) {
		stateData := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
		}
		actionsData := map[interface{}]interface{}{
			"foo": "int",
			"ban": "string",
			"baz": "string",
			"bar": "string",
		}

		actualErrors := validateTypeAndActionWithSameName(stateData, actionsData)
		expectedErrors := []error{
			newValidationErrorTypeAndActionWithSameName("foo"),
			newValidationErrorTypeAndActionWithSameName("bar"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
