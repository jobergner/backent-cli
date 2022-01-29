package server

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteMessageKinds(t *testing.T) {
	t.Run("writes message kinds", func(t *testing.T) {
		sf := newServerFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeMessageKinds()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_MessageKindAction_addItemToPlayer_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
