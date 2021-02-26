package statefactory

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractStructNames(t *testing.T) {
	t.Run("should find all struct names in file", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			// input_person_type,
			// input_child_type,
			// input_name_type,
		})

		smb := newStateMachineBuilder(input)
		actual := extractStructNames(smb.input)
		expected := []string{"person", "name", "child"}

		assert.Equal(t, expected, actual)
	})
}

func extractStructNames(sm *ast.File) []string {
	return []string{}
}
