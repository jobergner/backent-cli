package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
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
			_AnyOfPlayer_PositionID_type,
			_AnyOfPlayer_ZoneItemID_type,
			_AnyOfItem_Player_ZoneItemID_type,
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
			equipmentSetCore_type,
			_EquipmentSet_type,
			gearScoreCore_type,
			_GearScore_type,
			itemCore_type,
			_Item_type,
			playerCore_type,
			_Player_type,
			positionCore_type,
			_Position_type,
			zoneCore_type,
			_Zone_type,
			zoneItemCore_type,
			_ZoneItem_type,
			equipmentSetEquipmentRefCore_type,
			_EquipmentSetEquipmentRef_type,
			itemBoundToRefCore_type,
			_ItemBoundToRef_type,
			playerEquipmentSetRefCore_type,
			_PlayerEquipmentSetRef_type,
			playerGuildMemberRefCore_type,
			_PlayerGuildMemberRef_type,
			playerTargetRefCore_type,
			_PlayerTargetRef_type,
			playerTargetedByRefCore_type,
			_PlayerTargetedByRef_type,
			anyOfPlayer_PositionCore_type,
			_AnyOfPlayer_Position_type,
			anyOfPlayer_ZoneItemCore_type,
			_AnyOfPlayer_ZoneItem_type,
			anyOfItem_Player_ZoneItemCore_type,
			_AnyOfItem_Player_ZoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
