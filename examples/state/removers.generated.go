package state

func (_zone Zone) RemovePlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]PlayerID, len(zone.zone.Players))
		copy(cp, zone.zone.Players)
		zone.zone.Players = cp
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

func (_zone Zone) RemoveItem(itemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]ZoneItemID, len(zone.zone.Items))
		copy(cp, zone.zone.Items)
		zone.zone.Items = cp
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

func (_zone Zone) RemoveInteractableItem(itemToRemove ItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; ItemID(childID) != itemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveInteractablePlayer(playerToRemove PlayerID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; PlayerID(childID) != playerToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_zone Zone) RemoveInteractableZoneItem(zoneItemToRemove ZoneItemID) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]AnyOfItem_Player_ZoneItemID, len(zone.zone.Interactables))
		copy(cp, zone.zone.Interactables)
		zone.zone.Interactables = cp
	}

	for i, id := range zone.zone.Interactables {
		if childID := zone.zone.engine.anyOfItem_Player_ZoneItem(id).anyOfItem_Player_ZoneItem.ChildID; ZoneItemID(childID) != zoneItemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = 0
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(id, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone

		break
	}

	return zone
}

func (_player Player) RemoveAction(actionToRemove AttackEventID) Player {
	player := _player.player.engine.Player(_player.player.ID)

	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]AttackEventID, len(player.player.Action))
		copy(cp, player.player.Action)
		player.player.Action = cp
	}

	for i, attackEventID := range player.player.Action {
		if attackEventID != actionToRemove {
			continue
		}

		player.player.Action[i] = player.player.Action[len(player.player.Action)-1]
		player.player.Action[len(player.player.Action)-1] = 0
		player.player.Action = player.player.Action[:len(player.player.Action)-1]
		player.player.engine.deleteAttackEvent(attackEventID)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveItem(itemToRemove ItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)

	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]ItemID, len(player.player.Items))
		copy(cp, player.player.Items)
		player.player.Items = cp
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

func (_player Player) RemoveEquipmentSet(equipmentSetToRemove EquipmentSetID) Player {
	player := _player.player.engine.Player(_player.player.ID)

	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerEquipmentSetRefID, len(player.player.EquipmentSets))
		copy(cp, player.player.EquipmentSets)
		player.player.EquipmentSets = cp
	}

	for i, id := range player.player.EquipmentSets {
		if childID := player.player.engine.playerEquipmentSetRef(id).playerEquipmentSetRef.ChildID; EquipmentSetID(childID) != equipmentSetToRemove {
			continue
		}

		player.player.EquipmentSets[i] = player.player.EquipmentSets[len(player.player.EquipmentSets)-1]
		player.player.EquipmentSets[len(player.player.EquipmentSets)-1] = 0
		player.player.EquipmentSets = player.player.EquipmentSets[:len(player.player.EquipmentSets)-1]
		player.player.engine.deletePlayerEquipmentSetRef(id)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_player Player) RemoveGuildMember(guildMemberToRemove PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)

	if player.player.OperationKind == OperationKindDelete {
		return player
	}

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerGuildMemberRefID, len(player.player.GuildMembers))
		copy(cp, player.player.GuildMembers)
		player.player.GuildMembers = cp
	}

	for i, id := range player.player.GuildMembers {
		if childID := player.player.engine.playerGuildMemberRef(id).playerGuildMemberRef.ChildID; PlayerID(childID) != guildMemberToRemove {
			continue
		}

		player.player.GuildMembers[i] = player.player.GuildMembers[len(player.player.GuildMembers)-1]
		player.player.GuildMembers[len(player.player.GuildMembers)-1] = 0
		player.player.GuildMembers = player.player.GuildMembers[:len(player.player.GuildMembers)-1]
		player.player.engine.deletePlayerGuildMemberRef(id)

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

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}

	for i, id := range player.player.TargetedBy {
		if childID := player.player.engine.playerTargetedByRef(id).playerTargetedByRef.ChildID; ZoneItemID(childID) != zoneItemToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(id)

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

	if _, ok := player.player.engine.Patch.Player[player.player.ID]; !ok {
		cp := make([]PlayerTargetedByRefID, len(player.player.TargetedBy))
		copy(cp, player.player.TargetedBy)
		player.player.TargetedBy = cp
	}

	for i, id := range player.player.TargetedBy {
		if childID := player.player.engine.playerTargetedByRef(id).playerTargetedByRef.ChildID; PlayerID(childID) != playerToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = 0
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(id)

		player.player.OperationKind = OperationKindUpdate
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
}

func (_zone Zone) RemoveTag(tagToRemove string) Zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)

	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}

	if _, ok := zone.zone.engine.Patch.Zone[zone.zone.ID]; !ok {
		cp := make([]StringValueID, len(zone.zone.Tags))
		copy(cp, zone.zone.Tags)
		zone.zone.Tags = cp
	}

	for i, valID := range zone.zone.Tags {
		if zone.zone.engine.stringValue(valID).Value != tagToRemove {
			continue
		}

		zone.zone.Tags[i] = zone.zone.Tags[len(zone.zone.Tags)-1]
		zone.zone.Tags[len(zone.zone.Tags)-1] = 0
		zone.zone.Tags = zone.zone.Tags[:len(zone.zone.Tags)-1]
		zone.zone.engine.deleteStringValue(valID)

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

	if _, ok := equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID]; !ok {
		cp := make([]EquipmentSetEquipmentRefID, len(equipmentSet.equipmentSet.Equipment))
		copy(cp, equipmentSet.equipmentSet.Equipment)
		equipmentSet.equipmentSet.Equipment = cp
	}

	for i, id := range equipmentSet.equipmentSet.Equipment {
		if childID := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(id).equipmentSetEquipmentRef.ChildID; ItemID(childID) != equipmentToRemove {
			continue
		}

		equipmentSet.equipmentSet.Equipment[i] = equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1] = 0
		equipmentSet.equipmentSet.Equipment = equipmentSet.equipmentSet.Equipment[:len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(id)

		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet

		break
	}

	return equipmentSet
}
