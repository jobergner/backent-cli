package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteReference(t *testing.T) {
	t.Run("writes reference", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeReference()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Get_AttackEventTargetRef_func,
			_IsSet_AttackEventTargetRef_func,
			_Unset_AttackEventTargetRef_func,
			_Get_EquipmentSetEquipmentRef_func,
			_IsSet_ItemBoundToRef_func,
			_Unset_ItemBoundToRef_func,
			_Get_ItemBoundToRef_func,
			_Get_PlayerEquipmentSetRef_func,
			_Get_PlayerGuildMemberRef_func,
			_IsSet_PlayerTargetRef_func,
			_Unset_PlayerTargetRef_func,
			_Get_PlayerTargetRef_func,
			_Get_PlayerTargetedByRef_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes dereference", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeDereference()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			dereferenceAttackEventTargetRefs_Engine_func,
			dereferenceEquipmentSetEquipmentRefs_Engine_func,
			dereferenceItemBoundToRefs_Engine_func,
			dereferencePlayerEquipmentSetRefs_Engine_func,
			dereferencePlayerGuildMemberRefs_Engine_func,
			dereferencePlayerTargetRefsPlayer_Engine_func,
			dereferencePlayerTargetRefsZoneItem_Engine_func,
			dereferencePlayerTargetedByRefsPlayer_Engine_func,
			dereferencePlayerTargetedByRefsZoneItem_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
