package enginefactory

import (
	"bar-cli/ast"

	. "github.com/dave/jennifer/jen"
)

func (s *stateFactory) writeOperationKind() *stateFactory {
	decls := newDeclSet()

	decls.file.Type().Id("OperationKind").String()

	decls.file.Const().Defs(
		Id("OperationKindDelete").Id("OperationKind").Op("=").Lit("DELETE"),
		Id("OperationKindUpdate").Op("=").Lit("UPDATE"),
	)

	decls.render(s.buf)
	return s
}

func (s *stateFactory) writeEngine() *stateFactory {
	decls := newDeclSet()

	decls.file.Type().Id("Engine").Struct(
		Id("State").Id("State"),
		Id("Patch").Id("State"),
		Id("IDgen").Int(),
	)

	decls.file.Func().Id("newEngine").Params().Id("*Engine").Block(
		Return(Id("&Engine").Values(Dict{
			Id("State"): Id("newState").Call(),
			Id("Patch"): Id("newState").Call(),
			Id("IDgen"): Lit(1),
		})),
	)

	decls.render(s.buf)
	return s
}

func (s *stateFactory) writeGenerateID() *stateFactory {
	decls := newDeclSet()

	decls.file.Func().Params(Id("se").Id("*Engine")).Id("GenerateID").Params().Int().Block(
		Id("newID").Op(":=").Id("se").Dot("IDgen"),
		Id("se").Dot("IDgen").Op("=").Id("se").Dot("IDgen").Op("+").Lit(1),
		Return(Id("newID")),
	)

	decls.render(s.buf)
	return s
}

func (s *stateFactory) writeUpdateState() *stateFactory {
	decls := newDeclSet()

	u := updateState{}

	decls.file.Func().Params(u.receiverParams()).Id("UpdateState").Params().Block(
		forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.t = &configType
			return For(u.loopPatchElementsConditions()).Block(
				If(u.isOperationKindDelete()).Block(
					u.deleteElement(),
				).Else().Block(
					u.updateElement(),
				),
			)
		}),
		forEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.t = &configType
			return For(u.loopPatchKeysConditions()).Block(
				u.clearElementFromPatch(),
			)
		}),
	)

	decls.render(s.buf)
	return s
}

type updateState struct {
	t *ast.ConfigType
}

func (u updateState) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (u updateState) loopPatchElementsConditions() *Statement {
	return List(Id("_"), Id(u.t.Name)).Op(":=").Range().Id("se").Dot("Patch").Dot(title(u.t.Name))
}

func (u updateState) isOperationKindDelete() *Statement {
	return Id(u.t.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
}

func (u updateState) deleteElement() *Statement {
	return Id("delete").Call(Id("se").Dot("State").Dot(title(u.t.Name)), Id(u.t.Name).Dot("ID"))
}

func (u updateState) updateElement() *Statement {
	return Id("se").Dot("State").Dot(title(u.t.Name)).Index(Id(u.t.Name).Dot("ID")).Op("=").Id(u.t.Name)
}

func (u updateState) loopPatchKeysConditions() *Statement {
	return List(Id("key")).Op(":=").Range().Id("se").Dot("Patch").Dot(title(u.t.Name))
}

func (u updateState) clearElementFromPatch() *Statement {
	return Id("delete").Call(Id("se").Dot("Patch").Dot(title(u.t.Name)), Id("key"))
}
