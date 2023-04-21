package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeImportUpdate() *Factory {
	s.file.Export().Function("import_Update").Param(Param{Id: "current", Type: Id("Tree")}, Param{Id: "update", Type: Id("Tree")}).FuncBody(s.rangeTypes(func(configType ast.ConfigType) *Code {
		if configType.IsEvent {
			return Empty()
		}
		return If(Id("update").Dot(configType.Name).EqualsNot().Null().And().Id("update").Dot(configType.Name).EqualsNot().Undf()).Block(
			If(Id("current").Dot(configType.Name).Equals().Null().Or().Id("current").Dot(configType.Name).Equals().Undf()).Block(
				Id("current").Dot(configType.Name).Assign().Id("{}").Sc(),
			),
			ForIn(Const("id"), Id("update").Dot(configType.Name)).Block(
				If(Id("update").Dot(configType.Name).Index(Id("id")).Dot("operationKind").Equals().Id("OperationKind").Dot("OperationKindDelete")).Block(
					Delete().Id("current").Dot(configType.Name).Index(Id("id")).Sc(),
				).Id(" else").Block(
					Id("current").Dot(configType.Name).Index(Id("id")).Assign().Id("import"+Title(configType.Name)).Call(Id("current").Dot(configType.Name).Index(Id("id")), Id("update").Dot(configType.Name).Index(Id("id"))).Sc(),
				),
			),
		)
	})...)

	return s
}
