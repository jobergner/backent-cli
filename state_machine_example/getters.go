package statemachine

func (sm *stateMachine) GetPlayer(playerID playerID) player {
	patchingPlayer, ok := sm.patch.player[playerID]
	if ok {
		return patchingPlayer
	}
	currentPlayer := sm.state.player[playerID]
	return currentPlayer
}

func (e player) GetItems(sm *stateMachine) []item {
	var items []item
	for _, itemID := range e.items {
		items = append(items, sm.GetItem(itemID))
	}
	return items
}

func (e player) GetGearScore(sm *stateMachine) gearScore {
	patchingGearScore, ok := sm.patch.gearScore[e.gearScore]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.state.gearScore[e.gearScore]
	return currentGearScore
}

func (e player) GetPosition(sm *stateMachine) position {
	patchingPosition, ok := sm.patch.position[e.position]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.state.position[e.position]
	return currentPosition
}

func (sm *stateMachine) GetGearScore(gearScoreID gearScoreID) gearScore {
	patchingGearScore, ok := sm.patch.gearScore[gearScoreID]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.state.gearScore[gearScoreID]
	return currentGearScore
}

func (e gearScore) GetLevel() int {
	return e.level
}

func (e gearScore) GetScore() int {
	return e.score
}

func (sm *stateMachine) GetItem(itemID itemID) item {
	patchingItem, ok := sm.patch.item[itemID]
	if ok {
		return patchingItem
	}
	currentItem := sm.state.item[itemID]
	return currentItem
}

func (e item) GetGearScore(sm *stateMachine) gearScore {
	patchingGearScore, ok := sm.patch.gearScore[e.gearScore]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.state.gearScore[e.gearScore]
	return currentGearScore
}

func (sm *stateMachine) GetPosition(positionID positionID) position {
	patchingPosition, ok := sm.patch.position[positionID]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.state.position[positionID]
	return currentPosition
}

func (e position) GetX() float64 {
	return e.x
}

func (e position) GetY() float64 {
	return e.y
}

func (sm *stateMachine) GetZoneItem(zoneItemID zoneItemID) zoneItem {
	patchingZoneItem, ok := sm.patch.zoneItem[zoneItemID]
	if ok {
		return patchingZoneItem
	}
	currentZoneItem := sm.state.zoneItem[zoneItemID]
	return currentZoneItem
}

func (e zoneItem) GetPosition(sm *stateMachine) position {
	patchingPosition, ok := sm.patch.position[e.position]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.state.position[e.position]
	return currentPosition
}

func (e zoneItem) GetItem(sm *stateMachine) item {
	patchingItem, ok := sm.patch.item[e.item]
	if ok {
		return patchingItem
	}
	currentItem := sm.state.item[e.item]
	return currentItem
}

func (sm *stateMachine) GetZone(zoneID zoneID) zone {
	patchingZone, ok := sm.patch.zone[zoneID]
	if ok {
		return patchingZone
	}
	currentZone := sm.state.zone[zoneID]
	return currentZone
}

func (e zone) GetPlayers(sm *stateMachine) []player {
	var players []player
	for _, playerID := range e.players {
		players = append(players, sm.GetPlayer(playerID))
	}
	return players
}

func (e zone) GetZoneItems(sm *stateMachine) []zoneItem {
	var items []zoneItem
	for _, zoneItemID := range e.items {
		items = append(items, sm.GetZoneItem(zoneItemID))
	}
	return items
}
