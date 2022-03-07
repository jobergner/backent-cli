package engine

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

type updateStateWriter struct {
	typeName func() string
	t        *ast.ConfigType
}

func (u updateStateWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (u updateStateWriter) loopPatchElementsConditions() *Statement {
	return List(Id("_"), Id(u.typeName())).Op(":=").Range().Id("engine").Dot("Patch").Dot(Title(u.typeName()))
}

func (u updateStateWriter) isOperationKindDelete() *Statement {
	return Id(u.typeName()).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (u updateStateWriter) deleteElement() *Statement {
	return Id("delete").Call(Id("engine").Dot("State").Dot(Title(u.typeName())), Id(u.typeName()).Dot("ID"))
}

func (u updateStateWriter) setOperationKindUnchanged() *Statement {
	return Id(u.typeName()).Dot("OperationKind").Op("=").Id("OperationKindUnchanged")
}

func (u updateStateWriter) unsignElement() *Statement {
	return Id(u.typeName()).Dot("Meta").Dot("unsign").Call()
}

func (u updateStateWriter) emptyEvents() *Statement {
	if u.t == nil {
		return Empty()
	}

	return ForEachFieldInType(*u.t, func(field ast.Field) *Statement {
		if !field.ValueType().IsEvent {
			return Empty()
		}
		return Id(u.typeName()).Dot(Title(field.Name)).Op("=").Id(u.typeName()).Dot(Title(field.Name)).Index(Empty(), Lit(0))
	})
}

func (u updateStateWriter) updateElement() *Statement {
	return Id("engine").Dot("State").Dot(Title(u.typeName())).Index(Id(u.typeName()).Dot("ID")).Op("=").Id(u.typeName())
}

func (u updateStateWriter) loopPatchKeysConditions() *Statement {
	return List(Id("key")).Op(":=").Range().Id("engine").Dot("Patch").Dot(Title(u.typeName()))
}

func (u updateStateWriter) clearElementFromPatch() *Statement {
	return Id("delete").Call(Id("engine").Dot("Patch").Dot(Title(u.typeName())), Id("key"))
}
