package statemachine

// deep deleting
func (sm *stateMachine) DeletePlayer(playerID playerID) {
	element := sm.GetPlayer(playerID)
	element.operationKind = operationKindDelete
	sm.patch.player[element.id] = element
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
