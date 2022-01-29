package serverfactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteResponses(t *testing.T) {
	t.Run("writes responses", func(t *testing.T) {
		sf := newServerFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeResponses()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AddItemToPlayerResponse_type,
			_SpawnZoneItemsResponse_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
