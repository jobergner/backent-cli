package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteElementKind(t *testing.T) {
	t.Run("writes ElementKind", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeElementKind()

		actual := sf.file.String()
		expected := strings.Join([]string{
			enum_ElementKind + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
