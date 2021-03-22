package enginefactory

import (
	"text/template"
)

const treeTemplateString string = `
type Tree struct {<( range .Types )>
	<( toTitleCase .Name )> map[<( toTitleCase .Name )>ID]_<( .Name )> ` + "`" + `json:"<( .Name )>"` + "`" + `
<( end )>}

func newTree() Tree {
	return Tree{
	<(- range .Types -)>
	<( toTitleCase .Name )>: make(map[<( toTitleCase .Name )>ID]_<( .Name )>)<( doNotWriteOnIndex $.Types . -1 ", ")>
	<(- end -)>
	}
}
`

var treeTemplate *template.Template = newTemplateFrom("treeTemplate", treeTemplateString)

func (s *stateFactory) writeTree() *stateFactory {
	treeTemplate.Execute(s.buf, s.ast)
	return s
}

const treeElementTemplateString string = `
<( range .Types )>
type _<( .Name )> struct {
	ID <( toTitleCase .Name )>ID ` + "`" + `json:"id"` + "`" + `
	<(- range .Fields )>
	<( toTitleCase .Name )><( print " " )> 
	<(- if not .ValueType.IsBasicType -)>
		<( if not .HasSliceValue -)>
			*
		<(- end )>
	<(- end )>
	<(- if .HasSliceValue -)>
		[]
	<(- end -)>
	<( if not .ValueType.IsBasicType -)>
		_
	<(- end )>
	<(- .ValueType.Name )> ` + "`" + `json:"<( .Name )>"` + "`" + `
<( end )>
	OperationKind_ OperationKind ` + "`" + `json:"operationKind_"` + "`" + `
}
<( end )>
`

var treeElementTemplate *template.Template = newTemplateFrom("treeElementTemplate", treeElementTemplateString)

func (s *stateFactory) writeTreeElements() *stateFactory {
	treeElementTemplate.Execute(s.buf, s.ast)
	return s
}
