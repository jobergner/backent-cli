package enginefactory

import (
	"text/template"
)

const deleterTemplateString string = `
<( range .Types )>
func (se *Engine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( if not .IsRootType -)>
		<( .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
		if <( .Name )>.HasParent {
			return
		}
	<( end -)>
	se.delete<( toTitleCase .Name )>(<( .Name )>ID)
}
func (se *Engine) delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( .Name )>.OperationKind = OperationKindDelete
	se.Patch.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name -)>
	<( $Type := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( $Type.Name )>.<( toTitleCase .Name )> {
				se.delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			se.delete<( toTitleCase .ValueType.Name )>(<( $Type.Name )>.<( toTitleCase .ValueType.Name )>)
		<(- end -)>
	<(- end )><(- end )>
}
<( end )>
`

var deleterTemplate *template.Template = newTemplateFrom("deleterTemplate", deleterTemplateString)

func (s *stateFactory) writeDeleters() *stateFactory {
	deleterTemplate.Execute(s.buf, s.ast)
	return s
}
