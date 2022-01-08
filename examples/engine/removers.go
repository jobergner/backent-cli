package state

func (_zone Zone) RemovePlayers(playersToRemove ...PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for _, playerID := range playersToRemove {
		_, ok := zone.zone.Players[playerID]
		if !ok {
			continue
		}
		delete(zone.zone.Players, playerID)
		zone.zone.engine.deletePlayer(playerID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_zone Zone) RemoveItems(itemsToRemove ...ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for _, zoneItemID := range itemsToRemove {
		_, ok := zone.zone.Items[zoneItemID]
		if !ok {
			continue
		}
		delete(zone.zone.Items, zoneItemID)
		zone.zone.engine.deleteZoneItem(zoneItemID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_zone Zone) RemoveInteractablesItem(itemsToRemove ...ItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[ItemID]AnyOfItem_Player_ZoneItemID)
	for wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindItem {
			continue
		}
		wrappers[wrapper.Item().ID()] = wrapperID
	}

	for _, itemID := range itemsToRemove {
		wrapperID, ok := wrappers[itemID]
		if !ok {
			continue
		}
		delete(zone.zone.Interactables, wrapperID)
		zone.zone.engine.deleteItem(itemID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_zone Zone) RemoveInteractablesPlayer(itemsToRemove ...PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[PlayerID]AnyOfItem_Player_ZoneItemID)
	for wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindPlayer {
			continue
		}
		wrappers[wrapper.Player().ID()] = wrapperID
	}

	for _, playerID := range itemsToRemove {
		wrapperID, ok := wrappers[playerID]
		if !ok {
			continue
		}
		delete(zone.zone.Interactables, wrapperID)
		zone.zone.engine.deletePlayer(playerID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_zone Zone) RemoveInteractablesZoneItem(itemsToRemove ...ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[ZoneItemID]AnyOfItem_Player_ZoneItemID)
	for wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindZoneItem {
			continue
		}
		wrappers[wrapper.ZoneItem().ID()] = wrapperID
	}

	for _, zoneItemID := range itemsToRemove {
		wrapperID, ok := wrappers[zoneItemID]
		if !ok {
			continue
		}
		delete(zone.zone.Interactables, wrapperID)
		zone.zone.engine.deleteZoneItem(zoneItemID)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_player Player) RemoveItems(itemsToRemove ...ItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	for _, itemID := range itemsToRemove {
		_, ok := player.player.Items[itemID]
		if !ok {
			continue
		}
		delete(player.player.Items, itemID)
		player.player.engine.deleteItem(itemID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
	}

	return player
}

func (_player Player) RemoveEquipmentSets(equipmentSetsToRemove ...EquipmentSetID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	wrappers := make(map[EquipmentSetID]PlayerEquipmentSetRefID)
	for wrapperID := range player.player.EquipmentSets {
		wrapper := player.player.engine.playerEquipmentSetRef(wrapperID)
		wrappers[wrapper.playerEquipmentSetRef.ReferencedElementID] = wrapperID
	}

	for _, equipmentSetID := range equipmentSetsToRemove {
		wrapperID, ok := wrappers[equipmentSetID]
		if !ok {
			continue
		}
		delete(player.player.EquipmentSets, wrapperID)
		player.player.engine.deletePlayerEquipmentSetRef(wrapperID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
	}

	return player
}

func (_player Player) RemoveGuildMembers(guildMembersToRemove ...PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	wrappers := make(map[PlayerID]PlayerGuildMemberRefID)
	for wrapperID := range player.player.GuildMembers {
		wrapper := player.player.engine.playerGuildMemberRef(wrapperID)
		wrappers[wrapper.playerGuildMemberRef.ReferencedElementID] = wrapperID
	}

	for _, playerID := range guildMembersToRemove {
		wrapperID, ok := wrappers[playerID]
		if !ok {
			continue
		}
		delete(player.player.GuildMembers, wrapperID)
		player.player.engine.deletePlayerGuildMemberRef(wrapperID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
	}

	return player
}

func (_player Player) RemoveTargetedByZoneItem(zoneItemsToRemove ...ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	refs := make(map[AnyOfPlayer_ZoneItemID]PlayerTargetedByRefID)
	for refID := range player.player.TargetedBy {
		ref := player.player.engine.playerTargetedByRef(refID)
		refs[ref.playerTargetedByRef.ReferencedElementID] = refID
	}

	wrappers := make(map[ZoneItemID]PlayerTargetedByRefID)
	for wrapperID, refID := range refs {
		wrapper := player.player.engine.anyOfPlayer_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindZoneItem {
			continue
		}
		wrappers[wrapper.ZoneItem().ID()] = refID
	}

	for _, zoneItemID := range zoneItemsToRemove {
		wrapperID, ok := wrappers[zoneItemID]
		if !ok {
			continue
		}
		delete(player.player.TargetedBy, wrapperID)
		player.player.engine.deletePlayerTargetedByRef(wrapperID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
	}

	return player
}

func (_player Player) RemoveTargetedByPlayer(playersToRemove ...PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	refs := make(map[AnyOfPlayer_ZoneItemID]PlayerTargetedByRefID)
	for refID := range player.player.TargetedBy {
		ref := player.player.engine.playerTargetedByRef(refID)
		refs[ref.playerTargetedByRef.ReferencedElementID] = refID
	}

	wrappers := make(map[PlayerID]PlayerTargetedByRefID)
	for wrapperID, refID := range refs {
		wrapper := player.player.engine.anyOfPlayer_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindPlayer {
			continue
		}
		wrappers[wrapper.Player().ID()] = refID
	}

	for _, playerID := range playersToRemove {
		wrapperID, ok := wrappers[playerID]
		if !ok {
			continue
		}
		delete(player.player.TargetedBy, wrapperID)
		player.player.engine.deletePlayerTargetedByRef(wrapperID)
		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player
	}

	return player
}

func (_zone Zone) RemoveTags(tagsToRemove ...string) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for _, tag := range tagsToRemove {
		_, ok := zone.zone.Tags[tag]
		if !ok {
			continue
		}
		delete(zone.zone.Tags, tag)
		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	}

	return zone
}

func (_equipmentSet EquipmentSet) RemoveEquipment(equipmentToRemove ...ItemID) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}

	wrappers := make(map[ItemID]EquipmentSetEquipmentRefID)
	for wrapperID := range equipmentSet.equipmentSet.Equipment {
		wrapper := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(wrapperID)
		wrappers[wrapper.equipmentSetEquipmentRef.ReferencedElementID] = wrapperID
	}

	for _, itemID := range equipmentToRemove {
		wrapperID, ok := wrappers[itemID]
		if !ok {
			continue
		}
		delete(equipmentSet.equipmentSet.Equipment, wrapperID)
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(wrapperID)
		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	}

	return equipmentSet
}
