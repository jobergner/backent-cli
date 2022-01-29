package serverfactory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/examples/configs"
	"github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/testutils"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, configs.ActionsConfig, configs.ResponsesConfig)
	return simpleAST
}

func TestWriteParameters(t *testing.T) {
	t.Run("writes parameters", func(t *testing.T) {
		sf := newServerFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeParameters()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			_AddItemToPlayerParams_type,
			_MovePlayerParams_type,
			_SpawnZoneItemsParams_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
