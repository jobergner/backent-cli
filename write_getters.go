package statefactory

import (
	"text/template"
)

const getterTemplateString string = `
<( range .Decls )>
func (sm *StateMachine) Get<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) <( toTitleCase .Name )> {
	patching<( toTitleCase .Name )>, ok := sm.Patch.<( toTitleCase .Name )>[<( .Name )>ID]
	if !ok {
		return <( toTitleCase .Name )>{patching<( toTitleCase .Name )>} 
	}
	current<( toTitleCase .Name )> := sm.State.<( toTitleCase .Name )>[<( .Name )>ID]
	return <( toTitleCase .Name )>{current<( toTitleCase .Name )>} 
}
<( $Decl := . )><( range .Fields )>
func (_e <( toTitleCase $Decl.Name )>) Get<( toTitleCase .Name )>(sm *StateMachine) <( if .ValueType.IsBasicType )><( .ValueType.Name )><( else )><( toTitleCase .ValueType.Name )><( end )> {
	e := sm.Get<( toTitleCase $Decl.Name )>(_e.<( $Decl.Name )>.ID)
	if e.<( $Decl.Name )>.OperationKind == OperationKindDelete {
		return<( if not .ValueType.IsBasicType )> <( toTitleCase .ValueType.Name )>{}<( end )>
	}
	<( if not .ValueType.IsBasicType )><( .ValueType.Name )> := sm.create<( toTitleCase .ValueType.Name )>(true)
	<( end )>e.<( $Decl.Name )>.<( toTitleCase .Name )> = append(e.<( $Decl.Name )>.<( toTitleCase .Name )>, <(if .ValueType.IsBasicType )><( .Name )>...<( else )><( .ValueType.Name )>.<( .ValueType.Name )>.ID<( end )>)
	e.<( $Decl.Name )>.OperationKind = OperationKindUpdate
	sm.Patch.<( toTitleCase $Decl.Name )>[e.<( $Decl.Name )>.ID] = e.<( $Decl.Name )><( if not .ValueType.IsBasicType )>
	return <( .ValueType.Name )><( end )>
}
<( end )><( end )>
`

var getterTemplate *template.Template = newTemplateFrom("getterTemplate", getterTemplateString)

func (s *stateFactory) writeGetters() *stateFactory {
	getterTemplate.Execute(&s.buf, s.ast)
	return s
}
