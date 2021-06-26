package state

type messageKind string

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
	client  *Client
}

func printMessage(msg message) string {
	b, err := msg.MarshalJSON()
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}
