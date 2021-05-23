package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteDeleters(t *testing.T) {
	t.Run("writes deleters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeleters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_DeleteEquipmentSet_Engine_func,
			deleteEquipmentSet_Engine_func,
			_DeleteGearScore_Engine_func,
			deleteGearScore_Engine_func,
			_DeleteItem_Engine_func,
			deleteItem_Engine_func,
			_DeletePlayer_Engine_func,
			deletePlayer_Engine_func,
			_DeletePosition_Engine_func,
			deletePosition_Engine_func,
			_DeleteZone_Engine_func,
			deleteZone_Engine_func,
			_DeleteZoneItem_Engine_func,
			deleteZoneItem_Engine_func,
			deleteEquipmentSetEquipmentRef_Engine_func,
			deleteItemBoundToRef_Engine_func,
			deletePlayerEquipmentSetRef_Engine_func,
			deletePlayerGuildMemberRef_Engine_func,
			deletePlayerTargetRef_Engine_func,
			deletePlayerTargetedByRef_Engine_func,
			deleteAnyOfPlayerPosition_Engine_func,
			deleteAnyOfPlayerZoneItem_Engine_func,
			deleteAnyOfItemPlayerZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
