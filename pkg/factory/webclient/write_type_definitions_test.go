package webclient

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
	"github.com/stretchr/testify/assert"
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

		fmt.Println(actual)
		assert.Equal(t, expected, actual)
	})
}
