package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOutStructTree(t *testing.T) {
	t.Run("should replace object references with ids", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.evenOutStructTree())
		expected := []string{`
type person struct {
	name nameID
}`, `
type name struct {
	first string
	last string
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) evenOutStructTree() *stateMachine {
	return sm
}
