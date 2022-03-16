package client

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		sf := newFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeActions()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AddItemToPlayer_Client_func,
			_MovePlayer_Client_func,
			_SpawnZoneItems_Client_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
