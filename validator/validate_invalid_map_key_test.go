package validator

import (
	"testing"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/stretchr/testify/assert"
)

func TestValidateDataInvalidMapKey(t *testing.T) {
	t.Run("should not fail on usage of valid map keys", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "string",
			"bal": "map[foo]int",
			"baz": map[interface{}]interface{}{
				"bal": "map[foo]int",
				"ban": "map[[2]foo]int",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of reference type directly as map key", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "string",
			"bar": "map[*foo]int",
			"buf": "map[map[int]bool]string",
			"baz": map[interface{}]interface{}{
				"ban": "map[[]foo]int",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorInvalidMapKey("*foo", "map[*foo]int"),
			newValidationErrorInvalidMapKey("map[int]bool", "map[map[int]bool]string"),
			newValidationErrorInvalidMapKey("[]foo", "map[[]foo]int"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of reference type as map key", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo":  "[]string",
			"ban":  "*int",
			"bunt": "map[int]string",
			"bar":  "map[foo]int",
			"baz": map[interface{}]interface{}{
				"bal": "map[ban]int",
				"buf": "map[bunt]int",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorInvalidMapKey("foo", "map[foo]int"),
			newValidationErrorInvalidMapKey("ban", "map[ban]int"),
			newValidationErrorInvalidMapKey("bunt", "map[bunt]int"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})

	t.Run("should fail on usage of reference type as map key in nested map", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "[]string",
			"bar": "map[int]map[foo]int",
			"baz": map[interface{}]interface{}{
				"bal": "map[bar]int",
			},
		}

		actualErrors := logicalValidation(data)
		expectedErrors := []error{
			newValidationErrorInvalidMapKey("foo", "map[int]map[foo]int"),
			newValidationErrorInvalidMapKey("bar", "map[bar]int"),
		}

		missingErrors, redundantErrors := matchErrors(actualErrors, expectedErrors)

		assert.Empty(t, missingErrors)
		assert.Empty(t, redundantErrors)
	})
}

func TestExtractMapKeys(t *testing.T) {
	t.Run("should extract map keys from value strings", func(t *testing.T) {
		assert.Equal(t, extractMapKeys("map[int]string"), []string{"int"})
		assert.Equal(t, extractMapKeys("map[*int]string"), []string{"*int"})
		assert.Equal(t, extractMapKeys("map[[]int]string"), []string{"[]int"})
		assert.Equal(t, extractMapKeys("map[int]map[float]map[string]bool"), []string{"int", "float", "string"})
		assert.Equal(t, extractMapKeys("map[map[map[bool]int]string]float"), []string{"map[map[bool]int]string"})
		assert.Equal(t, extractMapKeys("map[map[map[bool]int]string]map[map[bool]int]float"), []string{"map[map[bool]int]string", "map[bool]int"})
		assert.Equal(t, extractMapKeys("map[int][]map[float]int"), []string{"int", "float"})
		assert.Equal(t, extractMapKeys("map[int][][][]map[float]int"), []string{"int", "float"})
		assert.Equal(t, extractMapKeys("map[int]*[][]map[float]int"), []string{"int", "float"})
	})
}

func TestExtracpMapDeclExpression(t *testing.T) {
	t.Run("should extract map decl", func(t *testing.T) {
		mockSrc := `
	package main
	type mockType map[int]string
	`

		file, _ := parser.ParseFile(token.NewFileSet(), "", mockSrc, 0)
		assert.Equal(t, extracpMapDeclExpression(file).(*ast.MapType).Key.(*ast.Ident).Name, "int")
	})
}

func TestContainsInvalidMapKeys(t *testing.T) {
	t.Run("should not contain illegal map keys (1/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "string",
			"bar": "map[foo]int",
		}

		illegalMapKeys := findIllegalMapKeys("map[foo]int", data)

		assert.Equal(t, len(illegalMapKeys), 0)
	})
	t.Run("should not contain illegal map keys (2/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[foo]int",
			"foo": map[interface{}]interface{}{
				"bal": "string",
			},
		}

		illegalMapKeys := findIllegalMapKeys("map[foo]int", data)

		assert.Equal(t, len(illegalMapKeys), 0)
	})
	t.Run("should not contain illegal map keys (3/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[foo]int",
			"foo": map[interface{}]interface{}{
				"bal": "ban",
			},
			"ban": map[interface{}]interface{}{
				"baf": "int",
			},
		}

		illegalMapKeys := findIllegalMapKeys("map[foo]int", data)

		assert.Equal(t, len(illegalMapKeys), 0)
	})
	t.Run("should contain illegal map keys (1/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[*string]int",
		}

		illegalMapKeys := findIllegalMapKeys("map[*string]int", data)

		assert.Equal(t, illegalMapKeys, []string{"*string"})
	})
	t.Run("should contain illegal map keys (2/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": "[]string",
			"bar": "map[foo]int",
		}

		illegalMapKeys := findIllegalMapKeys("map[foo]int", data)

		assert.Equal(t, illegalMapKeys, []string{"foo"})
	})
	t.Run("should contain illegal map keys (3/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[foo]int",
			"foo": map[interface{}]interface{}{
				"bal": "*int",
			},
		}

		illegalMapKeys := findIllegalMapKeys("map[foo]int", data)

		assert.Equal(t, illegalMapKeys, []string{"foo"})
	})
	t.Run("should contain illegal nested map keys (1/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[map[bool]float]int",
		}

		illegalMapKeys := findIllegalMapKeys("map[map[bool]float]int", data)

		assert.Equal(t, illegalMapKeys, []string{"map[bool]float"})
	})
	t.Run("should contain illegal nested map keys (2/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[*string]map[foo]bool",
			"foo": map[interface{}]interface{}{
				"bal": "*int",
			},
		}

		illegalMapKeys := findIllegalMapKeys("map[*string]map[foo]bool", data)

		assert.Equal(t, illegalMapKeys, []string{"*string", "foo"})
	})
	t.Run("should contain illegal nested map keys (3/3)", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"bar": "map[ban]bool",
			"foo": map[interface{}]interface{}{
				"bal": "*int",
			},
			"ban": map[interface{}]interface{}{
				"bunt": "foo",
			},
		}

		illegalMapKeys := findIllegalMapKeys("map[ban]bool", data)

		assert.Equal(t, illegalMapKeys, []string{"ban"})
	})
}
