package serverfactory

import (
	"strings"
	"testing"

	"github.com/Java-Jonas/bar-cli/ast"
	"github.com/Java-Jonas/bar-cli/examples/configs"
	"github.com/Java-Jonas/bar-cli/testutils"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(configs.StateConfig, configs.ActionsConfig)
	return simpleAST
}

func TestWriteParameters(t *testing.T) {
	t.Run("writes parameters", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeParameters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			_AddItemToPlayerParams_type,
			_MovePlayerParams_type,
			_SpawnZoneItemsParams_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
