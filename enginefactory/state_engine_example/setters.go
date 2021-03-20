package state

func (_e GearScore) SetLevel(se *Engine, newLevel int) GearScore {
	e := se.GearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Level = newLevel
	e.gearScore.OperationKind = OperationKindUpdate
	se.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e GearScore) SetScore(se *Engine, newScore int) GearScore {
	e := se.GearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Score = newScore
	e.gearScore.OperationKind = OperationKindUpdate
	se.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e Position) SetX(se *Engine, newX float64) Position {
	e := se.Position(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.X = newX
	e.position.OperationKind = OperationKindUpdate
	se.Patch.Position[e.position.ID] = e.position
	return e
}

func (_e Position) SetY(se *Engine, newY float64) Position {
	e := se.Position(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.Y = newY
	e.position.OperationKind = OperationKindUpdate
	se.Patch.Position[e.position.ID] = e.position
	return e
}
