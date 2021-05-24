package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeOperationKind() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("OperationKind").String()

	decls.File.Const().Defs(
		Id("OperationKindDelete").Id("OperationKind").Op("=").Lit("DELETE"),
		Id("OperationKindUpdate").Id("OperationKind").Op("=").Lit("UPDATE"),
		Id("OperationKindUnchanged").Id("OperationKind").Op("=").Lit("UNCHANGED"),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeEngine() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Type().Id("Engine").Struct(
		Id("State").Id("State"),
		Id("Patch").Id("State"),
		Id("Tree").Id("Tree"),
		Id("PathTrack").Id("pathTrack"),
		Id("IDgen").Int(),
	)

	decls.File.Func().Id("newEngine").Params().Id("*Engine").Block(
		Return(Id("&Engine").Values(Dict{
			Id("State"):     Id("newState").Call(),
			Id("Patch"):     Id("newState").Call(),
			Id("Tree"):      Id("newTree").Call(),
			Id("PathTrack"): Id("newPathTrack").Call(),
			Id("IDgen"):     Lit(1),
		})),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeGenerateID() *EngineFactory {
	decls := NewDeclSet()

	decls.File.Func().Params(Id("engine").Id("*Engine")).Id("GenerateID").Params().Int().Block(
		Id("newID").Op(":=").Id("engine").Dot("IDgen"),
		Id("engine").Dot("IDgen").Op("=").Id("engine").Dot("IDgen").Op("+").Lit(1),
		Return(Id("newID")),
	)

	decls.Render(s.buf)
	return s
}

func (s *EngineFactory) writeUpdateState() *EngineFactory {
	decls := NewDeclSet()

	u := updateStateWriter{}

	decls.File.Func().Params(u.receiverParams()).Id("UpdateState").Params().Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			return writeUpdateElement(u)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return field.ValueTypeName
			}
			return writeUpdateElement(u)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return anyNameByField(field)
			}
			return writeUpdateElement(u)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			return For(u.loopPatchKeysConditions()).Block(
				u.clearElementFromPatch(),
			)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return field.ValueTypeName
			}
			return For(u.loopPatchKeysConditions()).Block(
				u.clearElementFromPatch(),
			)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return anyNameByField(field)
			}
			return For(u.loopPatchKeysConditions()).Block(
				u.clearElementFromPatch(),
			)
		}),
	)

	decls.Render(s.buf)
	return s
}

func writeUpdateElement(u updateStateWriter) *Statement {
	return For(u.loopPatchElementsConditions()).Block(
		If(u.isOperationKindDelete()).Block(
			u.deleteElement(),
		).Else().Block(
			u.setOperationKindUnchanged(),
			u.updateElement(),
		),
	)
}

type updateStateWriter struct {
	typeName func() string
}

func (u updateStateWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (u updateStateWriter) loopPatchElementsConditions() *Statement {
	return List(Id("_"), Id(u.typeName())).Op(":=").Range().Id("engine").Dot("Patch").Dot(title(u.typeName()))
}

func (u updateStateWriter) isOperationKindDelete() *Statement {
	return Id(u.typeName()).Dot("OperationKind").Op("==").Id("OperationKindDelete")
}

func (u updateStateWriter) deleteElement() *Statement {
	return Id("delete").Call(Id("engine").Dot("State").Dot(title(u.typeName())), Id(u.typeName()).Dot("ID"))
}

func (u updateStateWriter) setOperationKindUnchanged() *Statement {
	return Id(u.typeName()).Dot("OperationKind").Op("=").Id("OperationKindUnchanged")
}

func (u updateStateWriter) updateElement() *Statement {
	return Id("engine").Dot("State").Dot(title(u.typeName())).Index(Id(u.typeName()).Dot("ID")).Op("=").Id(u.typeName())
}

func (u updateStateWriter) loopPatchKeysConditions() *Statement {
	return List(Id("key")).Op(":=").Range().Id("engine").Dot("Patch").Dot(title(u.typeName()))
}

func (u updateStateWriter) clearElementFromPatch() *Statement {
	return Id("delete").Call(Id("engine").Dot("Patch").Dot(title(u.typeName())), Id("key"))
}
