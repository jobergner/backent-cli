package state

type OperationKind string

const (
	OperationKindDelete OperationKind = "DELETE"
	OperationKindUpdate OperationKind = "UPDATE"
)

type Engine struct {
	State State
	Patch State
	IDgen int
}

func newEngine() *Engine {
	return &Engine{
		IDgen: 1,
		Patch: newState(),
		State: newState(),
	}
}

func (se *Engine) GenerateID() int {
	newID := se.IDgen
	se.IDgen = se.IDgen + 1
	return newID
}

func (se *Engine) updateReferenceRelationships() {
	for _, player := range se.Patch.Player {
		if player.OperationKind_ == OperationKindDelete {
			// delete all references of this element in other elements
			for _, itemID := range se.allItemIDs() {
				_item := se.Item(itemID)
				if _item.BoundTo(se).ID(se) == player.ID {
					_item.BoundTo(se).Unset(se)
				}
			}
			for _, playerID := range se.allPlayerIDs() {
				_player := se.Player(playerID)
				_player.RemoveGuildMembers(se, playerID)
			}
		} else {
			// update all elements referencing this element
			for _, itemID := range se.allItemIDs() {
				_item := se.Item(itemID)
				if _item.BoundTo(se).ID(se) == player.ID {
					se.updateItemUpstream(_item.item)
				}
			}
			for _, playerID := range se.allPlayerIDs() {
				_player := se.Player(playerID)
				se.updatePlayerUpstream(_player.player, []PlayerID{})
			}
		}
	}
}

func (se *Engine) UpdateState() {
	for _, gearScore := range se.Patch.GearScore {
		if gearScore.OperationKind_ == OperationKindDelete {
			delete(se.State.GearScore, gearScore.ID)
		} else {
			se.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range se.Patch.Item {
		if item.OperationKind_ == OperationKindDelete {
			delete(se.State.Item, item.ID)
		} else {
			se.State.Item[item.ID] = item
		}
	}
	for _, player := range se.Patch.Player {
		if player.OperationKind_ == OperationKindDelete {
			delete(se.State.Player, player.ID)
		} else {
			se.State.Player[player.ID] = player
		}
	}
	for _, position := range se.Patch.Position {
		if position.OperationKind_ == OperationKindDelete {
			delete(se.State.Position, position.ID)
		} else {
			se.State.Position[position.ID] = position
		}
	}
	for _, zone := range se.Patch.Zone {
		if zone.OperationKind_ == OperationKindDelete {
			delete(se.State.Zone, zone.ID)
		} else {
			se.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if zoneItem.OperationKind_ == OperationKindDelete {
			delete(se.State.ZoneItem, zoneItem.ID)
		} else {
			se.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	// TODO dont forget to update, delete references

	for key := range se.Patch.GearScore {
		delete(se.Patch.GearScore, key)
	}
	for key := range se.Patch.Item {
		delete(se.Patch.Item, key)
	}
	for key := range se.Patch.Player {
		delete(se.Patch.Player, key)
	}
	for key := range se.Patch.Position {
		delete(se.Patch.Position, key)
	}
	for key := range se.Patch.Zone {
		delete(se.Patch.Zone, key)
	}
	for key := range se.Patch.ZoneItem {
		delete(se.Patch.ZoneItem, key)
	}
}
