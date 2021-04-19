package state

func (se *Engine) Player(playerID PlayerID) player {
	patchingPlayer, ok := se.Patch.Player[playerID]
	if ok {
		return player{player: patchingPlayer}
	}
	currentPlayer, ok := se.State.Player[playerID]
	if ok {
		return player{player: currentPlayer}
	}
	return player{player: playerCore{OperationKind_: OperationKindDelete}}
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

func (_player player) GuildMembers(se *Engine) []playerSliceRef {
	player := se.Player(_player.player.ID)
	return player.player.GuildMembers
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
	currentGearScore, ok := se.State.GearScore[gearScoreID]
	if ok {
		return gearScore{gearScore: currentGearScore}
	}
	return gearScore{gearScore: gearScoreCore{OperationKind_: OperationKindDelete}}
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
	currentItem, ok := se.State.Item[itemID]
	if ok {
		return item{item: currentItem}
	}
	return item{item: itemCore{OperationKind_: OperationKindDelete}}
}

func (_item item) ID(se *Engine) ItemID {
	return _item.item.ID
}

func (_item item) GearScore(se *Engine) gearScore {
	item := se.Item(_item.item.ID)
	return se.GearScore(item.item.GearScore)
}

func (_item item) BoundTo(se *Engine) itemBoundToRef {
	item := se.Item(_item.item.ID)
	return item.item.BoundTo
}

func (se *Engine) Position(positionID PositionID) position {
	patchingPosition, ok := se.Patch.Position[positionID]
	if ok {
		return position{position: patchingPosition}
	}
	currentPosition, ok := se.State.Position[positionID]
	if ok {
		return position{position: currentPosition}
	}
	return position{position: positionCore{OperationKind_: OperationKindDelete}}
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
	currentZoneItem, ok := se.State.ZoneItem[zoneItemID]
	if ok {
		return zoneItem{zoneItem: currentZoneItem}
	}
	return zoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
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
	currentZone, ok := se.State.Zone[zoneID]
	if ok {
		return zone{zone: currentZone}
	}
	return zone{zone: zoneCore{OperationKind_: OperationKindDelete}}
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
