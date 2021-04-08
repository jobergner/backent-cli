package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteDeleters(t *testing.T) {
	t.Run("writes deleters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeleters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_DeleteGearScore_Engine_func,
			deleteGearScore_Engine_func,
			_DeleteItem_Engine_func,
			deleteItem_Engine_func,
			_DeletePlayer_Engine_func,
			deletePlayer_Engine_func,
			_DeletePosition_Engine_func,
			deletePosition_Engine_func,
			_DeleteZone_Engine_func,
			deleteZone_Engine_func,
			_DeleteZoneItem_Engine_func,
			deleteZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
