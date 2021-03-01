package statefactory

import (
	"strings"
	"testing"
)

func TestWriteStateMachine(t *testing.T) {
	t.Run("writes operationKind", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeOperationKind()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			OperationKind_type,
			OperationKindDelete_type,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes stateMachine", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeStateMachine()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			StateMachine_type,
			GenerateID_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
	t.Run("writes stateMachine", func(t *testing.T) {
		sf := newStateFactory(newSimpleASTExample())
		sf.writeUpdateState()

		actual := normalizeWhitespace(sf.buf.String())
		expected := normalizeWhitespace(strings.Join([]string{
			UpdateState_StateMachine_func,
		}, "\n"))

		if expected != actual {
			t.Errorf(diff(actual, expected))
		}
	})
}
