package main

type messageKind int

const (
	messageKindAction messageKind = iota
)

type message struct {
	kind    messageKind
	content []byte
}

type actionMessageContent struct {
	actionName   string
	paramContent []byte
}

// generate for each message
