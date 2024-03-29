package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataUnavailableFieldName(t *testing.T) {
	t.Run("should not fail on usage of available field names", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar":           "[]int32",
				"ban":           "int",
				"operationKind": "string",
				"path":          "int",
				"queryFoo":      "int",
			},
		}

		actualErrors := validateUnavailableFieldName(data)
		expectedErrors := []error{
			newValidationErrorUnavailableFieldName("operationKind"),
			newValidationErrorUnavailableFieldName("path"),
			newValidationErrorUnavailableFieldName("queryFoo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

}
