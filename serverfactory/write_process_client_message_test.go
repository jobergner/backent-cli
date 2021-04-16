package serverfactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteProcessClientMessage(t *testing.T) {
	t.Run("writes processClientMessage", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeProcessClientMessage()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			processClientMessage_Room_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
