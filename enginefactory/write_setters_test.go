package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeSetters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_SetName_equipmentSet_func,
			_SetLevel_gearScore_func,
			_SetScore_gearScore_func,
			_SetName_item_func,
			_SetX_position_func,
			_SetY_position_func,
			_SetBoundTo_item_func,
			_SetTargetPlayer_player_func,
			_SetTargetZoneItem_player_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
