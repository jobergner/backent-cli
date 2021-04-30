package state

func (_zone zone) AddPlayer() player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return player{player: playerCore{OperationKind_: OperationKindDelete}}
	}
	player := zone.zone.engine.createPlayer(true)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone zone) AddItem() zoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return zoneItem{zoneItem: zoneItemCore{OperationKind_: OperationKindDelete}}
	}
	zoneItem := zone.zone.engine.createZoneItem(true)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind_ = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone zone) AddTags(tags ...string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
		return
	}
	zone.zone.Tags = append(zone.zone.Tags, tags...)
	zone.zone.OperationKind_ = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
}

func (_player player) AddItem() item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return item{item: itemCore{OperationKind_: OperationKindDelete}}
	}
	item := player.player.engine.createItem(true)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind_ = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}

func (_player player) AddGuildMember(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind_ == OperationKindDelete {
		return
	}
	ref := player.player.engine.createPlayerGuildMemberRef(playerID, player.player.ID)
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddTargetedByPlayer(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind_ == OperationKindDelete {
		return
	}
	anyContainer := player.player.engine.createAnyOfPlayerZoneItem(false).anyOfPlayerZoneItem
	anyContainer.setPlayer(playerID)
	ref := player.player.engine.createPlayerTargetedByRef(anyContainer.ID, player.player.ID)
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddTargetedByZoneItem(zoneItemID ZoneItemID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind_ == OperationKindDelete {
		return
	}
	anyContainer := player.player.engine.createAnyOfPlayerZoneItem(false).anyOfPlayerZoneItem
	anyContainer.setZoneItem(zoneItemID)
	ref := player.player.engine.createPlayerTargetedByRef(anyContainer.ID, player.player.ID)
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player player) AddEquipmentSet(equipmentSetID EquipmentSetID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return
	}
	if player.player.engine.EquipmentSet(equipmentSetID).equipmentSet.OperationKind_ == OperationKindDelete {
		return
	}
	ref := player.player.engine.createPlayerEquipmentSetRef(equipmentSetID, player.player.ID)
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind_ = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_equipmentSet equipmentSet) AddEquipment(itemID ItemID) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind_ == OperationKindDelete {
		return
	}
	if equipmentSet.equipmentSet.engine.Item(itemID).item.OperationKind_ == OperationKindDelete {
		return
	}
	ref := equipmentSet.equipmentSet.engine.createEquipmentSetEquipmentRef(itemID, equipmentSet.equipmentSet.ID)
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind_ = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}
