package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeSetters()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			setBoolValue_Engine_func,
			setFloatValue_Engine_func,
			setIntValue_Engine_func,
			setStringValue_Engine_func,
			_SetName_EquipmentSet_func,
			_SetLevel_GearScore_func,
			_SetScore_GearScore_func,
			_SetName_Item_func,
			_SetX_Position_func,
			_SetY_Position_func,
			_SetTarget_AttackEvent_func,
			_SetBoundTo_Item_func,
			_SetTargetPlayer_Player_func,
			_SetTargetZoneItem_Player_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
