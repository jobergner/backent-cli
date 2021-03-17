package statefactory

import (
	"strings"
	"testing"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes tree elements", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTreeElements()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			_gearScore_type,
			_item_type,
			_player_type,
			_position_type,
			_zone_type,
			_zoneItem_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes tree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeTree()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			Tree_type,
			newTree_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}