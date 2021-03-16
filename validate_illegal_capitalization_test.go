package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataIllegalCapitalization(t *testing.T) {
	t.Run("should fail on usage of illegally capitalized field names or types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"Bar": "string",
			},
			"Baz": map[interface{}]interface{}{
				"ban": "int",
				"Bal": "string",
				"buf": "bool",
			},
		}

		actualErrors := thematicalValidation(data, false, true)
		expectedErrors := []error{
			newValidationErrorIllegalCapitalization("Bar", literalKindFieldName),
			newValidationErrorIllegalCapitalization("Baz", literalKindType),
			newValidationErrorIllegalCapitalization("Bal", literalKindFieldName),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
