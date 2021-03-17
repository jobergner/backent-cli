package statefactory

import (
	"text/template"
)

const assembleTreeTemplateString string = `
func (sm *StateMachine) assembleTree() Tree {
	tree := newTree()
	<(- range .Types )>
	<( if .IsRootType -)>
	for _, <( .Name )> := range sm.Patch.<( toTitleCase .Name )> {
		tree<( toTitleCase .Name )>, hasUpdated := sm.assemble<( toTitleCase .Name )>(<( .Name )>.ID)
		if hasUpdated {
			tree.<( toTitleCase .Name )>[<( .Name )>.ID] = tree<( toTitleCase .Name )>
		}
	}
	<(- else -)>
	for _, <( .Name )> := range sm.Patch.<( toTitleCase .Name )> {
		if !<( .Name )>.HasParent {
			tree<( toTitleCase .Name )>, hasUpdated := sm.assemble<( toTitleCase .Name )>(<( .Name )>.ID)
			if hasUpdated {
				tree.<( toTitleCase .Name )>[<( .Name )>.ID] = tree<( toTitleCase .Name )>
			}
		}
	}
	<(- end )>
	<(- end )>
	<(- range .Types )>
	<(- if .IsRootType )>
	for _, <( .Name )> := range sm.State.<( toTitleCase .Name )> {
		if _, ok := tree.<( toTitleCase .Name )>[<( .Name )>.ID]; !ok {
			tree<( toTitleCase .Name )>, hasUpdated := sm.assemble<( toTitleCase .Name )>(<( .Name )>.ID)
			if hasUpdated {
				tree.<( toTitleCase .Name )>[<( .Name )>.ID] = tree<( toTitleCase .Name )>
			}
		}
	}
	<(- else )>
	for _, <( .Name )> := range sm.State.<( toTitleCase .Name )> {
		if !<( .Name )>.HasParent {
			if _, ok := tree.<( toTitleCase .Name )>[<( .Name )>.ID]; !ok {
				tree<( toTitleCase .Name )>, hasUpdated := sm.assemble<( toTitleCase .Name )>(<( .Name )>.ID)
				if hasUpdated {
					tree.<( toTitleCase .Name )>[<( .Name )>.ID] = tree<( toTitleCase .Name )>
				}
			}
		}
	}
	<(- end )>
	<(- end )>
	return tree
}
`

var assembleTreeTemplate *template.Template = newTemplateFrom("assembleTreeTemplate", assembleTreeTemplateString)

func (s *stateFactory) writeAssembleTree() *stateFactory {
	assembleTreeTemplate.Execute(s.buf, s.ast)
	return s
}

const assembleTreeElementTemplateString string = `
<(- range .Types )><( $Type := . )>
func (sm *StateMachine) assemble<( toTitleCase .Name )>(<( .Name )>ID <( toTitleCase .Name )>ID) (_<( .Name )>, bool) {
	<( .Name )>, hasUpdated := sm.Patch.<( toTitleCase .Name )>[<( .Name )>ID]
	if !hasUpdated {
		<( if .IsLeafType -)>
			return _<( .Name )>{}, false
		<(- else -)>
			<( .Name )> = sm.State.<( toTitleCase .Name )>[<( .Name )>ID]
		<(- end )>
	}
	var tree<( toTitleCase .Name )> _<( .Name )><( range .Fields -)>
	<( if not .ValueType.IsBasicType -)>
		<( if .HasSliceValue )>
			for _, <( .ValueType.Name )>ID := range deduplicate<( toTitleCase .ValueType.Name )>IDs(sm.State.<( toTitleCase $Type.Name )>[<( $Type.Name )>.ID].<( toTitleCase .Name )>, sm.Patch.<( toTitleCase $Type.Name )>[<( $Type.Name )>.ID].<( toTitleCase .Name )>) {
				if tree<( toTitleCase .ValueType.Name )>, <( .ValueType.Name )>HasUpdated := sm.assemble<( toTitleCase .ValueType.Name )>(<( .ValueType.Name )>ID); <( .ValueType.Name )>HasUpdated {
					hasUpdated = true
					tree<( toTitleCase $Type.Name )>.<( toTitleCase .Name )> = append(tree<( toTitleCase $Type.Name )>.<( toTitleCase .Name )>, tree<( toTitleCase .ValueType.Name )>)
				}
			}
		<(- else )>
			if tree<( toTitleCase .Name )>, <( .Name )>HasUpdated := sm.assemble<( toTitleCase .Name )>(<( $Type.Name )>.<( toTitleCase .Name )>); <( .Name )>HasUpdated {
				hasUpdated = true
				tree<( toTitleCase $Type.Name )>.<( toTitleCase .Name )> = &tree<( toTitleCase .ValueType.Name )>
			}
		<(- end -)>
	<(- end )>
	<(- end )>
	tree<( toTitleCase .Name )>.ID = <( .Name )>.ID
	tree<( toTitleCase .Name )>.OperationKind = <( .Name )>.OperationKind
	<(- range .Fields )>
	<(- if .ValueType.IsBasicType )>
		tree<( toTitleCase $Type.Name )>.<( toTitleCase .Name )> = <( $Type.Name )>.<( toTitleCase .Name )>
	<(- end -)>
	<(- end )>
	return tree<( toTitleCase .Name )>, <( if .IsLeafType )>true<( else )>hasUpdated<( end )>
}
<(- end )>
`

var assembleTreeElementTemplate *template.Template = newTemplateFrom("assembleTreeElementTemplate", assembleTreeElementTemplateString)

func (s *stateFactory) writeAssembleTreeElement() *stateFactory {
	assembleTreeElementTemplate.Execute(s.buf, s.ast)
	return s
}
