package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes tree elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTreeElements()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			tGearScore_type,
			tItem_type,
			tPlayer_type,
			tPosition_type,
			tZone_type,
			tZoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes tree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTree()

		actual := utils.FormatCode(sf.buf.String())
		expected := utils.FormatCode(strings.Join([]string{
			_Tree_type,
			newTree_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
