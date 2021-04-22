package state

func (_zone zone) RemovePlayers(se *Engine, playersToRemove ...PlayerID) zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
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
				se.deletePlayer(element)
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
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_zone zone) RemoveItems(se *Engine, itemsToRemove ...ZoneItemID) zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
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
				se.deleteZoneItem(element)
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
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_player player) RemoveItems(se *Engine, itemsToRemove ...ItemID) player {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
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
				se.deleteItem(element)
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
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player player) RemoveGuildMembers(se *Engine, guildMembersToRemove ...PlayerID) player {
	player := se.Player(_player.player.ID)
	if player.player.OperationKind_ == OperationKindDelete {
		return player
	}
	var wereElementsAltered bool
	var newElements []PlayerGuildMemberRefID
	for _, refElement := range player.player.GuildMembers {
		element := se.playerGuildMemberRef(refElement).playerGuildMemberRef.ReferencedElementID
		var toBeRemoved bool
		for _, elementToRemove := range guildMembersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				se.deletePlayerGuildMemberRef(refElement)
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
	player.player.OperationKind_ = OperationKindUpdate
	se.Patch.Player[player.player.ID] = player.player
	return player
}

func (_zone zone) RemoveTags(se *Engine, tagsToRemove ...string) zone {
	zone := se.Zone(_zone.zone.ID)
	if zone.zone.OperationKind_ == OperationKindDelete {
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
	zone.zone.OperationKind_ = OperationKindUpdate
	se.Patch.Zone[zone.zone.ID] = zone.zone
	return zone
}

func (_equipmentSet equipmentSet) RemoveEquipment(se *Engine, itemsToRemove ...ItemID) equipmentSet {
	equipmentSet := se.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind_ == OperationKindDelete {
		return equipmentSet
	}
	var wereElementsAltered bool
	var newElements []EquipmentSetEquipmentRefID
	for _, refElement := range equipmentSet.equipmentSet.Equipment {
		element := se.equipmentSetEquipmentRef(refElement).equipmentSetEquipmentRef.ReferencedElementID
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				wereElementsAltered = true
				se.deleteEquipmentSetEquipmentRef(refElement)
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
	equipmentSet.equipmentSet.OperationKind_ = OperationKindUpdate
	se.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	return equipmentSet
}
