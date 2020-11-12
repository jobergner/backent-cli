package inputexample

type person struct {
	name     name
	children []child
	age      int
}

type name struct {
	first string
	last  string
}

type child struct {
	name name
}
