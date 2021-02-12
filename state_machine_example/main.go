package statemachineexample

import (
	"github.com/jinzhu/copier"
	"time"
)

type operationKind string

type personID int
type nameID int
type childID int

type entityKind string

const (
	entityKindPerson entityKind = "person"
	entityKindName              = "name"
	entityKindChild             = "child"
)

const (
	operationKindCreate operationKind = "CREATE"
	operationKindDelete               = "DELETE"
	operationKindUpdate               = "UPDATE"
)

type state struct {
	person map[personID]person
	name   map[nameID]name
	child  map[childID]child
}

type stateMachine struct {
	state *state
	patch *state
	idgen int
}

func newState() *state {
	return &state{
		person: make(map[personID]person),
		child:  make(map[childID]child),
		name:   make(map[nameID]name),
	}
}

type parentage []parentInfo
type parentInfo struct {
	kind entityKind
	id   int
}

func (sm *stateMachine) generateID() int {
	newID := sm.idgen
	sm.idgen = sm.idgen + 1
	return newID
}

func (sm *stateMachine) updateState() {
	for _, person := range sm.patch.person {
		if person.operationKind == operationKindDelete {
			delete(sm.state.person, person.id)
		} else {
			sm.state.person[person.id] = person
		}
	}
	for _, name := range sm.patch.name {
		if name.operationKind == operationKindDelete {
			delete(sm.state.name, name.id)
		} else {
			sm.state.name[name.id] = name
		}
	}
	sm.patch = newState()
}

type person struct {
	id            personID
	name          nameID
	children      []childID
	age           int
	lastModified  int64
	operationKind operationKind
	parentage     parentage
}

type name struct {
	id            nameID
	first         string
	last          string
	lastModified  int64
	operationKind operationKind
	parentage     parentage
}

type child struct {
	id            childID
	name          nameID
	lastModified  int64
	operationKind operationKind
	parentage     parentage
}

func (sm *stateMachine) CreatePerson() person {
	var person person
	personName := sm.CreateName()
	person.name = personName.id
	person.id = personID(sm.generateID())
	person.operationKind = operationKindCreate
	person.lastModified = time.Now().UnixNano()
	sm.patch.person[person.id] = person
	return person
}

func (sm *stateMachine) CreateName() name {
	var name name
	name.id = nameID(sm.generateID())
	name.operationKind = operationKindCreate
	name.lastModified = time.Now().UnixNano()
	sm.patch.name[name.id] = name
	return name
}

func (sm *stateMachine) CreateChild() child {
	var child child
	childName := sm.CreateName()
	child.name = childName.id
	child.id = childID(sm.generateID())
	child.operationKind = operationKindCreate
	child.lastModified = time.Now().UnixNano()
	sm.patch.child[child.id] = child
	return child
}

func (sm *stateMachine) GetPerson(personID personID) person {
	patchingPerson, ok := sm.patch.person[personID]
	if ok {
		return patchingPerson
	}
	currentPerson := sm.state.person[personID]
	personCopy := person{}
	copier.Copy(&personCopy, &currentPerson)
	return personCopy
}

func (sm *stateMachine) GetChild(childID childID) child {
	patchingChild, ok := sm.patch.child[childID]
	if ok {
		return patchingChild
	}
	currentChild := sm.state.child[childID]
	childCopy := child{}
	copier.Copy(&childCopy, &currentChild)
	return childCopy
}

func (sm *stateMachine) GetName(nameID nameID) name {
	patchingName, ok := sm.patch.name[nameID]
	if ok {
		return patchingName
	}
	currentName := sm.state.name[nameID]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}

func (p person) GetName(sm *stateMachine) name {
	patchingName, ok := sm.patch.name[p.name]
	if ok {
		return patchingName
	}
	currentName := sm.state.name[p.name]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}

func (p person) GetChildren(sm *stateMachine) []child {
	var children []child
	for _, childID := range p.children {
		children = append(children, sm.GetChild(childID))
	}
	return children
}

func (c child) GetName(sm *stateMachine) name {
	patchingName, ok := sm.patch.name[c.name]
	if ok {
		return patchingName
	}
	currentName := sm.state.name[c.name]
	nameCopy := name{}
	copier.Copy(&nameCopy, &currentName)
	return nameCopy
}

func (p person) GetAge() int {
	return p.age
}

func (n name) GetFirst() string {
	return n.first
}

func (n name) GetLast() string {
	return n.last
}

func (sm *stateMachine) DeletePerson(personID personID) {
	person := sm.GetPerson(personID)
	person.operationKind = operationKindDelete
	person.lastModified = time.Now().UnixNano()
	sm.patch.person[person.id] = person
}

func (sm *stateMachine) DeleteChild(childID childID) {
	child := sm.GetChild(childID)
	child.operationKind = operationKindDelete
	child.lastModified = time.Now().UnixNano()
	sm.patch.child[child.id] = child
}

func (sm *stateMachine) DeleteName(nameID nameID) {
	name := sm.GetName(nameID)
	name.operationKind = operationKindDelete
	name.lastModified = time.Now().UnixNano()
	sm.patch.name[name.id] = name
}

func (p person) RemoveChild(childID childID, sm *stateMachine) {
	var indexToRemove int
	for i, _childID := range p.children {
		if _childID == childID {
			indexToRemove = i
			break
		}
	}
	p.children = append(p.children[:indexToRemove], p.children[indexToRemove+1:]...)
	p.operationKind = operationKindUpdate
	p.lastModified = time.Now().UnixNano()
	sm.patch.person[p.id] = p
}

func (p person) AddChild(childID childID, sm *stateMachine) person {
	p.children = append(p.children, childID)
	p.operationKind = operationKindUpdate
	p.lastModified = time.Now().UnixNano()
	sm.patch.person[p.id] = p
	return p
}

func (p person) SetAge(val int, sm *stateMachine) person {
	p.age = val
	p.operationKind = operationKindUpdate
	p.lastModified = time.Now().UnixNano()
	sm.patch.person[p.id] = p
	return p
}

func (n name) SetFirst(val string, sm *stateMachine) name {
	n.first = val
	n.operationKind = operationKindUpdate
	n.lastModified = time.Now().UnixNano()
	sm.patch.name[n.id] = n
	return n
}

func (n name) SetLast(val string, sm *stateMachine) name {
	n.last = val
	n.operationKind = operationKindUpdate
	n.lastModified = time.Now().UnixNano()
	sm.patch.name[n.id] = n
	return n
}
