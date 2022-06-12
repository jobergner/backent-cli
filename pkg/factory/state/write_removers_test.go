package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeRemovers()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_RemoveAction_Player_func,
			_RemoveEquipment_EquipmentSet_func,
			_RemoveEquipmentSet_Player_func,
			_RemoveGuildMember_Player_func,
			_RemoveItem_Player_func,
			_RemoveTargetedByPlayer_Player_func,
			_RemoveTargetedByZoneItem_Player_func,
			_RemoveInteractableItem_Zone_func,
			_RemoveInteractablePlayer_Zone_func,
			_RemoveInteractableZoneItem_Zone_func,
			_RemoveItem_Zone_func,
			_RemovePlayer_Zone_func,
			_RemoveTag_Zone_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
