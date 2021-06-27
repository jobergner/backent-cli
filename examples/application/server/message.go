package state

import (
	"fmt"
)

type MessageKind string

const (
	MessageKindError        MessageKind = "error"
	MessageKindCurrentState MessageKind = "currentState"
	MessageKindUpdate       MessageKind = "update"
)

type Message struct {
	Kind    MessageKind `json:"kind"`
	Content []byte      `json:"content"`
	client  *Client
}

func printMessage(msg Message) string {
	b, err := msg.MarshalJSON()
	if err != nil {
		return err.Error()
	} else {
		return string(b)
	}
}

func messageUnmarshallingError(msgContent []byte, err error) []byte {
	return []byte(fmt.Sprintf("error unmarshalling received message content `%s`: %s", msgContent, err))
}

func responseMarshallingError(msgContent []byte, err error) []byte {
	return []byte(fmt.Sprintf("error marshalling response to `%s`: %s", msgContent, err))
}
