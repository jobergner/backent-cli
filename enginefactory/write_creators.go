package enginefactory

import (
	"text/template"
)

const creatorTemplateString string = `
<(- range .Types )>
func (se *Engine) Create<( toTitleCase .Name )>() <( toTitleCase .Name )> {
	return se.create<( toTitleCase .Name )>(<( if not .IsRootType )>false<( end )>)
}
func (se *Engine) create<( toTitleCase .Name )>(<( if not .IsRootType )>hasParent bool<( end )>) <( toTitleCase .Name )> {
	var e <( .Name )>Core
	e.ID = <( toTitleCase .Name )>ID(se.GenerateID())
	<(- if not .IsRootType )>
		e.HasParent = hasParent
	<(- end )>
	<(- range .Fields -)>
		<(- if not .HasSliceValue )><( if not .ValueType.IsBasicType )>
			element<( toTitleCase .ValueType.Name )> := se.create<( toTitleCase .ValueType.Name )>(true)
			e.<( toTitleCase .Name )> = element<( toTitleCase .ValueType.Name )>.<( .ValueType.Name )>.ID
		<(- end )><( end -)>
	<( end )>
	e.OperationKind = OperationKindUpdate
	se.Patch.<( toTitleCase .Name )>[e.ID] = e
	return <( toTitleCase .Name )>{<( .Name )>: e}
}
<( end )>
`

var creatorTemplate *template.Template = newTemplateFrom("creatorTemplate", creatorTemplateString)

func (s *stateFactory) writeCreators() *stateFactory {
	creatorTemplate.Execute(s.buf, s.ast)
	return s
}
