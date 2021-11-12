package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
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
			equipmentSetReference_type,
			_GearScore_type,
			gearScoreReference_type,
			_Item_type,
			itemReference_type,
			_Player_type,
			playerReference_type,
			_Position_type,
			positionReference_type,
			_Zone_type,
			zoneReference_type,
			_ZoneItem_type,
			zoneItemReference_type,
			anyOfPlayer_ZoneItemReference_type,
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
	t.Run("writes assembleCache", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleCache()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleCache_type,
			newAssembleCache_func,
			equipmentSetCacheElement_type,
			gearScoreCacheElement_type,
			itemCacheElement_type,
			playerCacheElement_type,
			positionCacheElement_type,
			zoneCacheElement_type,
			zoneItemCacheElement_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
