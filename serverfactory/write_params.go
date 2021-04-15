package serverfactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *ServerFactory) writeParameters() *ServerFactory {
	decls := NewDeclSet()
	s.config.RangeActions(func(action ast.Action) {

		p := paramsWriter{
			a: action,
		}

		decls.File.Type().Id("_" + action.Name + "Params").Struct(ForEachParamInAction(action, func(param ast.Field) *Statement {
			p.p = &param
			return Id(p.fieldName()).Id(p.paramType()).Id(p.fieldTag())
		}))
	})

	decls.Render(s.buf)
	return s
}

type paramsWriter struct {
	a ast.Action
	p *ast.Field
}

func (p paramsWriter) fieldName() string {
	return title(p.p.Name)
}

func (p paramsWriter) paramType() string {
	var s string
	if p.p.HasSliceValue {
		s += "[]"
	}
	if p.p.ValueType.IsBasicType {
		return s + p.p.ValueType.Name
	}
	return s + title(p.p.ValueType.Name)
}

func (p paramsWriter) fieldTag() string {
	return "`json:\"" + p.p.Name + "\"`"
}
