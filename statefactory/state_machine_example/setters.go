package state

func (_e GearScore) SetLevel(sm *Engine, newLevel int) GearScore {
	e := sm.GearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Level = newLevel
	e.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e GearScore) SetScore(sm *Engine, newScore int) GearScore {
	e := sm.GearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Score = newScore
	e.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e Position) SetX(sm *Engine, newX float64) Position {
	e := sm.Position(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.X = newX
	e.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[e.position.ID] = e.position
	return e
}

func (_e Position) SetY(sm *Engine, newY float64) Position {
	e := sm.Position(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.Y = newY
	e.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[e.position.ID] = e.position
	return e
}
