package enginefactory

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

const EngineTemplateString string = `
type Engine struct {
	State State
	Patch State
	IDgen int
}

func newEngine() *Engine {
	return &Engine{State: newState(), Patch: newState(), IDgen: 1}
}
`

var EngineTemplate *template.Template = newTemplateFrom("EngineTemplate", EngineTemplateString)

func (s *stateFactory) writeEngine() *stateFactory {
	EngineTemplate.Execute(s.buf, s.ast)
	return s
}

const generateIDTemplateString string = `
func (se *Engine) GenerateID() int {
	newID := se.IDgen
	se.IDgen = se.IDgen + 1
	return newID
}
`

var generateIDTemplate *template.Template = newTemplateFrom("generateIDTemplate", generateIDTemplateString)

func (s *stateFactory) writeGenerateID() *stateFactory {
	generateIDTemplate.Execute(s.buf, s.ast)
	return s
}

const updateStateTemplateString string = `
func (se *Engine) UpdateState() {
<( range .Types )>
	for _, <( encrypt .Name )> := range se.Patch.<( toTitleCase .Name )> {
		if <( encrypt .Name )>.OperationKind == OperationKindDelete {
			delete(se.State.<( toTitleCase .Name )>, <( encrypt .Name )>.ID)
		} else {
			se.State.<( toTitleCase .Name )>[<( encrypt .Name )>.ID] = <( encrypt .Name )>
		}
	}
<( end )>
<( range .Types )>
	for key := range se.Patch.<( toTitleCase .Name )> {
		delete(se.Patch.<( toTitleCase .Name )>, key)
	}
<(- end )>
}
`

var updateStateTemplate *template.Template = newTemplateFrom("updateStateTemplate", updateStateTemplateString)

func (s *stateFactory) writeUpdateState() *stateFactory {
	updateStateTemplate.Execute(s.buf, s.ast)
	return s
}
