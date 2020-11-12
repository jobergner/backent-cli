package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCreaters(t *testing.T) {
	t.Run("adds creators", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			output_person_type,
			output_name_type,
			output_child_type,
		})

		actual := splitPrintedDeclarations(input.addCreaters())
		expected := []string{
			output_person_type,
			output_name_type,
			output_CreatePerson_stateMachine_func,
			output_CreateName_stateMachine_func,
			output_CreateChild_stateMachine_func,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addCreaters() *stateMachine {
	return sm
}
