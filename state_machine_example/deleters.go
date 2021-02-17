package statemachine

func (sm *stateMachine) DeletePlayer(playerID playerID) {
	player := sm.GetPlayer(playerID)
	player.operationKind = operationKindDelete
	sm.patch.player[player.id] = player
	for _, itemID := range player.items {
		sm.DeleteItem(itemID)
	}
	sm.DeleteGearScore(player.gearScore)
	sm.DeletePosition(player.position)
}

func (sm *stateMachine) DeleteGearScore(gearScoreID gearScoreID) {
	gearScore := sm.GetGearScore(gearScoreID)
	gearScore.operationKind = operationKindDelete
	sm.patch.gearScore[gearScore.id] = gearScore
}

func (sm *stateMachine) DeletePosition(positionID positionID) {
	position := sm.GetPosition(positionID)
	position.operationKind = operationKindDelete
	sm.patch.position[position.id] = position
}

func (sm *stateMachine) DeleteItem(itemID itemID) {
	item := sm.GetItem(itemID)
	item.operationKind = operationKindDelete
	sm.patch.item[item.id] = item
}

func (sm *stateMachine) DeleteZoneItem(zoneItemID zoneItemID) {
	zoneItem := sm.GetZoneItem(zoneItemID)
	zoneItem.operationKind = operationKindDelete
	sm.patch.zoneItem[zoneItem.id] = zoneItem
}

func (sm *stateMachine) DeleteZone(zoneID zoneID) {
	zone := sm.GetZone(zoneID)
	zone.operationKind = operationKindDelete
	sm.patch.zone[zone.id] = zone
	for _, playerID := range zone.players {
		sm.DeletePlayer(playerID)
	}
	for _, zoneItemID := range zone.items {
		sm.DeleteZoneItem(zoneItemID)
	}
}
