package state

func (_gearScore gearScore) SetLevel(se *Engine, newLevel int) gearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_gearScore gearScore) SetScore(se *Engine, newScore int) gearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	gearScore.gearScore.OperationKind_ = OperationKindUpdate
	se.Patch.GearScore[gearScore.gearScore.ID] = gearScore.gearScore
	return gearScore
}

func (_position position) SetX(se *Engine, newX float64) position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.X = newX
	position.position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.position.ID] = position.position
	return position
}

func (_position position) SetY(se *Engine, newY float64) position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.Y = newY
	position.position.OperationKind_ = OperationKindUpdate
	se.Patch.Position[position.position.ID] = position.position
	return position
}

func (_item item) SetName(se *Engine, nameName string) item {
	item := se.Item(_item.item.ID)
	if item.item.OperationKind_ == OperationKindDelete {
		return item
	}
	item.item.Name = nameName
	item.item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.item.ID] = item.item
	return item
}

func (_item item) SetBoundTo(se *Engine, playerID PlayerID) item {
	item := se.Item(_item.item.ID)
	if item.item.OperationKind_ == OperationKindDelete {
		return item
	}
	ref := se.createItemBoundToRef()
	item.item.BoundTo = ref.ID
	item.item.OperationKind_ = OperationKindUpdate
	se.Patch.Item[item.item.ID] = item.item
	return item
}

func (_equipmentSet equipmentSet) SetName(se *Engine, nameName string) equipmentSet {
	equipmentSet := se.EquipmentSet(_equipmentSet.equipmentSet.ID)
	if equipmentSet.equipmentSet.OperationKind_ == OperationKindDelete {
		return equipmentSet
	}
	equipmentSet.equipmentSet.Name = nameName
	equipmentSet.equipmentSet.OperationKind_ = OperationKindUpdate
	se.Patch.EquipmentSet[equipmentSet.equipmentSet.ID] = equipmentSet.equipmentSet
	return equipmentSet
}
