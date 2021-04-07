package enginefactory

import (
	"bar-cli/ast"
	"bar-cli/utils"
	"io/ioutil"
	"testing"
)

var testData = map[interface{}]interface{}{
	"player": map[interface{}]interface{}{
		"items":     "[]item",
		"gearScore": "gearScore",
		"position":  "position",
	},
	"zone": map[interface{}]interface{}{
		"items":   "[]zoneItem",
		"players": "[]player",
		"tags":    "[]string",
	},
	"zoneItem": map[interface{}]interface{}{
		"position": "position",
		"item":     "item",
	},
	"position": map[interface{}]interface{}{
		"x": "float64",
		"y": "float64",
	},
	"item": map[interface{}]interface{}{
		"gearScore": "gearScore",
	},
	"gearScore": map[interface{}]interface{}{
		"level": "int",
		"score": "int",
	},
}

func TestStateFactory(t *testing.T) {
	t.Run("doesnt crash", func(t *testing.T) {
		expected, err := ioutil.ReadFile("testdata/golden.go")
		if err != nil {
			t.Fatalf("Error loading golden file: %s", err)
		}
		actual := WriteEngineFrom(testData)

		if string(expected) != string(actual) {
			t.Errorf(utils.Diff(string(actual), string(expected)))
		}
	})
}

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(testData, map[interface{}]interface{}{})
	return simpleAST
}
