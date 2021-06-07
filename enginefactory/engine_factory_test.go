package enginefactory

import (
	"bar-cli/ast"
	"bar-cli/examples"
	// "bar-cli/testutils"
	// "io/ioutil"
	// "testing"
)

// DISABLED DUE TO NOT BEING VERY USEFUL ATM!!
// func TestStateFactory(t *testing.T) {
// 	t.Run("doesnt crash", func(t *testing.T) {
// 		expected, err := ioutil.ReadFile("testdata/golden.go")
// 		if err != nil {
// 			t.Fatalf("Error loading golden file: %s", err)
// 		}
// 		actual := WriteEngineFrom(examples.StateConfig)

// 		if string(expected) != string(actual) {
// 			t.Errorf(testutils.Diff(string(actual), string(expected)))
// 		}
// 	})
// }

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(examples.StateConfig, map[interface{}]interface{}{})
	return simpleAST
}
