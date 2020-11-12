package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGetters(t *testing.T) {
	t.Run("adds getters", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			output_person_type,
			output_name_type,
			output_child_type,
		})

		actual := splitPrintedDeclarations(input.addGetters())
		expected := []string{
			output_person_type,
			output_name_type,
			output_child_type,
			output_GetPerson_stateMachine_func,
			output_GetChild_stateMachine_func,
			output_GetName_stateMachine_func,
			output_GetName_person_func,
			output_GetChildren_person_func,
			output_GetName_child_func,
			output_GetAge_person_func,
			output_GetFirst_name_func,
			output_GetLast_name_func,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addGetters() *stateMachine {
	return sm
}
