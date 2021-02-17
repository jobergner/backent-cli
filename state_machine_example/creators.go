package statemachine

func (sm *stateMachine) CreateGearScore(parentage ...parentInfo) gearScore {
	var element gearScore
	element.id = gearScoreID(sm.generateID())
	element.parentage = append(element.parentage, parentage...)
	element.operationKind = operationKindUpdate
	sm.patch.gearScore[element.id] = element
	return element
}

func (sm *stateMachine) CreatePosition(parentage ...parentInfo) position {
	var element position
	element.id = positionID(sm.generateID())
	element.parentage = append(element.parentage, parentage...)
	element.operationKind = operationKindUpdate
	sm.patch.position[element.id] = element
	return element
}

func (sm *stateMachine) CreateItem(parentage ...parentInfo) item {
	var element item
	element.id = itemID(sm.generateID())
	element.parentage = append(element.parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(element.parentage, parentInfo{entityKindItem, int(element.id)})...)
	element.gearScore = elementGearScore.id
	element.operationKind = operationKindUpdate
	sm.patch.item[element.id] = element
	return element
}

func (sm *stateMachine) CreateZoneItem(parentage ...parentInfo) zoneItem {
	var element zoneItem
	element.id = zoneItemID(sm.generateID())
	element.parentage = append(element.parentage, parentage...)
	elementItem := sm.CreateItem(append(element.parentage, parentInfo{entityKindZoneItem, int(element.id)})...)
	element.item = elementItem.id
	elementPosition := sm.CreatePosition(append(element.parentage, parentInfo{entityKindZoneItem, int(element.id)})...)
	element.position = elementPosition.id
	element.operationKind = operationKindUpdate
	sm.patch.zoneItem[element.id] = element
	return element
}

func (sm *stateMachine) CreatePlayer(parentage ...parentInfo) player {
	var element player
	element.id = playerID(sm.generateID())
	element.parentage = append(element.parentage, parentage...)
	elementGearScore := sm.CreateGearScore(append(element.parentage, parentInfo{entityKindPlayer, int(element.id)})...)
	element.gearScore = elementGearScore.id
	elementPosition := sm.CreatePosition(append(element.parentage, parentInfo{entityKindPlayer, int(element.id)})...)
	element.position = elementPosition.id
	element.operationKind = operationKindUpdate
	sm.patch.player[element.id] = element
	return element
}

func (sm *stateMachine) CreateZone() zone {
	var element zone
	element.id = zoneID(sm.generateID())
	element.operationKind = operationKindUpdate
	sm.patch.zone[element.id] = element
	return element
}
