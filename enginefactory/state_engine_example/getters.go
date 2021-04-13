package state

func (se *Engine) Player(playerID PlayerID) player {
	patchingPlayer, ok := se.Patch.Player[playerID]
	if ok {
		return player{player: patchingPlayer}
	}
	currentPlayer := se.State.Player[playerID]
	return player{player: currentPlayer}
}

func (_player player) ID(se *Engine) PlayerID {
	return _player.player.ID
}

func (_player player) Items(se *Engine) []item {
	player := se.Player(_player.player.ID)
	var items []item
	for _, itemID := range player.player.Items {
		items = append(items, se.Item(itemID))
	}
	return items
}

func (_player player) GearScore(se *Engine) gearScore {
	player := se.Player(_player.player.ID)
	return se.GearScore(player.player.GearScore)
}

func (_player player) Position(se *Engine) position {
	player := se.Player(_player.player.ID)
	return se.Position(player.player.Position)
}

func (se *Engine) GearScore(gearScoreID GearScoreID) gearScore {
	patchingGearScore, ok := se.Patch.GearScore[gearScoreID]
	if ok {
		return gearScore{gearScore: patchingGearScore}
	}
	currentGearScore := se.State.GearScore[gearScoreID]
	return gearScore{gearScore: currentGearScore}
}

func (_gearScore gearScore) ID(se *Engine) GearScoreID {
	return _gearScore.gearScore.ID
}

func (_gearScore gearScore) Level(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Level
}

func (_gearScore gearScore) Score(se *Engine) int {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	return gearScore.gearScore.Score
}

func (se *Engine) Item(itemID ItemID) item {
	patchingItem, ok := se.Patch.Item[itemID]
	if ok {
		return item{item: patchingItem}
	}
	currentItem := se.State.Item[itemID]
	return item{item: currentItem}
}

func (_item item) ID(se *Engine) ItemID {
	return _item.item.ID
}

func (_item item) GearScore(se *Engine) gearScore {
	item := se.Item(_item.item.ID)
	return se.GearScore(item.item.GearScore)
}

func (se *Engine) Position(positionID PositionID) position {
	patchingPosition, ok := se.Patch.Position[positionID]
	if ok {
		return position{position: patchingPosition}
	}
	currentPosition := se.State.Position[positionID]
	return position{position: currentPosition}
}

func (_position position) ID(se *Engine) PositionID {
	return _position.position.ID
}

func (_position position) X(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.X
}

func (_position position) Y(se *Engine) float64 {
	position := se.Position(_position.position.ID)
	return position.position.Y
}

func (se *Engine) ZoneItem(zoneItemID ZoneItemID) zoneItem {
	patchingZoneItem, ok := se.Patch.ZoneItem[zoneItemID]
	if ok {
		return zoneItem{zoneItem: patchingZoneItem}
	}
	currentZoneItem := se.State.ZoneItem[zoneItemID]
	return zoneItem{zoneItem: currentZoneItem}
}

func (_zoneItem zoneItem) ID(se *Engine) ZoneItemID {
	return _zoneItem.zoneItem.ID
}

func (_zoneItem zoneItem) Position(se *Engine) position {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Position(zoneItem.zoneItem.Position)
}

func (_zoneItem zoneItem) Item(se *Engine) item {
	zoneItem := se.ZoneItem(_zoneItem.zoneItem.ID)
	return se.Item(zoneItem.zoneItem.Item)
}

func (se *Engine) Zone(zoneID ZoneID) zone {
	patchingZone, ok := se.Patch.Zone[zoneID]
	if ok {
		return zone{zone: patchingZone}
	}
	currentZone := se.State.Zone[zoneID]
	return zone{zone: currentZone}
}

func (_zone zone) ID(se *Engine) ZoneID {
	return _zone.zone.ID
}

func (_zone zone) Players(se *Engine) []player {
	zone := se.Zone(_zone.zone.ID)
	var players []player
	for _, playerID := range zone.zone.Players {
		players = append(players, se.Player(playerID))
	}
	return players
}

func (_zone zone) Items(se *Engine) []zoneItem {
	zone := se.Zone(_zone.zone.ID)
	var items []zoneItem
	for _, zoneItemID := range zone.zone.Items {
		items = append(items, se.ZoneItem(zoneItemID))
	}
	return items
}

func (_zone zone) Tags(se *Engine) []string {
	zone := se.Zone(_zone.zone.ID)
	var tags []string
	for _, element := range zone.zone.Tags {
		tags = append(tags, element)
	}
	return tags
}
