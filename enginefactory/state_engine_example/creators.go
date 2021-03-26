package state

func (se *Engine) CreateGearScore() GearScore {
	return se.createGearScore(false)
}

func (se *Engine) createGearScore(hasParent bool) GearScore {
	var gearScore gearScoreCore
	gearScore.ID = GearScoreID(se.GenerateID())
	gearScore.HasParent_ = hasParent
	gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.ID] = gearScore
	return GearScore{gearScore: gearScore}
}

func (se *Engine) CreatePosition() Position {
	return se.createPosition(false)
}

func (se *Engine) createPosition(hasParent bool) Position {
	var position positionCore
	position.ID = PositionID(se.GenerateID())
	position.HasParent_ = hasParent
	position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.ID] = position
	return Position{position: position}
}

func (se *Engine) CreateItem() Item {
	return se.createItem(false)
}

func (se *Engine) createItem(hasParent bool) Item {
	var item itemCore
	item.ID = ItemID(se.GenerateID())
	item.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	item.GearScore = elementGearScore.gearScore.ID
	item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.ID] = item
	return Item{item: item}
}

func (se *Engine) CreateZoneItem() ZoneItem {
	return se.createZoneItem(false)
}

func (se *Engine) createZoneItem(hasParent bool) ZoneItem {
	var zoneItem zoneItemCore
	zoneItem.ID = ZoneItemID(se.GenerateID())
	zoneItem.HasParent_ = hasParent
	elementItem := se.createItem(true)
	zoneItem.Item = elementItem.item.ID
	elementPosition := se.createPosition(true)
	zoneItem.Position = elementPosition.position.ID
	zoneItem.OperationKind_ = OperationKindUpdate
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
	return ZoneItem{zoneItem: zoneItem}
}

func (se *Engine) CreatePlayer() Player {
	return se.createPlayer(false)
}

func (se *Engine) createPlayer(hasParent bool) Player {
	var player playerCore
	player.ID = PlayerID(se.GenerateID())
	player.HasParent_ = hasParent
	elementGearScore := se.createGearScore(true)
	player.GearScore = elementGearScore.gearScore.ID
	elementPosition := se.createPosition(true)
	player.Position = elementPosition.position.ID
	player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.ID] = player
	return Player{player: player}
}

func (se *Engine) CreateZone() Zone {
	return se.createZone()
}

func (se *Engine) createZone() Zone {
	var zone zoneCore
	zone.ID = ZoneID(se.GenerateID())
	zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.ID] = zone
	return Zone{zone: zone}
}
