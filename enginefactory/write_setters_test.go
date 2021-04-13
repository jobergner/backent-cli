package enginefactory

import (
	"bar-cli/testutils"
	"strings"
	"testing"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeSetters()

		actual := sf.buf.String()
		expected := testutils.FormatCode(strings.Join([]string{
			_SetLevel_gearScore_func,
			_SetScore_gearScore_func,
			_SetX_position_func,
			_SetY_position_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
