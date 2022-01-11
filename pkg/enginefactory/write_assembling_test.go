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
			assembleUpdateTree_Engine_func,
			assembleFullTree_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble branch", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeReference()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleEquipmentSetPath_Engine_func,
			assembleGearScorePath_Engine_func,
			assembleItemPath_Engine_func,
			assemblePlayerPath_Engine_func,
			assemblePositionPath_Engine_func,
			assembleZoneItemPath_Engine_func,
			assembleZoneItemPath_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
