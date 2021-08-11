package serverfactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteResponses(t *testing.T) {
	t.Run("writes parameters", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeResponses()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_AddItemToPlayerResponse_type,
			_SpawnZoneItemsResponse_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
