package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyCombination(t *testing.T) {
	t.Run("should build anyOfTypeCombinator", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "anyOf<yee,oot>",
			},
			"soo": map[interface{}]interface{}{
				"saz": "anyOf<wii>",
				"sar": "string",
			},
		}

		cmb := newAnyOfTypeCombinator()
		cmb.build(data)

		expectedData := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
			},
			"soo": map[interface{}]interface{}{
				"sar": "string",
			},
		}

		assert.Equal(t, expectedData, cmb.data)

		assert.Contains(t, cmb.anyOfTypes, anyOfTypeIterator{"foo", "baz", []string{"yee", "oot"}, 0})
		assert.Contains(t, cmb.anyOfTypes, anyOfTypeIterator{"soo", "saz", []string{"wii"}, 0})
	})
}
