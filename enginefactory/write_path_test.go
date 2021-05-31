package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWritePath(t *testing.T) {
	t.Run("writes path track", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePathTrack()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			pathTrack_type,
			newPathTrack_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
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
}
