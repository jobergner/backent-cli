package statemachine

func (sm *StateMachine) GetPlayer(playerID PlayerID) Player {
	patchingPlayer, ok := sm.Patch.Player[playerID]
	if ok {
		return Player{patchingPlayer}
	}
	currentPlayer := sm.State.Player[playerID]
	return Player{currentPlayer}
}

func (_e Player) GetID(sm *StateMachine) PlayerID {
	return _e.player.ID
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
	return sm.GetGearScore(e.player.GearScore)
}

func (_e Player) GetPosition(sm *StateMachine) Position {
	e := sm.GetPlayer(_e.player.ID)
	return sm.GetPosition(e.player.Position)
}

func (sm *StateMachine) GetGearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[gearScoreID]
	return GearScore{currentGearScore}
}

func (_e GearScore) GetID(sm *StateMachine) GearScoreID {
	return _e.gearScore.ID
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

func (_e Item) GetID(sm *StateMachine) ItemID {
	return _e.item.ID
}

func (_e Item) GetGearScore(sm *StateMachine) GearScore {
	e := sm.GetItem(_e.item.ID)
	return sm.GetGearScore(e.item.GearScore)
}

func (sm *StateMachine) GetPosition(positionID PositionID) Position {
	patchingPosition, ok := sm.Patch.Position[positionID]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[positionID]
	return Position{currentPosition}
}

func (_e Position) GetID(sm *StateMachine) PositionID {
	return _e.position.ID
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

func (_e ZoneItem) GetID(sm *StateMachine) ZoneItemID {
	return _e.zoneItem.ID
}

func (_e ZoneItem) GetPosition(sm *StateMachine) Position {
	e := sm.GetZoneItem(_e.zoneItem.ID)
	return sm.GetPosition(e.zoneItem.Position)
}

func (_e ZoneItem) GetItem(sm *StateMachine) Item {
	e := sm.GetZoneItem(_e.zoneItem.ID)
	return sm.GetItem(e.zoneItem.Item)
}

func (sm *StateMachine) GetZone(zoneID ZoneID) Zone {
	patchingZone, ok := sm.Patch.Zone[zoneID]
	if ok {
		return Zone{patchingZone}
	}
	currentZone := sm.State.Zone[zoneID]
	return Zone{currentZone}
}

func (_e Zone) GetID(sm *StateMachine) ZoneID {
	return _e.zone.ID
}

func (_e Zone) GetPlayers(sm *StateMachine) []Player {
	e := sm.GetZone(_e.zone.ID)
	var players []Player
	for _, playerID := range e.zone.Players {
		players = append(players, sm.GetPlayer(playerID))
	}
	return players
}

func (_e Zone) GetItems(sm *StateMachine) []ZoneItem {
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
	for _, element := range e.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
