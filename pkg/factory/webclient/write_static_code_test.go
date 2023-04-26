package webclient

import (
	"strings"
	"testing"

	"github.com/jobergner/backent-cli/pkg/factory/testutils"
)

func TestWriteStaticCode(t *testing.T) {
	t.Run("writes static code", func(t *testing.T) {
		sf := NewFactory(testutils.NewSimpleASTExample())
		sf.writeStaticCode()

		actual := sf.file.String()
		expected := strings.Join([]string{
			type_EventListener,
			class_EventEmitter,
			const_ErrResponseTimeout,
			const_responseTimeout,
			const_elementRegistrar,
			const_eventEmitter,
			enum_ReferencedDataStatus,
			enum_OperationKind,
			const_currentState,
			interface_WebSocketMessage,
			function_generateID + "\n",
		}, "\n")

		if actual != expected {
			diffs := testutils.PrettyDiffText(actual, expected)
			t.Errorf(diffs)
		}
	})
}
