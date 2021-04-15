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
			return Id(p.fieldName()).Id(p.paramType(s)).Id(p.fieldTag())
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

func (p paramsWriter) paramType(s *ServerFactory) string {
	var typeName string
	if p.p.HasSliceValue {
		typeName += "[]"
	}
	if s.isIDTypeOfType(p.p.ValueType.Name) || !p.p.ValueType.IsBasicType {
		return typeName + title(p.p.ValueType.Name)
	}
	return typeName + p.p.ValueType.Name
}

func (p paramsWriter) fieldTag() string {
	return "`json:\"" + p.p.Name + "\"`"
}
