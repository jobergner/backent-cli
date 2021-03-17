package statefactory

import (
	"bar-cli/utils"
	"strings"
	"testing"
)

func TestWriteStateMachine(t *testing.T) {
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
	t.Run("writes stateMachine", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeStateMachine()

		actual := utils.NormalizeWhitespace(sf.buf.String())
		expected := utils.NormalizeWhitespace(strings.Join([]string{
			StateMachine_type,
			newStateMachine_func,
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
			GenerateID_StateMachine_func,
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
			UpdateState_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(utils.Diff(actual, expected))
		}
	})
}
