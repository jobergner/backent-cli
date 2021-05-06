package state

func (_any anyOfPlayerZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayerZoneItem.engine.anyOfPlayerZoneItem(_any.anyOfPlayerZoneItem.ID)
	return any.anyOfPlayerZoneItem.ElementKind
}

func (_any anyOfPlayerZoneItem) SetZoneItem() zoneItem {
	zoneItem := _any.anyOfPlayerZoneItem.engine.createZoneItem(true)
	_any.anyOfPlayerZoneItem.setZoneItem(zoneItem.ID())
	return zoneItem
}

func (_any anyOfPlayerZoneItemCore) setZoneItem(zoneItemID ZoneItemID) {
	any := _any.engine.anyOfPlayerZoneItem(_any.ID).anyOfPlayerZoneItem
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfPlayerZoneItem[any.ID] = any
}

func (_any anyOfPlayerZoneItem) SetPlayer() player {
	player := _any.anyOfPlayerZoneItem.engine.createPlayer(true)
	_any.anyOfPlayerZoneItem.setPlayer(player.ID())
	return player
}

func (_any anyOfPlayerZoneItemCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayerZoneItem(_any.ID).anyOfPlayerZoneItem
	if any.ZoneItem != 0 {
		any.engine.deleteZoneItem(any.ZoneItem)
		any.ZoneItem = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayerZoneItem[any.ID] = any
}

func (_any anyOfPlayerZoneItemCore) deleteChild() {
	any := _any.engine.anyOfPlayerZoneItem(_any.ID).anyOfPlayerZoneItem
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(any.ZoneItem)
	}
}

func (_any anyOfPlayerPosition) Kind() ElementKind {
	any := _any.anyOfPlayerPosition.engine.anyOfPlayerPosition(_any.anyOfPlayerPosition.ID)
	return any.anyOfPlayerPosition.ElementKind
}

func (_any anyOfPlayerPosition) SetPosition() position {
	position := _any.anyOfPlayerPosition.engine.createPosition(true)
	_any.anyOfPlayerPosition.setPosition(position.ID())
	return position
}

func (_any anyOfPlayerPositionCore) setPosition(positionID PositionID) {
	any := _any.engine.anyOfPlayerPosition(_any.ID).anyOfPlayerPosition
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindPosition
	any.Position = positionID
	any.engine.Patch.AnyOfPlayerPosition[any.ID] = any
}

func (_any anyOfPlayerPositionCore) deleteChild() {
	any := _any.engine.anyOfPlayerPosition(_any.ID).anyOfPlayerPosition
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindPosition:
		any.engine.deletePosition(any.Position)
	}
}

func (_any anyOfPlayerPosition) SetPlayer() player {
	player := _any.anyOfPlayerPosition.engine.createPlayer(true)
	_any.anyOfPlayerPosition.setPlayer(player.ID())
	return player
}

func (_any anyOfPlayerPositionCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayerPosition(_any.ID).anyOfPlayerPosition
	if any.Position != 0 {
		any.engine.deletePosition(any.Position)
		any.Position = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayerPosition[any.ID] = any
}

func (_any anyOfItemPlayerZoneItem) Kind() ElementKind {
	any := _any.anyOfItemPlayerZoneItem.engine.anyOfItemPlayerZoneItem(_any.anyOfItemPlayerZoneItem.ID)
	return any.anyOfItemPlayerZoneItem.ElementKind
}

func (_any anyOfItemPlayerZoneItem) SetZoneItem() zoneItem {
	zoneItem := _any.anyOfItemPlayerZoneItem.engine.createZoneItem(true)
	_any.anyOfItemPlayerZoneItem.setZoneItem(zoneItem.ID())
	return zoneItem
}

func (_any anyOfItemPlayerZoneItemCore) setZoneItem(zoneItemID ZoneItemID) {
	any := _any.engine.anyOfItemPlayerZoneItem(_any.ID).anyOfItemPlayerZoneItem
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
	any.engine.Patch.AnyOfItemPlayerZoneItem[any.ID] = any
}

func (_any anyOfItemPlayerZoneItem) SetPlayer() player {
	player := _any.anyOfItemPlayerZoneItem.engine.createPlayer(true)
	_any.anyOfItemPlayerZoneItem.setPlayer(player.ID())
	return player
}

func (_any anyOfItemPlayerZoneItemCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfItemPlayerZoneItem(_any.ID).anyOfItemPlayerZoneItem
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
	any.engine.Patch.AnyOfItemPlayerZoneItem[any.ID] = any
}

func (_any anyOfItemPlayerZoneItem) SetItem() item {
	item := _any.anyOfItemPlayerZoneItem.engine.createItem(true)
	_any.anyOfItemPlayerZoneItem.setItem(item.ID())
	return item
}

func (_any anyOfItemPlayerZoneItemCore) setItem(itemID ItemID) {
	any := _any.engine.anyOfItemPlayerZoneItem(_any.ID).anyOfItemPlayerZoneItem
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
	any.engine.Patch.AnyOfItemPlayerZoneItem[any.ID] = any
}

func (_any anyOfItemPlayerZoneItemCore) deleteChild() {
	any := _any.engine.anyOfItemPlayerZoneItem(_any.ID).anyOfItemPlayerZoneItem
	switch any.ElementKind {
	case ElementKindItem:
		any.engine.deleteItem(any.Item)
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindZoneItem:
		any.engine.deleteZoneItem(any.ZoneItem)
	}
}
