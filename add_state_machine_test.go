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
	patchReceiver chan state
	idgen int
}`, `
func (*sm) generateID() int {
	newID := sm.idgen
	sm.idgen = sm.idgen + 1
	return newID
}
`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addStateMachineDeclaration() *stateMachine {
	return sm
}
