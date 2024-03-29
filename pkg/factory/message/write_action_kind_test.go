package message

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteMessageKinds(t *testing.T) {
	t.Run("writes message kinds", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeMessageKinds()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_MessageKindAction_addItemToPlayer_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
