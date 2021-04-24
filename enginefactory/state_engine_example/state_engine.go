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

func (se *Engine) UpdateState() {
	for _, equipmentSet := range se.Patch.EquipmentSet {
		if equipmentSet.OperationKind_ == OperationKindDelete {
			delete(se.State.EquipmentSet, equipmentSet.ID)
		} else {
			se.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
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
	for _, equipmentSetEquipmentRef := range se.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind_ == OperationKindDelete {
			delete(se.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			se.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range se.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind_ == OperationKindDelete {
			delete(se.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			se.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range se.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind_ == OperationKindDelete {
			delete(se.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			se.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range se.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind_ == OperationKindDelete {
			delete(se.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			se.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}

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
