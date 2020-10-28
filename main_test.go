package statefactory

import (
	"bytes"
	"go/ast"
	"go/parser"
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

func unsafePrintFile(file *ast.File) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, token.NewFileSet(), file)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func compareFile(file *ast.File) string {
	return normalizeWhitespace(unsafePrintFile(file))
}

func unsafeParseCode(code string) *ast.File {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+code, 0)
	if err != nil {
		panic(err)
	}
	return file
}

func TestFlattenStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseCode(`
type person struct {
	name residency
}
type residency struct {
	location location
}
type location struct {
	x float
	y float
}`)

		actual := flattenStructTree(input)
		expected := unsafeParseCode(`
type person struct {
	residency residencyID
}
type residency struct {
	location locationID
}
type residency struct {
	x float
	y float
}`)

		assert.Equal(t, compareFile(expected), compareFile(actual))
	})
}

func flattenStructTree(file *ast.File) *ast.File {
	return file
}

func TestExtractDeclaredStructNames(t *testing.T) {
	t.Run("should find all struct names in file", func(t *testing.T) {
		input := unsafeParseCode(`
type person struct {
	name residency
}
type residency struct {
	location location
}
type location struct {
	x float
	y float
}`)

		actual := extractDeclaredStructNames(input)
		expected := []string{"person", "residency", "location"}

		assert.Equal(t, expected, actual)
	})
}

func extractDeclaredStructNames(file *ast.File) []string {
	return []string{}
}

func TestEmbedStructMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseCode(`
type person struct {
	residency residencyID
}
type residency struct {
	location locationID
}
type residency struct {
	x float
	y float
}`)

		actual := embedStructMetaFields(input)
		expected := unsafeParseCode(`
type person struct {
	residency residencyID
	lastUpdated int
}
type residency struct {
	location locationID
	lastUpdated int
}
type residency struct {
	x float
	y float
	lastUpdated int
}`)

		assert.Equal(t, compareFile(expected), compareFile(actual))
	})
}

func embedStructMetaFields(file *ast.File) *ast.File {
	return file
}
