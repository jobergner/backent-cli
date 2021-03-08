package statefactory

import (
	"strings"
	"testing"
)

func TestWriteGetters(t *testing.T) {
	t.Run("writes getters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeGetters()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			GetGearScore_StateMachine_func,
			GetLevel_GearScore_func,
			GetScore_GearScore_func,
			GetItem_StateMachine_func,
			GetGearScore_Item_func,
			GetPlayer_StateMachine_func,
			GetGearScore_Player_func,
			GetItems_Player_func,
			GetPosition_Player_func,
			GetPosition_StateMachine_func,
			GetX_Position_func,
			GetY_Position_func,
			GetZone_StateMachine_func,
			GetItems_Zone_func,
			GetPlayers_Zone_func,
			GetTags_Zone_func,
			GetZoneItem_StateMachine_func,
			GetItem_ZoneItem_func,
			GetPosition_ZoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
