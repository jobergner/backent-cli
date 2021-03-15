package statefactory

import (
	"text/template"
)

const entityKindTemplateString string = `
type EntityKind string

const (<( range .Types )>
	EntityKind<( toTitleCase .Name )> <( writeOnIndex $.Types . 0 "EntityKind" )> = "<( .Name )>"<(end)>
)`

var entityKindTemplate *template.Template = newTemplateFrom("entityKindTemplate", entityKindTemplateString)

func (s *stateFactory) writeEntityKinds() *stateFactory {
	entityKindTemplate.Execute(s.buf, s.ast)
	return s
}

const idTemplateString string = `<( range .Types )>
type <( toTitleCase .Name )>ID int<( end )>
`

var idTemplate *template.Template = newTemplateFrom("idTemplate", idTemplateString)

func (s *stateFactory) writeIDs() *stateFactory {
	idTemplate.Execute(s.buf, s.ast)
	return s
}

const stateTemplateString string = `
type State struct {<( range .Types )>
	<( toTitleCase .Name )> map[<( toTitleCase .Name )>ID]<( .Name )>Core ` + "`" + `json:"<( .Name )>"` + "`" + `
<( end )>}

func newState() State {
	return State{<( range .Types )><( toTitleCase .Name )>: make(map[<( toTitleCase .Name )>ID]<( .Name )>Core)<( doNotWriteOnIndex $.Types . -1 ", ")><( end )>}
}
`

var stateTemplate *template.Template = newTemplateFrom("stateTemplate", stateTemplateString)

func (s *stateFactory) writeState() *stateFactory {
	stateTemplate.Execute(s.buf, s.ast)
	return s
}

const elementTemplateString string = `
<( define "elementFieldValue" )>
	<(- if .HasSliceValue -)>
		[]
	<(- end -)>
	<(- if .ValueType.IsBasicType -)>
		<( .ValueType.Name )>
	<(- else -)>
		<( toTitleCase .ValueType.Name )>ID	
	<(- end -)>
<( end )>
<( range .Types )>
type <( .Name )>Core struct {
	ID <( toTitleCase .Name )>ID ` + "`" + `json:"id"` + "`" + `
<( range .Fields )> <( toTitleCase .Name )> <( template "elementFieldValue" . )>  ` + "`" + `json:"<( .Name )>"` + "`" + `
<( end )>
	OperationKind OperationKind ` + "`" + `json:"operationKind"` + "`" + `
<( if not .IsRootType )> HasParent bool ` + "`" + `json:"hasParent"` + "`" + `<( end )>
}
type <( toTitleCase .Name )> struct{ <( .Name )> <( .Name )>Core }
<( end )>
`

var elementTemplate *template.Template = newTemplateFrom("elementTemplate", elementTemplateString)

func (s *stateFactory) writeElements() *stateFactory {
	elementTemplate.Execute(s.buf, s.ast)
	return s
}
