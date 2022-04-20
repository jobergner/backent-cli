package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteState(t *testing.T) {
	t.Run("writes ids", func(t *testing.T) {
		sf := NewFactory(newSimpleASTExample())
		sf.writeIDs()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_BoolValueID_type,
			_FloatValueID_type,
			_IntValueID_type,
			_StringValueID_type,
			_AttackEventID_type,
			_EquipmentSetID_type,
			_GearScoreID_type,
			_ItemID_type,
			_PlayerID_type,
			_PositionID_type,
			_ZoneID_type,
			_ZoneItemID_type,
			_AttackEventTargetRefID_type,
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

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes state", func(t *testing.T) {
		sf := NewFactory(newSimpleASTExample())
		sf.writeState()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_State_type,
			newState_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes isEmpty", func(t *testing.T) {
		sf := NewFactory(newSimpleASTExample())
		sf.writeIsEmpty()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_IsEmpty_State_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes metaData", func(t *testing.T) {
		sf := NewFactory(newSimpleASTExample())
		sf.writeMetaData()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			metaData_type,
			unsign_metaData_func,
			sign_metaData_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes elements", func(t *testing.T) {
		sf := NewFactory(newSimpleASTExample())
		sf.writeElements()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			boolValue_type,
			floatValue_type,
			intValue_type,
			stringValue_type,
			attackEventCore_type,
			_AttackEvent_type,
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
			attackEventTargetRefCore_type,
			_AttackEventTargetRef_type,
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

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
