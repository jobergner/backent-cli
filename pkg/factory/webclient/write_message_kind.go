package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeMessageKind() *Factory {

	var enumFields []EnumField

	enumFields = append(enumFields,
		EnumField{
			Name:  "ID",
			Value: Id(Lit("id")),
		},
		EnumField{
			Name:  "Error",
			Value: Id(Lit("error")),
		},
		EnumField{
			Name:  "Update",
			Value: Id(Lit("update")),
		},
		EnumField{
			Name:  "CurrentState",
			Value: Id(Lit("currentState")),
		},
		EnumField{
			Name:  "Global",
			Value: Id(Lit("global")),
		},
	)

	s.config.RangeActions(func(action ast.Action) {
		enumFields = append(enumFields, EnumField{
			Name:  "Action" + Title(action.Name),
			Value: Id(Lit(action.Name)),
		})
	})

	s.file.Export().Enum("MessageKind", enumFields...)

	return s
}
