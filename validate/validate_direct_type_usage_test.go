package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataDirectTypeUsage(t *testing.T) {
	t.Run("should not fail on usage of standard types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "int32",
				"ban": "[]int",
				"baz": "ran",
				"bau": "ranID",
			},
			"ran": map[interface{}]interface{}{
				"lan": "[]foo",
				"wan": "[]fooID",
			},
		}

		actualErrors := validateDirectTypeUsage(data, []string{"fooID", "ranID"})
		expectedErrors := []error{
			newValidationErrorDirectTypeUsage("foo", "ran"),
			newValidationErrorDirectTypeUsage("ran", "[]foo"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
