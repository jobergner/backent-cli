package statefactory

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
			CreateGearScore_StateMachine_func,
			createGearScore_StateMachine_func,
			CreateItem_StateMachine_func,
			createItem_StateMachine_func,
			CreatePlayer_StateMachine_func,
			createPlayer_StateMachine_func,
			CreatePosition_StateMachine_func,
			createPosition_StateMachine_func,
			CreateZone_StateMachine_func,
			createZone_StateMachine_func,
			CreateZoneItem_StateMachine_func,
			createZoneItem_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
