package state

func (_zone zone) RemovePlayers(playersToRemove ...PlayerID) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []PlayerID
	for _, element := range zone.zone.Players {
		var toBeRemoved bool
		for _, elementToRemove := range playersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				zone.zone.engine.deletePlayer(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Players = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_zone zone) RemoveItems(itemsToRemove ...ZoneItemID) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []ZoneItemID
	for _, element := range zone.zone.Items {
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				zone.zone.engine.deleteZoneItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Items = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_zone zone) RemoveInteracableItem(itemsToRemove ...ItemID) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []AnyOfItemPlayerZoneItemID
	for _, anyContainerID := range zone.zone.Interactables {
		anyContainer := zone.zone.engine.anyOfItemPlayerZoneItem(anyContainerID)
		element := anyContainer.Item().ID()
		if element == 0 {
			continue
		}
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				zone.zone.engine.deleteItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, anyContainer.anyOfItemPlayerZoneItem.ID)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Interactables = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_zone zone) RemoveInteracablePlayer(playersToRemove ...PlayerID) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []AnyOfItemPlayerZoneItemID
	for _, anyContainerID := range zone.zone.Interactables {
		anyContainer := zone.zone.engine.anyOfItemPlayerZoneItem(anyContainerID)
		element := anyContainer.Player().ID()
		if element == 0 {
			continue
		}
		var toBeRemoved bool
		for _, elementToRemove := range playersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				zone.zone.engine.deletePlayer(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, anyContainer.anyOfItemPlayerZoneItem.ID)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Interactables = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_zone zone) RemoveInteracableZoneItem(zoneItemsToRemove ...ZoneItemID) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []AnyOfItemPlayerZoneItemID
	for _, anyContainerID := range zone.zone.Interactables {
		anyContainer := zone.zone.engine.anyOfItemPlayerZoneItem(anyContainerID)
		element := anyContainer.ZoneItem().ID()
		if element == 0 {
			continue
		}
		var toBeRemoved bool
		for _, elementToRemove := range zoneItemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				zone.zone.engine.deleteZoneItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, anyContainer.anyOfItemPlayerZoneItem.ID)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Interactables = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_player player) RemoveItems(itemsToRemove ...ItemID) player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []ItemID
	for _, element := range player.player.Items {
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				player.player.engine.deleteItem(element)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.Items = newElements
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player player) RemoveEquipmentSets(equipmentSetsToRemove ...EquipmentSetID) player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []PlayerEquipmentSetRefID
	for _, refElement := range player.player.EquipmentSets {
		element := player.player.engine.playerEquipmentSetRef(refElement).playerEquipmentSetRef.ReferencedElementID
		var toBeRemoved bool
		for _, elementToRemove := range equipmentSetsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				player.player.engine.deletePlayerEquipmentSetRef(refElement)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, refElement)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.EquipmentSets = newElements
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player player) RemoveGuildMembers(guildMembersToRemove ...PlayerID) player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []PlayerGuildMemberRefID
	for _, refElement := range player.player.GuildMembers {
		element := player.player.engine.playerGuildMemberRef(refElement).playerGuildMemberRef.ReferencedElementID
		var toBeRemoved bool
		for _, elementToRemove := range guildMembersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				player.player.engine.deletePlayerGuildMemberRef(refElement)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, refElement)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.GuildMembers = newElements
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player player) RemoveTargetedByZoneItem(zoneItemsToRemove ...ZoneItemID) player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []PlayerTargetedByRefID
	for _, refElement := range player.player.TargetedBy {
		anyContainer := player.player.engine.playerTargetedByRef(refElement).Get()
		element := anyContainer.ZoneItem().ID()
		if element == 0 {
			continue
		}
		var toBeRemoved bool
		for _, elementToRemove := range zoneItemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				player.player.engine.deletePlayerTargetedByRef(refElement)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, refElement)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.TargetedBy = newElements
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player player) RemoveTargetedByPlayer(playersToRemove ...PlayerID) player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []PlayerTargetedByRefID
	for _, refElement := range player.player.TargetedBy {
		anyContainer := player.player.engine.playerTargetedByRef(refElement).Get()
		element := anyContainer.Player().ID()
		if element == 0 {
			continue
		}
		var toBeRemoved bool
		for _, elementToRemove := range playersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				player.player.engine.deletePlayerTargetedByRef(refElement)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, refElement)
		}
	}
	if !wereElementsAltered {
		return player
	}
	player.player.TargetedBy = newElements
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_zone zone) RemoveTags(tagsToRemove ...string) zone {
	zone := _zone.zone.engine.Zone(_zone.zone.ID)
	if zone.zone.OperationKind == OperationKindDelete {
		return zone
	}
	var wereElementsAltered bool
	var newElements []string
	for _, element := range zone.zone.Tags {
		var toBeRemoved bool
		for _, elementToRemove := range tagsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !wereElementsAltered {
		return zone
	}
	zone.zone.Tags = newElements
	zone.zone.OperationKind = OperationKindUpdate
	zone.zone.engine.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_equipmentSet equipmentSet) RemoveEquipment(equipmentToRemove ...ItemID) equipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	var wereElementsAltered bool
	var newElements []EquipmentSetEquipmentRefID
	for _, refElement := range equipmentSet.equipmentSet.Equipment {
		element := equipmentSet.equipmentSet.engine.equipmentSetEquipmentRef(refElement).equipmentSetEquipmentRef.ReferencedElementID
		var toBeRemoved bool
		for _, elementToRemove := range equipmentToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				equipmentSet.equipmentSet.engine.deleteEquipmentSetEquipmentRef(refElement)
				break
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, refElement)
		}
	}
	if !wereElementsAltered {
		return equipmentSet
	}
	equipmentSet.equipmentSet.Equipment = newElements
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	return equipmentSet
}
