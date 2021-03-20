package actionsfactory

import (
	"testing"
)

func TestActionsFactory(t *testing.T) {
	t.Run("doesnt crash", func(t *testing.T) {
		WriteActionsFrom(testActionsConfig, "ownmod", "ownpackage")
	})
}
