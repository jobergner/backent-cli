package statemachine

type operationKind string

type entityKind string

const (
	operationKindDelete = "DELETE"
	operationKindUpdate = "UPDATE"
)

type stateMachine struct {
	state state
	patch state
	idgen int
}

func (sm *stateMachine) generateID() int {
	newID := sm.idgen
	sm.idgen = sm.idgen + 1
	return newID
}

func (sm *stateMachine) updateState() {
	for _, player := range sm.patch.player {
		if player.operationKind == operationKindDelete {
			delete(sm.state.player, player.id)
		} else {
			sm.state.player[player.id] = player
		}
	}
	for _, zone := range sm.patch.zone {
		if zone.operationKind == operationKindDelete {
			delete(sm.state.zone, zone.id)
		} else {
			sm.state.zone[zone.id] = zone
		}
	}
	for _, zoneItem := range sm.patch.zoneItem {
		if zoneItem.operationKind == operationKindDelete {
			delete(sm.state.zoneItem, zoneItem.id)
		} else {
			sm.state.zoneItem[zoneItem.id] = zoneItem
		}
	}
	for _, position := range sm.patch.position {
		if position.operationKind == operationKindDelete {
			delete(sm.state.position, position.id)
		} else {
			sm.state.position[position.id] = position
		}
	}
	for _, item := range sm.patch.item {
		if item.operationKind == operationKindDelete {
			delete(sm.state.item, item.id)
		} else {
			sm.state.item[item.id] = item
		}
	}
	for _, gearScore := range sm.patch.gearScore {
		if gearScore.operationKind == operationKindDelete {
			delete(sm.state.gearScore, gearScore.id)
		} else {
			sm.state.gearScore[gearScore.id] = gearScore
		}
	}
	sm.patch = newState()
}
