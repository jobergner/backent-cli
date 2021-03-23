package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteAssembling(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTree()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			assembleTree_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble tree element", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTreeElement()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			assembleGearScore_Engine_func,
			assembleItem_Engine_func,
			assemblePlayer_Engine_func,
			assemblePosition_Engine_func,
			assembleZone_Engine_func,
			assembleZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
