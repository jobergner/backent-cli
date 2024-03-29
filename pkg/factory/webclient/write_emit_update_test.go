package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteEmitUpdate(t *testing.T) {
	t.Run("writes emit update", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeEmitUpdate()

		actual := sf.file.String()
		expected := strings.Join([]string{
			function_emit_Update + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
