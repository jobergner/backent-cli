package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWrite(t *testing.T) {
	t.Run("writes web client", func(t *testing.T) {
		f := NewFactory(testutils.NewSimpleASTExample())
		f.Write()

		s := f.file.String()
		for name, decl := range decl_to_string_decl_collection {
			if !strings.Contains(s, decl) {
				t.Errorf("missing decl %s", name)
			}
		}

		for _, decl := range decl_to_string_decl_collection {
			s = strings.Replace(s, decl, "", 1)
		}

		remaining := strings.TrimSpace(s)
		if len(remaining) != 0 {
			t.Errorf("found redundant decls:\n%s", remaining)
		}

	})
}
