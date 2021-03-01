package statefactory

import (
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestWriteState(t *testing.T) {
	t.Run("writes entityKinds", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeEntityKinds()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			EntityKind_type,
			EntityKindGearScore_type,
		}, "\n"))

		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual, expected, true)

		dmp.DiffPrettyText(diffs)

		if expected != actual {
			t.Errorf(dmp.DiffPrettyText(diffs))
		}
	})
}
