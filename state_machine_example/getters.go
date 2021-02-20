package statemachine

func (sm *StateMachine) GetPlayer(playerID PlayerID) Player {
	patchingPlayer, ok := sm.Patch.Player[playerID]
	if ok {
		return Player{patchingPlayer}
	}
	currentPlayer := sm.State.Player[playerID]
	return Player{currentPlayer}
}

func (e Player) GetItems(sm *StateMachine) []Item {
	var items []Item
	for _, itemID := range e.player.Items {
		items = append(items, sm.GetItem(itemID))
	}
	return items
}

func (e Player) GetGearScore(sm *StateMachine) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[e.player.GearScore]
	if ok {
		return GearScore{patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[e.player.GearScore]
	return GearScore{currentGearScore}
}

func (e Player) GetPosition(sm *StateMachine) Position {
	patchingPosition, ok := sm.Patch.Position[e.player.Position]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[e.player.Position]
	return Position{currentPosition}
}

func (sm *StateMachine) GetGearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[gearScoreID]
	return GearScore{gearScore: currentGearScore}
}

func (e GearScore) GetLevel() int {
	return e.gearScore.Level
}

func (e GearScore) GetScore() int {
	return e.gearScore.Score
}

func (sm *StateMachine) GetItem(itemID ItemID) Item {
	patchingItem, ok := sm.Patch.Item[itemID]
	if ok {
		return Item{patchingItem}
	}
	currentItem := sm.State.Item[itemID]
	return Item{currentItem}
}

func (e Item) GetGearScore(sm *StateMachine) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[e.item.GearScore]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[e.item.GearScore]
	return GearScore{gearScore: currentGearScore}
}

func (sm *StateMachine) GetPosition(positionID PositionID) Position {
	patchingPosition, ok := sm.Patch.Position[positionID]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[positionID]
	return Position{currentPosition}
}

func (e Position) GetX() float64 {
	return e.position.X
}

func (e Position) GetY() float64 {
	return e.position.Y
}

func (sm *StateMachine) GetZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := sm.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{patchingZoneItem}
	}
	currentZoneItem := sm.State.ZoneItem[zoneItemID]
	return ZoneItem{currentZoneItem}
}

func (e ZoneItem) GetPosition(sm *StateMachine) Position {
	patchingPosition, ok := sm.Patch.Position[e.zoneItem.Position]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[e.zoneItem.Position]
	return Position{currentPosition}
}

func (e ZoneItem) GetItem(sm *StateMachine) Item {
	patchingItem, ok := sm.Patch.Item[e.zoneItem.Item]
	if ok {
		return Item{patchingItem}
	}
	currentItem := sm.State.Item[e.zoneItem.Item]
	return Item{currentItem}
}

func (sm *StateMachine) GetZone(zoneID ZoneID) Zone {
	patchingZone, ok := sm.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone := sm.State.Zone[zoneID]
	return Zone{zone: currentZone}
}

func (e Zone) GetPlayers(sm *StateMachine) []Player {
	var players []Player
	for _, playerID := range e.zone.Players {
		players = append(players, sm.GetPlayer(playerID))
	}
	return players
}

func (e Zone) GetZoneItems(sm *StateMachine) []ZoneItem {
	var items []ZoneItem
	for _, zoneItemID := range e.zone.Items {
		items = append(items, sm.GetZoneItem(zoneItemID))
	}
	return items
}
