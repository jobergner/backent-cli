package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTypeNameConstraint(t *testing.T) {
	t.Run("should fail on violating type names", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"queryBan": map[interface{}]interface{}{},
			"foo":      map[interface{}]interface{}{},
			"fooID":    map[interface{}]interface{}{},
			"bar_baz":  map[interface{}]interface{}{},
		}

		actualErrors := validateTypeNameConstraintViolation(data)
		expectedErrors := []error{
			newValidationErrorTypeNameConstraintViolation("queryBan"),
			newValidationErrorTypeNameConstraintViolation("fooID"),
			newValidationErrorTypeNameConstraintViolation("bar_baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
