package statemachine

func (p person) AddChild(childID childID, sm *stateMachine) person {
	p.children = append(p.children, childID)
	p.operationKind = operationKindUpdate
	sm.patch.person[p.id] = p
	return p
}
