package statemachine

func (sm *stateMachine) CreatePerson() person {
	var person person
	personName := sm.CreateName()
	person.name = personName.id
	person.id = personID(sm.generateID())
	person.operationKind = operationKindCreate
	sm.patch.person[person.id] = person
	return person
}

func (sm *stateMachine) CreateName() name {
	var name name
	name.id = nameID(sm.generateID())
	name.operationKind = operationKindCreate
	sm.patch.name[name.id] = name
	return name
}

func (sm *stateMachine) CreateChild() child {
	var child child
	childName := sm.CreateName()
	child.name = childName.id
	child.id = childID(sm.generateID())
	child.operationKind = operationKindCreate
	sm.patch.child[child.id] = child
	return child
}
