package actionsfactory

import (
	"strings"
	"testing"
)

func TestWriteActions(t *testing.T) {
	t.Run("writes actions", func(t *testing.T) {
		data := map[interface{}]interface{}{
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

		ast := buildActionsConfigAST(data)
		af := newActionsFactory(ast)

		actual := normalizeWhitespace(string(af.writeActions().writtenSourceCode()))
		expected := normalizeWhitespace(strings.TrimSpace(`
func interactBaz(target bool) {}
func makeFoo(count int, entities []statemachine.Entity, origins []string) {}
func walkBar(altitude statemachine.Altitude, distance float64) {}
		`))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
