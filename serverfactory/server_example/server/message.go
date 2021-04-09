package state

type messageKind int

const (
	messageKindInit messageKind = iota + 1
)

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
}
