package state

func (_zone Zone) AddPlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]PlayerID, len(zone.zone.Players))
		copy(cp, zone.zone.Players)
		zone.zone.Players = cp
	}

	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_playersIdentifier)
	zone.zone.Players = append(zone.zone.Players, player.player.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone Zone) AddItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]ZoneItemID, len(zone.zone.Items))
		copy(cp, zone.zone.Items)
		zone.zone.Items = cp
	}

	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_itemsIdentifier)
	zone.zone.Items = append(zone.zone.Items, zoneItem.zoneItem.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone Zone) AddInteractablePlayer() Player {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Player{player: playerCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	player := zone.zone.engine.createPlayer(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(player.player.ID), ElementKindPlayer, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return player
}

func (_zone Zone) AddInteractableZoneItem() ZoneItem {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return ZoneItem{zoneItem: zoneItemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	zoneItem := zone.zone.engine.createZoneItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(zoneItem.zoneItem.ID), ElementKindZoneItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zoneItem
}

func (_zone Zone) AddInteractableItem() Item {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: zone.zone.engine}}
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	item := zone.zone.engine.createItem(zone.zone.Path, zone_interactablesIdentifier)
	anyContainer := zone.zone.engine.createAnyOfItem_Player_ZoneItem(int(zone.zone.ID), int(item.item.ID), ElementKindItem, zone.zone.Path, zone_interactablesIdentifier).anyOfItem_Player_ZoneItem
	zone.zone.Interactables = append(zone.zone.Interactables, anyContainer.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return item
}

func (_zone Zone) AddTag(tag string) {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]StringValueID, len(zone.zone.Tags))
		copy(cp, zone.zone.Tags)
		zone.zone.Tags = cp
	}

	tagValue := zone.zone.engine.createStringValue(zone.zone.Path, zone_tagsIdentifier, tag)
	zone.zone.Tags = append(zone.zone.Tags, tagValue.ID)
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
}

func (_player Player) AddItem() Item {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return Item{item: itemCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]ItemID, len(player.player.Items))
		copy(cp, player.player.Items)
		player.player.Items = cp
	}

	item := player.player.engine.createItem(player.player.Path, player_itemsIdentifier)
	player.player.Items = append(player.player.Items, item.item.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return item
}

func (_player Player) AddAction() AttackEvent {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return AttackEvent{attackEvent: attackEventCore{OperationKind: OperationKindDelete, engine: player.player.engine}}
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]AttackEventID, len(player.player.Action))
		copy(cp, player.player.Action)
		player.player.Action = cp
	}

	attackEvent := player.player.engine.createAttackEvent(player.player.Path, player_actionIdentifier)
	player.player.Action = append(player.player.Action, attackEvent.attackEvent.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return attackEvent
}

func (_player Player) AddGuildMember(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerGuildMemberRefID, len(player.player.GuildMembers))
		copy(cp, player.player.GuildMembers)
		player.player.GuildMembers = cp
	}

	for _, currentRefID := range player.player.GuildMembers {
		currentRef := player.player.engine.playerGuildMemberRef(currentRefID)
		if currentRef.playerGuildMemberRef.ReferencedElementID == playerID {
			return
		}
	}
	ref := player.player.engine.createPlayerGuildMemberRef(player.player.Path, player_guildMembersIdentifier, playerID, player.player.ID, int(playerID))
	player.player.GuildMembers = append(player.player.GuildMembers, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player Player) AddTargetedByPlayer(playerID PlayerID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}

	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if PlayerID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == playerID {
			return
		}
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem
	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player Player) AddTargetedByZoneItem(zoneItemID ZoneItemID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}

	for _, currentRefID := range player.player.TargetedBy {
		currentRef := player.player.engine.playerTargetedByRef(currentRefID)
		anyContainer := player.player.engine.anyOfPlayer_ZoneItem(currentRef.playerTargetedByRef.ReferencedElementID)
		if ZoneItemID(anyContainer.anyOfPlayer_ZoneItem.ChildID) == zoneItemID {
			return
		}
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetedByIdentifier).anyOfPlayer_ZoneItem
	ref := player.player.engine.createPlayerTargetedByRef(player.player.Path, player_targetedByIdentifier, anyContainer.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.TargetedBy = append(player.player.TargetedBy, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_player Player) AddEquipmentSet(equipmentSetID EquipmentSetID) {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return
	}
	if player.player.engine.EquipmentSet(equipmentSetID).equipmentSet.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerEquipmentSetRefID, len(player.player.EquipmentSets))
		copy(cp, player.player.EquipmentSets)
		player.player.EquipmentSets = cp
	}

	for _, currentRefID := range player.player.EquipmentSets {
		currentRef := player.player.engine.playerEquipmentSetRef(currentRefID)
		if currentRef.playerEquipmentSetRef.ReferencedElementID == equipmentSetID {
			return
		}
	}
	ref := player.player.engine.createPlayerEquipmentSetRef(player.player.Path, player_equipmentSetsIdentifier, equipmentSetID, player.player.ID, int(equipmentSetID))
	player.player.EquipmentSets = append(player.player.EquipmentSets, ref.ID)
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
}

func (_equipmentSet EquipmentSet) AddEquipment(itemID ItemID) {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return
	}
	if equipmentSet.equipmentSet.engine.Item(itemID).item.OperationKind == OperationKindDelete {
		return
	}

	if _, ok := equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID]; !ok {
		cp := make([]EquipmentSetEquipmentRefID, len(equipmentSet.equipmentSet.Equipment))
		copy(cp, equipmentSet.equipmentSet.Equipment)
		equipmentSet.equipmentSet.Equipment = cp
	}

	for _, currentRefID := range equipmentSet.equipmentSet.Equipment {
		currentRef := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(currentRefID)
		if currentRef.equipmentSetEquipmentRef.ReferencedElementID == itemID {
			return
		}
	}
	ref := equipmentSet.equipmentSet.engine.createEquipmentSetEquipmentRef(equipmentSet.equipmentSet.Path, equipmentSet_equipmentIdentifier, itemID, equipmentSet.equipmentSet.ID, int(itemID))
	equipmentSet.equipmentSet.Equipment = append(equipmentSet.equipmentSet.Equipment, ref.ID)
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
}
