package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteClient(t *testing.T) {
	t.Run("writes client", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeClient()

		actual := sf.file.String()
		expected := strings.Join([]string{
			class_Client + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
