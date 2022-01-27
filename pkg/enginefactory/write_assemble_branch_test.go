package enginefactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAssembleBranch(t *testing.T) {
	t.Run("writes assemblers", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssembleBranch()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
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
