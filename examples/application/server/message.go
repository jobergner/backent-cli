package state

type messageKind int

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
	source  *Client
}

type response struct {
	Content  []byte `json:"content"`
	receiver *Client
}
