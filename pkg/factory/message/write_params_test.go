package message

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteParameters(t *testing.T) {
	t.Run("writes parameters", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeParameters()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AddItemToPlayerParams_type,
			_MovePlayerParams_type,
			_SpawnZoneItemsParams_type,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
