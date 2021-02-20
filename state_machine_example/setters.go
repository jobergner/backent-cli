package statemachine

func (g GearScore) SetLevel(newLevel int, sm *StateMachine) GearScore {
	g.Level = newLevel
	g.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.ID] = g
	return g
}

func (g GearScore) SetScore(newScore int, sm *StateMachine) GearScore {
	g.Score = newScore
	g.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.ID] = g
	return g
}

func (p Position) SetX(newX float64, sm *StateMachine) Position {
	p.X = newX
	p.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.ID] = p
	return p
}

func (p Position) SetY(newY float64, sm *StateMachine) Position {
	p.X = newY
	p.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.ID] = p
	return p
}
