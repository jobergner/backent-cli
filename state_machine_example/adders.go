package statemachine

func (e Zone) AddPlayer(sm *StateMachine) Player {
	player := sm.CreatePlayer(ParentInfo{EntityKindZone, int(e.zone.ID)})
	e.zone.Players = append(e.zone.Players, player.player.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return player
}

func (e Zone) AddZoneItem(sm *StateMachine) ZoneItem {
	zoneItem := sm.CreateZoneItem(ParentInfo{EntityKindZone, int(e.zone.ID)})
	e.zone.Items = append(e.zone.Items, zoneItem.zoneItem.ID)
	e.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.zone.ID] = e.zone
	return zoneItem
}

func (e Player) AddItem(sm *StateMachine) Item {
	item := sm.CreateItem(append(e.player.Parentage, ParentInfo{EntityKindPlayer, int(e.player.ID)})...)
	e.player.Items = append(e.player.Items, item.item.ID)
	e.player.OperationKind = OperationKindUpdate
	sm.Patch.Player[e.player.ID] = e.player
	return item
}
