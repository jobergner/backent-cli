package main

import "fmt"

type id int

const (
	defaultName string = "Good Read"
)

const greeting string = `Dearest Person,
I'm sending you this letter to let you know`

type printable interface {
	print()
}

func main() {
	myBook := newBook(defaultName)
	myBook.addContent([]byte(`This is a book about
one young fella who tried to change the world`))
	myBook.print()
	fmt.Println(myBook.name)

	myLetter := newLetter([]byte(greeting + ` that at some point
I might be no more.`))
	myLetter.print()
}
