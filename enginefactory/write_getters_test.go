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
			_GearScore_Engine_func,
			_ID_gearScore_func,
			_Level_gearScore_func,
			_Score_gearScore_func,
			_Item_Engine_func,
			_ID_item_func,
			_GearScore_item_func,
			_Player_Engine_func,
			_ID_player_func,
			_GearScore_player_func,
			_Items_player_func,
			_Position_player_func,
			_Position_Engine_func,
			_ID_position_func,
			_X_position_func,
			_Y_position_func,
			_Zone_Engine_func,
			_ID_zone_func,
			_Items_zone_func,
			_Players_zone_func,
			_Tags_zone_func,
			_ZoneItem_Engine_func,
			_ID_zoneItem_func,
			_Item_zoneItem_func,
			_Position_zoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
