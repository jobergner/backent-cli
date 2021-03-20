package state

func (_e Zone) RemovePlayers(se *Engine, playersToRemove ...PlayerID) Zone {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []PlayerID
	for _, element := range e.zone.Players {
		var toBeRemoved bool
		for _, elementToRemove := range playersToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true
				se.deletePlayer(element)
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.zone.Players = newElements
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
	return e
}

func (_e Zone) RemoveItems(se *Engine, zoneItemsToRemove ...ZoneItemID) Zone {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []ZoneItemID
	for _, element := range e.zone.Items {
		var toBeRemoved bool
		for _, elementToRemove := range zoneItemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true
				se.deleteZoneItem(element)
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.zone.Items = newElements
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
	return e
}

func (_e Player) RemoveItems(se *Engine, itemsToRemove ...ItemID) Player {
	e := se.Player(_e.player.ID)
	if e.player.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []ItemID
	for _, element := range e.player.Items {
		var toBeRemoved bool
		for _, elementToRemove := range itemsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true
				se.deleteItem(element)
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.player.Items = newElements
	e.player.OperationKind = OperationKindUpdate
	se.Patch.Player[e.player.ID] = e.player
	return e
}

func (_e Zone) RemoveTags(se *Engine, tagsToRemove ...string) Zone {
	e := se.Zone(_e.zone.ID)
	if e.zone.OperationKind == OperationKindDelete {
		return e
	}
	var elementsAltered bool
	var newElements []string
	for _, element := range e.zone.Tags {
		var toBeRemoved bool
		for _, elementToRemove := range tagsToRemove {
			if element == elementToRemove {
				toBeRemoved = true
				elementsAltered = true
			}
		}
		if !toBeRemoved {
			newElements = append(newElements, element)
		}
	}
	if !elementsAltered {
		return e
	}
	e.zone.Tags = newElements
	e.zone.OperationKind = OperationKindUpdate
	se.Patch.Zone[e.zone.ID] = e.zone
	return e
}
