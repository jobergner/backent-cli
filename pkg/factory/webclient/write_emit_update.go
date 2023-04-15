package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeEmitUpdate() *Factory {
	s.file.Export().Function("emit_Update").Param(Param{Id: "update", Type: Id("Tree")}).Block(s.rangeTypes(func(configType ast.ConfigType) *Code {
		return If(Id("update").Dot(configType.Name).EqualsNot().Null().And().Id("update").Dot(configType.Name).EqualsNot().Undf()).Block(
			ForIn(Const("id"), Id("update").Dot(configType.Name)).Block(
				Id("em")
			),
		)
	})...)

	var fields []InterfaceField

	s.config.RangeTypes(func(configType ast.ConfigType) {
		fields = append(fields, InterfaceField{
			Optional: true,
			Name:     configType.Name,
			Type: ObjectSpaced(ObjectField{
				Id:   Index(Id("id").Is(Id("number"))),
				Type: Id(Title(configType.Name)),
			}),
		})
	})

	s.file.Export().Interface("Tree", fields...)

	return s
}

func (s *Factory) rangeTypes(fn func(configType ast.ConfigType) *Code) []*Code {
	var code []*Code
	s.config.RangeTypes(func(configType ast.ConfigType) {
		code = append(code, fn(configType))
	})
	return code
}
