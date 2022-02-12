package state

type OperationKind string

const (
	OperationKindDelete    OperationKind = "DELETE"
	OperationKindUpdate    OperationKind = "UPDATE"
	OperationKindUnchanged OperationKind = "UNCHANGED"
)

type Engine struct {
	State   *State
	Patch   *State
	Tree    *Tree
	planner *assemblePlanner
	IDgen   int
}

func newEngine() *Engine {
	return &Engine{
		IDgen:   1,
		Patch:   newState(),
		State:   newState(),
		Tree:    newTree(),
		planner: newAssemblePlanner(),
	}
}

func (engine *Engine) GenerateID() int {
	newID := engine.IDgen
	engine.IDgen = engine.IDgen + 1
	return newID
}

func (engine *Engine) UpdateState() {
	for _, attackEvent := range engine.Patch.AttackEvent {
		// so event and children will be deleted
		engine.deleteAttackEvent(attackEvent.ID)
	}

	for _, attackEvent := range engine.Patch.AttackEvent {
		if attackEvent.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEvent, attackEvent.ID)
		} else {
			attackEvent.OperationKind = OperationKindUnchanged
			engine.State.AttackEvent[attackEvent.ID] = attackEvent
		}
	}
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		if equipmentSet.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSet, equipmentSet.ID)
		} else {
			equipmentSet.OperationKind = OperationKindUnchanged
			engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range engine.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(engine.State.GearScore, gearScore.ID)
		} else {
			gearScore.OperationKind = OperationKindUnchanged
			engine.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range engine.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(engine.State.Item, item.ID)
		} else {
			item.OperationKind = OperationKindUnchanged
			engine.State.Item[item.ID] = item
		}
	}
	for _, player := range engine.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(engine.State.Player, player.ID)
		} else {
			player.Action = player.Action[:0]
			player.OperationKind = OperationKindUnchanged
			engine.State.Player[player.ID] = player
		}
	}
	for _, position := range engine.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(engine.State.Position, position.ID)
		} else {
			position.OperationKind = OperationKindUnchanged
			engine.State.Position[position.ID] = position
		}
	}
	for _, zone := range engine.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(engine.State.Zone, zone.ID)
		} else {
			zone.OperationKind = OperationKindUnchanged
			engine.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.ZoneItem, zoneItem.ID)
		} else {
			zoneItem.OperationKind = OperationKindUnchanged
			engine.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			equipmentSetEquipmentRef.OperationKind = OperationKindUnchanged
			engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind == OperationKindDelete {
			delete(engine.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			itemBoundToRef.OperationKind = OperationKindUnchanged
			engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			playerEquipmentSetRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			playerGuildMemberRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		if playerTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetRef, playerTargetRef.ID)
		} else {
			playerTargetRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
		}
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		if playerTargetedByRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetedByRef, playerTargetedByRef.ID)
		} else {
			playerTargetedByRef.OperationKind = OperationKindUnchanged
			engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
		}
	}
	for _, anyOfPlayer_Position := range engine.Patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_Position, anyOfPlayer_Position.ID)
		} else {
			anyOfPlayer_Position.OperationKind = OperationKindUnchanged
			engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
		}
	}
	for _, anyOfPlayer_ZoneItem := range engine.Patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItem.ID)
		} else {
			anyOfPlayer_ZoneItem.OperationKind = OperationKindUnchanged
			engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
		}
	}
	for _, anyOfItem_Player_ZoneItem := range engine.Patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItem.ID)
		} else {
			anyOfItem_Player_ZoneItem.OperationKind = OperationKindUnchanged
			engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
		}
	}

	for key := range engine.Patch.AttackEvent {
		delete(engine.Patch.AttackEvent, key)
	}
	for key := range engine.Patch.EquipmentSet {
		delete(engine.Patch.EquipmentSet, key)
	}
	for key := range engine.Patch.GearScore {
		delete(engine.Patch.GearScore, key)
	}
	for key := range engine.Patch.Item {
		delete(engine.Patch.Item, key)
	}
	for key := range engine.Patch.Player {
		delete(engine.Patch.Player, key)
	}
	for key := range engine.Patch.Position {
		delete(engine.Patch.Position, key)
	}
	for key := range engine.Patch.Zone {
		delete(engine.Patch.Zone, key)
	}
	for key := range engine.Patch.ZoneItem {
		delete(engine.Patch.ZoneItem, key)
	}
	for key := range engine.Patch.AttackEventTargetRef {
		delete(engine.Patch.AttackEventTargetRef, key)
	}
	for key := range engine.Patch.EquipmentSetEquipmentRef {
		delete(engine.Patch.EquipmentSetEquipmentRef, key)
	}
	for key := range engine.Patch.ItemBoundToRef {
		delete(engine.Patch.ItemBoundToRef, key)
	}
	for key := range engine.Patch.PlayerEquipmentSetRef {
		delete(engine.Patch.PlayerEquipmentSetRef, key)
	}
	for key := range engine.Patch.PlayerGuildMemberRef {
		delete(engine.Patch.PlayerGuildMemberRef, key)
	}
	for key := range engine.Patch.PlayerTargetRef {
		delete(engine.Patch.PlayerTargetRef, key)
	}
	for key := range engine.Patch.PlayerTargetedByRef {
		delete(engine.Patch.PlayerTargetedByRef, key)
	}
	for key := range engine.Patch.AnyOfPlayer_Position {
		delete(engine.Patch.AnyOfPlayer_Position, key)
	}
	for key := range engine.Patch.AnyOfPlayer_ZoneItem {
		delete(engine.Patch.AnyOfPlayer_ZoneItem, key)
	}
	for key := range engine.Patch.AnyOfItem_Player_ZoneItem {
		delete(engine.Patch.AnyOfItem_Player_ZoneItem, key)
	}
}
