package state

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/examples/configs"
	"github.com/jobergner/backent-cli/pkg/ast"
	// "github.com/jobergner/backent-cli/pkg/factory/testutils"
	// "io/ioutil"
	// "testing"
)

func TestEngineFactory(t *testing.T) {
	t.Run("builds successfully", func(t *testing.T) {

		f := jen.NewFile("state")
		Write(f, configs.StateConfig)

		buf := new(bytes.Buffer)
		f.Render(buf)

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
