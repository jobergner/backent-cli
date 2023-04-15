package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteMessageKind(t *testing.T) {
	t.Run("writes MessageKind", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeMessageKind()

		actual := sf.file.String()
		expected := strings.Join([]string{
			enum_MessageKind + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
