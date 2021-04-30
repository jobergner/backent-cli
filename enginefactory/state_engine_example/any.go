package state

func (_any anyOfPlayerZoneItem) Kind() ElementKind {
	any := _any.anyOfPlayerZoneItem.engine.anyOfPlayerZoneItem(_any.anyOfPlayerZoneItem.ID)
	return any.anyOfPlayerZoneItem.ElementKind
}

func (_any anyOfPlayerZoneItemCore) setZoneItem(zoneItemID ZoneItemID) {
	any := _any.engine.anyOfPlayerZoneItem(_any.ID).anyOfPlayerZoneItem
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
	}
	any.ElementKind = ElementKindZoneItem
	any.ZoneItem = zoneItemID
	any.engine.Patch.AnyOfPlayerZoneItem[any.ID] = any
}

func (_any anyOfPlayerZoneItemCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayerZoneItem(_any.ID).anyOfPlayerZoneItem
	if any.ZoneItem != 0 {
		any.engine.deleteZoneItem(any.ZoneItem)
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayerZoneItem[any.ID] = any
}

func (_any anyOfPlayerZone) Kind() ElementKind {
	any := _any.anyOfPlayerZone.engine.anyOfPlayerZone(_any.anyOfPlayerZone.ID)
	return any.anyOfPlayerZone.ElementKind
}

func (_any anyOfPlayerZoneCore) setZone(zoneItemID ZoneID) {
	any := _any.engine.anyOfPlayerZone(_any.ID).anyOfPlayerZone
	if any.Player != 0 {
		any.engine.deletePlayer(any.Player)
	}
	any.ElementKind = ElementKindZone
	any.Zone = zoneItemID
	any.engine.Patch.AnyOfPlayerZone[any.ID] = any
}

func (_any anyOfPlayerZoneCore) setPlayer(playerID PlayerID) {
	any := _any.engine.anyOfPlayerZone(_any.ID).anyOfPlayerZone
	if any.Zone != 0 {
		any.engine.deleteZone(any.Zone)
	}
	any.ElementKind = ElementKindPlayer
	any.Player = playerID
	any.engine.Patch.AnyOfPlayerZone[any.ID] = any
}
