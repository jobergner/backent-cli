package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteRemovers(t *testing.T) {
	t.Run("writes removers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeRemovers()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_RemoveItems_Player_func,
			_RemoveItems_Zone_func,
			_RemovePlayers_Zone_func,
			_RemoveTags_Zone_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
