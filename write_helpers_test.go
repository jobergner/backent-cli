package statefactory

import (
	"strings"
	"testing"
)

func TestWriteHelpers(t *testing.T) {
	t.Run("writes deduplicate", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeduplicate()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			deduplicateGearScoreIDs_func,
			deduplicateItemIDs_func,
			deduplicatePlayerIDs_func,
			deduplicatePositionIDs_func,
			deduplicateZoneIDs_func,
			deduplicateZoneItemIDs_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
