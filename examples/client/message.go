package client

import "github.com/jobergner/backent-cli/examples/message"

type Message struct {
	ID      int          `json:"id"`
	Kind    message.Kind `json:"kind"`
	Content []byte       `json:"content"`
}
