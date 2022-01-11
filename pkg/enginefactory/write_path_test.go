package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWritePath(t *testing.T) {
	t.Run("writes path identifiers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeIdentifiers()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			equipmentSetIdentifier_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes path", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePath()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			path_type,
			newPath_func,
			toJSONPath_path_func,
			pathIdentifierToString_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
