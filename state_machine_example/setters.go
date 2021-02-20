package statemachine

func (g GearScore) SetLevel(newLevel int, sm *StateMachine) GearScore {
	g.gearScore.Level = newLevel
	g.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.gearScore.ID] = g.gearScore
	return g
}

func (g GearScore) SetScore(newScore int, sm *StateMachine) GearScore {
	g.gearScore.Score = newScore
	g.gearScore.OperationKind = OperationKindUpdate
	sm.Patch.GearScore[g.gearScore.ID] = g.gearScore
	return g
}

func (p Position) SetX(newX float64, sm *StateMachine) Position {
	p.position.X = newX
	p.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.position.ID] = p.position
	return p
}

func (p Position) SetY(newY float64, sm *StateMachine) Position {
	p.position.X = newY
	p.position.OperationKind = OperationKindUpdate
	sm.Patch.Position[p.position.ID] = p.position
	return p
}
