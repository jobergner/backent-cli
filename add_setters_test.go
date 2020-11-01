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
	lastModified int
}`, `
type name struct {
	id string
	first string
	last string
	lastModified int
}`,
		})

		metaFields := []metaField{{"lastModified", "int"}, {"id", "string"}}
		actual := splitPrintedDeclarations(input.addSetters(metaFields))
		expected := []string{`
type person struct {
	id string
	name nameID
	age int
	lastModified int
}`, `
type name struct {
	id string
	first string
	last string
	lastModified int
}`, `
func (p *person) prepareModification(sm *stateMachine) *person {
	_, ok := sm.patch.person[p.id]
	if ok {
		return p
	}
	sm.patch.person[p.id] = *p
	patchingPerson := &sm.patch.person[p.id] 
	return patchingPerson
}`, `
func (n *name) prepareModification(sm *stateMachine) *name {
	_, ok := sm.patch.name[n.id]
	if ok {
		return n
	}
	sm.patch.name[n.id] = *n
	patchingName := &sm.patch.name[n.id] 
	return patchingName
}`, `
func (p person) SetAge(val int, sm *stateMachine) person {
	patchingPerson := sm.patch.person[p.id]
	patchingPerson.age = val
	sm.patch.person[p.id] = patchingPerson
	return patchingPerson
}`, `
func (n name) SetFirst(val string, sm *stateMachine) name {
	patchingName := sm.patch.name[p.id]
	patchingName.first = val
	sm.patch.name[p.id] = patchingName
	return patchingName
}`, `
func (n name) SetLast(val string, sm *stateMachine) name {
	patchingName := sm.patch.name[p.id]
	patchingName.last = val
	sm.patch.name[p.id] = patchingName
	return patchingName
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

type x int

func (x *x) set() *x {
	return x
}
func (x x) get() x {
	return x
}

func (sm *stateMachine) addSetters(metaFields []metaField) *stateMachine {
	return sm
}
