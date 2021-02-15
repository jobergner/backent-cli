package statemachine

func (sm *stateMachine) GetPerson(personID personID) person {
	patchingPerson, ok := sm.patch.person[personID]
	if ok {
		return patchingPerson
	}
	currentPerson := sm.state.person[personID]
	return currentPerson
}

func (sm *stateMachine) GetChild(childID childID) child {
	patchingChild, ok := sm.patch.child[childID]
	if ok {
		return patchingChild
	}
	currentChild := sm.state.child[childID]
	return currentChild
}

func (sm *stateMachine) GetName(nameID nameID) name {
	patchingName, ok := sm.patch.name[nameID]
	if ok {
		return patchingName
	}
	currentName := sm.state.name[nameID]
	return currentName
}

func (p person) GetName(sm *stateMachine) name {
	patchingName, ok := sm.patch.name[p.name]
	if ok {
		return patchingName
	}
	currentName := sm.state.name[p.name]
	return currentName
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
	return currentName
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
