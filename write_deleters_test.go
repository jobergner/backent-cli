package statefactory

import (
	"strings"
	"testing"
)

func TestWriteDeleters(t *testing.T) {
	t.Run("writes deleters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeDeleters()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			DeleteGearScore_StateMachine_func,
			DeleteItem_StateMachine_func,
			DeletePlayer_StateMachine_func,
			DeletePosition_StateMachine_func,
			DeleteZone_StateMachine_func,
			DeleteZoneItem_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
