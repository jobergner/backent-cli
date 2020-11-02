package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.embedMetaFields())
		expected := []string{`
type person struct {
	person personID
	name name
	lastModified int64
	operationKind operationKind
}`, `
type name struct {
	person nameID
	first string
	last string
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
