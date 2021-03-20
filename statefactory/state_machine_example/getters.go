package state

func (sm *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := sm.Patch.Player[playerID]
	if ok {
		return Player{patchingPlayer}
	}
	currentPlayer := sm.State.Player[playerID]
	return Player{currentPlayer}
}

func (_e Player) ID(sm *Engine) PlayerID {
	return _e.player.ID
}

func (_e Player) Items(sm *Engine) []Item {
	e := sm.Player(_e.player.ID)
	var items []Item
	for _, itemID := range e.player.Items {
		items = append(items, sm.Item(itemID))
	}
	return items
}

func (_e Player) GearScore(sm *Engine) GearScore {
	e := sm.Player(_e.player.ID)
	return sm.GearScore(e.player.GearScore)
}

func (_e Player) Position(sm *Engine) Position {
	e := sm.Player(_e.player.ID)
	return sm.Position(e.player.Position)
}

func (sm *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := sm.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{patchingGearScore}
	}
	currentGearScore := sm.State.GearScore[gearScoreID]
	return GearScore{currentGearScore}
}

func (_e GearScore) ID(sm *Engine) GearScoreID {
	return _e.gearScore.ID
}

func (_e GearScore) Level(sm *Engine) int {
	e := sm.GearScore(_e.gearScore.ID)
	return e.gearScore.Level
}

func (_e GearScore) Score(sm *Engine) int {
	e := sm.GearScore(_e.gearScore.ID)
	return e.gearScore.Score
}

func (sm *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := sm.Patch.Item[itemID]
	if ok {
		return Item{patchingItem}
	}
	currentItem := sm.State.Item[itemID]
	return Item{currentItem}
}

func (_e Item) ID(sm *Engine) ItemID {
	return _e.item.ID
}

func (_e Item) GearScore(sm *Engine) GearScore {
	e := sm.Item(_e.item.ID)
	return sm.GearScore(e.item.GearScore)
}

func (sm *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := sm.Patch.Position[positionID]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := sm.State.Position[positionID]
	return Position{currentPosition}
}

func (_e Position) ID(sm *Engine) PositionID {
	return _e.position.ID
}

func (_e Position) X(sm *Engine) float64 {
	e := sm.Position(_e.position.ID)
	return e.position.X
}

func (_e Position) Y(sm *Engine) float64 {
	e := sm.Position(_e.position.ID)
	return e.position.Y
}

func (sm *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := sm.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{patchingZoneItem}
	}
	currentZoneItem := sm.State.ZoneItem[zoneItemID]
	return ZoneItem{currentZoneItem}
}

func (_e ZoneItem) ID(sm *Engine) ZoneItemID {
	return _e.zoneItem.ID
}

func (_e ZoneItem) Position(sm *Engine) Position {
	e := sm.ZoneItem(_e.zoneItem.ID)
	return sm.Position(e.zoneItem.Position)
}

func (_e ZoneItem) Item(sm *Engine) Item {
	e := sm.ZoneItem(_e.zoneItem.ID)
	return sm.Item(e.zoneItem.Item)
}

func (sm *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := sm.Patch.Zone[zoneID]
	if ok {
		return Zone{patchingZone}
	}
	currentZone := sm.State.Zone[zoneID]
	return Zone{currentZone}
}

func (_e Zone) ID(sm *Engine) ZoneID {
	return _e.zone.ID
}

func (_e Zone) Players(sm *Engine) []Player {
	e := sm.Zone(_e.zone.ID)
	var players []Player
	for _, playerID := range e.zone.Players {
		players = append(players, sm.Player(playerID))
	}
	return players
}

func (_e Zone) Items(sm *Engine) []ZoneItem {
	e := sm.Zone(_e.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range e.zone.Items {
		items = append(items, sm.ZoneItem(zoneItemID))
	}
	return items
}

func (_e Zone) Tags(sm *Engine) []string {
	e := sm.Zone(_e.zone.ID)
	var tags []string
	for _, element := range e.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
