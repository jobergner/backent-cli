package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateYamlDataTypeNotFound(t *testing.T) {
	t.Run("should not fail on usage of standard types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baf": "[]string",
			"bal": "map[string]int",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bunt": "[]int",
				"bap":  "map[int16]string",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not fail on usage of declared types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"bar": "string",
			"baf": "[]foo",
			"bal": "map[foo]bar",
			"bum": "*int",
			"baz": map[interface{}]interface{}{
				"ban":  "int32",
				"bam":  "bar",
				"bunt": "[]baf",
				"bap":  "map[bar]foo",
				"bal":  "***bar",
				"slap": "**[]**baf",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of types declared in object", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
				"bar": "ban",
			},
			"boo": "ban",
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("ban", "baz"),
			newValidationErrorTypeNotFound("ban", "root"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of unknown types", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"fof": "schtring",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
				"bam": "bar",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("schtring", "root"),
			newValidationErrorTypeNotFound("bar", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of unknown types in slices", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"fof": "[]schtring",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
				"bam": "[]bar",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("schtring", "root"),
			newValidationErrorTypeNotFound("bar", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of unknown types in maps", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "int",
			"fof": "[int]schtring",
			"boo": "[schtring]int",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
				"bam": "[int]bar",
				"bal": "[bar]int",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("schtring", "root"),
			newValidationErrorTypeNotFound("schtring", "root"),
			newValidationErrorTypeNotFound("bar", "baz"),
			newValidationErrorTypeNotFound("bar", "baz"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail with multiple errors of multiple undefined types are used in declaration", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "map[bar]map[ban]baz",
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorTypeNotFound("bar", "root"),
			newValidationErrorTypeNotFound("ban", "root"),
			newValidationErrorTypeNotFound("baz", "root"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should not fail when type is used before declared", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"fof": "foo",
			"foo": "int",
			"baz": map[interface{}]interface{}{
				"ban": "int32",
				"bam": "bar",
			},
			"bar": "string",
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

}

func TestExtractTypes(t *testing.T) {
	t.Run("should extract a basic type", func(t *testing.T) {
		input := "string"

		actualOutput := extractTypes(input)
		expectedOutput := []string{"string"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
	t.Run("should extract a slice type", func(t *testing.T) {
		input := "[]int"

		actualOutput := extractTypes(input)
		expectedOutput := []string{"int"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
	t.Run("should extract both types form a map declaration", func(t *testing.T) {
		input := "map[string]int16"

		actualOutput := extractTypes(input)
		expectedOutput := []string{"string", "int16"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
	t.Run("should extract all types from a complicated declaration (1/2)", func(t *testing.T) {
		input := "map[string]map[uint][][]bool"

		actualOutput := extractTypes(input)
		expectedOutput := []string{"string", "uint", "bool"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
	t.Run("should extract all types from a complicated declaration (2/2)", func(t *testing.T) {
		input := "map[[23]int]map[float][][][12]string"

		actualOutput := extractTypes(input)
		expectedOutput := []string{"int", "float", "string"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
}

func TestFindUndefinedTypesIn(t *testing.T) {
	t.Run("should find all undefined types", func(t *testing.T) {
		definedTypesInput := []string{"foo", "bar"}
		usedTypesInput := []string{"foo", "bar", "baz", "string", "uint16", "bool", "bam"}

		actualOutput := findUndefinedTypesIn(usedTypesInput, definedTypesInput)
		expectedOutput := []string{"baz", "bam"}

		assert.Equal(t, expectedOutput, actualOutput)
	})
}
