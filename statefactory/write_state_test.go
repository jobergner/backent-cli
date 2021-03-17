package statefactory

import (
	"strings"
	"testing"
)

func TestWriteState(t *testing.T) {
	t.Run("writes entityKinds", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeEntityKinds()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			EntityKind_type,
			EntityKindGearScore_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes ids", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeIDs()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			GearScoreID_type,
			ItemID_type,
			PlayerID_type,
			PositionID_type,
			ZoneID_type,
			ZoneItemID_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes state", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeState()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			State_type,
			newState_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeElements()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
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
			t.Errorf(diff(actual, expected))
		}
	})
}
