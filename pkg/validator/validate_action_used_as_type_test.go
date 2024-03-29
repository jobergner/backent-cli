package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateActionUsedAsType(t *testing.T) {
	t.Run("should fail on usage of two field names with same singular", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "foo",
			},
		}

		actualErrors := validateActionUsedAsType(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("foo", "foo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
