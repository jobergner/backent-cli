package statemachine

func (z zone) RemovePlayer(playerID playerID, sm *stateMachine) zone {
	var indexToRemove int
	for i, _playerID := range z.players {
		if _playerID == playerID {
			indexToRemove = i
			break
		}
	}
	z.players = append(z.players[:indexToRemove], z.players[indexToRemove+1:]...)
	z.operationKind = operationKindUpdate
	sm.patch.zone[z.id] = z
	sm.DeletePlayer(playerID)
	return z
}

func (z zone) RemoveZoneItem(zoneItemID zoneItemID, sm *stateMachine) zone {
	var indexToRemove int
	for i, _zoneItemID := range z.items {
		if _zoneItemID == zoneItemID {
			indexToRemove = i
			break
		}
	}
	z.items = append(z.items[:indexToRemove], z.items[indexToRemove+1:]...)
	z.operationKind = operationKindUpdate
	sm.patch.zone[z.id] = z
	sm.DeleteZoneItem(zoneItemID)
	return z
}

func (p player) RemoveItem(itemID itemID, sm *stateMachine) player {
	var indexToRemove int
	for i, _itemID := range p.items {
		if _itemID == itemID {
			indexToRemove = i
			break
		}
	}
	p.items = append(p.items[:indexToRemove], p.items[indexToRemove+1:]...)
	p.operationKind = operationKindUpdate
	sm.patch.player[p.id] = p
	sm.DeleteItem(itemID)
	return p
}
