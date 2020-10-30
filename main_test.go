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

func unsafePrintFile(file *stateMachine) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, token.NewFileSet(), file)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func stringifyFile(sm *stateMachine) string {
	return normalizeWhitespace(unsafePrintFile(sm))
}

func unsafeParseDecls(decls string) *stateMachine {
	file, err := parser.ParseFile(token.NewFileSet(), "", "package foo\n"+decls, 0)
	if err != nil {
		panic(err)
	}
	x := stateMachine(*file)
	return &x
}

func TestFlattenStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		inputStateMachine := stateMachine(*input)
		actual := inputStateMachine.flattenStructTree()
		expected := unsafeParseDecls(`
type person struct {
	name nameID
}
type name struct {
	first string
	last string
} `)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) flattenStructTree() *stateMachine {
	return sm
}

func TestAddIDTypes(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		inputStateMachine := stateMachine(*input)
		actual := inputStateMachine.addIdTypes()
		expected := unsafeParseDecls(`
type personID string
type nameID string
type person struct {
	name nameID
}
type name struct {
	first string
	last string
} `)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) addIdTypes() *stateMachine {
	return sm
}

func TestExtractDeclaredStructNames(t *testing.T) {
	t.Run("should find all struct names in file", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		actual := input.extractDeclaredStructNames()
		expected := []string{"person", "name"}

		assert.Equal(t, expected, actual)
	})
}

func (sm *stateMachine) extractDeclaredStructNames() *stateMachine {
	return sm
}

type metaField struct {
	name        string
	typeLiteral string
}

func TestEmbedStructMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		actual := input.embedStructMetaFields([]metaField{{"lastModified", "int"}})
		expected := unsafeParseDecls(`
type person struct {
	name name
	lastModified int
}
type name struct {
	first string
	last string
	lastModified int
}`)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) embedStructMetaFields(metaFields []metaField) *stateMachine {
	return sm
}

func TestEmbedParentageDeclaration(t *testing.T) {
	t.Run("should add parentage declaration", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		actual := input.embedParentageDeclaration()
		expected := unsafeParseDecls(`
type person struct {
	name name
	parentage Parentage
}
type name struct {
	first string
	last string
	parentage Parentage
}
type parentage []parentInfo
type parentInfo struct {
	kind entityKind
	id string
}`)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) embedParentageDeclaration() *stateMachine {
	return sm
}

func TestAddStateMachine(t *testing.T) {
	t.Run("should add state machine declaration", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}`)

		actual := input.addStateMachineDeclaration()
		expected := unsafeParseDecls(`
type person struct {
	name name
}
type name struct {
	first string
	last string
}
type state struct {
	person map[personID]person
	name map[nameID]name
}
type stateMachine struct {
	state state
	patch state
	patchReceiver chan state
}
`)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) addStateMachineDeclaration() *stateMachine {
	return sm
}

func TestAddGetters(t *testing.T) {
	t.Run("adds getters", func(t *testing.T) {
		input := unsafeParseDecls(`
type person struct {
	name nameID
	age int
}
type name struct {
	first string
	last string
}`)

		actual := input.addGetters([]metaField{{"lastModified", "int"}})
		expected := unsafeParseDecls(`
type person struct {
	name nameID
	age int
	lastModified int
}
type name struct {
	first string
	last string
	lastModified int
}
func (sm *stateMachine) GetPerson(personID personID) person {
	person, ok := sm.patch.person[personID]
	if ok {
		return person
	}
	currentPerson := sm.state.person[personID]
	personCopy := person{}
	copier.Copy(&personCopy, &currentPerson)
	return personCopy
}
func (sm *stateMachine) GetName(nameID nameID) name {
	name, ok := sm.patch.name[nameID]
	if ok {
		return name
	}
	currentName := sm.state.name[nameID]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}
func (p person) GetName(sm *stateMachine) name {
	name, ok := sm.patch.name[p.name]
	if ok {
		return name
	}
	currentName := sm.state.name[p.name]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}
func (p person) GetAge() int {
	return p.age
}
func (n name) GetFirst() string {
	return n.first
}
func (n name) GetLast() string {
	return n.last
}
`)

		assert.Equal(t, stringifyFile(expected), stringifyFile(actual))
	})
}

func (sm *stateMachine) addGetters(metaFields []metaField) *stateMachine {
	return sm
}
