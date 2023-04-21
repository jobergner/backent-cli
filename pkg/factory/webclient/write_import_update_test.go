package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteImportUpdate(t *testing.T) {
	t.Run("writes import update", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeImportUpdate()

		actual := sf.file.String()
		expected := strings.Join([]string{
			function_import_Update + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
