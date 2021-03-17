package statefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes tree elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTreeElements()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			_gearScore_type,
			_item_type,
			_player_type,
			_position_type,
			_zone_type,
			_zoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes tree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTree()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			Tree_type,
			newTree_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
