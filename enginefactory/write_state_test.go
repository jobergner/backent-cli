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
			GearScoreID_type,
			ItemID_type,
			PlayerID_type,
			PositionID_type,
			ZoneID_type,
			ZoneItemID_type,
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
			State_type,
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
			GearScore_type,
			itemCore_type,
			Item_type,
			playerCore_type,
			Player_type,
			positionCore_type,
			Position_type,
			zoneCore_type,
			Zone_type,
			zoneItemCore_type,
			ZoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
