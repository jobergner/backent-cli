package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSetters(t *testing.T) {
	t.Run("adds setters", func(t *testing.T) {
		input := unsafeParseDecls([]string{`
type person struct {
	id string
	name nameID
	age int
	operationKind operationKind
}`, `
type name struct {
	id string
	first string
	last string
	lastModified int64
	operationKind operationKind
}`,
		})

		metaFields := []metaField{{"lastModified", "int64"}, {"id", "string"}, {"operationKind", "operationKind"}}
		actual := splitPrintedDeclarations(input.addSetters(metaFields))
		expected := []string{`
type person struct {
	id string
	name nameID
	age int
	lastModified int64
	operationKind operationKind
}`, `
type name struct {
	id string
	first string
	last string
	lastModified int64
	operationKind operationKind
}`, `
func (p person) SetAge(val int, sm *stateMachine) person {
	p.age = val
	p.lastModified = time.Now().UnixNano()
	p.operationKind = operationKindUpdate
	sm.patch.person[p.id] = p
	return p
}`, `
func (n name) SetFirst(val string, sm *stateMachine) name {
	n.first = val
	n.lastModified = time.Now().UnixNano()
	n.operationKind = operationKindUpdate
	sm.patch.name[p.id] = n
	return n
}`, `
func (n name) SetLast(val string, sm *stateMachine) name {
	n.last = val
	n.lastModified = time.Now().UnixNano()
	n.operationKind = operationKindUpdate
	sm.patch.name[p.id] = n
	return n
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addSetters(metaFields []metaField) *stateMachine {
	return sm
}
