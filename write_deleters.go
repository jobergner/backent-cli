package statefactory

import (
	"text/template"
)

const deleterTemplateString string = `
<( range .Decls )>
func (sm *StateMachine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( if not .IsRootType -)>
		<( .Name )> := sm.Get<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
		if <( .Name )>.HasParent {
			return
		}
	<( end -)>
	sm.delete<( toTitleCase .Name )>(<( .Name )>ID)
}
func (sm *StateMachine) delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( .Name )> := sm.Get<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( .Name )>.OperationKind = OperationKindDelete
	sm.Patch.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name -)>
	<( $Decl := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( $Decl.Name )>.<( toTitleCase .Name )> {
				sm.delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			sm.delete<( toTitleCase .ValueType.Name )>(<( $Decl.Name )>.<( toTitleCase .ValueType.Name )>)
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
