package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRemovers(t *testing.T) {
	t.Run("adds removers", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			// input_person_type,
			// input_child_type,
			// input_name_type,
		})

		smb := newStateMachineBuilder(input)
		smb.addRemovers()
		actual := splitPrintedDeclarations(smb.stateMachine)
		expected := []string{
			// input_person_type,
			// input_child_type,
			// input_name_type,
			// output_RemovePerson_stateMachine_func,
			// output_RemoveName_stateMachine_func,
			// output_RemoveChild_stateMachine_func,
			// output_RemoveChild_person_func,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachineBuilder) addRemovers() *stateMachineBuilder {
	return sm
}
