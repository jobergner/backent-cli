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
	operationKindTemplate.Execute(s.buf, s.ast)
	return s
}

const stateMachineTemplateString string = `
type StateMachine struct {
	State State
	Patch State
	IDgen int
}

func newStateMachine() *StateMachine {
	return &StateMachine{State: newState(), Patch: newState(), IDgen: 1}
}
`

var stateMachineTemplate *template.Template = newTemplateFrom("stateMachineTemplate", stateMachineTemplateString)

func (s *stateFactory) writeStateMachine() *stateFactory {
	stateMachineTemplate.Execute(s.buf, s.ast)
	return s
}

const generateIDTemplateString string = `
func (sm *StateMachine) GenerateID() int {
	newID := sm.IDgen
	sm.IDgen = sm.IDgen + 1
	return newID
}
`

var generateIDTemplate *template.Template = newTemplateFrom("generateIDTemplate", generateIDTemplateString)

func (s *stateFactory) writeGenerateID() *stateFactory {
	generateIDTemplate.Execute(s.buf, s.ast)
	return s
}

const updateStateTemplateString string = `
func (sm *StateMachine) UpdateState() {
<( range .Types )>
	for _, <( .Name )> := range sm.Patch.<( toTitleCase .Name )> {
		if <( .Name )>.OperationKind == OperationKindDelete {
			delete(sm.State.<( toTitleCase .Name )>, <( .Name )>.ID)
		} else {
			sm.State.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name )>
		}
	}
<( end )>
<( range .Types )>
	for key := range sm.Patch.<( toTitleCase .Name )> {
		delete(sm.Patch.<( toTitleCase .Name )>, key)
	}
<(- end )>
}
`

var updateStateTemplate *template.Template = newTemplateFrom("updateStateTemplate", updateStateTemplateString)

func (s *stateFactory) writeUpdateState() *stateFactory {
	updateStateTemplate.Execute(s.buf, s.ast)
	return s
}
