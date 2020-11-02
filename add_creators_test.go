package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCreaters(t *testing.T) {
	t.Run("adds creators", func(t *testing.T) {
		input := unsafeParseDecls([]string{`
type person struct {
	id personID
	name nameID
	age int
	lastModified int64
	operationKind operationKind
}`, `
type name struct {
	id nameID
	first string
	last string
	lastModified int64
	operationKind operationKind
}`,
		})

		metaFields := []metaField{{"lastModified", "int64"}, {"id", "int"}, {"operationKind", "operationKind"}}
		actual := splitPrintedDeclarations(input.addCreaters(metaFields))
		expected := []string{`
type person struct {
	id personID
	name nameID
	age int
	lastModified int64
	operationKind operationKind
}`, `
type name struct {
	id nameID
	first string
	last string
	lastModified int64
	operationKind operationKind
}`, `
func (sm *stateMachine) CreatePerson(personID personID) person {
	var person person
	person.id = personID(sm.generateID())
	person.lastModified = time.Now().UnixNano()
	person.operationKind = operationKindCreate
	sm.patch.person[person.id] = person
	return person
}`, `
func (sm *stateMachine) CreateName(nameID nameID) name {
	var name name
	name.id = nameID(sm.generateID())
	name.lastModified = time.Now().UnixNano()
	name.operationKind = operationKindCreate
	sm.patch.name[name.id] = name
	return name
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addCreaters(metaFields []metaField) *stateMachine {
	return sm
}
