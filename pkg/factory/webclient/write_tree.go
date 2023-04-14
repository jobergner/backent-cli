package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeTree() *Factory {

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
