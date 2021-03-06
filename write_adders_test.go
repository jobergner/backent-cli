package statefactory

import (
	"strings"
	"testing"
)

func TestWriteAdders(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAdders()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			AddItem_Player_func,
			AddZoneItem_Zone_func,
			AddPlayer_Zone_func,
			AddTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
