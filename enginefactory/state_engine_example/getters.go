package state

func (se *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := se.Patch.Player[playerID]
	if ok {
		return Player{patchingPlayer}
	}
	currentPlayer := se.State.Player[playerID]
	return Player{currentPlayer}
}

func (_e Player) ID(se *Engine) PlayerID {
	return _e.player.ID
}

func (_e Player) Items(se *Engine) []Item {
	e := se.Player(_e.player.ID)
	var items []Item
	for _, itemID := range e.player.Items {
		items = append(items, se.Item(itemID))
	}
	return items
}

func (_e Player) GearScore(se *Engine) GearScore {
	e := se.Player(_e.player.ID)
	return se.GearScore(e.player.GearScore)
}

func (_e Player) Position(se *Engine) Position {
	e := se.Player(_e.player.ID)
	return se.Position(e.player.Position)
}

func (se *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := se.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{patchingGearScore}
	}
	currentGearScore := se.State.GearScore[gearScoreID]
	return GearScore{currentGearScore}
}

func (_e GearScore) ID(se *Engine) GearScoreID {
	return _e.gearScore.ID
}

func (_e GearScore) Level(se *Engine) int {
	e := se.GearScore(_e.gearScore.ID)
	return e.gearScore.Level
}

func (_e GearScore) Score(se *Engine) int {
	e := se.GearScore(_e.gearScore.ID)
	return e.gearScore.Score
}

func (se *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := se.Patch.Item[itemID]
	if ok {
		return Item{patchingItem}
	}
	currentItem := se.State.Item[itemID]
	return Item{currentItem}
}

func (_e Item) ID(se *Engine) ItemID {
	return _e.item.ID
}

func (_e Item) GearScore(se *Engine) GearScore {
	e := se.Item(_e.item.ID)
	return se.GearScore(e.item.GearScore)
}

func (se *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := se.Patch.Position[positionID]
	if ok {
		return Position{patchingPosition}
	}
	currentPosition := se.State.Position[positionID]
	return Position{currentPosition}
}

func (_e Position) ID(se *Engine) PositionID {
	return _e.position.ID
}

func (_e Position) X(se *Engine) float64 {
	e := se.Position(_e.position.ID)
	return e.position.X
}

func (_e Position) Y(se *Engine) float64 {
	e := se.Position(_e.position.ID)
	return e.position.Y
}

func (se *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := se.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{patchingZoneItem}
	}
	currentZoneItem := se.State.ZoneItem[zoneItemID]
	return ZoneItem{currentZoneItem}
}

func (_e ZoneItem) ID(se *Engine) ZoneItemID {
	return _e.zoneItem.ID
}

func (_e ZoneItem) Position(se *Engine) Position {
	e := se.ZoneItem(_e.zoneItem.ID)
	return se.Position(e.zoneItem.Position)
}

func (_e ZoneItem) Item(se *Engine) Item {
	e := se.ZoneItem(_e.zoneItem.ID)
	return se.Item(e.zoneItem.Item)
}

func (se *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := se.Patch.Zone[zoneID]
	if ok {
		return Zone{patchingZone}
	}
	currentZone := se.State.Zone[zoneID]
	return Zone{currentZone}
}

func (_e Zone) ID(se *Engine) ZoneID {
	return _e.zone.ID
}

func (_e Zone) Players(se *Engine) []Player {
	e := se.Zone(_e.zone.ID)
	var players []Player
	for _, playerID := range e.zone.Players {
		players = append(players, se.Player(playerID))
	}
	return players
}

func (_e Zone) Items(se *Engine) []ZoneItem {
	e := se.Zone(_e.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range e.zone.Items {
		items = append(items, se.ZoneItem(zoneItemID))
	}
	return items
}

func (_e Zone) Tags(se *Engine) []string {
	e := se.Zone(_e.zone.ID)
	var tags []string
	for _, element := range e.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
