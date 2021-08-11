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
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfPlayer_ZoneItem.engine.createZoneItem(any.anyOfPlayer_ZoneItem.ChildElementPath, false)
	any.anyOfPlayer_ZoneItem.setZoneItem(zoneItem.ID(), true)
	return zoneItem
}

func (_any anyOfPlayer_ZoneItemCore) setZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	if deleteCurrentChild {
		if any.Player != 0 {
			any.engine.deletePlayer(any.Player)
			any.Player = 0
		}
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfPlayer_ZoneItem[any.ID] = any
}

func (_any anyOfPlayer_ZoneItem) SetPlayer() player {
	any := _any.anyOfPlayer_ZoneItem.engine.anyOfPlayer_ZoneItem(_any.anyOfPlayer_ZoneItem.ID)
	if any.anyOfPlayer_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_ZoneItem.engine.createPlayer(any.anyOfPlayer_ZoneItem.ChildElementPath, false)
	any.anyOfPlayer_ZoneItem.setPlayer(player.ID(), true)
	return player
}

func (_any anyOfPlayer_ZoneItemCore) setPlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_ZoneItem(_any.ID).anyOfPlayer_ZoneItem
	if deleteCurrentChild {
		if any.ZoneItem != 0 {
			any.engine.deleteZoneItem(any.ZoneItem)
			any.ZoneItem = 0
		}
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
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPosition || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Position()
	}
	position := any.anyOfPlayer_Position.engine.createPosition(any.anyOfPlayer_Position.ChildElementPath, false)
	any.anyOfPlayer_Position.setPosition(position.ID(), true)
	return position
}

func (_any anyOfPlayer_PositionCore) setPosition(positionID PositionID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	if deleteCurrentChild {
		if any.Player != 0 {
			any.engine.deletePlayer(any.Player)
			any.Player = 0
		}
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
	any := _any.anyOfPlayer_Position.engine.anyOfPlayer_Position(_any.anyOfPlayer_Position.ID)
	if any.anyOfPlayer_Position.ElementKind == ElementKindPlayer || any.anyOfPlayer_Position.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfPlayer_Position.engine.createPlayer(any.anyOfPlayer_Position.ChildElementPath, false)
	any.anyOfPlayer_Position.setPlayer(player.ID(), true)
	return player
}

func (_any anyOfPlayer_PositionCore) setPlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfPlayer_Position(_any.ID).anyOfPlayer_Position
	if deleteCurrentChild {
		if any.Position != 0 {
			any.engine.deletePosition(any.Position)
			any.Position = 0
		}
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
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindZoneItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.ZoneItem()
	}
	zoneItem := any.anyOfItem_Player_ZoneItem.engine.createZoneItem(any.anyOfItem_Player_ZoneItem.ChildElementPath, false)
	any.anyOfItem_Player_ZoneItem.setZoneItem(zoneItem.ID(), true)
	return zoneItem
}

func (_any anyOfItem_Player_ZoneItemCore) setZoneItem(zoneItemID ZoneItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if deleteCurrentChild {
		if any.Item != 0 {
			any.engine.deleteItem(any.Item)
			any.Item = 0
		}
		if any.Player != 0 {
			any.engine.deletePlayer(any.Player)
			any.Player = 0
		}
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItem) SetPlayer() player {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindPlayer || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Player()
	}
	player := any.anyOfItem_Player_ZoneItem.engine.createPlayer(any.anyOfItem_Player_ZoneItem.ChildElementPath, false)
	any.anyOfItem_Player_ZoneItem.setPlayer(player.ID(), true)
	return player
}

func (_any anyOfItem_Player_ZoneItemCore) setPlayer(playerID PlayerID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if deleteCurrentChild {
		if any.Item != 0 {
			any.engine.deleteItem(any.Item)
			any.Item = 0
		}
		if any.ZoneItem != 0 {
			any.engine.deleteZoneItem(any.ZoneItem)
			any.ZoneItem = 0
		}
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfItem_Player_ZoneItem[any.ID] = any
}

func (_any anyOfItem_Player_ZoneItem) SetItem() item {
	any := _any.anyOfItem_Player_ZoneItem.engine.anyOfItem_Player_ZoneItem(_any.anyOfItem_Player_ZoneItem.ID)
	if any.anyOfItem_Player_ZoneItem.ElementKind == ElementKindItem || any.anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
		return any.Item()
	}
	item := any.anyOfItem_Player_ZoneItem.engine.createItem(any.anyOfItem_Player_ZoneItem.ChildElementPath, false)
	any.anyOfItem_Player_ZoneItem.setItem(item.ID(), true)
	return item
}

func (_any anyOfItem_Player_ZoneItemCore) setItem(itemID ItemID, deleteCurrentChild bool) {
	any := _any.engine.anyOfItem_Player_ZoneItem(_any.ID).anyOfItem_Player_ZoneItem
	if deleteCurrentChild {
		if any.Player != 0 {
			any.engine.deletePlayer(any.Player)
			any.Player = 0
		}
		if any.ZoneItem != 0 {
			any.engine.deleteZoneItem(any.ZoneItem)
			any.ZoneItem = 0
		}
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
