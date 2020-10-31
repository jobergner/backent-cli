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

const (
	_personDeclaration string = `
type person struct {
	name name
}
	`
	_nameDeclaration = `
type name struct {
	first string
	last string
}
	`
)

// "ab c  de\nf" => "ab c de f"
func normalizeWhitespace(_str string) string {
	str := strings.TrimSpace(_str)
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

func splitPrintedDeclarations(sm *stateMachine) []string {
	printedDeclarations := make([]string, 0)
	for _, decl := range sm.Decls {
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

func unsafeParseDecls(decls []string) *stateMachine {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+strings.Join(decls, "\n"), 0)
	if err != nil {
		panic(err)
	}
	x := stateMachine(*file)
	return &x
}

func TestFlattenStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.flattenStructTree())
		expected := []string{`
type person struct {
	name nameID
}`, `
type name struct {
	first string
	last string
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) flattenStructTree() *stateMachine {
	return sm
}

func TestAddIDTypes(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.addIdTypes())
		expected := []string{
			`type personID string`,
			`type nameID string`,
			_personDeclaration,
			_nameDeclaration,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addIdTypes() *stateMachine {
	return sm
}

func TestExtractDeclaredStructNames(t *testing.T) {
	t.Run("should find all struct names in file", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.extractDeclaredStructNames())
		expected := []string{"person", "name"}

		assert.Equal(t, expected, actual)
	})
}

// TODO not method
func (sm *stateMachine) extractDeclaredStructNames() *stateMachine {
	return sm
}

type metaField struct {
	name        string
	typeLiteral string
}

func TestEmbedStructMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.embedStructMetaFields([]metaField{{"lastModified", "int"}}))
		expected := []string{`
type person struct {
	name name
	lastModified int
}`, `
type name struct {
	first string
	last string
	lastModified int
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) embedStructMetaFields(metaFields []metaField) *stateMachine {
	return sm
}

func TestEmbedParentageDeclaration(t *testing.T) {
	t.Run("should add parentage declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.embedParentageDeclaration())
		expected := []string{`
type person struct {
	name name
	parentage Parentage
}`, `
type name struct {
	first string
	last string
	parentage Parentage
}`,
			`type parentage []parentInfo`, `
type parentInfo struct {
	kind entityKind
	id string
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) embedParentageDeclaration() *stateMachine {
	return sm
}

func TestAddStateMachine(t *testing.T) {
	t.Run("should add state machine declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.addStateMachineDeclaration())
		expected := []string{
			_personDeclaration,
			_nameDeclaration, `
type state struct {
	person map[personID]person
	name map[nameID]name
}`, `
type stateMachine struct {
	state state
	patch state
	patchReceiver chan state
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addStateMachineDeclaration() *stateMachine {
	return sm
}

func TestAddGetters(t *testing.T) {
	t.Run("adds getters", func(t *testing.T) {
		input := unsafeParseDecls([]string{`
type person struct {
	name nameID
	age int
}`,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.addGetters([]metaField{{"lastModified", "int"}}))
		expected := []string{`
type person struct {
	name nameID
	age int
	lastModified int
}`, `
type name struct {
	first string
	last string
	lastModified int
}`, `
func (sm *stateMachine) GetPerson(personID personID) person {
	person, ok := sm.patch.person[personID]
	if ok {
		return person
	}
	currentPerson := sm.state.person[personID]
	personCopy := person{}
	copier.Copy(&personCopy, &currentPerson)
	return personCopy
}`, `
func (sm *stateMachine) GetName(nameID nameID) name {
	name, ok := sm.patch.name[nameID]
	if ok {
		return name
	}
	currentName := sm.state.name[nameID]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}`, `
func (p person) GetName(sm *stateMachine) name {
	name, ok := sm.patch.name[p.name]
	if ok {
		return name
	}
	currentName := sm.state.name[p.name]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}`, `
func (p person) GetAge() int {
	return p.age
}`, `
func (n name) GetFirst() string {
	return n.first
}`, `
func (n name) GetLast() string {
	return n.last
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addGetters(metaFields []metaField) *stateMachine {
	return sm
}
