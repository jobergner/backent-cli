package server

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteProcessClientMessage(t *testing.T) {
	t.Run("writes processClientMessage", func(t *testing.T) {
		sf := newServerFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeProcessClientMessage()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			processClientMessage_Room_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
