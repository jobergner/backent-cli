package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataNonObjectType(t *testing.T) {
	t.Run("should fail on usage of named type", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
			},
		}

		actualErrors := validateNonObjectType(data)
		expectedErrors := []error{
			newValidationErrorNonObjectType("foo"),
			newValidationErrorNonObjectType("bar"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
