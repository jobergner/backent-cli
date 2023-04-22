package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteImportElement(t *testing.T) {
	t.Run("writes import element", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeImportElement()

		actual := sf.file.String()
		expected := strings.Join([]string{
			function_importEquipmentSet,
			function_importGearScore,
			function_importItem,
			function_importPlayer,
			function_importPosition,
			function_importZone,
			function_importZoneItem,
			function_importElementReference + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
