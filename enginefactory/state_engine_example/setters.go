package state

func (_gearScore gearScore) SetLevel(se *Engine, newLevel int) gearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Level = newLevel
	se.updateGearScore(gearScore.gearScore)
	return gearScore
}

func (_gearScore gearScore) SetScore(se *Engine, newScore int) gearScore {
	gearScore := se.GearScore(_gearScore.gearScore.ID)
	if gearScore.gearScore.OperationKind_ == OperationKindDelete {
		return gearScore
	}
	gearScore.gearScore.Score = newScore
	se.updateGearScore(gearScore.gearScore)
	return gearScore
}

func (_position position) SetX(se *Engine, newX float64) position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.X = newX
	se.updatePosition(position.position)
	return position
}

func (_position position) SetY(se *Engine, newY float64) position {
	position := se.Position(_position.position.ID)
	if position.position.OperationKind_ == OperationKindDelete {
		return position
	}
	position.position.Y = newY
	se.updatePosition(position.position)
	return position
}
