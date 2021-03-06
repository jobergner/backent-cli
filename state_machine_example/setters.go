package statemachine

func (_e GearScore) SetLevel(sm *StateMachine, newLevel int) GearScore {
	e := sm.GetGearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Level = newLevel
	e.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e GearScore) SetScore(sm *StateMachine, newScore int) GearScore {
	e := sm.GetGearScore(_e.gearScore.ID)
	if e.gearScore.OperationKind == OperationKindDelete {
		return e
	}
	e.gearScore.Score = newScore
	e.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[e.gearScore.ID] = e.gearScore
	return e
}

func (_e Position) SetX(sm *StateMachine, newX float64) Position {
	e := sm.GetPosition(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.X = newX
	e.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[e.position.ID] = e.position
	return e
}

func (_e Position) SetY(sm *StateMachine, newY float64) Position {
	e := sm.GetPosition(_e.position.ID)
	if e.position.OperationKind == OperationKindDelete {
		return e
	}
	e.position.Y = newY
	e.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[e.position.ID] = e.position
	return e
}
