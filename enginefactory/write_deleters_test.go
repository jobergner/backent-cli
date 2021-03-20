package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteDeleters(t *testing.T) {
	t.Run("writes deleters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeleters()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			DeleteGearScore_Engine_func,
			deleteGearScore_Engine_func,
			DeleteItem_Engine_func,
			deleteItem_Engine_func,
			DeletePlayer_Engine_func,
			deletePlayer_Engine_func,
			DeletePosition_Engine_func,
			deletePosition_Engine_func,
			DeleteZone_Engine_func,
			deleteZone_Engine_func,
			DeleteZoneItem_Engine_func,
			deleteZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
