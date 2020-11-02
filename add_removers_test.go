package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRemovers(t *testing.T) {
	t.Run("adds removers", func(t *testing.T) {
		input := unsafeParseDecls([]string{`
type person struct {
	id personID
	children []childID
	lastModified int64
	operationKind operationKind
}`, `
type child struct {
	id childID
	name string
	lastModified int64
	operationKind operationKind
}`,
		})

		actual := splitPrintedDeclarations(input.addRemovers())
		expected := []string{`
type person struct {
	id personID
	children []childID
	lastModified int64
	operationKind operationKind
}`, `
type child struct {
	id childID
	name string
	lastModified int64
	operationKind operationKind
}`, `
func (sm *stateMachine) RemovePerson(personID personID) {
	person := sm.GetPerson(personID)
	person.lastModified = time.Now().UnixNano()
	person.operationKind = operationKindDelete
	sm.patch.person[person.id] = person
}`, `
func (sm *stateMachine) RemoveName(nameID nameID) {
	name := sm.GetName(nameID)
	name.lastModified = time.Now().UnixNano()
	name.operationKind = operationKindDelete
	sm.patch.name[name.id] = name
}`, `
func (p person) RemoveChild(childID childID, sm *stateMachine) {
	var indexToRemove int
	for i, _childID := p.children {
		if _childID == childID {
			indexToRemove = i	
			break
		}
	}
	p.children = append(p.children[:indexToRemove], p.children[indexToRemove+1:]...)
	p.lastModified = time.Now().UnixNano()
	p.operationKind = operationKindUpdate
	sm.patch.person[p.id] = p
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addRemovers() *stateMachine {
	return sm
}
