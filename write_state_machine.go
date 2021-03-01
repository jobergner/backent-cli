package statefactory

import (
	"text/template"
)

const operationKindTemplateString string = `
type OperationKind string

const (
	OperationKindDelete = "DELETE"
	OperationKindUpdate = "UPDATE"
)
`

var operationKindTemplate *template.Template = newTemplateFrom("operationKindTemplate", operationKindTemplateString)

func (s *stateFactory) writeOperationKind() *stateFactory {
	operationKindTemplate.Execute(&s.buf, s.ast)
	return s
}

const stateMachineTemplateString string = `
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
`

var stateMachineTemplate *template.Template = newTemplateFrom("stateMachineTemplate", stateMachineTemplateString)

func (s *stateFactory) writeStateMachine() *stateFactory {
	stateMachineTemplate.Execute(&s.buf, s.ast)
	return s
}

const updateStateTemplateString string = `
func (sm *StateMachine) UpdateState() {
<( range .Decls )>
	for _, <( .Name )> := range sm.Patch.<( toTitleCase .Name )> {
		if <( .Name )>.OperationKind == OperationKindDelete {
			delete(sm.State.<( toTitleCase .Name )>, <( .Name )>.ID)
		} else {
			sm.State.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name )>
		}
	}
<( end )>
	sm.Patch = newState()
}
`

var updateStateTemplate *template.Template = newTemplateFrom("updateStateTemplate", updateStateTemplateString)

func (s *stateFactory) writeUpdateState() *stateFactory {
	updateStateTemplate.Execute(&s.buf, s.ast)
	return s
}
