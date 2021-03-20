package state

func (_e Zone) AddPlayer(sm *Engine) Player {
	e := sm.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return Player{playerCore{OperationKind: OperationKindDelete}}
	}
	player := sm.createPlayer(true)
	e.zone.Players = append(e.zone.Players, player.player.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return player
}

func (_e Zone) AddItem(sm *Engine) ZoneItem {
	e := sm.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItemCore{OperationKind: OperationKindDelete}}
	}
	zoneItem := sm.createZoneItem(true)
	e.zone.Items = append(e.zone.Items, zoneItem.zoneItem.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return zoneItem
}

func (_e Zone) AddTags(sm *Engine, tags ...string) {
	e := sm.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return
	}
	e.zone.Tags = append(e.zone.Tags, tags...)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
}

func (_e Player) AddItem(sm *Engine) Item {
	e := sm.Player(_e.player.ID)
	if e.player.OperationKind == OperationKindDelete {
		return Item{itemCore{OperationKind: OperationKindDelete}}
	}
	item := sm.createItem(true)
	e.player.Items = append(e.player.Items, item.item.ID)
	e.player.OperationKind = OperationKindUpdate
	sm.Patch.Player[e.player.ID] = e.player
	return item
}
