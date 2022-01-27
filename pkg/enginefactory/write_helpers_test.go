package enginefactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteHelpers(t *testing.T) {
	t.Run("writes deduplicate", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeDeduplicate()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			deduplicateEquipmentSetIDs_func,
			deduplicateGearScoreIDs_func,
			deduplicateItemIDs_func,
			deduplicatePlayerIDs_func,
			deduplicatePositionIDs_func,
			deduplicateZoneIDs_func,
			deduplicateZoneItemIDs_func,
			deduplicateEquipmentSetEquipmentRefIDs_func,
			deduplicateItemBoundToRefIDs_func,
			deduplicatePlayerEquipmentSetRefIDs_func,
			deduplicatePlayerGuildMemberRefIDs_func,
			deduplicatePlayerTargetRefIDs_func,
			deduplicatePlayerTargetedByRefIDs_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes allIDs methods", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAllIDsMethod()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			allEquipmentSetIDs_Engine_func,
			allGearScoreIDs_Engine_func,
			allItemIDs_Engine_func,
			allPlayerIDs_Engine_func,
			allPositionIDs_Engine_func,
			allZoneIDs_Engine_func,
			allZoneItemIDs_Engine_func,
			allEquipmentSetEquipmentRefIDs_Engine_func,
			allItemBoundToRefIDs_Engine_func,
			allPlayerEquipmentSetRefIDs_Engine_func,
			allPlayerGuildMemberRefIDs_Engine_func,
			allPlayerTargetRefIDs_Engine_func,
			allPlayerTargetedByRefIDs_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
