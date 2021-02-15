package statemachine

func (sm *stateMachine) DeletePerson(personID personID) {
	person := sm.GetPerson(personID)
	person.operationKind = operationKindDelete
	sm.patch.person[person.id] = person
}

func (sm *stateMachine) DeleteChild(childID childID) {
	child := sm.GetChild(childID)
	child.operationKind = operationKindDelete
	sm.patch.child[child.id] = child
}

func (sm *stateMachine) DeleteName(nameID nameID) {
	name := sm.GetName(nameID)
	name.operationKind = operationKindDelete
	sm.patch.name[name.id] = name
}
