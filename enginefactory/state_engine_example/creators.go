package state

func (se *Engine) CreateGearScore() GearScore {
	return se.createGearScore(false)
}

func (se *Engine) createGearScore(hasParent bool) GearScore {
	var e gearScoreCore
	e.ID = GearScoreID(se.GenerateID())
	e.HasParent = hasParent
	e.OperationKind = OperationKindUpdate
	se.Patch.GearScore[e.ID] = e
	return GearScore{gearScore: e}
}

func (se *Engine) CreatePosition() Position {
	return se.createPosition(false)
}

func (se *Engine) createPosition(hasParent bool) Position {
	var e positionCore
	e.ID = PositionID(se.GenerateID())
	e.HasParent = hasParent
	e.OperationKind = OperationKindUpdate
	se.Patch.Position[e.ID] = e
	return Position{position: e}
}

func (se *Engine) CreateItem() Item {
	return se.createItem(false)
}

func (se *Engine) createItem(hasParent bool) Item {
	var e itemCore
	e.ID = ItemID(se.GenerateID())
	e.HasParent = hasParent
	elementGearScore := se.createGearScore(true)
	e.GearScore = elementGearScore.gearScore.ID
	e.OperationKind = OperationKindUpdate
	se.Patch.Item[e.ID] = e
	return Item{item: e}
}

func (se *Engine) CreateZoneItem() ZoneItem {
	return se.createZoneItem(false)
}

func (se *Engine) createZoneItem(hasParent bool) ZoneItem {
	var e zoneItemCore
	e.ID = ZoneItemID(se.GenerateID())
	e.HasParent = hasParent
	elementItem := se.createItem(true)
	e.Item = elementItem.item.ID
	elementPosition := se.createPosition(true)
	e.Position = elementPosition.position.ID
	e.OperationKind = OperationKindUpdate
	se.Patch.ZoneItem[e.ID] = e
	return ZoneItem{zoneItem: e}
}

func (se *Engine) CreatePlayer() Player {
	return se.createPlayer(false)
}

func (se *Engine) createPlayer(hasParent bool) Player {
	var e playerCore
	e.ID = PlayerID(se.GenerateID())
	e.HasParent = hasParent
	elementGearScore := se.createGearScore(true)
	e.GearScore = elementGearScore.gearScore.ID
	elementPosition := se.createPosition(true)
	e.Position = elementPosition.position.ID
	e.OperationKind = OperationKindUpdate
	se.Patch.Player[e.ID] = e
	return Player{player: e}
}

func (se *Engine) CreateZone() Zone {
	return se.createZone()
}

func (se *Engine) createZone() Zone {
	var e zoneCore
	e.ID = ZoneID(se.GenerateID())
	e.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.ID] = e
	return Zone{zone: e}
}
