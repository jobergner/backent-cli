package statefactory

import (
	"text/template"
)

const creatorTemplateString string = `
<(- range .Decls )>
func (sm *StateMachine) Create<( toTitleCase .Name )>() <( toTitleCase .Name )> {
	return sm.create<( toTitleCase .Name )>(<( if not .IsRootType )>false<( end )>)
}
func (sm *StateMachine) create<( toTitleCase .Name )>(<( if not .IsRootType )>hasParent bool<( end )>) <( toTitleCase .Name )> {
	var e <( .Name )>Core
	e.ID = <( toTitleCase .Name )>ID(sm.GenerateID())
	<(- if not .IsRootType )>
		e.HasParent = hasParent
	<(- end )>
	<(- range .Fields -)>
		<(- if not .HasSliceValue )><( if not .ValueType.IsBasicType )>
			element<( toTitleCase .ValueType.Name )> := sm.create<( toTitleCase .ValueType.Name )>(true)
			e.<( toTitleCase .Name )> = element<( toTitleCase .ValueType.Name )>.<( .ValueType.Name )>.ID
		<(- end )><( end -)>
	<( end )>
	e.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase .Name )>[e.ID] = e
	return <( toTitleCase .Name )>{<( .Name )>: e}
}
<( end )>
`

var creatorTemplate *template.Template = newTemplateFrom("creatorTemplate", creatorTemplateString)

func (s *stateFactory) writeCreators() *stateFactory {
	creatorTemplate.Execute(&s.buf, s.ast)
	return s
}
