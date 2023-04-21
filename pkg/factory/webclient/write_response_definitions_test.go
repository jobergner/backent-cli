package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteResponseDefinitions(t *testing.T) {
	t.Run("writes response definitions", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeResponseDefinitions()

		actual := sf.file.String()
		expected := strings.Join([]string{
			interface_AddItemToPlayerResponse,
			interface_SpawnZoneItemsResponse + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
