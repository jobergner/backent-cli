package actionsfactory

import (
	"text/template"
)

const actionTemplateString string = `
<(- define "parameter" -)>
	<(- if .IsBasicType -)>
		<( .Name )> <( if .IsSliceValue )>[]<( end )><( .TypeLiteral )>
	<(- else -)>
		<( .Name )> <( if .IsSliceValue )>[]<( end )>statemachine.<( toTitleCase .TypeLiteral )>
	<(- end -)>
<(- end -)>
<( range .Actions )><( $Action := . )>
func <( .Name )>(<( range .Params )><( template "parameter" . )><( doNotWriteOnIndex $Action.Params . -1 ", " )><( end )>) {}
<(- end -)>
`

var actionTemplate *template.Template = newTemplateFrom("actionTemplate", actionTemplateString)

func (a *actionsFactory) writeActions() *actionsFactory {
	actionTemplate.Execute(a.buf, a.ast)
	return a
}
