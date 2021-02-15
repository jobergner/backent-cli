package statemachine

type state struct {
	person map[personID]person
	name   map[nameID]name
	child  map[childID]child
}

func newState() state {
	return state{
		person: make(map[personID]person),
		child:  make(map[childID]child),
		name:   make(map[nameID]name),
	}
}

type person struct {
	id            personID
	name          nameID
	children      []childID
	age           int
	operationKind operationKind
	parentage     parentage
}

type name struct {
	id            nameID
	first         string
	last          string
	operationKind operationKind
	parentage     parentage
}

type child struct {
	id            childID
	name          nameID
	operationKind operationKind
	parentage     parentage
}
