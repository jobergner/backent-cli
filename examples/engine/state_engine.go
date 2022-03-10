package state

type OperationKind string

const (
	OperationKindDelete    OperationKind = "DELETE"
	OperationKindUpdate    OperationKind = "UPDATE"
	OperationKindUnchanged OperationKind = "UNCHANGED"
)

type Engine struct {
	State                *State
	Patch                *State
	Tree                 *Tree
	broadcastingClientID string
	thisClientID         string
	planner              *assemblePlanner
	IDgen                int
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

	for _, boolValue := range engine.Patch.BoolValue {
		if boolValue.OperationKind == OperationKindDelete {
			delete(engine.State.BoolValue, boolValue.ID)
		} else {
			boolValue.OperationKind = OperationKindUnchanged
			boolValue.Meta.unsign()
			engine.State.BoolValue[boolValue.ID] = boolValue
		}
	}
	for _, floatValue := range engine.Patch.FloatValue {
		if floatValue.OperationKind == OperationKindDelete {
			delete(engine.State.FloatValue, floatValue.ID)
		} else {
			floatValue.OperationKind = OperationKindUnchanged
			floatValue.Meta.unsign()
			engine.State.FloatValue[floatValue.ID] = floatValue
		}
	}
	for _, intValue := range engine.Patch.IntValue {
		if intValue.OperationKind == OperationKindDelete {
			delete(engine.State.IntValue, intValue.ID)
		} else {
			intValue.OperationKind = OperationKindUnchanged
			intValue.Meta.unsign()
			engine.State.IntValue[intValue.ID] = intValue
		}
	}
	for _, stringValue := range engine.Patch.StringValue {
		if stringValue.OperationKind == OperationKindDelete {
			delete(engine.State.StringValue, stringValue.ID)
		} else {
			stringValue.OperationKind = OperationKindUnchanged
			stringValue.Meta.unsign()
			engine.State.StringValue[stringValue.ID] = stringValue
		}
	}

	for _, attackEvent := range engine.Patch.AttackEvent {
		if attackEvent.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEvent, attackEvent.ID)
		} else {
			attackEvent.OperationKind = OperationKindUnchanged
			attackEvent.Meta.unsign()
			engine.State.AttackEvent[attackEvent.ID] = attackEvent
		}
	}
	for _, equipmentSet := range engine.Patch.EquipmentSet {
		if equipmentSet.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSet, equipmentSet.ID)
		} else {
			equipmentSet.OperationKind = OperationKindUnchanged
			equipmentSet.Meta.unsign()
			engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
		}
	}
	for _, gearScore := range engine.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(engine.State.GearScore, gearScore.ID)
		} else {
			gearScore.OperationKind = OperationKindUnchanged
			gearScore.Meta.unsign()
			engine.State.GearScore[gearScore.ID] = gearScore
		}
	}
	for _, item := range engine.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(engine.State.Item, item.ID)
		} else {
			item.OperationKind = OperationKindUnchanged
			item.Meta.unsign()
			engine.State.Item[item.ID] = item
		}
	}
	for _, player := range engine.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(engine.State.Player, player.ID)
		} else {
			player.Action = player.Action[:0]
			player.OperationKind = OperationKindUnchanged
			player.Meta.unsign()
			engine.State.Player[player.ID] = player
		}
	}
	for _, position := range engine.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(engine.State.Position, position.ID)
		} else {
			position.OperationKind = OperationKindUnchanged
			position.Meta.unsign()
			engine.State.Position[position.ID] = position
		}
	}
	for _, zone := range engine.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(engine.State.Zone, zone.ID)
		} else {
			zone.OperationKind = OperationKindUnchanged
			zone.Meta.unsign()
			engine.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range engine.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.ZoneItem, zoneItem.ID)
		} else {
			zoneItem.OperationKind = OperationKindUnchanged
			zoneItem.Meta.unsign()
			engine.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, attackEventTargetRef := range engine.Patch.AttackEventTargetRef {
		if attackEventTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.AttackEventTargetRef, attackEventTargetRef.ID)
		} else {
			attackEventTargetRef.OperationKind = OperationKindUnchanged
			attackEventTargetRef.Meta.unsign()
			engine.State.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
		}
	}
	for _, equipmentSetEquipmentRef := range engine.Patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.OperationKind == OperationKindDelete {
			delete(engine.State.EquipmentSetEquipmentRef, equipmentSetEquipmentRef.ID)
		} else {
			equipmentSetEquipmentRef.OperationKind = OperationKindUnchanged
			equipmentSetEquipmentRef.Meta.unsign()
			engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
		}
	}
	for _, itemBoundToRef := range engine.Patch.ItemBoundToRef {
		if itemBoundToRef.OperationKind == OperationKindDelete {
			delete(engine.State.ItemBoundToRef, itemBoundToRef.ID)
		} else {
			itemBoundToRef.OperationKind = OperationKindUnchanged
			itemBoundToRef.Meta.unsign()
			engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
		}
	}
	for _, playerEquipmentSetRef := range engine.Patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerEquipmentSetRef, playerEquipmentSetRef.ID)
		} else {
			playerEquipmentSetRef.OperationKind = OperationKindUnchanged
			playerEquipmentSetRef.Meta.unsign()
			engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
		}
	}
	for _, playerGuildMemberRef := range engine.Patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerGuildMemberRef, playerGuildMemberRef.ID)
		} else {
			playerGuildMemberRef.OperationKind = OperationKindUnchanged
			playerGuildMemberRef.Meta.unsign()
			engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
		}
	}
	for _, playerTargetRef := range engine.Patch.PlayerTargetRef {
		if playerTargetRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetRef, playerTargetRef.ID)
		} else {
			playerTargetRef.OperationKind = OperationKindUnchanged
			playerTargetRef.Meta.unsign()
			engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
		}
	}
	for _, playerTargetedByRef := range engine.Patch.PlayerTargetedByRef {
		if playerTargetedByRef.OperationKind == OperationKindDelete {
			delete(engine.State.PlayerTargetedByRef, playerTargetedByRef.ID)
		} else {
			playerTargetedByRef.OperationKind = OperationKindUnchanged
			playerTargetedByRef.Meta.unsign()
			engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
		}
	}
	for _, anyOfPlayer_Position := range engine.Patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_Position, anyOfPlayer_Position.ID)
		} else {
			anyOfPlayer_Position.OperationKind = OperationKindUnchanged
			anyOfPlayer_Position.Meta.unsign()
			engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
		}
	}
	for _, anyOfPlayer_ZoneItem := range engine.Patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfPlayer_ZoneItem, anyOfPlayer_ZoneItem.ID)
		} else {
			anyOfPlayer_ZoneItem.OperationKind = OperationKindUnchanged
			anyOfPlayer_ZoneItem.Meta.unsign()
			engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
		}
	}
	for _, anyOfItem_Player_ZoneItem := range engine.Patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.OperationKind == OperationKindDelete {
			delete(engine.State.AnyOfItem_Player_ZoneItem, anyOfItem_Player_ZoneItem.ID)
		} else {
			anyOfItem_Player_ZoneItem.OperationKind = OperationKindUnchanged
			anyOfItem_Player_ZoneItem.Meta.unsign()
			engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
		}
	}

	for key := range engine.Patch.BoolValue {
		delete(engine.Patch.BoolValue, key)
	}
	for key := range engine.Patch.FloatValue {
		delete(engine.Patch.FloatValue, key)
	}
	for key := range engine.Patch.IntValue {
		delete(engine.Patch.IntValue, key)
	}
	for key := range engine.Patch.StringValue {
		delete(engine.Patch.StringValue, key)
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

func (engine *Engine) importPatch(patch *State) {
	for _, boolValue := range patch.BoolValue {
		if boolValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		boolValue.Meta.unsign()
		engine.State.BoolValue[boolValue.ID] = boolValue
	}
	for _, floatValue := range patch.FloatValue {
		if floatValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		floatValue.Meta.unsign()
		engine.State.FloatValue[floatValue.ID] = floatValue
	}
	for _, intValue := range patch.IntValue {
		if intValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		intValue.Meta.unsign()
		engine.State.IntValue[intValue.ID] = intValue
	}
	for _, stringValue := range patch.StringValue {
		if stringValue.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		stringValue.Meta.unsign()
		engine.State.StringValue[stringValue.ID] = stringValue
	}

	for _, attackEvent := range patch.AttackEvent {
		if attackEvent.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		attackEvent.Meta.unsign()
		engine.State.AttackEvent[attackEvent.ID] = attackEvent
	}
	for _, equipmentSet := range patch.EquipmentSet {
		if equipmentSet.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		equipmentSet.Meta.unsign()
		engine.State.EquipmentSet[equipmentSet.ID] = equipmentSet
	}
	for _, gearScore := range patch.GearScore {
		if gearScore.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		gearScore.Meta.unsign()
		engine.State.GearScore[gearScore.ID] = gearScore
	}
	for _, item := range patch.Item {
		if item.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		item.Meta.unsign()
		engine.State.Item[item.ID] = item
	}
	for _, player := range patch.Player {
		if player.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		player.Action = player.Action[:0]
		player.Meta.unsign()
	}
	for _, position := range patch.Position {
		if position.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		position.Meta.unsign()
		engine.State.Position[position.ID] = position
	}
	for _, zone := range patch.Zone {
		if zone.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		zone.Meta.unsign()
		engine.State.Zone[zone.ID] = zone
	}
	for _, zoneItem := range patch.ZoneItem {
		if zoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		zoneItem.Meta.unsign()
		engine.State.ZoneItem[zoneItem.ID] = zoneItem
	}
	for _, attackEventTargetRef := range patch.AttackEventTargetRef {
		if attackEventTargetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		attackEventTargetRef.Meta.unsign()
		engine.State.AttackEventTargetRef[attackEventTargetRef.ID] = attackEventTargetRef
	}
	for _, equipmentSetEquipmentRef := range patch.EquipmentSetEquipmentRef {
		if equipmentSetEquipmentRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		equipmentSetEquipmentRef.Meta.unsign()
		engine.State.EquipmentSetEquipmentRef[equipmentSetEquipmentRef.ID] = equipmentSetEquipmentRef
	}
	for _, itemBoundToRef := range patch.ItemBoundToRef {
		if itemBoundToRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		itemBoundToRef.Meta.unsign()
		engine.State.ItemBoundToRef[itemBoundToRef.ID] = itemBoundToRef
	}
	for _, playerEquipmentSetRef := range patch.PlayerEquipmentSetRef {
		if playerEquipmentSetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerEquipmentSetRef.Meta.unsign()
		engine.State.PlayerEquipmentSetRef[playerEquipmentSetRef.ID] = playerEquipmentSetRef
	}
	for _, playerGuildMemberRef := range patch.PlayerGuildMemberRef {
		if playerGuildMemberRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerGuildMemberRef.Meta.unsign()
		engine.State.PlayerGuildMemberRef[playerGuildMemberRef.ID] = playerGuildMemberRef
	}
	for _, playerTargetRef := range patch.PlayerTargetRef {
		if playerTargetRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerTargetRef.Meta.unsign()
		engine.State.PlayerTargetRef[playerTargetRef.ID] = playerTargetRef
	}
	for _, playerTargetedByRef := range patch.PlayerTargetedByRef {
		if playerTargetedByRef.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		playerTargetedByRef.Meta.unsign()
		engine.State.PlayerTargetedByRef[playerTargetedByRef.ID] = playerTargetedByRef
	}
	for _, anyOfPlayer_Position := range patch.AnyOfPlayer_Position {
		if anyOfPlayer_Position.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfPlayer_Position.Meta.unsign()
		engine.State.AnyOfPlayer_Position[anyOfPlayer_Position.ID] = anyOfPlayer_Position
	}
	for _, anyOfPlayer_ZoneItem := range patch.AnyOfPlayer_ZoneItem {
		if anyOfPlayer_ZoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfPlayer_ZoneItem.Meta.unsign()
		engine.State.AnyOfPlayer_ZoneItem[anyOfPlayer_ZoneItem.ID] = anyOfPlayer_ZoneItem
	}
	for _, anyOfItem_Player_ZoneItem := range patch.AnyOfItem_Player_ZoneItem {
		if anyOfItem_Player_ZoneItem.Meta.BroadcastedBy == engine.thisClientID {
			continue
		}
		anyOfItem_Player_ZoneItem.Meta.unsign()
		engine.State.AnyOfItem_Player_ZoneItem[anyOfItem_Player_ZoneItem.ID] = anyOfItem_Player_ZoneItem
	}
}
