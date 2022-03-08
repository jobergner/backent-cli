package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteComplexID(t *testing.T) {
	t.Run("writes ComplexID", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeComplexID()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_ComplexID_type,
			complexIDStructCache_type,
			_MarshalJSON_ComplexID_func,
			_UnmarshalJSON_ComplexID_func,
			_MarshalJSON_AttackEventTargetRefID_func,
			_UnmarshalJSON_AttackEventTargetRefID_func,
			_MarshalJSON_PlayerGuildMemberRefID_func,
			_UnmarshalJSON_PlayerGuildMemberRefID_func,
			_MarshalJSON_ItemBoundToRefID_func,
			_UnmarshalJSON_ItemBoundToRefID_func,
			_MarshalJSON_EquipmentSetEquipmentRefID_func,
			_UnmarshalJSON_EquipmentSetEquipmentRefID_func,
			_MarshalJSON_PlayerEquipmentSetRefID_func,
			_UnmarshalJSON_PlayerEquipmentSetRefID_func,
			_MarshalJSON_AnyOfItem_Player_ZoneItemID_func,
			_UnmarshalJSON_AnyOfItem_Player_ZoneItemID_func,
			_MarshalJSON_AnyOfPlayer_ZoneItemID_func,
			_UnmarshalJSON_AnyOfPlayer_ZoneItemID_func,
			_MarshalJSON_AnyOfPlayer_PositionID_func,
			_UnmarshalJSON_AnyOfPlayer_PositionID_func,
			_MarshalJSON_PlayerTargetRefID_func,
			_UnmarshalJSON_PlayerTargetRefID_func,
			_MarshalJSON_PlayerTargetedByRefID_func,
			_UnmarshalJSON_PlayerTargetedByRefID_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
