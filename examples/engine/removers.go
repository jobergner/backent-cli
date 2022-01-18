package state

func (_zone Zone) RemovePlayers(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for i, playerID := range zone.zone.Players {
		if playerID != playerToRemove {
			continue
		}

		zone.zone.Players[i] = zone.zone.Players[len(zone.zone.Players)-1]
		zone.zone.Players[len(zone.zone.Players)-1] = 0
		zone.zone.Players = zone.zone.Players[:len(zone.zone.Players)-1]
		zone.zone.engine.deletePlayer(playerID)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveItems(itemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for i, zoneItemID := range zone.zone.Items {
		if zoneItemID != itemToRemove {
			continue
		}

		zone.zone.Items[i] = zone.zone.Items[len(zone.zone.Items)-1]
		zone.zone.Items[len(zone.zone.Items)-1] = 0
		zone.zone.Items = zone.zone.Items[:len(zone.zone.Items)-1]
		zone.zone.engine.deleteZoneItem(zoneItemID)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveInteractablesItem(itemToRemove ItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[AnyOfItem_Player_ZoneItemID]ItemID)
	for _, wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindItem {
			continue
		}
		wrappers[wrapperID] = wrapper.Item().ID()
	}

	for i, wrapperID := range zone.zone.Interactables {
		itemID := wrappers[wrapperID]
		if itemID != itemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(wrapperID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveInteractablesPlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[AnyOfItem_Player_ZoneItemID]PlayerID)
	for _, wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindPlayer {
			continue
		}
		wrappers[wrapperID] = wrapper.Player().ID()
	}

	for i, wrapperID := range zone.zone.Interactables {
		playerID := wrappers[wrapperID]
		if playerID != playerToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(wrapperID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveInteractablesZoneItem(zoneItemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	wrappers := make(map[AnyOfItem_Player_ZoneItemID]ZoneItemID)
	for _, wrapperID := range zone.zone.Interactables {
		wrapper := zone.zone.engine.anyOfItem_Player_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindZoneItem {
			continue
		}
		wrappers[wrapperID] = wrapper.ZoneItem().ID()
	}

	for i, wrapperID := range zone.zone.Interactables {
		zoneItemID := wrappers[wrapperID]
		if zoneItemID != zoneItemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(wrapperID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_player Player) RemoveItems(itemToRemove ItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	for i, itemID := range player.player.Items {
		if itemID != itemToRemove {
			continue
		}

		player.player.Items[i] = player.player.Items[len(player.player.Items)-1]
		player.player.Items[len(player.player.Items)-1] = 0
		player.player.Items = player.player.Items[:len(player.player.Items)-1]
		player.player.engine.deleteItem(itemID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveEquipmentSets(equipmentSetToRemove EquipmentSetID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	wrappers := make(map[PlayerEquipmentSetRefID]EquipmentSetID)
	for _, wrapperID := range player.player.EquipmentSets {
		wrapper := player.player.engine.playerEquipmentSetRef(wrapperID)
		wrappers[wrapperID] = wrapper.playerEquipmentSetRef.ReferencedElementID
	}

	for i, wrapperID := range player.player.EquipmentSets {
		equipmentSetID := wrappers[wrapperID]
		if equipmentSetID != equipmentSetToRemove {
			continue
		}

		player.player.EquipmentSets[i] = player.player.EquipmentSets[len(player.player.EquipmentSets)-1]
		player.player.EquipmentSets[len(player.player.EquipmentSets)-1] = 0
		player.player.EquipmentSets = player.player.EquipmentSets[:len(player.player.EquipmentSets)-1]
		player.player.engine.deletePlayerEquipmentSetRef(wrapperID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveGuildMembers(guildMemberToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	wrappers := make(map[PlayerGuildMemberRefID]PlayerID)
	for _, wrapperID := range player.player.GuildMembers {
		wrapper := player.player.engine.playerGuildMemberRef(wrapperID)
		wrappers[wrapperID] = wrapper.playerGuildMemberRef.ReferencedElementID
	}

	for i, wrapperID := range player.player.GuildMembers {
		playerID := wrappers[wrapperID]
		if playerID != guildMemberToRemove {
			continue
		}

		player.player.GuildMembers[i] = player.player.GuildMembers[len(player.player.GuildMembers)-1]
		player.player.GuildMembers[len(player.player.GuildMembers)-1] = 0
		player.player.GuildMembers = player.player.GuildMembers[:len(player.player.GuildMembers)-1]
		player.player.engine.deletePlayerGuildMemberRef(wrapperID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveTargetedByZoneItem(zoneItemToRemove ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	refs := make(map[PlayerTargetedByRefID]AnyOfPlayer_ZoneItemID)
	for _, refID := range player.player.TargetedBy {
		ref := player.player.engine.playerTargetedByRef(refID)
		refs[refID] = ref.playerTargetedByRef.ReferencedElementID
	}

	wrappers := make(map[PlayerTargetedByRefID]ZoneItemID)
	for refID, wrapperID := range refs {
		wrapper := player.player.engine.anyOfPlayer_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindZoneItem {
			continue
		}
		wrappers[refID] = wrapper.ZoneItem().ID()
	}

	for i, wrapperID := range player.player.TargetedBy {
		zoneItemID := wrappers[wrapperID]
		if zoneItemID != zoneItemToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(wrapperID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveTargetedByPlayer(playerToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	refs := make(map[PlayerTargetedByRefID]AnyOfPlayer_ZoneItemID)
	for _, refID := range player.player.TargetedBy {
		ref := player.player.engine.playerTargetedByRef(refID)
		refs[refID] = ref.playerTargetedByRef.ReferencedElementID
	}

	wrappers := make(map[PlayerTargetedByRefID]PlayerID)
	for refID, wrapperID := range refs {
		wrapper := player.player.engine.anyOfPlayer_ZoneItem(wrapperID)
		if wrapper.Kind() != ElementKindPlayer {
			continue
		}
		wrappers[refID] = wrapper.Player().ID()
	}

	for i, wrapperID := range player.player.TargetedBy {
		playerID := wrappers[wrapperID]
		if playerID != playerToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(wrapperID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_zone Zone) RemoveTags(tagToRemove string) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	for i, val := range zone.zone.Tags {
		if val != tagToRemove {
			continue
		}

		zone.zone.Tags[i] = zone.zone.Tags[len(zone.zone.Tags)-1]
		zone.zone.Tags[len(zone.zone.Tags)-1] = ""
		zone.zone.Tags = zone.zone.Tags[:len(zone.zone.Tags)-1]

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_equipmentSet EquipmentSet) RemoveEquipment(equipmentToRemove ItemID) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}

	wrappers := make(map[EquipmentSetEquipmentRefID]ItemID)
	for _, wrapperID := range equipmentSet.equipmentSet.Equipment {
		wrapper := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(wrapperID)
		wrappers[wrapperID] = wrapper.equipmentSetEquipmentRef.ReferencedElementID
	}

	for i, wrapperID := range equipmentSet.equipmentSet.Equipment {
		itemID := wrappers[wrapperID]
		if itemID != equipmentToRemove {
			continue
		}

		equipmentSet.equipmentSet.Equipment[i] = equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1] = 0
		equipmentSet.equipmentSet.Equipment = equipmentSet.equipmentSet.Equipment[:len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(wrapperID)

		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet

		break
	}

	return equipmentSet
}
