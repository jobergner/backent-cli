package state

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
	// "io/ioutil"
	// "testing"
)

func TestFactory(t *testing.T) {
	t.Run("builds successfully", func(t *testing.T) {
		actual := NewFactory(testutils.NewSimpleASTExample()).Write()

		if unmatchedDecls, ok := testutils.FindUnmatchedDecls(actual, decl_to_string_decl_collection); ok {
			t.Errorf("actual output had missing decls: %s", strings.Join(unmatchedDecls, ", "))
		} else if redundantDecls, ok := testutils.FindRedundantDecls(actual, decl_to_string_decl_collection); ok {
			t.Errorf("actual output had redundant delcs: %s", strings.Join(redundantDecls, ", "))
		}
	})
}
