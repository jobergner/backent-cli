package state

func (se *Engine) Player(playerID PlayerID) Player {
	patchingPlayer, ok := se.Patch.Player[playerID]
	if ok {
		return Player{player: patchingPlayer}
	}
	currentPlayer := se.State.Player[playerID]
	return Player{player: currentPlayer}
}

func (_player Player) ID(se *Engine) PlayerID {
	return _player.player.ID
}

func (_player Player) Items(se *Engine) []Item {
	player := se.Player(_player.player.ID)
	var items []Item
	for _, itemID := range player.player.Items {
		items = append(items, se.Item(itemID))
	}
	return items
}

func (_player Player) GearScore(se *Engine) GearScore {
	player := se.Player(_player.player.ID)
	return se.GearScore(player.player.GearScore)
}

func (_player Player) Position(se *Engine) Position {
	player := se.Player(_player.player.ID)
	return se.Position(player.player.Position)
}

func (se *Engine) GearScore(gearScoreID GearScoreID) GearScore {
	patchingGearScore, ok := se.Patch.GearScore[gearScoreID]
	if ok {
		return GearScore{gearScore: patchingGearScore}
	}
	currentGearScore := se.State.GearScore[gearScoreID]
	return GearScore{gearScore: currentGearScore}
}

func (_gearScore GearScore) ID(se *Engine) GearScoreID {
	return _gearScore.gearScore.ID
}

func (_gearScore GearScore) Level(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Level
}

func (_gearScore GearScore) Score(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Score
}

func (se *Engine) Item(itemID ItemID) Item {
	patchingItem, ok := se.Patch.Item[itemID]
	if ok {
		return Item{item: patchingItem}
	}
	currentItem := se.State.Item[itemID]
	return Item{item: currentItem}
}

func (_item Item) ID(se *Engine) ItemID {
	return _item.item.ID
}

func (_item Item) GearScore(se *Engine) GearScore {
	item := se.Item(_item.item.ID)
	return se.GearScore(item.item.GearScore)
}

func (se *Engine) Position(positionID PositionID) Position {
	patchingPosition, ok := se.Patch.Position[positionID]
	if ok {
		return Position{position: patchingPosition}
	}
	currentPosition := se.State.Position[positionID]
	return Position{position: currentPosition}
}

func (_position Position) ID(se *Engine) PositionID {
	return _position.position.ID
}

func (_position Position) X(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.X
}

func (_position Position) Y(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.Y
}

func (se *Engine) ZoneItem(zoneItemID ZoneItemID) ZoneItem {
	patchingZoneItem, ok := se.Patch.ZoneItem[zoneItemID]
	if ok {
		return ZoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem := se.State.ZoneItem[zoneItemID]
	return ZoneItem{zoneItem: currentZoneItem}
}

func (_zoneItem ZoneItem) ID(se *Engine) ZoneItemID {
	return _zoneItem.zoneItem.ID
}

func (_zoneItem ZoneItem) Position(se *Engine) Position {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Position(zoneItem.zoneItem.Position)
}

func (_zoneItem ZoneItem) Item(se *Engine) Item {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Item(zoneItem.zoneItem.Item)
}

func (se *Engine) Zone(zoneID ZoneID) Zone {
	patchingZone, ok := se.Patch.Zone[zoneID]
	if ok {
		return Zone{zone: patchingZone}
	}
	currentZone := se.State.Zone[zoneID]
	return Zone{zone: currentZone}
}

func (_zone Zone) ID(se *Engine) ZoneID {
	return _zone.zone.ID
}

func (_zone Zone) Players(se *Engine) []Player {
	zone := se.Zone(_zone.zone.ID)
	var players []Player
	for _, playerID := range zone.zone.Players {
		players = append(players, se.Player(playerID))
	}
	return players
}

func (_zone Zone) Items(se *Engine) []ZoneItem {
	zone := se.Zone(_zone.zone.ID)
	var items []ZoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, se.ZoneItem(zoneItemID))
	}
	return items
}

func (_zone Zone) Tags(se *Engine) []string {
	zone := se.Zone(_zone.zone.ID)
	var tags []string
	for _, element := range zone.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
