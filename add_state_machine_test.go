package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStateMachine(t *testing.T) {
	t.Run("should add state machine declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			input_person_type,
			input_child_type,
			input_name_type,
		})

		actual := splitPrintedDeclarations(input.addStateMachineDeclaration())
		expected := []string{
			input_person_type,
			input_child_type,
			input_name_type,
			output_operationKind_type,
			output_operationKindCreate_type,
			output_state_type,
			output_stateMachine_type,
			output_generateID_stateMachine_func,
			output_updateState_stateMachine_func,
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
