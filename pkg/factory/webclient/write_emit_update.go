package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeEmitUpdate() *Factory {
	s.file.Export().Function("emit_Update").Param(Param{Id: "update", Type: Id("Tree")}).FuncBody(s.rangeTypes(func(configType ast.ConfigType) *Code {
		return If(Id("update").Dot(configType.Name).EqualsNot().Null().And().Id("update").Dot(configType.Name).EqualsNot().Undf()).Block(
			ForIn(Const("id"), Id("update").Dot(configType.Name)).Block(
				Id("emit" + Title(configType.Name)).Call(Id("update").Dot(configType.Name).Index(Id("id"))).Sc(),
			),
		)
	})...)

	return s
}
