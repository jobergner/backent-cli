package enginefactory

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/jobergner/backent-cli/examples/configs"
	"github.com/jobergner/backent-cli/pkg/ast"
	// "github.com/jobergner/backent-cli/pkg/testutils"
	// "io/ioutil"
	// "testing"
)

func TestEngineFactory(t *testing.T) {
	t.Run("builds successfully", func(t *testing.T) {

		buf := new(bytes.Buffer)
		buf.WriteString("package foo\n")
		WriteEngine(buf, configs.StateConfig)

		err := ioutil.WriteFile("tmp/out.go", buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

	})
}

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, map[interface{}]interface{}{}, map[interface{}]interface{}{})
	return simpleAST
}
