package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteSetters(t *testing.T) {
	t.Run("writes setters", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeSetters()

		actual := sf.buf.String()
		expected := utils.FormatCode(strings.Join([]string{
			_SetLevel_GearScore_func,
			_SetScore_GearScore_func,
			_SetX_Position_func,
			_SetY_Position_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
