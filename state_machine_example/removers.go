package statemachine

func (z Zone) RemovePlayer(playerID PlayerID, sm *StateMachine) Zone {
	var indexToRemove int
	for i, _playerID := range z.Players {
		if _playerID == playerID {
			indexToRemove = i
			break
		}
	}
	z.Players = append(z.Players[:indexToRemove], z.Players[indexToRemove+1:]...)
	z.OperationKind = OperationKindUpdate
	sm.Patch.Zone[z.ID] = z
	sm.DeletePlayer(playerID)
	return z
}

func (z Zone) RemoveZoneItem(zoneItemID ZoneItemID, sm *StateMachine) Zone {
	var indexToRemove int
	for i, _zoneItemID := range z.Items {
		if _zoneItemID == zoneItemID {
			indexToRemove = i
			break
		}
	}
	z.Items = append(z.Items[:indexToRemove], z.Items[indexToRemove+1:]...)
	z.OperationKind = OperationKindUpdate
	sm.Patch.Zone[z.ID] = z
	sm.DeleteZoneItem(zoneItemID)
	return z
}

func (p Player) RemoveItem(itemID ItemID, sm *StateMachine) Player {
	var indexToRemove int
	for i, _itemID := range p.Items {
		if _itemID == itemID {
			indexToRemove = i
			break
		}
	}
	p.Items = append(p.Items[:indexToRemove], p.Items[indexToRemove+1:]...)
	p.OperationKind = OperationKindUpdate
	sm.Patch.Player[p.ID] = p
	sm.DeleteItem(itemID)
	return p
}
