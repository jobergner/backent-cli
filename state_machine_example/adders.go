package statemachine

func (e zone) AddPlayer(sm *stateMachine) player {
	player := sm.CreatePlayer(parentInfo{entityKindZone, int(e.id)})
	e.players = append(e.players, player.id)
	e.operationKind = operationKindUpdate
	sm.patch.zone[e.id] = e
	return player
}

func (e zone) AddZoneItem(sm *stateMachine) zoneItem {
	zoneItem := sm.CreateZoneItem(parentInfo{entityKindZone, int(e.id)})
	e.items = append(e.items, zoneItem.id)
	e.operationKind = operationKindUpdate
	sm.patch.zone[e.id] = e
	return zoneItem
}

func (e player) AddItem(sm *stateMachine) item {
	item := sm.CreateItem(append(e.parentage, parentInfo{entityKindPlayer, int(e.id)})...)
	e.items = append(e.items, item.id)
	e.operationKind = operationKindUpdate
	sm.patch.player[e.id] = e
	return item
}
