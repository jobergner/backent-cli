package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteTypeDefinitions(t *testing.T) {
	t.Run("writes type definitions", func(t *testing.T) {
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
			interface_ElementReference + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
