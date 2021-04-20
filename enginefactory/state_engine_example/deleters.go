package state

func (se *Engine) DeletePlayer(playerID PlayerID) {
	player := se.Player(playerID).player
	if player.HasParent_ {
		return
	}
	se.deletePlayer(playerID)
}
func (se *Engine) deletePlayer(playerID PlayerID) {
	player := se.Player(playerID).player
	for _, id := range se.allItemIDs() {
		item := se.Item(id)
		if item.item.BoundTo.id == playerID {
			item.item.BoundTo.Unset(se)
		}
	}
	for _, id := range se.allPlayerIDs() {
		player := se.Player(id)
		player.RemoveGuildMembers(se, playerID)
	}

	player.OperationKind_ = OperationKindDelete
	se.Patch.Player[player.ID] = player

	se.deleteGearScore(player.GearScore)
	for _, itemID := range player.Items {
		se.deleteItem(itemID)
	}
	se.deletePosition(player.Position)
}

func (se *Engine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := se.GearScore(gearScoreID).gearScore
	if gearScore.HasParent_ {
		return
	}
	se.deleteGearScore(gearScoreID)
}
func (se *Engine) deleteGearScore(gearScoreID GearScoreID) {
	gearScore := se.GearScore(gearScoreID).gearScore
	gearScore.OperationKind_ = OperationKindDelete
	se.Patch.GearScore[gearScore.ID] = gearScore
}

func (se *Engine) DeletePosition(positionID PositionID) {
	position := se.Position(positionID).position
	if position.HasParent_ {
		return
	}
	se.deletePosition(positionID)
}
func (se *Engine) deletePosition(positionID PositionID) {
	position := se.Position(positionID).position
	position.OperationKind_ = OperationKindDelete
	se.Patch.Position[position.ID] = position
}

func (se *Engine) DeleteItem(itemID ItemID) {
	item := se.Item(itemID).item
	if item.HasParent_ {
		return
	}
	se.deleteItem(itemID)
}
func (se *Engine) deleteItem(itemID ItemID) {
	item := se.Item(itemID).item
	item.OperationKind_ = OperationKindDelete
	se.Patch.Item[item.ID] = item
	se.deleteGearScore(item.GearScore)
}

func (se *Engine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := se.ZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent_ {
		return
	}
	se.deleteZoneItem(zoneItemID)
}
func (se *Engine) deleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := se.ZoneItem(zoneItemID).zoneItem
	zoneItem.OperationKind_ = OperationKindDelete
	se.Patch.ZoneItem[zoneItem.ID] = zoneItem
	se.deleteItem(zoneItem.Item)
	se.deletePosition(zoneItem.Position)
}

func (se *Engine) DeleteZone(zoneID ZoneID) {
	se.deleteZone(zoneID)
}
func (se *Engine) deleteZone(zoneID ZoneID) {
	zone := se.Zone(zoneID).zone
	zone.OperationKind_ = OperationKindDelete
	se.Patch.Zone[zone.ID] = zone
	for _, zoneItemID := range zone.Items {
		se.deleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		se.deletePlayer(playerID)
	}
}
