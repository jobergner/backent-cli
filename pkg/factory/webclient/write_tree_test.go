package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteTree(t *testing.T) {
	t.Run("writes tree", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeTree()

		actual := sf.file.String()
		expected := strings.Join([]string{
			interface_Tree + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
