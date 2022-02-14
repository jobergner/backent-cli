package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteGetters(t *testing.T) {
	t.Run("writes getters", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeGetters()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Exists_AttackEvent_func,
			_QueryAttackEvents_Engine_func,
			_EveryAttackEvent_Engine_func,
			_AttackEvent_Engine_func,
			_ParentKind_AttackEvent_func,
			_ParentPlayer_AttackEvent_func,
			_ID_AttackEvent_func,
			_Path_AttackEvent_func,
			_Target_AttackEvent_func,
			_Exists_EquipmentSet_func,
			_QueryEquipmentSets_Engine_func,
			_EveryEquipmentSet_Engine_func,
			_EquipmentSet_Engine_func,
			_ID_EquipmentSet_func,
			_Path_EquipmentSet_func,
			_Equipment_EquipmentSet_func,
			_Name_EquipmentSet_func,
			_Exists_GearScore_func,
			_QueryGearScores_Engine_func,
			_EveryGearScore_Engine_func,
			_GearScore_Engine_func,
			_ParentKind_GearScore_func,
			_ParentItem_GearScore_func,
			_ParentPlayer_GearScore_func,
			_ID_GearScore_func,
			_Path_GearScore_func,
			_Level_GearScore_func,
			_Score_GearScore_func,
			_Exists_Item_func,
			_QueryItems_Engine_func,
			_EveryItem_Engine_func,
			_Item_Engine_func,
			_ParentKind_Item_func,
			_ParentPlayer_Item_func,
			_ParentZone_Item_func,
			_ParentZoneItem_Item_func,
			_ID_Item_func,
			_Path_Item_func,
			_BoundTo_Item_func,
			_GearScore_Item_func,
			_Name_Item_func,
			_Origin_Item_func,
			_Exists_Player_func,
			_QueryPlayers_Engine_func,
			_EveryPlayer_Engine_func,
			_Player_Engine_func,
			_ParentKind_Player_func,
			_ParentItem_Player_func,
			_ParentZone_Player_func,
			_ID_Player_func,
			_Path_Player_func,
			_Action_Player_func,
			_EquipmentSets_Player_func,
			_GearScore_Player_func,
			_GuildMembers_Player_func,
			_Items_Player_func,
			_Position_Player_func,
			_Target_Player_func,
			_TargetedBy_Player_func,
			_Exists_Position_func,
			_QueryPositions_Engine_func,
			_EveryPosition_Engine_func,
			_Position_Engine_func,
			_ParentKind_Position_func,
			_ParentItem_Position_func,
			_ParentPlayer_Position_func,
			_ParentZoneItem_Position_func,
			_ID_Position_func,
			_Path_Position_func,
			_X_Position_func,
			_Y_Position_func,
			_Exists_Zone_func,
			_QueryZones_Engine_func,
			_EveryZone_Engine_func,
			_Zone_Engine_func,
			_ID_Zone_func,
			_Path_Zone_func,
			_Interactables_Zone_func,
			_Items_Zone_func,
			_Players_Zone_func,
			_Tags_Zone_func,
			_Exists_ZoneItem_func,
			_QueryZoneItems_Engine_func,
			_EveryZoneItem_Engine_func,
			_ZoneItem_Engine_func,
			_ParentKind_ZoneItem_func,
			_ParentZone_ZoneItem_func,
			_ID_ZoneItem_func,
			_Path_ZoneItem_func,
			_Item_ZoneItem_func,
			_Position_ZoneItem_func,
			attackEventTargetRef_Engine_func,
			_ID_AttackEventTargetRef_func,
			equipmentSetEquipmentRef_Engine_func,
			_ID_EquipmentSetEquipmentRef_func,
			itemBoundToRef_Engine_func,
			_ID_ItemBoundToRef_func,
			playerEquipmentSetRef_Engine_func,
			_ID_PlayerEquipmentSetRef_func,
			playerGuildMemberRef_Engine_func,
			_ID_PlayerGuildMemberRef_func,
			playerTargetRef_Engine_func,
			_ID_PlayerTargetRef_func,
			playerTargetedByRef_Engine_func,
			_ID_PlayerTargetedByRef_func,
			anyOfPlayer_Position_Engine_func,
			_ID_AnyOfPlayer_Position_func,
			_Player_AnyOfPlayer_Position_func,
			_Position_AnyOfPlayer_Position_func,
			anyOfPlayer_ZoneItem_Engine_func,
			_ID_AnyOfPlayer_ZoneItem_func,
			_Player_AnyOfPlayer_ZoneItem_func,
			_ZoneItem_AnyOfPlayer_ZoneItem_func,
			anyOfItem_Player_ZoneItem_Engine_func,
			_ID_AnyOfItem_Player_ZoneItem_func,
			_Item_AnyOfItem_Player_ZoneItem_func,
			_Player_AnyOfItem_Player_ZoneItem_func,
			_ZoneItem_AnyOfItem_Player_ZoneItem_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
