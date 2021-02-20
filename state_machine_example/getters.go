package statemachine

func (sm *StateMachine) GetPlayer(playerID PlayerID) Player {
	patchingPlayer, ok := sm.Patch.Player[playerID]
	if ok {
		return patchingPlayer
	}
	currentPlayer := sm.State.Player[playerID]
	return currentPlayer
}

func (e Player) GetItems(sm *StateMachine) []Item {
	var items []Item
	for _, itemID := range e.Items {
		items = append(items, sm.GetItem(itemID))
	}
	return items
}

func (e Player) GetGearScore(sm *StateMachine) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[e.GearScore]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.State.GearScore[e.GearScore]
	return currentGearScore
}

func (e Player) GetPosition(sm *StateMachine) Position {
	patchingPosition, ok := sm.Patch.Position[e.Position]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.State.Position[e.Position]
	return currentPosition
}

func (sm *StateMachine) GetGearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[gearScoreID]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.State.GearScore[gearScoreID]
	return currentGearScore
}

func (e GearScore) GetLevel() int {
	return e.Level
}

func (e GearScore) GetScore() int {
	return e.Score
}

func (sm *StateMachine) GetItem(itemID ItemID) Item {
	patchingItem, ok := sm.Patch.Item[itemID]
	if ok {
		return patchingItem
	}
	currentItem := sm.State.Item[itemID]
	return currentItem
}

func (e Item) GetGearScore(sm *StateMachine) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[e.GearScore]
	if ok {
		return patchingGearScore
	}
	currentGearScore := sm.State.GearScore[e.GearScore]
	return currentGearScore
}

func (sm *StateMachine) GetPosition(positionID PositionID) Position {
	patchingPosition, ok := sm.Patch.Position[positionID]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.State.Position[positionID]
	return currentPosition
}

func (e Position) GetX() float64 {
	return e.X
}

func (e Position) GetY() float64 {
	return e.Y
}

func (sm *StateMachine) GetZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := sm.Patch.ZoneItem[zoneItemID]
	if ok {
		return patchingZoneItem
	}
	currentZoneItem := sm.State.ZoneItem[zoneItemID]
	return currentZoneItem
}

func (e ZoneItem) GetPosition(sm *StateMachine) Position {
	patchingPosition, ok := sm.Patch.Position[e.Position]
	if ok {
		return patchingPosition
	}
	currentPosition := sm.State.Position[e.Position]
	return currentPosition
}

func (e ZoneItem) GetItem(sm *StateMachine) Item {
	patchingItem, ok := sm.Patch.Item[e.Item]
	if ok {
		return patchingItem
	}
	currentItem := sm.State.Item[e.Item]
	return currentItem
}

func (sm *StateMachine) GetZone(zoneID ZoneID) Zone {
	patchingZone, ok := sm.Patch.Zone[zoneID]
	if ok {
		return patchingZone
	}
	currentZone := sm.State.Zone[zoneID]
	return currentZone
}

func (e Zone) GetPlayers(sm *StateMachine) []Player {
	var players []Player
	for _, playerID := range e.Players {
		players = append(players, sm.GetPlayer(playerID))
	}
	return players
}

func (e Zone) GetZoneItems(sm *StateMachine) []ZoneItem {
	var items []ZoneItem
	for _, zoneItemID := range e.Items {
		items = append(items, sm.GetZoneItem(zoneItemID))
	}
	return items
}
