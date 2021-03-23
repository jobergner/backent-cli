package enginefactory

import (
	"bytes"
	. "github.com/dave/jennifer/jen"
)

func abc(ast *stateConfigAST, buf *bytes.Buffer) {
	decls := newDeclSet()
	ast.orderedRange(func(configType stateConfigType) {
		configType.orderedRange(func(field stateConfigField) {

			if field.HasSliceValue || !field.ValueType.IsBasicType {
				return
			}

			receiverParams := Id("_e").Id(title(configType.Name))
			funcName := "Set" + title(field.Name)
			params := List(
				Id("se").Id("*Engine"),
				Id("new"+title(field.Name)).Id(field.ValueType.Name),
			)
			returns := title(configType.Name)
			isOperationKindDelete := Id("e").Dot(configType.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
			elementID := Id("e").Dot(configType.Name).Dot("ID")

			decls.file.Func().Params(receiverParams).Id(funcName).Params(params).Id(returns).Block(
				Id("e").Op(":=").Id("se").Dot(title(configType.Name)).Params(Id("_e").Dot(configType.Name).Dot("ID")),
				If(isOperationKindDelete).Block(
					Return(Id("e")),
				),
				Id("e").Dot(configType.Name).Dot(title(field.Name)).Op("=").Id("new"+title(field.Name)),
				Id("e").Dot(configType.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate"),
				Id("se").Dot("Patch").Dot(title(configType.Name)).Index(elementID).Op("=").Id("e").Dot(configType.Name),
				Return(Id("e")),
			)
		})
	})

	decls.render(buf)
}

func (s *stateFactory) writeSetters() *stateFactory {
	abc(s.ast, s.buf)
	return s
}
