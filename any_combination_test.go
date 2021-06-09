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
	t.Run("should generate combinations", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "anyOf<yee,oot>",
			},
			"soo": map[interface{}]interface{}{
				"saz": "anyOf<wii, loo>",
				"sar": "string",
				"sam": "anyOf<guu>",
			},
		}

		cmb := newAnyOfTypeCombinator()
		cmb.build(data)
		cmb.generateCombinations()

		assert.Equal(t, 4, len(cmb.dataCombinations))

		expectedData1 := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "yee",
			},
			"soo": map[interface{}]interface{}{
				"saz": "wii",
				"sar": "string",
				"sam": "guu",
			},
		}

		assert.Contains(t, cmb.dataCombinations, expectedData1)

		expectedData2 := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "oot",
			},
			"soo": map[interface{}]interface{}{
				"saz": "wii",
				"sar": "string",
				"sam": "guu",
			},
		}

		assert.Contains(t, cmb.dataCombinations, expectedData2)

		expectedData3 := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "yee",
			},
			"soo": map[interface{}]interface{}{
				"saz": "loo",
				"sar": "string",
				"sam": "guu",
			},
		}

		assert.Contains(t, cmb.dataCombinations, expectedData3)

		expectedData4 := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar": "string",
				"baz": "oot",
			},
			"soo": map[interface{}]interface{}{
				"saz": "loo",
				"sar": "string",
				"sam": "guu",
			},
		}

		assert.Contains(t, cmb.dataCombinations, expectedData4)
	})
}
