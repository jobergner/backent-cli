package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAdders(t *testing.T) {
	t.Run("writes adders", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeTypeDefinitions()

		actual := sf.file.String()
		expected := strings.Join([]string{
			interface_AttackEvent,
			interface_EquipmentSet,
			interface_GearScore,
			interface_Item,
			interface_Player,
			interface_Position,
			interface_Zone,
			interface_ZoneItem,
			interface_ElementReference,
		}, "\n")

		diffs := testutils.PrettyDiffText(actual, expected)
		if len(diffs) > 0 {
			t.Errorf(diffs)
		}
	})
}
