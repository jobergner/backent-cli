package message

type Kind string

const (
	// server -> client
	MessageKindError        Kind = "error"
	MessageKindCurrentState Kind = "currentState"
	MessageKindUpdate       Kind = "update"
	MessageKindID           Kind = "id"
	// signifies a message which will not be sent to the client
	MessageKindNoResponse Kind = "noResponse"
	// responses to messages which fail to unmarshal
	// client -> server
	MessageKindGlobal Kind   = "global"
	MessageIDUnknown  string = "unknown"
)
