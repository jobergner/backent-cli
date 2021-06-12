package state

func (_any anyOfPlayer_ZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	return any.anyOfPlayer_ZoneItem.ElementKind
}

func (_any anyOfPlayer_ZoneItem) SetZoneItem() zoneItem {
	zoneItem := _any.anyOfPlayer_ZoneItem.engine.createZoneItem(true)
	_any.anyOfPlayer_ZoneItem.setZoneItem(zoneItem.ID())
	return zoneItem
}

func (_any anyOfPlayer_ZoneItemCore) setZoneItem(zoneItemID ZoneItemID) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}

func (_any anyOfPlayer_ZoneItem) SetPlayer() player {
	player := _any.anyOfPlayer_ZoneItem.engine.createPlayer(true)
	_any.anyOfPlayer_ZoneItem.setPlayer(player.ID())
	return player
}

func (_any anyOfPlayer_ZoneItemCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	if any.ZoneItem != 0 {
		any.engine.deleteZoneItem(any.ZoneItem)
		any.ZoneItem = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}

func (_any anyOfPlayer_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(any.ZoneItem)
	}
}

func (_any anyOfPlayer_Position) Kind() ElementKind {
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	return any.anyOfPlayer_Position.ElementKind
}

func (_any anyOfPlayer_Position) SetPosition() position {
	position := _any.anyOfPlayer_Position.engine.createPosition(true)
	_any.anyOfPlayer_Position.setPosition(position.ID())
	return position
}

func (_any anyOfPlayer_PositionCore) setPosition(positionID PositionID) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindPosition
	any.Position = positionID
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}

func (_any anyOfPlayer_PositionCore) deleteChild() {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindPosition:
		any.engine.deletePosition(any.Position)
	}
}

func (_any anyOfPlayer_Position) SetPlayer() player {
	player := _any.anyOfPlayer_Position.engine.createPlayer(true)
	_any.anyOfPlayer_Position.setPlayer(player.ID())
	return player
}

func (_any anyOfPlayer_PositionCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	if any.Position != 0 {
		any.engine.deletePosition(any.Position)
		any.Position = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayer_Position[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItem) Kind() ElementKind {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	return any.anyOfItem_Player_ZoneItem.ElementKind
}

func (_any anyOfItem_Player_ZoneItem) SetZoneItem() zoneItem {
	zoneItem := _any.anyOfItem_Player_ZoneItem.engine.createZoneItem(true)
	_any.anyOfItem_Player_ZoneItem.setZoneItem(zoneItem.ID())
	return zoneItem
}

func (_any anyOfItem_Player_ZoneItemCore) setZoneItem(zoneItemID ZoneItemID) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if any.Item != 0 {
		any.engine.deleteItem(any.Item)
		any.Item = 0
	}
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItem) SetPlayer() player {
	player := _any.anyOfItem_Player_ZoneItem.engine.createPlayer(true)
	_any.anyOfItem_Player_ZoneItem.setPlayer(player.ID())
	return player
}

func (_any anyOfItem_Player_ZoneItemCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if any.Item != 0 {
		any.engine.deleteItem(any.Item)
		any.Item = 0
	}
	if any.ZoneItem != 0 {
		any.engine.deleteZoneItem(any.ZoneItem)
		any.ZoneItem = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItem) SetItem() item {
	item := _any.anyOfItem_Player_ZoneItem.engine.createItem(true)
	_any.anyOfItem_Player_ZoneItem.setItem(item.ID())
	return item
}

func (_any anyOfItem_Player_ZoneItemCore) setItem(itemID ItemID) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	if any.ZoneItem != 0 {
		any.engine.deleteZoneItem(any.ZoneItem)
		any.ZoneItem = 0
	}
	any.ElementKind = ElementKindItem
	any.Item = itemID
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItemCore) deleteChild() {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	switch any.ElementKind {
	case ElementKindItem:
		any.engine.deleteItem(any.Item)
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(any.ZoneItem)
	}
}
