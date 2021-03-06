package statefactory

import (
	"text/template"
)

const setterTemplateString string = `
<( range .Decls )><( $Decl := . )><( range .Fields )><( if not .HasSliceValue )><( if .ValueType.IsBasicType)>
func (_e <( toTitleCase $Decl.Name )>) Set<( toTitleCase .Name )>(sm *StateMachine, new<( toTitleCase .Name )> <( .ValueType.Name )>) <( toTitleCase $Decl.Name )> {
	e := sm.Get<( toTitleCase $Decl.Name )>(_e.<( $Decl.Name )>.ID)
	if e.<( $Decl.Name )>.OperationKind == OperationKindDelete {
		return e
	}
	e.<( $Decl.Name )>.<( toTitleCase .Name )> = new<( toTitleCase .Name )>
	e.<( $Decl.Name )>.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase $Decl.Name )>[e.<( $Decl.Name )>.ID] = e.<( $Decl.Name )>
	return e
}
<( end )><( end )>
<( end )><( end )>
`

var setterTemplate *template.Template = newTemplateFrom("setterTemplate", setterTemplateString)

func (s *stateFactory) writeSetters() *stateFactory {
	setterTemplate.Execute(&s.buf, s.ast)
	return s
}
