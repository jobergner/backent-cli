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

			receiverParams := Id("_" + configType.Name).Id(title(configType.Name))
			funcName := "Set" + title(field.Name)
			newValueParam := "new" + title(field.Name)
			params := List(
				Id("se").Id("*Engine"),
				Id(newValueParam).Id(field.ValueType.Name),
			)
			returns := title(configType.Name)
			isOperationKindDelete := Id(configType.Name).Dot(configType.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
			elementID := Id(configType.Name).Dot(configType.Name).Dot("ID")

			reassignElement := Id(configType.Name).Op(":=").Id("se").Dot(title(configType.Name)).Params(Id("_" + configType.Name).Dot(configType.Name).Dot("ID"))
			setAttribute := Id(configType.Name).Dot(configType.Name).Dot(title(field.Name)).Op("=").Id(newValueParam)
			setOpKind := Id(configType.Name).Dot(configType.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
			updateElementInPatch := Id("se").Dot("Patch").Dot(title(configType.Name)).Index(elementID).Op("=").Id(configType.Name).Dot(configType.Name)

			decls.file.Func().Params(receiverParams).Id(funcName).Params(params).Id(returns).Block(
				reassignElement,
				If(isOperationKindDelete).Block(
					Return(Id(configType.Name)),
				),
				setAttribute,
				setOpKind,
				updateElementInPatch,
				Return(Id(configType.Name)),
			)
		})
	})

	decls.render(buf)
}

func (s *stateFactory) writeSetters() *stateFactory {
	abc(s.ast, s.buf)
	return s
}
