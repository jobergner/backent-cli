package state

type anyOfPlayer_PositionRef struct {
	anyOfPlayer_PositionWrapper anyOfPlayer_Position
	anyOfPlayer_Position        anyOfPlayer_PositionCore
}

func (_any anyOfPlayer_PositionRef) Kind() ElementKind {
	return _any.anyOfPlayer_PositionWrapper.Kind()
}

func (_any anyOfPlayer_PositionRef) Player() player {
	return _any.anyOfPlayer_PositionWrapper.Player()
}

func (_any anyOfPlayer_PositionRef) Position() position {
	return _any.anyOfPlayer_PositionWrapper.Position()
}

type anyOfPlayer_ZoneItemRef struct {
	anyOfPlayer_ZoneItemWrapper anyOfPlayer_ZoneItem
	anyOfPlayer_ZoneItem        anyOfPlayer_ZoneItemCore
}

func (_any anyOfPlayer_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfPlayer_ZoneItemWrapper.Kind()
}

func (_any anyOfPlayer_ZoneItemRef) Player() player {
	return _any.anyOfPlayer_ZoneItemWrapper.Player()
}

func (_any anyOfPlayer_ZoneItemRef) ZoneItem() zoneItem {
	return _any.anyOfPlayer_ZoneItemWrapper.ZoneItem()
}

type anyOfItem_Player_ZoneItemRef struct {
	anyOfItem_Player_ZoneItemWrapper anyOfItem_Player_ZoneItem
	anyOfItem_Player_ZoneItem        anyOfItem_Player_ZoneItemCore
}

func (_any anyOfItem_Player_ZoneItemRef) Kind() ElementKind {
	return _any.anyOfItem_Player_ZoneItemWrapper.Kind()
}

func (_any anyOfItem_Player_ZoneItemRef) Item() item {
	return _any.anyOfItem_Player_ZoneItemWrapper.Item()
}

func (_any anyOfItem_Player_ZoneItemRef) Player() player {
	return _any.anyOfItem_Player_ZoneItemWrapper.Player()
}

func (_any anyOfItem_Player_ZoneItemRef) ZoneItem() zoneItem {
	return _any.anyOfItem_Player_ZoneItemWrapper.ZoneItem()
}

func (_any anyOfPlayer_ZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	return any.anyOfPlayer_ZoneItem.ElementKind
}

func (_any anyOfPlayer_ZoneItem) SetZoneItem() zoneItem {
	if _any.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem {
		return _any.ZoneItem()
	}
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
	if _any.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer {
		return _any.Player()
	}
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
	if _any.anyOfPlayer_Position.ElementKind == ElementKindPosition {
		return _any.Position()
	}
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
	if _any.anyOfPlayer_Position.ElementKind == ElementKindPlayer {
		return _any.Player()
	}
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
	if _any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindZoneItem {
		return _any.ZoneItem()
	}
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
	if _any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindPlayer {
		return _any.Player()
	}
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
	if _any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindItem {
		return _any.Item()
	}
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
