package statemachine

func (sm *StateMachine) CreateGearScore() GearScore {
	return sm.createGearScore(false)
}

func (sm *StateMachine) createGearScore(hasParent bool) GearScore {
	var e gearScoreCore
	e.ID = GearScoreID(sm.GenerateID())
	e.HasParent = hasParent
	e.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[e.ID] = e
	return GearScore{gearScore: e}
}

func (sm *StateMachine) CreatePosition() Position {
	return sm.createPosition(false)
}

func (sm *StateMachine) createPosition(hasParent bool) Position {
	var e positionCore
	e.ID = PositionID(sm.GenerateID())
	e.HasParent = hasParent
	e.OperationKind = OperationKindUpdate
	sm.Patch.Position[e.ID] = e
	return Position{position: e}
}

func (sm *StateMachine) CreateItem() Item {
	return sm.createItem(false)
}

func (sm *StateMachine) createItem(hasParent bool) Item {
	var e itemCore
	e.ID = ItemID(sm.GenerateID())
	e.HasParent = hasParent
	elementGearScore := sm.createGearScore(true)
	e.GearScore = elementGearScore.gearScore.ID
	e.OperationKind = OperationKindUpdate
	sm.Patch.Item[e.ID] = e
	return Item{item: e}
}

func (sm *StateMachine) CreateZoneItem() ZoneItem {
	return sm.createZoneItem(false)
}

func (sm *StateMachine) createZoneItem(hasParent bool) ZoneItem {
	var e zoneItemCore
	e.ID = ZoneItemID(sm.GenerateID())
	e.HasParent = hasParent
	elementItem := sm.createItem(true)
	e.Item = elementItem.item.ID
	elementPosition := sm.createPosition(true)
	e.Position = elementPosition.position.ID
	e.OperationKind = OperationKindUpdate
	sm.Patch.ZoneItem[e.ID] = e
	return ZoneItem{zoneItem: e}
}

func (sm *StateMachine) CreatePlayer() Player {
	return sm.createPlayer(false)
}

func (sm *StateMachine) createPlayer(hasParent bool) Player {
	var e playerCore
	e.ID = PlayerID(sm.GenerateID())
	e.HasParent = hasParent
	elementGearScore := sm.createGearScore(true)
	e.GearScore = elementGearScore.gearScore.ID
	elementPosition := sm.createPosition(true)
	e.Position = elementPosition.position.ID
	e.OperationKind = OperationKindUpdate
	sm.Patch.Player[e.ID] = e
	return Player{player: e}
}

func (sm *StateMachine) CreateZone() Zone {
	return sm.createZone()
}

func (sm *StateMachine) createZone() Zone {
	var e zoneCore
	e.ID = ZoneID(sm.GenerateID())
	e.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.ID] = e
	return Zone{zone: e}
}
