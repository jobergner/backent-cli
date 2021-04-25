package state

func (_zone zone) AddPlayer(se *Engine) player {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return player{player: playerCore{OperationKind_: OperationKindDelete}}
	}
	player := se.createPlayer(true)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone zone) AddItem(se *Engine) zoneItem {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
	}
	zoneItem := se.createZoneItem(true)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone zone) AddTags(se *Engine, tags ...string) {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return
	}
	zone.zone.Tags = append(zone.zone.Tags, tags...)
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
}

func (_player player) AddItem(se *Engine) item {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return item{item: itemCore{OperationKind_: OperationKindDelete}}
	}
	item := se.createItem(true)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
	return item
}

func (_player player) AddGuildMember(se *Engine, playerID PlayerID) {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if se.Player(playerID).player.OperationKind_ == OperationKindDelete {
		return
	}
	ref := se.createPlayerGuildMemberRef(playerID, player.player.ID)
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddEquipmentSet(se *Engine, equipmentSetID EquipmentSetID) {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if se.EquipmentSet(equipmentSetID).equipmentSet.OperationKind_ == OperationKindDelete {
		return
	}
	ref := se.createPlayerEquipmentSetRef(equipmentSetID, player.player.ID)
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
}

func (_equipmentSet equipmentSet) AddEquipment(se *Engine, itemID ItemID) {
	equipmentSet := se.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind_ == OperationKindDelete {
		return
	}
	if se.Item(itemID).item.OperationKind_ == OperationKindDelete {
		return
	}
	ref := se.createEquipmentSetEquipmentRef(itemID, equipmentSet.equipmentSet.ID)
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind_ = OperationKindUpdate
	se.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}
