package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWriteAny(t *testing.T) {
	t.Run("writes any", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAny()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Kind_anyOfPlayer_Position_func,
			_SetPlayer_anyOfPlayer_Position_func,
			setPlayer_anyOfPlayer_PositionCore_func,
			_SetPosition_anyOfPlayer_Position_func,
			setPosition_anyOfPlayer_PositionCore_func,
			deleteChild_anyOfPlayer_PositionCore_func,
			_Kind_anyOfPlayer_ZoneItem_func,
			_SetPlayer_anyOfPlayer_ZoneItem_func,
			setPlayer_anyOfPlayer_ZoneItemCore_func,
			_SetZoneItem_anyOfPlayer_ZoneItem_func,
			setZoneItem_anyOfPlayer_ZoneItemCore_func,
			deleteChild_anyOfPlayer_ZoneItemCore_func,
			_Kind_anyOfItem_Player_ZoneItem_func,
			_SetItem_anyOfItem_Player_ZoneItem_func,
			setItem_anyOfItem_Player_ZoneItemCore_func,
			_SetPlayer_anyOfItem_Player_ZoneItem_func,
			setPlayer_anyOfItem_Player_ZoneItemCore_func,
			_SetZoneItem_anyOfItem_Player_ZoneItem_func,
			setZoneItem_anyOfItem_Player_ZoneItemCore_func,
			deleteChild_anyOfItem_Player_ZoneItemCore_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
