package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedParentage(t *testing.T) {
	t.Run("should add parentage declaration", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.embedParentage())
		expected := []string{`
type person struct {
	name name
	parentage Parentage
}`, `
type name struct {
	first string
	last string
	parentage Parentage
}`,
			`type parentage []parentInfo`, `
type parentInfo struct {
	kind entityKind
	id int
}`,
			// TODO; entityKind for each entity
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) embedParentage() *stateMachine {
	return sm
}
