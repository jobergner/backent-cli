package statemachine

func (z Zone) RemovePlayer(playerID PlayerID, sm *StateMachine) Zone {
	var indexToRemove int
	for i, _playerID := range z.zone.Players {
		if _playerID == playerID {
			indexToRemove = i
			break
		}
	}
	z.zone.Players = append(z.zone.Players[:indexToRemove], z.zone.Players[indexToRemove+1:]...)
	z.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[z.zone.ID] = z.zone
	sm.DeletePlayer(playerID)
	return z
}

func (z Zone) RemoveZoneItem(zoneItemID ZoneItemID, sm *StateMachine) Zone {
	var indexToRemove int
	for i, _zoneItemID := range z.zone.Items {
		if _zoneItemID == zoneItemID {
			indexToRemove = i
			break
		}
	}
	z.zone.Items = append(z.zone.Items[:indexToRemove], z.zone.Items[indexToRemove+1:]...)
	z.zone.OperationKind = OperationKindUpdate
	sm.Patch.Zone[z.zone.ID] = z.zone
	sm.DeleteZoneItem(zoneItemID)
	return z
}

func (p Player) RemoveItem(itemID ItemID, sm *StateMachine) Player {
	var indexToRemove int
	for i, _itemID := range p.player.Items {
		if _itemID == itemID {
			indexToRemove = i
			break
		}
	}
	p.player.Items = append(p.player.Items[:indexToRemove], p.player.Items[indexToRemove+1:]...)
	p.player.OperationKind = OperationKindUpdate
	sm.Patch.Player[p.player.ID] = p.player
	sm.DeleteItem(itemID)
	return p
}
