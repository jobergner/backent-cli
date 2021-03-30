package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteGetters(t *testing.T) {
	t.Run("writes getters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeGetters()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_GearScore_Engine_func,
			_ID_GearScore_func,
			_Level_GearScore_func,
			_Score_GearScore_func,
			_Item_Engine_func,
			_ID_Item_func,
			_GearScore_Item_func,
			_Player_Engine_func,
			_ID_Player_func,
			_GearScore_Player_func,
			_Items_Player_func,
			_Position_Player_func,
			_Position_Engine_func,
			_ID_Position_func,
			_X_Position_func,
			_Y_Position_func,
			_Zone_Engine_func,
			_ID_Zone_func,
			_Items_Zone_func,
			_Players_Zone_func,
			_Tags_Zone_func,
			_ZoneItem_Engine_func,
			_ID_ZoneItem_func,
			_Item_ZoneItem_func,
			_Position_ZoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
