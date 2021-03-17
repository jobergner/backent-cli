package statefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteGetters(t *testing.T) {
	t.Run("writes getters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeGetters()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			GetGearScore_StateMachine_func,
			GetID_GearScore_func,
			GetLevel_GearScore_func,
			GetScore_GearScore_func,
			GetItem_StateMachine_func,
			GetID_Item_func,
			GetGearScore_Item_func,
			GetPlayer_StateMachine_func,
			GetID_Player_func,
			GetGearScore_Player_func,
			GetItems_Player_func,
			GetPosition_Player_func,
			GetPosition_StateMachine_func,
			GetID_Position_func,
			GetX_Position_func,
			GetY_Position_func,
			GetZone_StateMachine_func,
			GetID_Zone_func,
			GetItems_Zone_func,
			GetPlayers_Zone_func,
			GetTags_Zone_func,
			GetZoneItem_StateMachine_func,
			GetID_ZoneItem_func,
			GetItem_ZoneItem_func,
			GetPosition_ZoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
