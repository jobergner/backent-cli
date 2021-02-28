package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteState(t *testing.T) {
	t.Run("writes entityKinds", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeEntityKinds()

		actual := splitDecls(sf.buf.String())
		expected := []string{
			EntityKind_type,
			EntityKindGearScore_type,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}
