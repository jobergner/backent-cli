package serverfactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeActions()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			messageKindAction_addItemToPlayer_type,
			actions_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
