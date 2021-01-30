package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

type book struct {
	id      id
	name    string
	content bytes.Buffer
}

func (b *book) addContent(content []byte) {
	b.content.Write(content)
}

func newBook(name string) book {
	return book{
		id:      id(rand.Intn(100)),
		name:    name,
		content: bytes.Buffer{},
	}
}

func (b book) print() {
	fmt.Println(b.content.String())
}
