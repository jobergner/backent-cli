package actionsfactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

var testActionsConfig = map[interface{}]interface{}{
	"makeFoo": map[interface{}]interface{}{
		"entities": "[]entity",
		"count":    "int",
		"origins":  "[]string",
	},
	"walkBar": map[interface{}]interface{}{
		"distance": "float64",
		"altitude": "altitude",
	},
	"interactBaz": map[interface{}]interface{}{
		"target": "bool",
	},
}

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {

		ast := buildActionsConfigAST(testActionsConfig)
		af := newActionsFactory(ast)

		actual := utils.FormatCode(string(af.writeActions().writtenSourceCode()))
		expected := utils.FormatCode(strings.TrimSpace(`
func interactBaz(target bool, sm *state.Engine) {}
func makeFoo(count int, entities []state.Entity, origins []string, sm *state.Engine) {}
func walkBar(altitude state.Altitude, distance float64, sm *state.Engine) {}
		`))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
