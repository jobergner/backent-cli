package state

func (_any anyOfPlayerZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayerZoneItem.engine.anyOfPlayerZoneItem(_any.anyOfPlayerZoneItem.ID)
	return any.anyOfPlayerZoneItem.ElementKind
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

func (_any anyOfPlayerZone) Kind() ElementKind {
	any := _any.anyOfPlayerZone.engine.anyOfPlayerZone(_any.anyOfPlayerZone.ID)
	return any.anyOfPlayerZone.ElementKind
}

func (_any anyOfPlayerZoneCore) setZone(zoneItemID ZoneID) {
	any := _any.engine.anyOfPlayerZone(_any.ID).anyOfPlayerZone
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
		any.Player = 0
	}
	any.ElementKind = ElementKindZone
	any.Zone = zoneItemID
	any.engine.Patch.AnyOfPlayerZone[any.ID] = any
}

func (_any anyOfPlayerZoneCore) deleteChild() {
	any := _any.engine.anyOfPlayerZone(_any.ID).anyOfPlayerZone
	switch any.ElementKind {
	case ElementKindPlayer:
		any.engine.deletePlayer(any.Player)
	case ElementKindZone:
		any.engine.deleteZone(any.Zone)
	}
}

func (_any anyOfPlayerZoneCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayerZone(_any.ID).anyOfPlayerZone
	if any.Zone != 0 {
		any.engine.deleteZone(any.Zone)
		any.Zone = 0
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayerZone[any.ID] = any
}

func (_any anyOfItemPlayerZoneItem) Kind() ElementKind {
	any := _any.anyOfItemPlayerZoneItem.engine.anyOfItemPlayerZoneItem(_any.anyOfItemPlayerZoneItem.ID)
	return any.anyOfItemPlayerZoneItem.ElementKind
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
