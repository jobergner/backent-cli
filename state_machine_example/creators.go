package statemachine

func (sm *stateMachine) CreateGearScore(parentage ...parentInfo) gearScore {
	var gearScore gearScore
	gearScore.id = gearScoreID(sm.generateID())
	gearScore.parentage = append(gearScore.parentage, parentage...)
	gearScore.operationKind = operationKindUpdate
	sm.patch.gearScore[gearScore.id] = gearScore
	return gearScore
}

func (sm *stateMachine) CreatePosition(parentage ...parentInfo) position {
	var position position
	position.id = positionID(sm.generateID())
	position.parentage = append(position.parentage, parentage...)
	position.operationKind = operationKindUpdate
	sm.patch.position[position.id] = position
	return position
}

func (sm *stateMachine) CreateItem(parentage ...parentInfo) item {
	var item item
	item.id = itemID(sm.generateID())
	item.parentage = append(item.parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(item.parentage, parentInfo{entityKindItem, int(item.id)})...)
	item.gearScore = elementGearScore.id
	item.operationKind = operationKindUpdate
	sm.patch.item[item.id] = item
	return item
}

func (sm *stateMachine) CreateZoneItem(parentage ...parentInfo) zoneItem {
	var zoneItem zoneItem
	zoneItem.id = zoneItemID(sm.generateID())
	zoneItem.parentage = append(zoneItem.parentage, parentage...)
	elementItem := sm.CreateItem(append(zoneItem.parentage, parentInfo{entityKindZoneItem, int(zoneItem.id)})...)
	zoneItem.item = elementItem.id
	elementPosition := sm.CreatePosition(append(zoneItem.parentage, parentInfo{entityKindZoneItem, int(zoneItem.id)})...)
	zoneItem.position = elementPosition.id
	zoneItem.operationKind = operationKindUpdate
	sm.patch.zoneItem[zoneItem.id] = zoneItem
	return zoneItem
}

func (sm *stateMachine) CreatePlayer(parentage ...parentInfo) player {
	var player player
	player.id = playerID(sm.generateID())
	player.parentage = append(player.parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(player.parentage, parentInfo{entityKindPlayer, int(player.id)})...)
	player.gearScore = elementGearScore.id
	elementPosition := sm.CreatePosition(append(player.parentage, parentInfo{entityKindPlayer, int(player.id)})...)
	player.position = elementPosition.id
	player.operationKind = operationKindUpdate
	sm.patch.player[player.id] = player
	return player
}

func (sm *stateMachine) CreateZone() zone {
	var zone zone
	zone.id = zoneID(sm.generateID())
	zone.operationKind = operationKindUpdate
	sm.patch.zone[zone.id] = zone
	return zone
}
