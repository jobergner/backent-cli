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
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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

	for i, complexID := range zone.zone.Interactables {
		if ItemID(complexID.ChildID) != itemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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

	for i, complexID := range zone.zone.Interactables {
		if PlayerID(complexID.ChildID) != playerToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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

	for i, complexID := range zone.zone.Interactables {
		if ZoneItemID(complexID.ChildID) != zoneItemToRemove {
			continue
		}

		zone.zone.Interactables[i] = zone.zone.Interactables[len(zone.zone.Interactables)-1]
		zone.zone.Interactables[len(zone.zone.Interactables)-1] = AnyOfItem_Player_ZoneItemID{}
		zone.zone.Interactables = zone.zone.Interactables[:len(zone.zone.Interactables)-1]
		zone.zone.engine.deleteAnyOfItem_Player_ZoneItem(complexID, true)

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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

	for i, attackEventID := range player.player.Action {
		if attackEventID != actionToRemove {
			continue
		}

		player.player.Action[i] = player.player.Action[len(player.player.Action)-1]
		player.player.Action[len(player.player.Action)-1] = 0
		player.player.Action = player.player.Action[:len(player.player.Action)-1]
		player.player.engine.deleteAttackEvent(attackEventID)

		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
		player.player.engine.Patch.Player[player.player.ID] = player.player

		break
	}

	return player
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
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
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

	for i, complexID := range player.player.EquipmentSets {
		if EquipmentSetID(complexID.ChildID) != equipmentSetToRemove {
			continue
		}

		player.player.EquipmentSets[i] = player.player.EquipmentSets[len(player.player.EquipmentSets)-1]
		player.player.EquipmentSets[len(player.player.EquipmentSets)-1] = PlayerEquipmentSetRefID{}
		player.player.EquipmentSets = player.player.EquipmentSets[:len(player.player.EquipmentSets)-1]
		player.player.engine.deletePlayerEquipmentSetRef(complexID)

		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
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

	for i, complexID := range player.player.GuildMembers {
		if PlayerID(complexID.ChildID) != guildMemberToRemove {
			continue
		}

		player.player.GuildMembers[i] = player.player.GuildMembers[len(player.player.GuildMembers)-1]
		player.player.GuildMembers[len(player.player.GuildMembers)-1] = PlayerGuildMemberRefID{}
		player.player.GuildMembers = player.player.GuildMembers[:len(player.player.GuildMembers)-1]
		player.player.engine.deletePlayerGuildMemberRef(complexID)

		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
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

	for i, complexID := range player.player.TargetedBy {
		if ZoneItemID(complexID.ChildID) != zoneItemToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = PlayerTargetedByRefID{}
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(complexID)

		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
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

	for i, complexID := range player.player.TargetedBy {
		if PlayerID(complexID.ChildID) != playerToRemove {
			continue
		}

		player.player.TargetedBy[i] = player.player.TargetedBy[len(player.player.TargetedBy)-1]
		player.player.TargetedBy[len(player.player.TargetedBy)-1] = PlayerTargetedByRefID{}
		player.player.TargetedBy = player.player.TargetedBy[:len(player.player.TargetedBy)-1]
		player.player.engine.deletePlayerTargetedByRef(complexID)

		player.player.OperationKind = OperationKindUpdate
		player.player.Meta.sign(player.player.engine.broadcastingClientID)
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

	for i, valID := range zone.zone.Tags {
		if zone.zone.engine.stringValue(valID).Value != tagToRemove {
			continue
		}

		zone.zone.Tags[i] = zone.zone.Tags[len(zone.zone.Tags)-1]
		zone.zone.Tags[len(zone.zone.Tags)-1] = 0
		zone.zone.Tags = zone.zone.Tags[:len(zone.zone.Tags)-1]

		zone.zone.OperationKind = OperationKindUpdate
		zone.zone.Meta.sign(zone.zone.engine.broadcastingClientID)
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

	for i, complexID := range equipmentSet.equipmentSet.Equipment {
		if ItemID(complexID.ChildID) != equipmentToRemove {
			continue
		}

		equipmentSet.equipmentSet.Equipment[i] = equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.Equipment[len(equipmentSet.equipmentSet.Equipment)-1] = EquipmentSetEquipmentRefID{}
		equipmentSet.equipmentSet.Equipment = equipmentSet.equipmentSet.Equipment[:len(equipmentSet.equipmentSet.Equipment)-1]
		equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(complexID)

		equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
		equipmentSet.equipmentSet.Meta.sign(equipmentSet.equipmentSet.engine.broadcastingClientID)
		equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet

		break
	}

	return equipmentSet
}
