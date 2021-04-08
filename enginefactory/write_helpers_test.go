package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteHelpers(t *testing.T) {
	t.Run("writes deduplicate", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeduplicate()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			deduplicateGearScoreIDs_func,
			deduplicateItemIDs_func,
			deduplicatePlayerIDs_func,
			deduplicatePositionIDs_func,
			deduplicateZoneIDs_func,
			deduplicateZoneItemIDs_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
