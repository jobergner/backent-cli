package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteAdders(t *testing.T) {
	t.Run("writes adders", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAdders()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_AddItem_Player_func,
			_AddItem_Zone_func,
			_AddPlayer_Zone_func,
			_AddTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
