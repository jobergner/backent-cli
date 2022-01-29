package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAssembleTree(t *testing.T) {
	t.Run("writes assembleTree", func(t *testing.T) {
		sf := newStateFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssembleTree()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			assembleUpdateTree_Engine_func,
			assembleFullTree_Engine_func,
			assembleTree_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
