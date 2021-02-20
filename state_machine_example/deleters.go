package statemachine

func (sm *StateMachine) DeletePlayer(playerID PlayerID) {
	player := sm.GetPlayer(playerID)
	player.OperationKind = OperationKindDelete
	sm.Patch.Player[player.ID] = player
	for _, itemID := range player.Items {
		sm.DeleteItem(itemID)
	}
	sm.DeleteGearScore(player.GearScore)
	sm.DeletePosition(player.Position)
}

func (sm *StateMachine) DeleteGearScore(gearScoreID GearScoreID) {
	gearScore := sm.GetGearScore(gearScoreID)
	gearScore.OperationKind = OperationKindDelete
	sm.Patch.GearScore[gearScore.ID] = gearScore
}

func (sm *StateMachine) DeletePosition(positionID PositionID) {
	position := sm.GetPosition(positionID)
	position.OperationKind = OperationKindDelete
	sm.Patch.Position[position.ID] = position
}

func (sm *StateMachine) DeleteItem(itemID ItemID) {
	item := sm.GetItem(itemID)
	item.OperationKind = OperationKindDelete
	sm.Patch.Item[item.ID] = item
}

func (sm *StateMachine) DeleteZoneItem(zoneItemID ZoneItemID) {
	zoneItem := sm.GetZoneItem(zoneItemID)
	zoneItem.OperationKind = OperationKindDelete
	sm.Patch.ZoneItem[zoneItem.ID] = zoneItem
}

func (sm *StateMachine) DeleteZone(zoneID ZoneID) {
	zone := sm.GetZone(zoneID)
	zone.OperationKind = OperationKindDelete
	sm.Patch.Zone[zone.ID] = zone
	for _, playerID := range zone.Players {
		sm.DeletePlayer(playerID)
	}
	for _, zoneItemID := range zone.Items {
		sm.DeleteZoneItem(zoneItemID)
	}
}
