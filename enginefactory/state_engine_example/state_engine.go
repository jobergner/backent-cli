package state

type OperationKind string

const (
	OperationKindDelete    OperationKind = "DELETE"
	OperationKindUpdate    OperationKind = "UPDATE"
	OperationKindUnchanged OperationKind = "UNCHANGED"
)

type Engine struct {
	State     State
	Patch     State
	Tree      Tree
	PathTrack pathTrack
	IDgen     int
}

func newEngine() *Engine {
	return &Engine{
		IDgen:     1,
		Patch:     newState(),
		State:     newState(),
		PathTrack: newPathTrack(),
		Tree:      newTree(),
	}
}

func (se *Engine) GenerateID() int {
	newID := se.IDgen
	se.IDgen = se.IDgen + 1
	return newID
}

func (se *Engine) UpdateState() {
	for _, equipmentSet := range se.Patch.EquipmentSet {
		if equipmentSet.OperationKind == OperationKindDelete {
			delete(se.State.EquipmentSet, equipmentSet.ID)
		} else {
			equipmentSet.OperationKind = OperationKindUnchanged
			se.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range se.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(se.State.GearScore, gearScore.ID)
		} else {
			gearScore.OperationKind = OperationKindUnchanged
			se.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range se.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(se.State.Item, item.ID)
		} else {
			item.OperationKind = OperationKindUnchanged
			se.State.Item[item.ID] = item
		}
	}
	for _, player := range se.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(se.State.Player, player.ID)
		} else {
			player.OperationKind = OperationKindUnchanged
			se.State.Player[player.ID] = player
		}
	}
	for _, position := range se.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(se.State.Position, position.ID)
		} else {
			position.OperationKind = OperationKindUnchanged
			se.State.Position[position.ID] = position
		}
	}
	for _, zone := range se.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(se.State.Zone, zone.ID)
		} else {
			zone.OperationKind = OperationKindUnchanged
			se.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range se.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(se.State.ZoneItem, zoneItem.ID)
		} else {
			zoneItem.OperationKind = OperationKindUnchanged
			se.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, equipmentSetEquipmentRef := range se.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
			delete(se.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			equipmentSetEquipmentRef.OperationKind = OperationKindUnchanged
			se.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range se.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind == OperationKindDelete {
			delete(se.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			itemBoundToRef.OperationKind = OperationKindUnchanged
			se.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range se.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind == OperationKindDelete {
			delete(se.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			playerEquipmentSetRef.OperationKind = OperationKindUnchanged
			se.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range se.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind == OperationKindDelete {
			delete(se.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			playerGuildMemberRef.OperationKind = OperationKindUnchanged
			se.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}
	for _, playerTargetRef := range se.Patch.PlayerTargetRef {
		if playerTargetRef.OperationKind == OperationKindDelete {
			delete(se.State.PlayerTargetRef, playerTargetRef.ID)
		} else {
			playerTargetRef.OperationKind = OperationKindUnchanged
			se.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
		}
	}
	for _, playerTargetedByRef := range se.Patch.PlayerTargetedByRef {
		if playerTargetedByRef.OperationKind == OperationKindDelete {
			delete(se.State.PlayerTargetedByRef, playerTargetedByRef.ID)
		} else {
			playerTargetedByRef.OperationKind = OperationKindUnchanged
			se.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
		}
	}
	for _, anyOfItemPlayerZoneItem := range se.Patch.AnyOfItemPlayerZoneItem {
		if anyOfItemPlayerZoneItem.OperationKind == OperationKindDelete {
			delete(se.State.AnyOfItemPlayerZoneItem, anyOfItemPlayerZoneItem.ID)
		} else {
			anyOfItemPlayerZoneItem.OperationKind = OperationKindUnchanged
			se.State.AnyOfItemPlayerZoneItem[anyOfItemPlayerZoneItem.ID] = anyOfItemPlayerZoneItem
		}
	}
	for _, anyOfPlayerPosition := range se.Patch.AnyOfPlayerPosition {
		if anyOfPlayerPosition.OperationKind == OperationKindDelete {
			delete(se.State.AnyOfPlayerPosition, anyOfPlayerPosition.ID)
		} else {
			anyOfPlayerPosition.OperationKind = OperationKindUnchanged
			se.State.AnyOfPlayerPosition[anyOfPlayerPosition.ID] = anyOfPlayerPosition
		}
	}
	for _, anyOfPlayerZoneItem := range se.Patch.AnyOfPlayerZoneItem {
		if anyOfPlayerZoneItem.OperationKind == OperationKindDelete {
			delete(se.State.AnyOfPlayerZoneItem, anyOfPlayerZoneItem.ID)
		} else {
			anyOfPlayerZoneItem.OperationKind = OperationKindUnchanged
			se.State.AnyOfPlayerZoneItem[anyOfPlayerZoneItem.ID] = anyOfPlayerZoneItem
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
	for key := range se.Patch.PlayerTargetRef {
		delete(se.Patch.PlayerTargetRef, key)
	}
	for key := range se.Patch.PlayerTargetedByRef {
		delete(se.Patch.PlayerTargetedByRef, key)
	}
	for key := range se.Patch.AnyOfItemPlayerZoneItem {
		delete(se.Patch.AnyOfItemPlayerZoneItem, key)
	}
	for key := range se.Patch.AnyOfPlayerPosition {
		delete(se.Patch.AnyOfPlayerPosition, key)
	}
	for key := range se.Patch.AnyOfPlayerZoneItem {
		delete(se.Patch.AnyOfPlayerZoneItem, key)
	}
}
