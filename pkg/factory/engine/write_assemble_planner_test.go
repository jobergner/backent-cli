package engine

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteAssemblePlanner(t *testing.T) {
	t.Run("writes assemble planner", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssemblePlanner()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			assemblePlanner_type,
			newAssemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner clear", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssemblePlannerClear()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			clear_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner plan", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssemblePlannerPlan()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			plan_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
	t.Run("writes assemble planner fill", func(t *testing.T) {
		sf := newEngineFactory(jen.NewFile(testutils.PackageName), newSimpleASTExample())
		sf.writeAssemblePlannerFill()

		buf := new(bytes.Buffer)
		sf.file.Render(buf)

		actual := testutils.FormatCode(buf.String())
		expected := testutils.FormatUnpackagedCode(strings.Join([]string{
			fill_assemblePlanner_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(testutils.Diff(actual, expected))
		}
	})
}
