package statemachine

func (_e Zone) RemovePlayers(sm *StateMachine, playersToRemove ...PlayerID) Zone {
	e := sm.GetZone(_e.zone.ID)
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
				sm.DeletePlayer(element)
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
	sm.Patch.Zone[e.zone.ID] = e.zone
	return e
}

func (_e Zone) RemoveZoneItems(sm *StateMachine, zoneItemsToRemove ...ZoneItemID) Zone {
	e := sm.GetZone(_e.zone.ID)
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
				sm.DeleteZoneItem(element)
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
	sm.Patch.Zone[e.zone.ID] = e.zone
	return e
}

func (_e Player) RemoveItems(sm *StateMachine, itemsToRemove ...ItemID) Player {
	e := sm.GetPlayer(_e.player.ID)
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
				sm.DeleteItem(element)
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
	sm.Patch.Player[e.player.ID] = e.player
	return e
}

func (_e Zone) RemoveTags(sm *StateMachine, tagsToRemove ...string) Zone {
	e := sm.GetZone(_e.zone.ID)
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
	sm.Patch.Zone[e.zone.ID] = e.zone
	return e
}
