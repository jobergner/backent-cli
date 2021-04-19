package state

func (se *Engine) CreateGearScore() gearScore {
	return se.createGearScore(false)
}

func (se *Engine) createGearScore(hasParent bool) gearScore {
	var element gearScoreCore
	element.ID = GearScoreID(se.GenerateID())
	element.HasParent_ = hasParent
	element.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[element.ID] = element
	return gearScore{gearScore: element}
}

func (se *Engine) CreatePosition() position {
	return se.createPosition(false)
}

func (se *Engine) createPosition(hasParent bool) position {
	var element positionCore
	element.ID = PositionID(se.GenerateID())
	element.HasParent_ = hasParent
	element.OperationKind_ = OperationKindUpdate
	se.Patch.Position[element.ID] = element
	return position{position: element}
}

func (se *Engine) CreateItem() item {
	return se.createItem(false)
}

func (se *Engine) createItem(hasParent bool) item {
	var element itemCore
	element.ID = ItemID(se.GenerateID())
	element.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	element.BoundTo = itemBoundToRef{parentID: element.ID}
	element.OperationKind_ = OperationKindUpdate
	se.Patch.Item[element.ID] = element
	return item{item: element}
}

func (se *Engine) CreateZoneItem() zoneItem {
	return se.createZoneItem(false)
}

func (se *Engine) createZoneItem(hasParent bool) zoneItem {
	var element zoneItemCore
	element.ID = ZoneItemID(se.GenerateID())
	element.HasParent_ = hasParent
	elementItem := se.createItem(true)
	element.Item = elementItem.item.ID
	elementPosition := se.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind_ = OperationKindUpdate
	se.Patch.ZoneItem[element.ID] = element
	return zoneItem{zoneItem: element}
}

func (se *Engine) CreatePlayer() player {
	return se.createPlayer(false)
}

func (se *Engine) createPlayer(hasParent bool) player {
	var element playerCore
	element.ID = PlayerID(se.GenerateID())
	element.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	element.GearScore = elementGearScore.gearScore.ID
	elementPosition := se.createPosition(true)
	element.Position = elementPosition.position.ID
	element.OperationKind_ = OperationKindUpdate
	se.Patch.Player[element.ID] = element
	return player{player: element}
}

func (se *Engine) CreateZone() zone {
	return se.createZone()
}

func (se *Engine) createZone() zone {
	var element zoneCore
	element.ID = ZoneID(se.GenerateID())
	element.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[element.ID] = element
	return zone{zone: element}
}
