package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIDTypes(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.addIdTypes())
		expected := []string{
			`type personID int`,
			`type nameID int`,
			_personDeclaration,
			_nameDeclaration,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) addIdTypes() *stateMachine {
	return sm
}
