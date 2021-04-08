package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRemovers()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_RemoveItems_Player_func,
			_RemoveItems_Zone_func,
			_RemovePlayers_Zone_func,
			_RemoveTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
