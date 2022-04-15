package message

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeParameters() *Factory {
	s.config.RangeActions(func(action ast.Action) {

		p := paramsWriter{
			a: action,
		}

		s.file.Type().Id(p.name()).Struct(
			ForEachParamInAction(action, func(param ast.Field) *Statement {
				p.p = &param
				return Id(p.fieldName()).Add(p.paramType(s)).Id(p.fieldTag())
			}),
		)
	})

	return s
}

type paramsWriter struct {
	a ast.Action
	p *ast.Field
}

func (p paramsWriter) name() string {
	return Title(p.a.Name) + "Params"
}

func (p paramsWriter) fieldName() string {
	return Title(p.p.Name)
}

func (p paramsWriter) paramType(s *Factory) *Statement {
	var optionalIndex string
	if p.p.HasSliceValue {
		optionalIndex += "[]"
	}

	switch {
	case s.isIDTypeOfType(p.p.ValueType().Name):
		return Id(optionalIndex + "state").Dot(Title(p.p.ValueType().Name))
	default:
		return Id(optionalIndex + p.p.ValueType().Name)
	}
}

func (p paramsWriter) fieldTag() string {
	return "`json:\"" + p.p.Name + "\"`"
}
