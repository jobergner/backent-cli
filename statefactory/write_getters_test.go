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
			GearScore_StateMachine_func,
			ID_GearScore_func,
			Level_GearScore_func,
			Score_GearScore_func,
			Item_StateMachine_func,
			ID_Item_func,
			GearScore_Item_func,
			Player_StateMachine_func,
			ID_Player_func,
			GearScore_Player_func,
			Items_Player_func,
			Position_Player_func,
			Position_StateMachine_func,
			ID_Position_func,
			X_Position_func,
			Y_Position_func,
			Zone_StateMachine_func,
			ID_Zone_func,
			Items_Zone_func,
			Players_Zone_func,
			Tags_Zone_func,
			ZoneItem_StateMachine_func,
			ID_ZoneItem_func,
			Item_ZoneItem_func,
			Position_ZoneItem_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
