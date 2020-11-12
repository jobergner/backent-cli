package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			input_person_type,
			input_child_type,
			input_name_type,
		})

		actual := splitPrintedDeclarations(input.embedMetaFields())
		expected := []string{
			`type person struct {
	name		name
	children	[]child
	age		int
	lastModified int64
	operationKind operationKind
}`, `
type name struct {
	first string
	last string
	lastModified int64
	operationKind operationKind
}`, `
type child struct {
	name name
	lastModified int64
	operationKind operationKind
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) embedMetaFields() *stateMachine {
	return sm
}
