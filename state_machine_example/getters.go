package statemachine

func (sm *stateMachine) GetPlayer(playerID playerID) player {
	patchingElement, ok := sm.patch.player[playerID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.player[playerID]
	return currentElement
}

func (e player) GetItems(sm *stateMachine) []item {
	var items []item
	for _, itemID := range e.items {
		items = append(items, sm.GetItem(itemID))
	}
	return items
}

func (e player) GetGearScore(sm *stateMachine) gearScore {
	patchingElement, ok := sm.patch.gearScore[e.gearScore]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.gearScore[e.gearScore]
	return currentElement
}

func (e player) GetPosition(sm *stateMachine) position {
	patchingElement, ok := sm.patch.position[e.position]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.position[e.position]
	return currentElement
}

func (sm *stateMachine) GetGearScore(gearScoreID gearScoreID) gearScore {
	patchingElement, ok := sm.patch.gearScore[gearScoreID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.gearScore[gearScoreID]
	return currentElement
}

func (e gearScore) GetLevel() int {
	return e.level
}

func (e gearScore) GetScore() int {
	return e.score
}

func (sm *stateMachine) GetItem(itemID itemID) item {
	patchingElement, ok := sm.patch.item[itemID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.item[itemID]
	return currentElement
}

func (e item) GetGearScore(sm *stateMachine) gearScore {
	patchingElement, ok := sm.patch.gearScore[e.gearScore]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.gearScore[e.gearScore]
	return currentElement
}

func (sm *stateMachine) GetPosition(positionID positionID) position {
	patchingElement, ok := sm.patch.position[positionID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.position[positionID]
	return currentElement
}

func (e position) GetX() float64 {
	return e.x
}

func (e position) GetY() float64 {
	return e.y
}

func (sm *stateMachine) GetZoneItem(zoneItemID zoneItemID) zoneItem {
	patchingElement, ok := sm.patch.zoneItem[zoneItemID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.zoneItem[zoneItemID]
	return currentElement
}

func (e zoneItem) GetPosition(sm *stateMachine) position {
	patchingElement, ok := sm.patch.position[e.position]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.position[e.position]
	return currentElement
}

func (e zoneItem) GetItem(sm *stateMachine) item {
	patchingElement, ok := sm.patch.item[e.item]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.item[e.item]
	return currentElement
}

func (sm *stateMachine) GetZone(zoneID zoneID) zone {
	patchingElement, ok := sm.patch.zone[zoneID]
	if ok {
		return patchingElement
	}
	currentElement := sm.state.zone[zoneID]
	return currentElement
}

func (e zone) GetPlayers(sm *stateMachine) []player {
	var elements []player
	for _, elementID := range e.players {
		elements = append(elements, sm.GetPlayer(elementID))
	}
	return elements
}

func (e zone) GetZoneItems(sm *stateMachine) []zoneItem {
	var elements []zoneItem
	for _, elementID := range e.items {
		elements = append(elements, sm.GetZoneItem(elementID))
	}
	return elements
}
