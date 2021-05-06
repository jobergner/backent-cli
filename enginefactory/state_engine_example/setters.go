package state

func (_gearScore gearScore) SetLevel(newLevel int) gearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	gearScore.gearScore.OperationKind = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_gearScore gearScore) SetScore(newScore int) gearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	gearScore.gearScore.OperationKind = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_position position) SetX(newX float64) position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.X = newX
	position.position.OperationKind = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_position position) SetY(newY float64) position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind == OperationKindDelete {
		return position
	}
	position.position.Y = newY
	position.position.OperationKind = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_item item) SetName(nameName string) item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	item.item.Name = nameName
	item.item.OperationKind = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}

func (_item item) SetBoundTo(playerID PlayerID) item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind == OperationKindDelete {
		return item
	}
	if item.item.engine.Player(playerID).player.OperationKind == OperationKindDelete {
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

func (_equipmentSet equipmentSet) SetName(nameName string) equipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.Name = nameName
	equipmentSet.equipmentSet.OperationKind = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	return equipmentSet
}

// func (_player player) SetTarget(playerID PlayerID) player {
// 	player := _player.player.engine.Player(_player.player.ID)
// 	if player.player.OperationKind == OperationKindDelete {
// 		return player
// 	}
// 	if player.player.engine.Player(playerID).player.OperationKind == OperationKindDelete {
// 		return player
// 	}
// 	ref := player.player.engine.createPlayerBoundToRef(playerID, player.player.ID)
// 	player.player.BoundTo = ref.ID
// 	player.player.OperationKind = OperationKindUpdate
// 	player.player.engine.Patch.Player[player.player.ID] = player.player
// 	return player
// }
