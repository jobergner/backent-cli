package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRemovers()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
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
