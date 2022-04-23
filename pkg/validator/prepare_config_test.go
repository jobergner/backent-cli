package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareConfig(t *testing.T) {
	t.Run("should remove __event__ field", func(t *testing.T) {
		data := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar":         "string",
				"__event__":   "true",
				"__prevent__": "true", // expected to remain
			},
		}

		// we ignore errors here as it's just the structural validation
		actualPreparedConfig := prepareStateConfig(data)

		expectedPrepareConfig := map[interface{}]interface{}{
			"foo": map[interface{}]interface{}{
				"bar":         "string",
				"__prevent__": "true",
			},
		}

		assert.Equal(t, expectedPrepareConfig, actualPreparedConfig)
	})
}
