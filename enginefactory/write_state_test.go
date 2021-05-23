package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteState(t *testing.T) {
	t.Run("writes ids", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeIDs()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_EquipmentSetID_type,
			_GearScoreID_type,
			_ItemID_type,
			_PlayerID_type,
			_PositionID_type,
			_ZoneID_type,
			_ZoneItemID_type,
			_EquipmentSetEquipmentRefID_type,
			_ItemBoundToRefID_type,
			_PlayerEquipmentSetRefID_type,
			_PlayerGuildMemberRefID_type,
			_PlayerTargetRefID_type,
			_PlayerTargetedByRefID_type,
			_AnyOfPlayerPositionID_type,
			_AnyOfPlayerZoneItemID_type,
			_AnyOfItemPlayerZoneItemID_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes state", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeState()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_State_type,
			newState_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeElements()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			gearScoreCore_type,
			gearScore_type,
			itemCore_type,
			item_type,
			playerCore_type,
			player_type,
			positionCore_type,
			position_type,
			zoneCore_type,
			zone_type,
			zoneItemCore_type,
			zoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
