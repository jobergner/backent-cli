package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAssembleTree(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssembleTree()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			assembleUpdateTree_Engine_func,
			assembleFullTree_Engine_func,
			assembleTree_Engine_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
