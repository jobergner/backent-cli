package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAssembling(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTree()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleConfig_type,
			assembleTree_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble tree element", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeElement()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleEquipmentSet_Engine_func,
			assembleGearScore_Engine_func,
			assembleItem_Engine_func,
			assemblePlayer_Engine_func,
			assemblePosition_Engine_func,
			assembleZone_Engine_func,
			assembleZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble tree reference", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeReference()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleEquipmentSetEquipmentRef_Engine_func,
			assembleItemBoundToRef_Engine_func,
			assemblePlayerEquipmentSetRef_Engine_func,
			assemblePlayerGuildMemberRef_Engine_func,
			assemblePlayerTargetRef_Engine_func,
			assemblePlayerTargetedByRef_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
