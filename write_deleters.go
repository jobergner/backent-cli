package statefactory

import (
	"text/template"
)

const deleterTemplateString string = `
<( range .Decls )>
func (sm *StateMachine) Delete<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) {
	<( .Name )> := sm.Get<( toTitleCase .Name )>(<( .Name )>ID).<( .Name )>
	<( if not .IsRootType -)>
	if <( .Name )>.HasParent {
		return
	}
	<( end -)>
	<( .Name )>.OperationKind = OperationKindDelete
	sm.Patch.<( toTitleCase .Name )>[<( .Name )>.ID] = <( .Name -)>
	<( $Decl := . )>
	<(- range .Fields )><( if not .ValueType.IsBasicType )>
		<(- if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range <( $Decl.Name )>.<( toTitleCase .Name )> {
				sm.Delete<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			}
		<(- else )>
			sm.Delete<( toTitleCase .ValueType.Name )>(<( $Decl.Name )>.<( toTitleCase .ValueType.Name )>)
		<(- end -)>
	<(- end )><(- end )>
}
<( end )>
`

var deleterTemplate *template.Template = newTemplateFrom("deleterTemplate", deleterTemplateString)

func (s *stateFactory) writeDeleters() *stateFactory {
	deleterTemplate.Execute(&s.buf, s.ast)
	return s
}
