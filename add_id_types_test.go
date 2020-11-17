package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIDTypes(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			input_person_type,
			input_child_type,
			input_name_type,
		})

		smb := newStateMachineBuilder(input)
		smb.addIdTypes()
		actual := splitPrintedDeclarations(smb.stateMachine)
		expected := []string{
			input_person_type,
			input_child_type,
			input_name_type,
			output_childID_type,
			output_personID_type,
			output_nameID_type,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachineBuilder) addIdTypes() *stateMachineBuilder {
	return sm
}
