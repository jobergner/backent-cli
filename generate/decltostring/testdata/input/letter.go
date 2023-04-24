package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

type letter struct {
	id      id
	content bytes.Buffer
}

func newLetter(content []byte) letter {
	return letter{
		id:      id(rand.Intn(100)),
		content: *bytes.NewBuffer(content),
	}
}

func (l letter) print() {
	fmt.Println(l.content.String())
}
