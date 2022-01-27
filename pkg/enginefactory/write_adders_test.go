package enginefactory

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAdders(t *testing.T) {
	t.Run("writes adders", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAdders()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		fmt.Println(buf.String())
		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AddEquipment_EquipmentSet_func,
			_AddEquipmentSet_Player_func,
			_AddGuildMember_Player_func,
			_AddItem_Player_func,
			_AddTargetedByPlayer_Player_func,
			_AddTargetedByZoneItem_Player_func,
			_AddInteractableItem_Zone_func,
			_AddInteractablePlayer_Zone_func,
			_AddInteractableZoneItem_Zone_func,
			_AddItem_Zone_func,
			_AddPlayer_Zone_func,
			_AddTag_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
