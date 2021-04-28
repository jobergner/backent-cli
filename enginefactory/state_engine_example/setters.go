package state

func (_gearScore gearScore) SetLevel(newLevel int) gearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_gearScore gearScore) SetScore(newScore int) gearScore {
	gearScore := _gearScore.gearScore.engine.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	gearScore.gearScore.engine.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_position position) SetX(newX float64) position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.X = newX
	position.position.OperationKind_ = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_position position) SetY(newY float64) position {
	position := _position.position.engine.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.Y = newY
	position.position.OperationKind_ = OperationKindUpdate
	position.position.engine.Patch.Position[position.position.ID] = position.position
	return position
}

func (_item item) SetName(nameName string) item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind_ == OperationKindDelete {
		return item
	}
	item.item.Name = nameName
	item.item.OperationKind_ = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}

func (_item item) SetBoundTo(playerID PlayerID) item {
	item := _item.item.engine.Item(_item.item.ID)
	if item.item.OperationKind_ == OperationKindDelete {
		return item
	}
	if item.item.engine.Player(playerID).player.OperationKind_ == OperationKindDelete {
		return item
	}
	ref := item.item.engine.createItemBoundToRef(playerID, item.item.ID)
	item.item.BoundTo = ref.ID
	item.item.OperationKind_ = OperationKindUpdate
	item.item.engine.Patch.Item[item.item.ID] = item.item
	return item
}

func (_equipmentSet equipmentSet) SetName(nameName string) equipmentSet {
	equipmentSet := _equipmentSet.equipmentSet.engine.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind_ == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.Name = nameName
	equipmentSet.equipmentSet.OperationKind_ = OperationKindUpdate
	equipmentSet.equipmentSet.engine.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	return equipmentSet
}
