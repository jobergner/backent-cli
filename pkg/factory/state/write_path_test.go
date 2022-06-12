package state

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWritePath(t *testing.T) {
	t.Run("writes path identifiers", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeIdentifiers()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			treeFieldIdentifier_type,
			attackEventIdentifier_type,
			toString_treeFieldIdentifier_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
	t.Run("writes path", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writePath()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			segment_type,
			path_type,
			newPath_func,
			extendAndCopy_path_func,
			toJSONPath_path_func,
			isSliceFieldIdentifier_func,
		}, "\n"))

		diff, hasDiff := testutils.Diff(actual, expected)
		if hasDiff {
			t.Errorf(diff)
		}
	})
}
