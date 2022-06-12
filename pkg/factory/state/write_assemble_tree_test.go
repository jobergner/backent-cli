package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAssembleTree(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeAssembleTree()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AssembleUpdateTree_Engine_func,
			_AssembleFullTree_Engine_func,
			assembleTree_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
