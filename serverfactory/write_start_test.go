package serverfactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/testutils"
)

func TestWriteStart(t *testing.T) {
	t.Run("writes start", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeStart()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_Start_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
