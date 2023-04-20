package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteEmitElement(t *testing.T) {
	t.Run("writes emit element", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeEmitElement()

		actual := sf.file.String()
		expected := strings.Join([]string{
			function_emitAttackEvent,
			function_emitEquipmentSet,
			function_emitGearScore,
			function_emitItem,
			function_emitPlayer,
			function_emitPosition,
			function_emitZone,
			function_emitZoneItem,
			function_emitElementReference + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
