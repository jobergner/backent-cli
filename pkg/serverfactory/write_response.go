package serverfactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeResponses() *ServerFactory {
	decls := NewDeclSet()
	s.config.RangeActions(func(action ast.Action) {

		if action.Response == nil {
			return
		}

		r := responseWriter{
			a: action,
		}

		decls.File.Type().Id(r.name()).Struct(ForEachResponseValueInAction(action, func(value ast.Field) *Statement {
			r.v = &value
			return Id(r.fieldName()).Id(r.paramType(s)).Id(r.fieldTag())
		}))
	})

	decls.Render(s.buf)
	return s
}

type responseWriter struct {
	a ast.Action
	v *ast.Field
}

func (r responseWriter) name() string {
	return Title(r.a.Name) + "Response"
}

func (r responseWriter) fieldName() string {
	return Title(r.v.Name)
}

func (r responseWriter) paramType(s *ServerFactory) string {
	var typeName string
	if r.v.HasSliceValue {
		typeName += "[]"
	}
	if s.isIDTypeOfType(r.v.ValueType().Name) || !r.v.ValueType().IsBasicType {
		return typeName + Title(r.v.ValueType().Name)
	}
	return typeName + r.v.ValueType().Name
}

func (r responseWriter) fieldTag() string {
	return "`json:\"" + r.v.Name + "\"`"
}
