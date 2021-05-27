package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteAdders(t *testing.T) {
	t.Run("writes adders", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAdders()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_AddEquipment_equipmentSet_func,
			_AddEquipmentSet_player_func,
			_AddGuildMember_player_func,
			_AddItem_player_func,
			_AddTargetedByPlayer_player_func,
			_AddTargetedByZoneItem_player_func,
			_AddInteractableItem_zone_func,
			_AddInteractablePlayer_zone_func,
			_AddInteractableZoneItem_zone_func,
			_AddItem_zone_func,
			_AddPlayer_zone_func,
			_AddTags_zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
