package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteGetters(t *testing.T) {
	t.Run("writes getters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeGetters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_EquipmentSet_Engine_func,
			_ID_equipmentSet_func,
			_Equipment_equipmentSet_func,
			_Name_equipmentSet_func,
			_GearScore_Engine_func,
			_ID_gearScore_func,
			_Level_gearScore_func,
			_Score_gearScore_func,
			_Item_Engine_func,
			_ID_item_func,
			_BoundTo_item_func,
			_GearScore_item_func,
			_Name_item_func,
			_Origin_item_func,
			_Player_Engine_func,
			_ID_player_func,
			_EquipmentSets_player_func,
			_GearScore_player_func,
			_GuildMembers_player_func,
			_Items_player_func,
			_Position_player_func,
			_Target_player_func,
			_TargetedBy_player_func,
			_Position_Engine_func,
			_ID_position_func,
			_X_position_func,
			_Y_position_func,
			_Zone_Engine_func,
			_ID_zone_func,
			_Interactables_zone_func,
			_Items_zone_func,
			_Players_zone_func,
			_Tags_zone_func,
			_ZoneItem_Engine_func,
			_ID_zoneItem_func,
			_Item_zoneItem_func,
			_Position_zoneItem_func,
			equipmentSetEquipmentRef_Engine_func,
			_ID_equipmentSetEquipmentRef_func,
			itemBoundToRef_Engine_func,
			_ID_itemBoundToRef_func,
			playerEquipmentSetRef_Engine_func,
			_ID_playerEquipmentSetRef_func,
			playerGuildMemberRef_Engine_func,
			_ID_playerGuildMemberRef_func,
			playerTargetRef_Engine_func,
			_ID_playerTargetRef_func,
			playerTargetedByRef_Engine_func,
			_ID_playerTargetedByRef_func,
			anyOfPlayerPosition_Engine_func,
			_ID_anyOfPlayerPosition_func,
			anyOfPlayerZoneItem_Engine_func,
			_ID_anyOfPlayerZoneItem_func,
			anyOfItemPlayerZoneItem_Engine_func,
			_ID_anyOfItemPlayerZoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
