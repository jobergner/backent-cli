package server

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteTriggerAction(t *testing.T) {
	t.Run("writes triggerAction", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeTriggerAction()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			triggerAction_Room_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
