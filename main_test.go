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

type stateMachine ast.File

func simplifyIfWhitespace(ch rune) rune {
	if ch == '\n' {
		return ch
	}
	if unicode.IsSpace(ch) {
		return ' '
	}
	return ch
}

func normalizeWhitespace(_str string) string {
	str := strings.TrimSpace(_str)
	var b strings.Builder
	b.Grow(len(str))

	var lastWrittenRune rune = '1'

	for _, _ch := range str {
		ch := simplifyIfWhitespace(_ch)
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			lastWrittenRune = ch
		} else {
			if lastWrittenRune != ch {
				b.WriteRune(ch)
				lastWrittenRune = ch
			}
		}
	}

	return b.String()
}

func findDeclarationIn(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func removeDeclarationFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func normalizeDeclarations(declarations []string) []string {
	var normalizedDeclarations []string
	for _, def := range declarations {
		normalizedDeclarations = append(normalizedDeclarations, normalizeWhitespace(def))
	}
	return normalizedDeclarations
}

func matchDeclarations(actualDeclarations, expectedDeclarations []string) (leftoverDeclarations, redundantDeclarations []string) {
	// redefine redunantDeclarations so it never returns as nil (which happens when there are no redunant strings)
	// so comparing string slices becomes more conventient
	redundantDeclarations = make([]string, 0)
	leftoverDeclarations = make([]string, len(expectedDeclarations))
	copy(leftoverDeclarations, expectedDeclarations)

	actualDeclarations = normalizeDeclarations(actualDeclarations)
	leftoverDeclarations = normalizeDeclarations(leftoverDeclarations)

	for _, actualDeclaration := range actualDeclarations {
		leftoverDeclarationIndex, isFound := findDeclarationIn(actualDeclaration, leftoverDeclarations)
		if isFound {
			leftoverDeclarations = removeDeclarationFromSlice(leftoverDeclarations, leftoverDeclarationIndex)
		} else {
			redundantDeclarations = append(redundantDeclarations, actualDeclaration)
		}
	}

	return leftoverDeclarations, redundantDeclarations
}

func TestMatchDeclarations(t *testing.T) {
	t.Run("should ignore leading/trailing whitespace", func(t *testing.T) {
		actualDeclarations := []string{
			`   

			type foo struct {
				a string
			}

			  `,
		}
		expectedDeclarations := []string{
			`type foo struct {
				a string
			}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actualDeclarations, expectedDeclarations)

		assert.Equal(t, 0, len(missingDeclarations))
		assert.Equal(t, 0, len(redundantDeclarations))
	})
	t.Run("should match all declarations", func(t *testing.T) {
		actualDeclarations := []string{
			`type foo struct {
				a string
			}`,
		}
		expectedDeclarations := []string{
			`type foo struct {
				a string
			}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actualDeclarations, expectedDeclarations)

		assert.Equal(t, 0, len(missingDeclarations))
		assert.Equal(t, 0, len(redundantDeclarations))
	})
	t.Run("should find redundant declaration", func(t *testing.T) {
		actualDeclarations := []string{
			`type foo struct {
				a string
			}`,
			`func (a int) int {
				return 1+a
			}`,
		}
		expectedDeclarations := []string{
			`type foo struct {
				a string
			}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actualDeclarations, expectedDeclarations)

		assert.Equal(t, 0, len(missingDeclarations))
		assert.Equal(t, 1, len(redundantDeclarations))
		assert.Equal(t, []string{normalizeWhitespace(`func (a int) int { return 1+a }`)}, redundantDeclarations)
	})
	t.Run("should find missing declaration", func(t *testing.T) {
		actualDeclarations := []string{
			`type foo struct {
				a string
			}`,
		}
		expectedDeclarations := []string{
			`type foo struct {
				a string
			}`,
			`func (a int) int {
				return 1+a
			}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actualDeclarations, expectedDeclarations)

		assert.Equal(t, 1, len(missingDeclarations))
		assert.Equal(t, 0, len(redundantDeclarations))
		assert.Equal(t, []string{normalizeWhitespace(`func (a int) int { return 1+a }`)}, missingDeclarations)
	})
}

func splitDecls(decls string) []string {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+decls, 0)
	if err != nil {
		panic(err)
	}
	return splitPrintedDeclarations(file)
}

func splitPrintedDeclarations(f *ast.File) []string {
	printedDeclarations := make([]string, 0)
	for _, decl := range f.Decls {
		printedDeclarations = append(printedDeclarations, unsafePrintDeclaration(decl))
	}
	return printedDeclarations
}

func unsafePrintDeclaration(decl ast.Decl) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, token.NewFileSet(), decl)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func unsafeParseDecls(decls []string) *ast.File {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+strings.Join(decls, "\n"), 0)
	if err != nil {
		panic(err)
	}
	return file
}

func newSimpleASTExample() simpleAST {
	data := map[interface{}]interface{}{
		"player": map[interface{}]interface{}{
			"items":     "[]item",
			"gearScore": "gearScore",
			"position":  "position",
		},
		"zone": map[interface{}]interface{}{
			"items":   "[]zoneItem",
			"players": "[]player",
		},
		"zoneItem": map[interface{}]interface{}{
			"position": "position",
			"item":     "item",
		},
		"position": map[interface{}]interface{}{
			"x": "float64",
			"y": "float64",
		},
		"item": map[interface{}]interface{}{
			"gearScore": "gearScore",
		},
		"gearScore": map[interface{}]interface{}{
			"level": "int",
			"score": "int",
		},
	}

	// TODO: make prettier
	simpleAST := buildRudimentarySimpleAST(data)
	simpleAST.fillInReferences().fillInParentalInfo()

	return simpleAST
}
