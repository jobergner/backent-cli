package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataRecursiveTypeUsage(t *testing.T) {
	t.Run("should not fail on usage of non-recursive types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
			},
			"baz": map[interface{}]interface{}{
				"ban": "*bar",
			},
			"bal": map[interface{}]interface{}{
				"bam": "baz",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail when type is used in own declaration", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "bar",
			"baz": map[interface{}]interface{}{
				"ban": "baz",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"bar", "bar"}),
			newValidationErrorRecursiveTypeUsage([]string{"baz.ban", "baz"}),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of recursive types (1/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "baz",
			},
			"baz": map[interface{}]interface{}{
				"ban": "bar",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"bar.foo", "baz.ban", "bar"}),
			newValidationErrorRecursiveTypeUsage([]string{"baz.ban", "bar.foo", "baz"}),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of recursive types (2/2)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": map[interface{}]interface{}{
				"foo": "bam",
			},
			"baz": map[interface{}]interface{}{
				"ban": "bar",
			},
			"bam": map[interface{}]interface{}{
				"baf": "baz",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"bam.baf", "baz.ban", "bar.foo", "bam"}),
			newValidationErrorRecursiveTypeUsage([]string{"baz.ban", "bar.foo", "bam.baf", "baz"}),
			newValidationErrorRecursiveTypeUsage([]string{"bar.foo", "bam.baf", "baz.ban", "bar"}),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not trigger recursive errors when references are used", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "*foo",
			"bar": "[]bar",
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of directly recursive types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "foo",
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorRecursiveTypeUsage([]string{"foo", "foo"}),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not trigger recursive errors when field has same name as type in object", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"id": "string",
			"person": map[interface{}]interface{}{
				"id": "id",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not trigger recursive errors when field has same name as parent object", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"person": map[interface{}]interface{}{
				"person": "string",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}
