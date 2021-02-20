package statemachine

type OperationKind string

type EntityKind string

const (
	OperationKindDelete = "DELETE"
	OperationKindUpdate = "UPDATE"
)

type StateMachine struct {
	State State
	Patch State
	IDgen int
}

func (sm *StateMachine) GenerateID() int {
	newID := sm.IDgen
	sm.IDgen = sm.IDgen + 1
	return newID
}

func (sm *StateMachine) UpdateState() {
	for _, player := range sm.Patch.Player {
		if player.OperationKind == OperationKindDelete {
			delete(sm.State.Player, player.ID)
		} else {
			sm.State.Player[player.ID] = player
		}
	}
	for _, zone := range sm.Patch.Zone {
		if zone.OperationKind == OperationKindDelete {
			delete(sm.State.Zone, zone.ID)
		} else {
			sm.State.Zone[zone.ID] = zone
		}
	}
	for _, zoneItem := range sm.Patch.ZoneItem {
		if zoneItem.OperationKind == OperationKindDelete {
			delete(sm.State.ZoneItem, zoneItem.ID)
		} else {
			sm.State.ZoneItem[zoneItem.ID] = zoneItem
		}
	}
	for _, position := range sm.Patch.Position {
		if position.OperationKind == OperationKindDelete {
			delete(sm.State.Position, position.ID)
		} else {
			sm.State.Position[position.ID] = position
		}
	}
	for _, item := range sm.Patch.Item {
		if item.OperationKind == OperationKindDelete {
			delete(sm.State.Item, item.ID)
		} else {
			sm.State.Item[item.ID] = item
		}
	}
	for _, gearScore := range sm.Patch.GearScore {
		if gearScore.OperationKind == OperationKindDelete {
			delete(sm.State.GearScore, gearScore.ID)
		} else {
			sm.State.GearScore[gearScore.ID] = gearScore
		}
	}
	sm.Patch = newState()
}
