package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeSetters()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
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

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
