package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeRemovers()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_RemoveAction_Player_func,
			_RemoveEquipment_EquipmentSet_func,
			_RemoveEquipmentSets_Player_func,
			_RemoveGuildMembers_Player_func,
			_RemoveItems_Player_func,
			_RemoveTargetedByPlayer_Player_func,
			_RemoveTargetedByZoneItem_Player_func,
			_RemoveInteractablesItem_Zone_func,
			_RemoveInteractablesPlayer_Zone_func,
			_RemoveInteractablesZoneItem_Zone_func,
			_RemoveItems_Zone_func,
			_RemovePlayers_Zone_func,
			_RemoveTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
