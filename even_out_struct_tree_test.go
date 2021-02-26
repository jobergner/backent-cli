package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOutStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			// input_person_type,
			// input_child_type,
			// input_name_type,
		})

		smb := newStateMachineBuilder(input)
		smb.evenOutStructTree()
		actual := splitPrintedDeclarations(smb.stateMachine)
		expected := []string{`
type person struct {
	name nameID
	children []childID
}`, `
type name struct {
	first string
	last string
}`, `
type child struct {
	name nameID
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachineBuilder) evenOutStructTree() *stateMachineBuilder {
	return sm
}
