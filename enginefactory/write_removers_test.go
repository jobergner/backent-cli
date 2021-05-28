package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRemovers()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_RemoveEquipment_equipmentSet_func,
			_RemoveEquipmentSets_player_func,
			_RemoveGuildMembers_player_func,
			_RemoveItems_player_func,
			_RemoveTargetedByPlayer_player_func,
			_RemoveTargetedByZoneItem_player_func,
			_RemoveItems_zone_func,
			_RemovePlayers_zone_func,
			_RemoveTags_zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
