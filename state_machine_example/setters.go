package statemachine

func (p person) SetAge(val int, sm *stateMachine) person {
	p.age = val
	p.operationKind = operationKindUpdate
	sm.patch.person[p.id] = p
	return p
}

func (n name) SetFirst(val string, sm *stateMachine) name {
	n.first = val
	n.operationKind = operationKindUpdate
	sm.patch.name[n.id] = n
	return n
}

func (n name) SetLast(val string, sm *stateMachine) name {
	n.last = val
	n.operationKind = operationKindUpdate
	sm.patch.name[n.id] = n
	return n
}
