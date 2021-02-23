package validator

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	t.Run("should unmarshal without errors", func(t *testing.T) {
		yamlDataBytes := []byte(
			`foo: string
bar: 
  baz: int
  ban: "[]foo"`,
		)

		file, errs := Unmarshal(yamlDataBytes)

		assert.Equal(t, errs, []error{})

		output := normalizeWhitespace(printDecls(file))
		expectedOutput := normalizeWhitespace(
			`type bar struct {
				ban []foo
				baz int
			}
			type foo string`,
		)

		assert.Equal(t, output, expectedOutput)
	})

	t.Run("should return error", func(t *testing.T) {
		yamlDataBytes := []byte(
			`foo: string
bar: 
  baz: boo
  ban: "[]bool"`,
		)

		decls, errs := Unmarshal(yamlDataBytes)
		assert.Equal(t, errs, []error{newValidationErrorTypeNotFound("boo", "bar")})
		var expectedFile []ast.Decl
		assert.Equal(t, decls, expectedFile)
	})
}
