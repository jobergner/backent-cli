package state

func (_gearScore GearScore) SetLevel(newLevel int) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	if gearScore.gearScore.Level == newLevel {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	gearScore.gearScore.OperationKind = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_gearScore GearScore) SetScore(newScore int) GearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	if gearScore.gearScore.Score == newScore {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	gearScore.gearScore.OperationKind = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_position Position) SetX(newX float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	if position.position.X == newX {
		return position
	}
	position.position.X = newX
	position.position.OperationKind = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_position Position) SetY(newY float64) Position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	if position.position.Y == newY {
		return position
	}
	position.position.Y = newY
	position.position.OperationKind = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_item Item) SetName(newName string) Item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.Name == newName {
		return item
	}
	item.item.Name = newName
	item.item.OperationKind = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
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
	if item.item.engine.itemBoundToRef(item.item.BoundTo).itemBoundToRef.ReferencedElementID == playerID {
		return item
	}
	if item.item.BoundTo != 0 {
		item.item.engine.deleteItemBoundToRef(item.item.BoundTo)
	}
	ref := item.item.engine.createItemBoundToRef(playerID, item.item.ID)
	item.item.BoundTo = ref.ID
	item.item.OperationKind = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}

func (_equipmentSet EquipmentSet) SetName(newName string) EquipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	if equipmentSet.equipmentSet.Name == newName {
		return equipmentSet
	}
	equipmentSet.equipmentSet.Name = newName
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
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
	if player.player.engine.anyOfPlayer_ZoneItem(player.player.engine.playerTargetRef(player.player.Target).playerTargetRef.ReferencedElementID).anyOfPlayer_ZoneItem.Player == playerID {
		return player
	}
	if player.player.Target != 0 {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(false, nil)
	anyContainer.anyOfPlayer_ZoneItem.setPlayer(playerID, false)
	ref := player.player.engine.createPlayerTargetRef(anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID)
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
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
	if player.player.engine.anyOfPlayer_ZoneItem(player.player.engine.playerTargetRef(player.player.Target).playerTargetRef.ReferencedElementID).anyOfPlayer_ZoneItem.ZoneItem == zoneItemID {
		return player
	}
	if player.player.Target != 0 {
		player.player.engine.deletePlayerTargetRef(player.player.Target)
	}
	anyContainer := player.player.engine.createAnyOfPlayer_ZoneItem(false, nil)
	anyContainer.anyOfPlayer_ZoneItem.setZoneItem(zoneItemID, false)
	ref := player.player.engine.createPlayerTargetRef(anyContainer.anyOfPlayer_ZoneItem.ID, player.player.ID)
	player.player.Target = ref.ID
	player.player.OperationKind = OperationKindUpdate
	player.player.engine.Patch.Player[player.player.ID] = player.player
	return player
}
