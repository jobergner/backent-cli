package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateYamlDataUnknwonMethod(t *testing.T) {
	t.Run("should generally fail when values are literals sequenced with dots", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int.bar",
			"bar": "float64.foo",
			"baz": map[interface{}]interface{}{
				"ban": "foo.bar",
				"bal": "string.int.bool",
				"buf": "map[foo.int]string",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorUnknownMethod("int", "bar"),
			newValidationErrorUnknownMethod("float64", "foo"),
			newValidationErrorUnknownMethod("foo", "bar"),
			newValidationErrorUnknownMethod("string", "int"),
			newValidationErrorUnknownMethod("foo", "int"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
