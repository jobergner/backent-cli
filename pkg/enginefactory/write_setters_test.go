package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeSetters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_SetName_EquipmentSet_func,
			_SetLevel_GearScore_func,
			_SetScore_GearScore_func,
			_SetName_Item_func,
			_SetX_Position_func,
			_SetY_Position_func,
			_SetBoundTo_Item_func,
			_SetTargetPlayer_Player_func,
			_SetTargetZoneItem_Player_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
