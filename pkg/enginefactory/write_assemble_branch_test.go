package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAssembleBranch(t *testing.T) {
	t.Run("writes adders", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleBranch()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleGearScorePath_Engine_func,
			assembleEquipmentSetPath_Engine_func,
			assembleGearScorePath_Engine_func,
			assembleItemPath_Engine_func,
			assemblePlayerPath_Engine_func,
			assemblePositionPath_Engine_func,
			assembleZonePath_Engine_func,
			assembleZoneItemPath_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
