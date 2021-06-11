package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataIllegalPointerParameter(t *testing.T) {
	t.Run("should not fail on usage of standard types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "int32",
				"ban": "[]int",
				"baz": "*ran",
				"bau": "*ranID",
				"lau": "*[]ranID",
			},
		}

		actualErrors := validateIllegalPointerParameter(data)
		expectedErrors := []error{
			newValidationErrorIllegalPointerParameter("foo", "baz"),
			newValidationErrorIllegalPointerParameter("foo", "bau"),
			newValidationErrorIllegalPointerParameter("foo", "lau"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
