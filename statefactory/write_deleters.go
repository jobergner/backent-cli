package statefactory

import (
	"text/template"
)

const deleterTemplateString string = `
<( range .Types )>
func (sm *StateMachine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( if not .IsRootType -)>
		<( .Name )> := sm.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
		if <( .Name )>.HasParent {
			return
		}
	<( end -)>
	sm.delete<( toTitleCase .Name )>(<( .Name )>ID)
}
func (sm *StateMachine) delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( .Name )> := sm.<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( .Name )>.OperationKind = OperationKindDelete
	sm.Patch.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name -)>
	<( $Type := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( $Type.Name )>.<( toTitleCase .Name )> {
				sm.delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			sm.delete<( toTitleCase .ValueType.Name )>(<( $Type.Name )>.<( toTitleCase .ValueType.Name )>)
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
