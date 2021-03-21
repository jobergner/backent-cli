package actionsfactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteStart(t *testing.T) {
	t.Run("writes server start", func(t *testing.T) {

		ast := buildActionsConfigAST(testActionsConfig)
		af := newActionsFactory(ast)

		actual := utils.NormalizeWhitespace(string(af.writeStart().writtenSourceCode()))
		expected := utils.NormalizeWhitespace(strings.TrimSpace(`
func main() {
	err := state.Start(
		interactBaz,
		makeFoo,
		walkBar,
	)

	if err != nil {
		panic(err)
	}
}
		`))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
