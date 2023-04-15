package server

import "github.com/jobergner/backent-cli/examples/message"

type Message struct {
	ID      int64        `json:"id"`
	Kind    message.Kind `json:"kind"`
	Content []byte       `json:"content"`
	client  *Client
}
