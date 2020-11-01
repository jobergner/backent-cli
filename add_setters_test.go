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
}`, `
type name struct {
	id string
	first string
	last string
	lastModified int64
}`, `
func (p person) SetAge(val int, sm *stateMachine) person {
	patchingPerson := sm.patch.person[p.id]
	patchingPerson.age = val
	patchingPerson.lastModified = time.Now().UnixNano()
	patchingPerson.operationKind = operationKindUpdate
	sm.patch.person[p.id] = patchingPerson
	return patchingPerson
}`, `
func (n name) SetFirst(val string, sm *stateMachine) name {
	patchingName := sm.patch.name[p.id]
	patchingName.first = val
	patchingName.lastModified = time.Now().UnixNano()
	patchingName.operationKind = operationKindUpdate
	sm.patch.name[p.id] = patchingName
	return patchingName
}`, `
func (n name) SetLast(val string, sm *stateMachine) name {
	patchingName := sm.patch.name[p.id]
	patchingName.last = val
	patchingName.lastModified = time.Now().UnixNano()
	patchingName.operationKind = operationKindUpdate
	sm.patch.name[p.id] = patchingName
	return patchingName
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
