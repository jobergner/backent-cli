package statemachine

func (g gearScore) SetLevel(newLevel int, sm *stateMachine) gearScore {
	g.level = newLevel
	g.operationKind = operationKindUpdate
	sm.patch.gearScore[g.id] = g
	return g
}

func (g gearScore) SetScore(newScore int, sm *stateMachine) gearScore {
	g.score = newScore
	g.operationKind = operationKindUpdate
	sm.patch.gearScore[g.id] = g
	return g
}

func (p position) SetX(newX float64, sm *stateMachine) position {
	p.x = newX
	p.operationKind = operationKindUpdate
	sm.patch.position[p.id] = p
	return p
}

func (p position) SetY(newY float64, sm *stateMachine) position {
	p.x = newY
	p.operationKind = operationKindUpdate
	sm.patch.position[p.id] = p
	return p
}
