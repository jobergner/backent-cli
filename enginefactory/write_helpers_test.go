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
	t.Run("writes allIDs methods", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAllIDsMethod()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			allEquipmentSetEquipmentRefIDs_Engine_func,
			allItemBoundToRefIDs_Engine_func,
			allPlayerEquipmentSetRefIDs_Engine_func,
			allPlayerGuildMemberRefIDs_Engine_func,
			allPlayerTargetRefIDs_Engine_func,
			allPlayerTargetedByRefIDs_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
