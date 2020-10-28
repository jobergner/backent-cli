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

func unsafeParseDecls(decls string) *ast.File {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+decls, 0)
	if err != nil {
		panic(err)
	}
	return file
}

func TestFlattenStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls(`
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
		expected := unsafeParseDecls(`
type person struct {
	residency residencyID
}
type residency struct {
	location locationID
}
type location struct {
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
		input := unsafeParseDecls(`
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
		input := unsafeParseDecls(`
type person struct {
	residency residencyID
}
type residency struct {
	location locationID
}
type location struct {
	x float
	y float
}`)

		actual := embedStructMetaFields(input)
		expected := unsafeParseDecls(`
type person struct {
	residency residencyID
	lastUpdated int
}
type residency struct {
	location locationID
	lastUpdated int
}
type location struct {
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

func TestEmbedParentageDeclaration(t *testing.T) {
	t.Run("should add parentage declaration", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	residency residencyID
}
type residency struct {
	location locationID
}
type location struct {
	x float
	y float
}`)

		actual := embedParentageDeclaration(input)
		expected := unsafeParseDecls(`
type person struct {
	residency residencyID
	parentage Parentage
}
type residency struct {
	location locationID
	parentage Parentage
}
type location struct {
	x float
	y float
	parentage Parentage
}
type parentage []parentInfo
type parentInfo struct {
	kind entityKind
	id string
}`)

		assert.Equal(t, compareFile(expected), compareFile(actual))
	})
}

func embedParentageDeclaration(file *ast.File) *ast.File {
	return file
}

func TestAddStateMachine(t *testing.T) {
	t.Run("should add state machine declaration", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	residency residencyID
	lastUpdated int
}
type residency struct {
	location locationID
	lastUpdated int
}
type location struct {
	x float
	y float
	lastUpdated int
}`)

		actual := addStateMachineDeclaration(input)
		expected := unsafeParseDecls(`
type person struct {
	residency residencyID
	lastUpdated int
}
type residency struct {
	location locationID
	lastUpdated int
}
type location struct {
	x float
	y float
	lastUpdated int
}
type state struct {
	residency map[residencyID]residency
	location map[locationID]location
	person map[personID]person
}
type stateMachine struct {
	state state
	patch state
	patchReceiver chan state
}
`)

		assert.Equal(t, compareFile(expected), compareFile(actual))
	})
}

func addStateMachineDeclaration(file *ast.File) *ast.File {
	return file
}
