package statefactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type metaField struct {
	name        string
	typeLiteral string
}

func TestEmbedMetaFields(t *testing.T) {
	t.Run("should embed meta fields in all structs", func(t *testing.T) {
		input := unsafeParseDecls([]string{
			_personDeclaration,
			_nameDeclaration,
		})

		actual := splitPrintedDeclarations(input.embedMetaFields([]metaField{{"lastModified", "int64"}}))
		expected := []string{`
type person struct {
	name name
	lastModified int64
}`, `
type name struct {
	first string
	last string
	lastModified int64
}`,
		}

		missingDeclarations, redundantDeclarations := matchDeclarations(actual, expected)

		assert.Equal(t, []string{}, missingDeclarations)
		assert.Equal(t, []string{}, redundantDeclarations)
	})
}

func (sm *stateMachine) embedMetaFields(metaFields []metaField) *stateMachine {
	return sm
}
