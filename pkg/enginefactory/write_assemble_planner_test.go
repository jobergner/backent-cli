package enginefactory

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/testutils"
)

func TestWriteAssemblePlanner(t *testing.T) {
	t.Run("writes assemble planner", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssemblePlanner()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			assemblePlanner_type,
			newAssemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner clear", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssemblePlannerClear()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			clear_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner plan", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssemblePlannerPlan()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			plan_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner fill", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeAssemblePlannerFill()

		actual := testutils.FormatCode(sf.buf.String())
		expected := testutils.FormatCode(strings.Join([]string{
			fill_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
