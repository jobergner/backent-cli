package statemachine

type operationKind string

type personID int
type nameID int
type childID int

type entityKind string

const (
	entityKindPerson entityKind = "person"
	entityKindName              = "name"
	entityKindChild             = "child"
)

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
	for _, child := range sm.patch.child {
		if child.operationKind == operationKindDelete {
			delete(sm.state.child, child.id)
		} else {
			sm.state.child[child.id] = child
		}
	}
	for _, name := range sm.patch.name {
		if name.operationKind == operationKindDelete {
			delete(sm.state.name, name.id)
		} else {
			sm.state.name[name.id] = name
		}
	}
	for _, person := range sm.patch.person {
		if person.operationKind == operationKindDelete {
			delete(sm.state.person, person.id)
		} else {
			sm.state.person[person.id] = person
		}
	}
	sm.patch = newState()
}
