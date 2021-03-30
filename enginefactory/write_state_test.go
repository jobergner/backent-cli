package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteState(t *testing.T) {
	t.Run("writes ids", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeIDs()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_GearScoreID_type,
			_ItemID_type,
			_PlayerID_type,
			_PositionID_type,
			_ZoneID_type,
			_ZoneItemID_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes state", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeState()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_State_type,
			newState_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeElements()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
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
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
