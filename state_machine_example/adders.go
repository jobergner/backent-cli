package statemachine

func (e Zone) AddPlayer(sm *StateMachine) Player {
	player := sm.CreatePlayer(ParentInfo{EntityKindZone, int(e.ID)})
	e.Players = append(e.Players, player.ID)
	e.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.ID] = e
	return player
}

func (e Zone) AddZoneItem(sm *StateMachine) ZoneItem {
	zoneItem := sm.CreateZoneItem(ParentInfo{EntityKindZone, int(e.ID)})
	e.Items = append(e.Items, zoneItem.ID)
	e.OperationKind = OperationKindUpdate
	sm.Patch.Zone[e.ID] = e
	return zoneItem
}

func (e Player) AddItem(sm *StateMachine) Item {
	item := sm.CreateItem(append(e.Parentage, ParentInfo{EntityKindPlayer, int(e.ID)})...)
	e.Items = append(e.Items, item.ID)
	e.OperationKind = OperationKindUpdate
	sm.Patch.Player[e.ID] = e
	return item
}
