package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedParentage(t *testing.T) {
	t.Run("should add parentage declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			input_person_type,
			input_child_type,
			input_name_type,
		})

		smb := newStateMachineBuilder(input)
		smb.embedParentage()
		actual := splitPrintedDeclarations(smb.stateMachine)
		expected := []string{
			`type person struct {
	name		name
	children	[]child
	age		int
	parentage Parentage
}`, `
type name struct {
	first string
	last string
	parentage Parentage
}`, `
type child struct {
	name name
	parentage Parentage
}`,
			output_parentInfo_type,
			output_parentage_type,
			output_entityKind_type,
			output_entityKindPerson_type,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachineBuilder) embedParentage() *stateMachineBuilder {
	return sm
}
