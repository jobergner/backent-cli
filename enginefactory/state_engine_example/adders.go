package state

func (_e Zone) AddPlayer(se *Engine) Player {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return Player{playerCore{OperationKind: OperationKindDelete}}
	}
	player := se.createPlayer(true)
	e.zone.Players = append(e.zone.Players, player.player.ID)
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
	return player
}

func (_e Zone) AddItem(se *Engine) ZoneItem {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItemCore{OperationKind: OperationKindDelete}}
	}
	zoneItem := se.createZoneItem(true)
	e.zone.Items = append(e.zone.Items, zoneItem.zoneItem.ID)
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
	return zoneItem
}

func (_e Zone) AddTags(se *Engine, tags ...string) {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return
	}
	e.zone.Tags = append(e.zone.Tags, tags...)
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
}

func (_e Player) AddItem(se *Engine) Item {
	e := se.Player(_e.player.ID)
	if e.player.OperationKind == OperationKindDelete {
		return Item{itemCore{OperationKind: OperationKindDelete}}
	}
	item := se.createItem(true)
	e.player.Items = append(e.player.Items, item.item.ID)
	e.player.OperationKind = OperationKindUpdate
	se.Patch.Player[e.player.ID] = e.player
	return item
}
