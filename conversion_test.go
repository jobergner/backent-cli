package validator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// "ab c  de\nf" => "ab c de f"
func normalizeWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))

	var wroteSpace bool = true

	for _, ch := range str {
		var isSpace bool = unicode.IsSpace(ch)

		if isSpace && wroteSpace {
			continue
		}

		if isSpace {
			b.WriteRune(' ')
		} else {
			b.WriteRune(ch)
		}

		if isSpace {
			wroteSpace = true
		} else {
			wroteSpace = false
		}
	}

	return b.String()
}

func printDecls(decls []ast.Decl) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, token.NewFileSet(), decls)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func printDeclsFromYamlData(inputYamlData map[interface{}]interface{}) string {
	golangAST := convertToAST(inputYamlData)
	golangDecls := printDecls(golangAST.Decls)
	return golangDecls
}

func TestConvertToASTBasicCases(t *testing.T) {
	t.Run("should convert named types", func(t *testing.T) {
		input := map[interface{}]interface{}{
			"foo": "string",
			"bar": "int",
		}
		expectedOutput := `
		type bar int
		type foo string`

		normalizedActualOutput := normalizeWhitespace(printDeclsFromYamlData(input))
		normalizedExpectedOutput := normalizeWhitespace(expectedOutput)

		assert.Equal(t, normalizedActualOutput, normalizedExpectedOutput)
	})
	t.Run("should convert struct types", func(t *testing.T) {
		input := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "int",
			},
		}
		expectedOutput := `
		type foo struct{ bar int }`

		normalizedActualOutput := normalizeWhitespace(printDeclsFromYamlData(input))
		normalizedExpectedOutput := normalizeWhitespace(expectedOutput)

		assert.Equal(t, normalizedActualOutput, normalizedExpectedOutput)
	})
}

func TestRangeInAlphabeticalOrder(t *testing.T) {
	t.Run("should loop in alphabetical range", func(t *testing.T) {
		input := map[interface{}]interface{}{
			"a": "1",
			"b": "2",
			"c": "3",
		}

		for i := 0; i < 100; i++ {
			var receivedKeys []string
			var receivedValues []string
			rangeInAlphabeticalOrder(input, func(key string, value interface{}) {
				_key := fmt.Sprintf("%v", key)
				receivedKeys = append(receivedKeys, _key)
				_value := fmt.Sprintf("%v", value)
				receivedValues = append(receivedValues, _value)
			})
			assert.Equal(t, receivedKeys, []string{"a", "b", "c"})
			assert.Equal(t, receivedValues, []string{"1", "2", "3"})
		}

	})
}
