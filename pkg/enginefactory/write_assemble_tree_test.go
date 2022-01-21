package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAssembleTree(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssembleTree()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assembleUpdateTree_Engine_func,
			assembleFullTree_Engine_func,
			assembleTree_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
