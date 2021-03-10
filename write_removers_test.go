package statefactory

import (
	"strings"
	"testing"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRemovers()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			RemoveItems_Player_func,
			RemoveItems_Zone_func,
			RemovePlayers_Zone_func,
			RemoveTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
