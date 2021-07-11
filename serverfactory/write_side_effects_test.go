package serverfactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/testutils"
)

func TestWriteSideEffects(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeSideEffects()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			`type SideEffects struct {
	OnDeploy    func(*Engine)
	OnFrameTick func(*Engine)
}`,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
