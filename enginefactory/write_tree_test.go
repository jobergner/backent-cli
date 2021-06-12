package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes ReferencedDataStatus", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeReferencedDataStatus()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_ReferencedDataStatus_type,
			_ReferencedDataModified_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes elementKinds", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeElementKinds()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_ElementKind_type,
			_ElementKindEquipmentSet_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes tree elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTreeElements()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_EquipmentSet_type,
			_EquipmentSetReference_type,
			_GearScore_type,
			_GearScoreReference_type,
			_Item_type,
			_ItemReference_type,
			_Player_type,
			_PlayerReference_type,
			_Position_type,
			_PositionReference_type,
			_Zone_type,
			_ZoneReference_type,
			_ZoneItem_type,
			_ZoneItemReference_type,
			_AnyOfPlayer_ZoneItemReference_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes tree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTree()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Tree_type,
			newTree_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes recursionCheck", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRecursionCheck()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			recursionCheck_type,
			newRecursionCheck_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
