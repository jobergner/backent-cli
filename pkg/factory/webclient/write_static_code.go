package webclient

import (
	"strings"
)

func (s *Factory) writeStaticCode() *Factory {
	static := strings.Join([]string{
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

	s.file.WriteString(static)

	return s
}
