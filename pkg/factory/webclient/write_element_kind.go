package webclient

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) writeElementKind() *Factory {

	var enumFields []EnumField

	enumFields = append(enumFields, EnumField{
		Name:  "ElementKind_Root",
		Value: Id(Lit("root")),
	})

	s.config.RangeTypes(func(configType ast.ConfigType) {
		enumFields = append(enumFields, EnumField{
			Name:  "ElementKind" + Title(configType.Name),
			Value: Id(Lit(Title(configType.Name))),
		})
	})

	s.file.Export().Enum("ElementKind", enumFields...)

	return s
}
