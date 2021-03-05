package statemachine

func (sm *StateMachine) GetPlayer(playerID PlayerID) Player {
	patchingPlayer, ok := sm.Patch.Player[playerID]
	if ok {
		return Player{patchingPlayer}
	}
	currentPlayer := sm.State.Player[playerID]
	return Player{currentPlayer}
}

func (_e Player) GetItems(sm *StateMachine) []Item {
	e := sm.GetPlayer(_e.player.ID)
	var items []Item
	for _, itemID := range e.player.Items {
		items = append(items, sm.GetItem(itemID))
	}
	return items
}

func (_e Player) GetGearScore(sm *StateMachine) GearScore {
	e := sm.GetPlayer(_e.player.ID)
	patchingGearScore, ok := sm.Patch.GearScore[e.player.GearScore]
	if ok {
		return GearScore{patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[e.player.GearScore]
	return GearScore{currentGearScore}
}

func (_e Player) GetPosition(sm *StateMachine) Position {
	e := sm.GetPlayer(_e.player.ID)
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

func (_e GearScore) GetLevel(sm *StateMachine) int {
	e := sm.GetGearScore(_e.gearScore.ID)
	return e.gearScore.Level
}

func (_e GearScore) GetScore(sm *StateMachine) int {
	e := sm.GetGearScore(_e.gearScore.ID)
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

func (_e Item) GetGearScore(sm *StateMachine) GearScore {
	e := sm.GetItem(_e.item.ID)
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

func (_e Position) GetX(sm *StateMachine) float64 {
	e := sm.GetPosition(_e.position.ID)
	return e.position.X
}

func (_e Position) GetY(sm *StateMachine) float64 {
	e := sm.GetPosition(_e.position.ID)
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

func (_e ZoneItem) GetPosition(sm *StateMachine) Position {
	e := sm.GetZoneItem(_e.zoneItem.ID)
	patchingPosition, ok := sm.Patch.Position[e.zoneItem.Position]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[e.zoneItem.Position]
	return Position{currentPosition}
}

func (_e ZoneItem) GetItem(sm *StateMachine) Item {
	e := sm.GetZoneItem(_e.zoneItem.ID)
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

func (_e Zone) GetPlayers(sm *StateMachine) []Player {
	e := sm.GetZone(_e.zone.ID)
	var players []Player
	for _, playerID := range e.zone.Players {
		players = append(players, sm.GetPlayer(playerID))
	}
	return players
}

func (_e Zone) GetZoneItems(sm *StateMachine) []ZoneItem {
	e := sm.GetZone(_e.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range e.zone.Items {
		items = append(items, sm.GetZoneItem(zoneItemID))
	}
	return items
}

func (_e Zone) GetTags(sm *StateMachine) []string {
	e := sm.GetZone(_e.zone.ID)
	var tags []string
	for _, tag := range e.zone.Tags {
		tags = append(tags, tag)
	}
	return tags
}
