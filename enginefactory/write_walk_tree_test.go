package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteWalkTree(t *testing.T) {
	t.Run("writes walk element", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeWalkElement()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			walkEquipmentSet_Engine_func,
			walkGearScore_Engine_func,
			walkItem_Engine_func,
			walkPlayer_Engine_func,
			walkPosition_Engine_func,
			walkZone_Engine_func,
			walkZoneItem_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
