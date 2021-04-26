package state

type OperationKind string

const (
	OperationKindDelete    OperationKind = "DELETE"
	OperationKindUpdate    OperationKind = "UPDATE"
	OperationKindUnchanged OperationKind = "UNCHANGED"
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
			equipmentSet.OperationKind_ = OperationKindUnchanged
			se.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range se.Patch.GearScore {
		if gearScore.OperationKind_ == OperationKindDelete {
			delete(se.State.GearScore, gearScore.ID)
		} else {
			gearScore.OperationKind_ = OperationKindUnchanged
			se.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range se.Patch.Item {
		if item.OperationKind_ == OperationKindDelete {
			delete(se.State.Item, item.ID)
		} else {
			item.OperationKind_ = OperationKindUnchanged
			se.State.Item[item.ID] = item
		}
	}
	for _, player := range se.Patch.Player {
		if player.OperationKind_ == OperationKindDelete {
			delete(se.State.Player, player.ID)
		} else {
			player.OperationKind_ = OperationKindUnchanged
			se.State.Player[player.ID] = player
		}
	}
	for _, position := range se.Patch.Position {
		if position.OperationKind_ == OperationKindDelete {
			delete(se.State.Position, position.ID)
		} else {
			position.OperationKind_ = OperationKindUnchanged
			se.State.Position[position.ID] = position
		}
	}
	for _, zone := range se.Patch.Zone {
		if zone.OperationKind_ == OperationKindDelete {
			delete(se.State.Zone, zone.ID)
		} else {
			zone.OperationKind_ = OperationKindUnchanged
			se.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if zoneItem.OperationKind_ == OperationKindDelete {
			delete(se.State.ZoneItem, zoneItem.ID)
		} else {
			zoneItem.OperationKind_ = OperationKindUnchanged
			se.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, equipmentSetEquipmentRef := range se.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind_ == OperationKindDelete {
			delete(se.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			equipmentSetEquipmentRef.OperationKind_ = OperationKindUnchanged
			se.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range se.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind_ == OperationKindDelete {
			delete(se.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			itemBoundToRef.OperationKind_ = OperationKindUnchanged
			se.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range se.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind_ == OperationKindDelete {
			delete(se.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			playerEquipmentSetRef.OperationKind_ = OperationKindUnchanged
			se.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range se.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind_ == OperationKindDelete {
			delete(se.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			playerGuildMemberRef.OperationKind_ = OperationKindUnchanged
			se.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}

	for key := range se.Patch.EquipmentSet {
		delete(se.Patch.EquipmentSet, key)
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
	for key := range se.Patch.EquipmentSetEquipmentRef {
		delete(se.Patch.EquipmentSetEquipmentRef, key)
	}
	for key := range se.Patch.ItemBoundToRef {
		delete(se.Patch.ItemBoundToRef, key)
	}
	for key := range se.Patch.PlayerEquipmentSetRef {
		delete(se.Patch.PlayerEquipmentSetRef, key)
	}
	for key := range se.Patch.PlayerGuildMemberRef {
		delete(se.Patch.PlayerGuildMemberRef, key)
	}
}
