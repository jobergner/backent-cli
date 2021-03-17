package statefactory

import (
	"strings"
	"bar-cli/utils"
	"testing"
)

func TestWriteAssembling(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTree()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			assembleTree_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble tree element", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeElement()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			assembleGearScore_StateMachine_func,
			assembleItem_StateMachine_func,
			assemblePlayer_StateMachine_func,
			assemblePosition_StateMachine_func,
			assembleZone_StateMachine_func,
			assembleZoneItem_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
