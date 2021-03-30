package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteCreators(t *testing.T) {
	t.Run("writes creators", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeCreators()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_CreateGearScore_Engine_func,
			createGearScore_Engine_func,
			_CreateItem_Engine_func,
			createItem_Engine_func,
			_CreatePlayer_Engine_func,
			createPlayer_Engine_func,
			_CreatePosition_Engine_func,
			createPosition_Engine_func,
			_CreateZone_Engine_func,
			createZone_Engine_func,
			_CreateZoneItem_Engine_func,
			createZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
