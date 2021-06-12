package enginefactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWritePath(t *testing.T) {
	t.Run("writes path track", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePathTrack()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			pathTrack_type,
			newPathTrack_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes path identifiers", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeIdentifiers()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			equipmentSetIdentifier_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes path segment", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePathSegments()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			equipmentSet_path_func,
			gearScore_path_func,
			item_path_func,
			origin_path_func,
			player_path_func,
			items_path_func,
			position_path_func,
			zone_path_func,
			interactables_path_func,
			players_path_func,
			zoneItem_path_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes path", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writePath()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			path_type,
			newPath_func,
			index_path_func,
			equals_path_func,
			toJSONPath_path_func,
			pathIdentifierToString_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
