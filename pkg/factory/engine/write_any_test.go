package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAny(t *testing.T) {
	t.Run("writes any", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAny()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Kind_AnyOfPlayer_Position_func,
			_BePlayer_AnyOfPlayer_Position_func,
			bePlayer_anyOfPlayer_PositionCore_func,
			_BePosition_AnyOfPlayer_Position_func,
			bePosition_anyOfPlayer_PositionCore_func,
			deleteChild_anyOfPlayer_PositionCore_func,
			_Kind_AnyOfPlayer_ZoneItem_func,
			_BePlayer_AnyOfPlayer_ZoneItem_func,
			bePlayer_anyOfPlayer_ZoneItemCore_func,
			_BeZoneItem_AnyOfPlayer_ZoneItem_func,
			beZoneItem_anyOfPlayer_ZoneItemCore_func,
			deleteChild_anyOfPlayer_ZoneItemCore_func,
			_Kind_AnyOfItem_Player_ZoneItem_func,
			_BeItem_AnyOfItem_Player_ZoneItem_func,
			beItem_anyOfItem_Player_ZoneItemCore_func,
			_BePlayer_AnyOfItem_Player_ZoneItem_func,
			bePlayer_anyOfItem_Player_ZoneItemCore_func,
			_BeZoneItem_AnyOfItem_Player_ZoneItem_func,
			beZoneItem_anyOfItem_Player_ZoneItemCore_func,
			deleteChild_anyOfItem_Player_ZoneItemCore_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes anyRefs", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAnyRefs()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			anyOfPlayer_PositionRef_type,
			_Kind_anyOfPlayer_PositionRef_func,
			_Player_anyOfPlayer_PositionRef_func,
			_Position_anyOfPlayer_PositionRef_func,
			anyOfPlayer_ZoneItemRef_type,
			_Kind_anyOfPlayer_ZoneItemRef_func,
			_Player_anyOfPlayer_ZoneItemRef_func,
			_ZoneItem_anyOfPlayer_ZoneItemRef_func,
			anyOfItem_Player_ZoneItemRef_type,
			_Kind_anyOfItem_Player_ZoneItemRef_func,
			_Item_anyOfItem_Player_ZoneItemRef_func,
			_Player_anyOfItem_Player_ZoneItemRef_func,
			_ZoneItem_anyOfItem_Player_ZoneItemRef_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
