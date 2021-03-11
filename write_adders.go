package statefactory

import (
	"text/template"
)

const adderTemplateString string = `
<( range .Decls )><( $Decl := . )><( range .Fields )><( if .HasSliceValue )>
func (_e <( toTitleCase $Decl.Name )>) Add
<(- if .ValueType.IsBasicType -)>
	<( toTitleCase .Name )>
<(- else -)>
	<( .Name | toSingular | toTitleCase )>
<(- end -)>
(sm *StateMachine<(if .ValueType.IsBasicType )>, <( .Name)> ...<( .ValueType.Name )><( end )>) <( if not .ValueType.IsBasicType )><( toTitleCase .ValueType.Name )><( end )> {
	e := sm.Get<( toTitleCase $Decl.Name )>(_e.<( $Decl.Name )>.ID)
	if e.<( $Decl.Name )>.OperationKind == OperationKindDelete {
		return<( if not .ValueType.IsBasicType )> <( toTitleCase .ValueType.Name )>{<( .ValueType.Name )>Core{OperationKind: OperationKindDelete}}<( end )>
	}
	<(- if not .ValueType.IsBasicType )>
		<( .ValueType.Name )> := sm.create<( toTitleCase .ValueType.Name )>(true)
	<(- end )>
	e.<( $Decl.Name )>.<( toTitleCase .Name )> = append(e.<( $Decl.Name )>.<( toTitleCase .Name )>,<(print " ")>
	<(- if .ValueType.IsBasicType -)>
		<( .Name )>...
	<(- else -)>
		<( .ValueType.Name )>.<( .ValueType.Name )>.ID
	<(- end -)>
	)
	e.<( $Decl.Name )>.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase $Decl.Name )>[e.<( $Decl.Name )>.ID] = e.<( $Decl.Name )><( if not .ValueType.IsBasicType )>
	return <( .ValueType.Name )><( end )>
}
<( end )>
<( end )><( end )>
`

var adderTemplate *template.Template = newTemplateFrom("adderTemplate", adderTemplateString)

func (s *stateFactory) writeAdders() *stateFactory {
	adderTemplate.Execute(s.buf, s.ast)
	return s
}
