package main

type messageKind int

const (
	messageKindInit messageKind = iota
)

type message struct {
	Kind    messageKind `json:"kind"`
	Content []byte      `json:"content"`
}
