package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGetters(t *testing.T) {
	t.Run("adds getters", func(t *testing.T) {
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
		actual := splitPrintedDeclarations(input.addGetters(metaFields))
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
func (p *person) GetName(sm *stateMachine) name {
	name, ok := sm.patch.name[p.name]
	if ok {
		return name
	}
	currentName := sm.state.name[p.name]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}`, `
func (p *person) GetAge() int {
	return p.age
}`, `
func (n *name) GetFirst() string {
	return n.first
}`, `
func (n *name) GetLast() string {
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
