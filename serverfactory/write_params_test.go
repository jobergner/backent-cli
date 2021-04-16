package serverfactory

import (
	"bar-cli/ast"
	"bar-cli/examples"
	"bar-cli/testutils"
	"strings"
	"testing"
)

func newSimpleASTExample() *ast.AST {
	simpleAST := ast.Parse(examples.StateConfig, examples.ActionsConfig)
	return simpleAST
}

func TestWriteParameters(t *testing.T) {
	t.Run("writes parameters", func(t *testing.T) {
		sf := newServerFactory(newSimpleASTExample())
		sf.writeParameters()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			__MovePlayerParams_type,
			__addItemToPlayerParams_type,
			__spawnZoneItemsParams_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
