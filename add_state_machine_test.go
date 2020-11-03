package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
type operationKind string
const (
	operationKindCreate operationKind = "CREATE"
	operationKindDelete = "DELETE"
	operationKindUpdate = "UPDATE"
)
type state struct {
	person map[personID]person
	name map[nameID]name
}`, `
type stateMachine struct {
	state state
	patch state
	idgen int
}`, `
func (*sm) generateID() int {
	newID := sm.idgen
	sm.idgen = sm.idgen + 1
	return newID
}
`, `
func (*sm) updateState() {
	for _, person := range sm.patch.person {
		if person.operationKind == operationKindDelete {
			delete(sm.state.person, person.id)
		} else {
			sm.state.person[person.id] = person
		}
	}
	for _, name := range sm.patch.name {
		if name.operationKind == operationKindDelete {
			delete(sm.state.name, name.id)
		} else {
			sm.state.name[name.id] = name
		}
	}
	sm.patch = newState()
}
`,
			// TODO newState()
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addStateMachineDeclaration() *stateMachine {
	return sm
}
