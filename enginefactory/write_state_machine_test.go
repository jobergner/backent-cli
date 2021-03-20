package enginefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteEngine(t *testing.T) {
	t.Run("writes operationKind", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeOperationKind()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			OperationKind_type,
			OperationKindDelete_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes Engine", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeEngine()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			Engine_type,
			newEngine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes generateID method", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeGenerateID()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			GenerateID_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
	t.Run("writes updateState method", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeUpdateState()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			UpdateState_Engine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
