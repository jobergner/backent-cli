package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWriteReference(t *testing.T) {
	t.Run("writes reference", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeReference()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Get_equipmentSetEquipmentRef_func,
			_IsSet_itemBoundToRef_func,
			_Unset_itemBoundToRef_func,
			_Get_itemBoundToRef_func,
			_Get_playerEquipmentSetRef_func,
			_Get_playerGuildMemberRef_func,
			_IsSet_playerTargetRef_func,
			_Unset_playerTargetRef_func,
			_Get_playerTargetRef_func,
			_Get_playerTargetedByRef_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes dereference", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDereference()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			dereferenceEquipmentSetEquipmentRefs_Engine_func,
			dereferenceItemBoundToRefs_Engine_func,
			dereferencePlayerEquipmentSetRefs_Engine_func,
			dereferencePlayerGuildMemberRefs_Engine_func,
			dereferencePlayerTargetRefsPlayer_Engine_func,
			dereferencePlayerTargetRefsZoneItem_Engine_func,
			dereferencePlayerTargetedByRefsPlayer_Engine_func,
			dereferencePlayerTargetedByRefsZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
