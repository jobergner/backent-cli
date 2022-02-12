package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes ReferencedDataStatus", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeReferencedDataStatus()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_ReferencedDataStatus_type,
			_ReferencedDataModified_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes elementKinds", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeElementKinds()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_ElementKind_type,
			_ElementKindAttackEvent_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes tree elements", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeTreeElements()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			attackEvent_type,
			equipmentSet_type,
			gearScore_type,
			item_type,
			player_type,
			position_type,
			zone_type,
			zoneItem_type,
			elementReference_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes tree", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeTree()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Tree_type,
			newTree_func,
			clear_Tree_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
