package statemachine

func (_e Zone) AddPlayer(sm *StateMachine) Player {
	e := sm.GetZone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return Player{}
	}
	player := sm.CreatePlayer(ParentInfo{EntityKindZone, int(e.zone.ID)})
	e.zone.Players = append(e.zone.Players, player.player.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return player
}

func (_e Zone) AddZoneItem(sm *StateMachine) ZoneItem {
	e := sm.GetZone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return ZoneItem{}
	}
	zoneItem := sm.CreateZoneItem(ParentInfo{EntityKindZone, int(e.zone.ID)})
	e.zone.Items = append(e.zone.Items, zoneItem.zoneItem.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return zoneItem
}

func (_e Player) AddItem(sm *StateMachine) Item {
	e := sm.GetPlayer(_e.player.ID)
	if e.player.OperationKind == OperationKindDelete {
		return Item{}
	}
	item := sm.CreateItem(append(e.player.Parentage, ParentInfo{EntityKindPlayer, int(e.player.ID)})...)
	e.player.Items = append(e.player.Items, item.item.ID)
	e.player.OperationKind = OperationKindUpdate
	sm.Patch.Player[e.player.ID] = e.player
	return item
}
