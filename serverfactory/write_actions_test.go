package serverfactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/testutils"
)

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeActions()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Actions_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
