package enginefactory

import (
	. "github.com/dave/jennifer/jen"
	"text/template"
)

func (s *stateFactory) writeTree() *stateFactory {
	decls := newDeclSet()
	decls.file.Type().Id("Tree").Struct(forEachTypeInAST(s.ast, func(configType stateConfigType) *Statement {
		s := treeWriter{configType}
		return Id(s.fieldName()).Map(s.mapKey()).Id(s.mapValue()).Id(s.fieldTag()).Line()
	}))

	decls.file.Func().Id("newTree").Params().Id("Tree").Block(
		Return(Id("Tree").Values(forEachTypeInAST(s.ast, func(configType stateConfigType) *Statement {
			s := treeWriter{configType}
			return Id(s.fieldName()).Id(":").Make(Map(s.mapKey()).Id(s.mapValue())).Id(",")
		}))),
	)

	decls.render(s.buf)
	return s
}

type treeWriter struct {
	t stateConfigType
}

func (s treeWriter) fieldName() string {
	return title(s.t.Name)
}

func (s treeWriter) mapKey() *Statement {
	return Id(title(s.t.Name) + "ID")
}

func (s treeWriter) mapValue() string {
	return "t" + title(s.t.Name)
}

func (s treeWriter) fieldTag() string {
	return "`json:\"" + s.t.Name + "\"`"
}

const treeElementTemplateString string = `
<( range .Types )>
type _<( .Name )> struct {
	ID <( toTitleCase .Name )>ID ` + "`" + `json:"id"` + "`" + `
	<(- range .Fields )>
	<( toTitleCase .Name )><( print " " )> 
	<(- if not .ValueType.IsBasicType -)>
		<( if not .HasSliceValue -)>
			*
		<(- end )>
	<(- end )>
	<(- if .HasSliceValue -)>
		[]
	<(- end -)>
	<( if not .ValueType.IsBasicType -)>
		_
	<(- end )>
	<(- .ValueType.Name )> ` + "`" + `json:"<( .Name )>"` + "`" + `
<( end )>
	OperationKind_ OperationKind ` + "`" + `json:"operationKind_"` + "`" + `
}
<( end )>
`

var treeElementTemplate *template.Template = newTemplateFrom("treeElementTemplate", treeElementTemplateString)

func (s *stateFactory) writeTreeElements() *stateFactory {
	decls := newDeclSet()

	s.ast.rangeTypes(func(configType stateConfigType) {

		e := treeElementWriter{
			t: configType,
		}

		decls.file.Type().Id(e.name()).Struct(
			Id("ID").Id(e.idType()).Id(e.metaFieldTag("id")).Line(),
			forEachFieldInType(configType, func(field stateConfigField) *Statement {
				e.f = &field
				return Id(e.fieldName()).Id(e.fieldValue()).Id(e.fieldTag()).Line()
			}),
			Id("OperationKind_").Id("OperationKind").Id(e.metaFieldTag("operationKind_")).Line(),
		)
	})

	decls.render(s.buf)
	return s
}

type treeElementWriter struct {
	t stateConfigType
	f *stateConfigField
}

func (e treeElementWriter) fieldValue() string {
	if e.f.ValueType.IsBasicType {
		if e.f.HasSliceValue {
			return "[]" + e.f.ValueType.Name
		}
		return e.f.ValueType.Name
	}

	if e.f.HasSliceValue {
		return "[]" + "t" + title(e.f.ValueType.Name)
	}

	return "*" + "t" + title(e.f.ValueType.Name)
}

func (e treeElementWriter) fieldTag() string {
	return "`json:\"" + e.f.Name + "\"`"
}

func (e treeElementWriter) metaFieldTag(name string) string {
	return "`json:\"" + name + "\"`"
}

func (e treeElementWriter) fieldName() string {
	return title(e.f.Name)
}

func (e treeElementWriter) name() string {
	return "t" + title(e.t.Name)
}

func (e treeElementWriter) idType() string {
	return title(e.t.Name) + "ID"
}
