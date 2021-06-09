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
	t.Run("writes merge IDs", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeMergeIDs()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			mergeEquipmentSetIDs_func,
			mergeGearScoreIDs_func,
			mergeItemIDs_func,
			mergePlayerIDs_func,
			mergePositionIDs_func,
			mergeZoneIDs_func,
			mergeZoneItemIDs_func,
			mergeEquipmentSetEquipmentRefIDs_func,
			mergeItemBoundToRefIDs_func,
			mergePlayerEquipmentSetRefIDs_func,
			mergePlayerGuildMemberRefIDs_func,
			mergePlayerTargetRefIDs_func,
			mergePlayerTargetedByRefIDs_func,
			mergeAnyOfPlayer_PositionIDs_func,
			mergeAnyOfPlayer_ZoneItemIDs_func,
			mergeAnyOfItem_Player_ZoneItemIDs_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
