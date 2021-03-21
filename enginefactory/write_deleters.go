package enginefactory

import (
	"text/template"
)

const deleterTemplateString string = `
<( range .Types )>
func (se *Engine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( if not .IsRootType -)>
		<( encrypt .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
		if <( encrypt .Name )>.HasParent {
			return
		}
	<( end -)>
	se.delete<( toTitleCase .Name )>(<( .Name )>ID)
}
func (se *Engine) delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( encrypt .Name )> := se.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( encrypt .Name )>.OperationKind = OperationKindDelete
	se.Patch.<( toTitleCase .Name )>[<( encrypt .Name )>.ID] = <( encrypt .Name -)>
	<( $Type := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( encrypt $Type.Name )>.<( toTitleCase .Name )> {
				se.delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			se.delete<( toTitleCase .ValueType.Name )>(<( encrypt $Type.Name )>.<( toTitleCase .ValueType.Name )>)
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
