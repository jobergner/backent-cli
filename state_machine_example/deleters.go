package statemachine

func (sm *StateMachine) DeletePlayer(playerID PlayerID) {
	player := sm.GetPlayer(playerID).player
	if player.HasParent {
		return
	}
	player.OperationKind = OperationKindDelete
	sm.Patch.Player[player.ID] = player
	sm.DeleteGearScore(player.GearScore)
	for _, itemID := range player.Items {
		sm.DeleteItem(itemID)
	}
	sm.DeletePosition(player.Position)
}

func (sm *StateMachine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := sm.GetGearScore(gearScoreID).gearScore
	if gearScore.HasParent {
		return
	}
	gearScore.OperationKind = OperationKindDelete
	sm.Patch.GearScore[gearScore.ID] = gearScore
}

func (sm *StateMachine) DeletePosition(positionID PositionID) {
	position := sm.GetPosition(positionID).position
	if position.HasParent {
		return
	}
	position.OperationKind = OperationKindDelete
	sm.Patch.Position[position.ID] = position
}

func (sm *StateMachine) DeleteItem(itemID ItemID) {
	item := sm.GetItem(itemID).item
	if item.HasParent {
		return
	}
	item.OperationKind = OperationKindDelete
	sm.Patch.Item[item.ID] = item
	sm.DeleteGearScore(item.GearScore)
}

func (sm *StateMachine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := sm.GetZoneItem(zoneItemID).zoneItem
	if zoneItem.HasParent {
		return
	}
	zoneItem.OperationKind = OperationKindDelete
	sm.Patch.ZoneItem[zoneItem.ID] = zoneItem
	sm.DeleteItem(zoneItem.Item)
	sm.DeletePosition(zoneItem.Position)
}

func (sm *StateMachine) DeleteZone(zoneID ZoneID) {
	zone := sm.GetZone(zoneID).zone
	zone.OperationKind = OperationKindDelete
	sm.Patch.Zone[zone.ID] = zone
	for _, zoneItemID := range zone.Items {
		sm.DeleteZoneItem(zoneItemID)
	}
	for _, playerID := range zone.Players {
		sm.DeletePlayer(playerID)
	}
}
