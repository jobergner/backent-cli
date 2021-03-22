package enginefactory

import (
	"text/template"
)

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
	OperationKind_ OperationKind ` + "`" + `json:"operationKind_"` + "`" + `
<( if not .IsRootType )> HasParent_ bool ` + "`" + `json:"hasParent_"` + "`" + `<( end )>
}
type <( toTitleCase .Name )> struct{ <( .Name )> <( .Name )>Core }
<( end )>
`

var elementTemplate *template.Template = newTemplateFrom("elementTemplate", elementTemplateString)

func (s *stateFactory) writeElements() *stateFactory {
	elementTemplate.Execute(s.buf, s.ast)
	return s
}
