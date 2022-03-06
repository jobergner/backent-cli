package state

func (engine *Engine) setBoolValue(id BoolValueID, val bool) {
	boolValue := engine.boolValue(id)
	if boolValue.OperationKind == OperationKindDelete {
		return
	}
	if boolValue.Value == val {
		return
	}
	boolValue.Value = val
	boolValue.OperationKind = OperationKindUpdate
	boolValue.Meta.sign(boolValue.engine.broadcastingClientID)
	engine.Patch.BoolValue[id] = boolValue
}

func (engine *Engine) setFloatValue(id FloatValueID, val float64) {
	floatValue := engine.floatValue(id)
	if floatValue.OperationKind == OperationKindDelete {
		return
	}
	if floatValue.Value == val {
		return
	}
	floatValue.Value = val
	floatValue.OperationKind = OperationKindUpdate
	floatValue.Meta.sign(floatValue.engine.broadcastingClientID)
	engine.Patch.FloatValue[id] = floatValue
}

func (engine *Engine) setIntValue(id IntValueID, val int64) {
	intValue := engine.intValue(id)
	if intValue.OperationKind == OperationKindDelete {
		return
	}
	if intValue.Value == val {
		return
	}
	intValue.Value = val
	intValue.OperationKind = OperationKindUpdate
	intValue.Meta.sign(intValue.engine.broadcastingClientID)
	engine.Patch.IntValue[id] = intValue
}

func (engine *Engine) setStringValue(id StringValueID, val string) {
	stringValue := engine.stringValue(id)
	if stringValue.OperationKind == OperationKindDelete {
		return
	}
	if stringValue.Value == val {
		return
	}
	stringValue.Value = val
	stringValue.OperationKind = OperationKindUpdate
	stringValue.Meta.sign(stringValue.engine.broadcastingClientID)
	engine.Patch.StringValue[id] = stringValue
}

func (_gearScore GearScore) SetLevel(newLevel int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Level, newLevel)
	return gearScore
}

func (_gearScore GearScore) SetScore(newScore int64) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.engine.setIntValue(gearScore.gearScore.Score, newScore)
	return gearScore
}

func (_position Position) SetX(newX float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.X, newX)
	return position
}

func (_position Position) SetY(newY float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.engine.setFloatValue(position.position.Y, newY)
	return position
}

func (_item Item) SetName(newName string) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	item.item.engine.setStringValue(item.item.Name, newName)
	return item
}

func (_item Item) SetBoundTo(playerID PlayerID) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return item
	}
	if PlayerID(item.item.BoundTo.ChildID) == playerID {
		return item
	}
	if item.item.BoundTo != (ItemBoundToRefID{}) {
		item.item.engine.deleteItemBoundToRef(item.item.BoundTo)
	}
	ref := item.item.engine.createItemBoundToRef(item.item.Path, item_boundToIdentifier, playerID, item.item.ID)
	item.item.BoundTo = ref.ID
	item.item.OperationKind = OperationKindUpdate
	item.item.Meta.sign(item.item.engine.broadcastingClientID)
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}

func (_attackEvent AttackEvent) SetTarget(playerID PlayerID) AttackEvent {
	attackEvent := _attackEvent.attackEvent.engine.AttackEvent(_attackEvent.attackEvent.ID)
	if attackEvent.attackEvent.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if attackEvent.attackEvent.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return attackEvent
	}
	if PlayerID(attackEvent.attackEvent.Target.ChildID) == playerID {
		return attackEvent
	}
	if attackEvent.attackEvent.Target != (AttackEventTargetRefID{}) {
		attackEvent.attackEvent.engine.deleteAttackEventTargetRef(attackEvent.attackEvent.Target)
	}
	ref := attackEvent.attackEvent.engine.createAttackEventTargetRef(attackEvent.attackEvent.Path, attackEvent_targetIdentifier, playerID, attackEvent.attackEvent.ID)
	attackEvent.attackEvent.Target = ref.ID
	attackEvent.attackEvent.OperationKind = OperationKindUpdate
	attackEvent.attackEvent.Meta.sign(attackEvent.attackEvent.engine.broadcastingClientID)
	attackEvent.attackEvent.engine.Patch.AttackEvent[attackEvent.attackEvent.ID] = attackEvent.attackEvent
	return attackEvent
}

func (_equipmentSet EquipmentSet) SetName(newName string) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.engine.setStringValue(equipmentSet.equipmentSet.Name, newName)
	return equipmentSet
}

func (_player Player) SetTargetPlayer(playerID PlayerID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
		return player
	}
	if PlayerID(player.player.Target.ChildID) == playerID {
		return player
	}
	if player.player.Target != (PlayerTargetRefID{}) {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(playerID), ElementKindPlayer, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindPlayer, int(playerID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}

func (_player Player) SetTargetZoneItem(zoneItemID ZoneItemID) Player {
	player := _player.player.engine.Player(_player.player.ID)
	if player.player.OperationKind == OperationKindDelete {
		return player
	}
	if player.player.engine.ZoneItem(zoneItemID).zoneItem.OperationKind == OperationKindDelete {
		return player
	}
	if ZoneItemID(player.player.Target.ChildID) == zoneItemID {
		return player
	}
	if player.player.Target != (PlayerTargetRefID{}) {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(int(player.player.ID), int(zoneItemID), ElementKindZoneItem, player.player.Path, player_targetIdentifier)
	ref := player.player.engine.createPlayerTargetRef(player.player.Path, player_targetIdentifier, anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID, ElementKindZoneItem, int(zoneItemID))
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.Meta.sign(player.player.engine.broadcastingClientID)
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}
