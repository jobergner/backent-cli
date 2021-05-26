package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteAny(t *testing.T) {
	t.Run("writes any", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAny()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Kind_anyOfPlayerPosition_func,
			_SetPlayer_anyOfPlayerPosition_func,
			setPlayer_anyOfPlayerPositionCore_func,
			_SetPosition_anyOfPlayerPosition_func,
			setPosition_anyOfPlayerPositionCore_func,
			deleteChild_anyOfPlayerPositionCore_func,
			_Kind_anyOfPlayerZoneItem_func,
			_SetPlayer_anyOfPlayerZoneItem_func,
			setPlayer_anyOfPlayerZoneItemCore_func,
			_SetZoneItem_anyOfPlayerZoneItem_func,
			setZoneItem_anyOfPlayerZoneItemCore_func,
			deleteChild_anyOfPlayerZoneItemCore_func,
			_Kind_anyOfItemPlayerZoneItem_func,
			_SetItem_anyOfItemPlayerZoneItem_func,
			setItem_anyOfItemPlayerZoneItemCore_func,
			_SetPlayer_anyOfItemPlayerZoneItem_func,
			setPlayer_anyOfItemPlayerZoneItemCore_func,
			_SetZoneItem_anyOfItemPlayerZoneItem_func,
			setZoneItem_anyOfItemPlayerZoneItemCore_func,
			deleteChild_anyOfItemPlayerZoneItemCore_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
