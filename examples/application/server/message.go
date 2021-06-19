package state

type messageKind int

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
	client  *Client
}
