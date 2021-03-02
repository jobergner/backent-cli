package statemachine

func (_g GearScore) SetLevel(newLevel int, sm *StateMachine) GearScore {
	g := sm.GetGearScore(_g.gearScore.ID)
	if g.gearScore.OperationKind == OperationKindDelete {
		return g
	}
	g.gearScore.Level = newLevel
	g.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.gearScore.ID] = g.gearScore
	return g
}

func (_g GearScore) SetScore(newScore int, sm *StateMachine) GearScore {
	g := sm.GetGearScore(_g.gearScore.ID)
	if g.gearScore.OperationKind == OperationKindDelete {
		return g
	}
	g.gearScore.Score = newScore
	g.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.gearScore.ID] = g.gearScore
	return g
}

func (_p Position) SetX(newX float64, sm *StateMachine) Position {
	p := sm.GetPosition(_p.position.ID)
	if p.position.OperationKind == OperationKindDelete {
		return p
	}
	p.position.X = newX
	p.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.position.ID] = p.position
	return p
}

func (_p Position) SetY(newY float64, sm *StateMachine) Position {
	p := sm.GetPosition(_p.position.ID)
	if p.position.OperationKind == OperationKindDelete {
		return p
	}
	p.position.X = newY
	p.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.position.ID] = p.position
	return p
}
