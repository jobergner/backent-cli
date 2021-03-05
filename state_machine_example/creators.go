package statemachine

func (sm *StateMachine) CreateGearScore() GearScore {
	return sm.createGearScore(false)
}

func (sm *StateMachine) createGearScore(hasParent bool) GearScore {
	var gearScore gearScoreCore
	gearScore.ID = GearScoreID(sm.GenerateID())
	gearScore.HasParent = hasParent
	gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[gearScore.ID] = gearScore
	return GearScore{gearScore: gearScore}
}

func (sm *StateMachine) CreatePosition() Position {
	return sm.createPosition(false)
}

func (sm *StateMachine) createPosition(hasParent bool) Position {
	var position positionCore
	position.ID = PositionID(sm.GenerateID())
	position.HasParent = hasParent
	position.OperationKind = OperationKindUpdate
	sm.Patch.Position[position.ID] = position
	return Position{position: position}
}

func (sm *StateMachine) CreateItem() Item {
	return sm.createItem(false)
}

func (sm *StateMachine) createItem(hasParent bool) Item {
	var item itemCore
	item.ID = ItemID(sm.GenerateID())
	item.HasParent = hasParent
	elementGearScore := sm.createGearScore(true)
	item.GearScore = elementGearScore.gearScore.ID
	item.OperationKind = OperationKindUpdate
	sm.Patch.Item[item.ID] = item
	return Item{item: item}
}

func (sm *StateMachine) CreateZoneItem() ZoneItem {
	return sm.createZoneItem(false)
}

func (sm *StateMachine) createZoneItem(hasParent bool) ZoneItem {
	var zoneItem zoneItemCore
	zoneItem.ID = ZoneItemID(sm.GenerateID())
	zoneItem.HasParent = hasParent
	elementItem := sm.createItem(true)
	zoneItem.Item = elementItem.item.ID
	elementPosition := sm.createPosition(true)
	zoneItem.Position = elementPosition.position.ID
	zoneItem.OperationKind = OperationKindUpdate
	sm.Patch.ZoneItem[zoneItem.ID] = zoneItem
	return ZoneItem{zoneItem: zoneItem}
}

func (sm *StateMachine) CreatePlayer() Player {
	return sm.createPlayer(false)
}

func (sm *StateMachine) createPlayer(hasParent bool) Player {
	var player playerCore
	player.ID = PlayerID(sm.GenerateID())
	player.HasParent = hasParent
	elementGearScore := sm.createGearScore(true)
	player.GearScore = elementGearScore.gearScore.ID
	elementPosition := sm.createPosition(true)
	player.Position = elementPosition.position.ID
	player.OperationKind = OperationKindUpdate
	sm.Patch.Player[player.ID] = player
	return Player{player: player}
}

func (sm *StateMachine) CreateZone() Zone {
	return sm.createZone()
}

func (sm *StateMachine) createZone() Zone {
	var zone zoneCore
	zone.ID = ZoneID(sm.GenerateID())
	zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[zone.ID] = zone
	return Zone{zone: zone}
}
