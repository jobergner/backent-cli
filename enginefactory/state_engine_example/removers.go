package state

func (_zone Zone) RemovePlayers(se *Engine, playersToRemove ...PlayerID) Zone {
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

func (_zone Zone) RemoveItems(se *Engine, itemsToRemove ...ZoneItemID) Zone {
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

func (_player Player) RemoveItems(se *Engine, itemsToRemove ...ItemID) Player {
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

func (_zone Zone) RemoveTags(se *Engine, tagsToRemove ...string) Zone {
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
