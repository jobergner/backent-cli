package statemachine

func (p person) RemoveChild(childID childID, sm *stateMachine) person {
	var indexToRemove int
	for i, _childID := range p.children {
		if _childID == childID {
			indexToRemove = i
			break
		}
	}
	p.children = append(p.children[:indexToRemove], p.children[indexToRemove+1:]...)
	p.operationKind = operationKindUpdate
	sm.patch.person[p.id] = p
	return p
}
