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
			deduplicateEquipmentSetEquipmentRefIDs_func,
			deduplicateItemBoundToRefIDs_func,
			deduplicatePlayerEquipmentSetRefIDs_func,
			deduplicatePlayerGuildMemberRefIDs_func,
			deduplicatePlayerTargetRefIDs_func,
			deduplicatePlayerTargetedByRefIDs_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
