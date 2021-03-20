package state

func (sm *Engine) DeletePlayer(playerID PlayerID) {
	player := sm.Player(playerID).player
	if player.HasParent {
		return
	}
	sm.deletePlayer(playerID)
}
func (sm *Engine) deletePlayer(playerID PlayerID) {
	player := sm.Player(playerID).player
	player.OperationKind = OperationKindDelete
	sm.Patch.Player[player.ID] = player
	sm.deleteGearScore(player.GearScore)
	for _, itemID := range player.Items {
		sm.deleteItem(itemID)
	}
	sm.deletePosition(player.Position)
}

func (sm *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := sm.GearScore(gearScoreID).gearScore
	if gearScore.HasParent {
		return
	}
	sm.deleteGearScore(gearScoreID)
}
func (sm *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := sm.GearScore(gearScoreID).gearScore
	gearScore.OperationKind = OperationKindDelete
	sm.Patch.GearScore[gearScore.ID] = gearScore
}

func (sm *Engine) DeletePosition(positionID PositionID) {
	position := sm.Position(positionID).position
	if position.HasParent {
		return
	}
	sm.deletePosition(positionID)
}
func (sm *Engine) deletePosition(positionID PositionID) {
	position := sm.Position(positionID).position
	position.OperationKind = OperationKindDelete
	sm.Patch.Position[position.ID] = position
}

func (sm *Engine) DeleteItem(itemID ItemID) {
	item := sm.Item(itemID).item
	if item.HasParent {
		return
	}
	sm.deleteItem(itemID)
}
func (sm *Engine) deleteItem(itemID ItemID) {
	item := sm.Item(itemID).item
	item.OperationKind = OperationKindDelete
	sm.Patch.Item[item.ID] = item
	sm.deleteGearScore(item.GearScore)
}

func (sm *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := sm.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent {
		return
	}
	sm.deleteZoneItem(zoneItemID)
}
func (sm *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := sm.ZoneItem(zoneItemID).zoneItem
	zoneItem.OperationKind = OperationKindDelete
	sm.Patch.ZoneItem[zoneItem.ID] = zoneItem
	sm.deleteItem(zoneItem.Item)
	sm.deletePosition(zoneItem.Position)
}

func (sm *Engine) DeleteZone(zoneID ZoneID) {
	sm.deleteZone(zoneID)
}
func (sm *Engine) deleteZone(zoneID ZoneID) {
	zone := sm.Zone(zoneID).zone
	zone.OperationKind = OperationKindDelete
	sm.Patch.Zone[zone.ID] = zone
	for _, zoneItemID := range zone.Items {
		sm.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		sm.deletePlayer(playerID)
	}
}
