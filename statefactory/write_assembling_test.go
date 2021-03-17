package statefactory

import (
	"strings"
	"testing"
)

func TestWriteAssembling(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTree()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			assembleTree_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes assemble tree element", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeElement()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			assembleGearScore_StateMachine_func,
			assembleItem_StateMachine_func,
			assemblePlayer_StateMachine_func,
			assemblePosition_StateMachine_func,
			assembleZone_StateMachine_func,
			assembleZoneItem_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
