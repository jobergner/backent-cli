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

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			CreateGearScore_Engine_func,
			createGearScore_Engine_func,
			CreateItem_Engine_func,
			createItem_Engine_func,
			CreatePlayer_Engine_func,
			createPlayer_Engine_func,
			CreatePosition_Engine_func,
			createPosition_Engine_func,
			CreateZone_Engine_func,
			createZone_Engine_func,
			CreateZoneItem_Engine_func,
			createZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
