package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractStructNames(t *testing.T) {
	t.Run("should find all struct names in file", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(extractStructNames(input))
		expected := []string{"person", "name"}

		assert.Equal(t, expected, actual)
	})
}

func extractStructNames(sm *stateMachine) *stateMachine {
	return sm
}
