package statefactory

import (
	"text/template"
)

const getterTemplateString string = `
<( define "returnValue" -)>
	<( if .HasSliceValue -)>
		[]
	<(- end -)>
	<( if .ValueType.IsBasicType -)>
		<( .ValueType.Name )>
	<(- else -)>
		<( toTitleCase .ValueType.Name )>
	<(- end )>
<(- end )>
<( range .Decls )>
func (sm *StateMachine) Get<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) <( toTitleCase .Name )> {
	patching<( toTitleCase .Name )>, ok := sm.Patch.<( toTitleCase .Name )>[<( .Name )>ID]
	if ok {
		return <( toTitleCase .Name )>{patching<( toTitleCase .Name )>}
	}
	current<( toTitleCase .Name )> := sm.State.<( toTitleCase .Name )>[<( .Name )>ID]
	return <( toTitleCase .Name )>{current<( toTitleCase .Name )>}
}
<( $Decl := . )><( range .Fields )>
func (_e <( toTitleCase $Decl.Name )>) Get<( toTitleCase .Name )>(sm *StateMachine) <( template "returnValue" . )> {
	e := sm.Get<( toTitleCase $Decl.Name )>(_e.<( $Decl.Name )>.ID)
	<( if .HasSliceValue -)>
		var <( .Name )> <( template "returnValue" . )>
		for _,<( print " " )>
		<(- if .ValueType.IsBasicType -)>
			element
		<(- else -)>
			<( .ValueType.Name )>ID
		<(- end -)>
		<( print " " )>:= range e.<( $Decl.Name )>.<( toTitleCase .Name )> {
			<( .Name )> = append(<( .Name )>,<( print " " )>
			<(- if .ValueType.IsBasicType -)>
				element
			<(- else -)>
				sm.Get<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID)
			<(- end -)>
			)
		}
		return <( .Name )>
	<(- else -)>
		<( if .ValueType.IsBasicType -)>
			return e.<( $Decl.Name )>.<( toTitleCase .Name )>
		<(- else -)>
			return sm.Get<( toTitleCase .Name )>(e.<( $Decl.Name )>.<( toTitleCase .Name )>)
		<(- end -)>
	<(- end )>
}
<( end )>
<( end )>
`

var getterTemplate *template.Template = newTemplateFrom("getterTemplate", getterTemplateString)

func (s *stateFactory) writeGetters() *stateFactory {
	getterTemplate.Execute(s.buf, s.ast)
	return s
}
