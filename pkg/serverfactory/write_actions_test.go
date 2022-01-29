package serverfactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		sf := newServerFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeActions()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_Actions_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
