package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeOperationKind() *Factory {

	s.file.Type().Id("OperationKind").String()

	s.file.Const().Defs(
		Id("OperationKindDelete").Id("OperationKind").Op("=").Lit("DELETE"),
		Id("OperationKindUpdate").Id("OperationKind").Op("=").Lit("UPDATE"),
		Id("OperationKindUnchanged").Id("OperationKind").Op("=").Lit("UNCHANGED"),
	)

	return s
}

func (s *Factory) writeEngine() *Factory {

	s.file.Type().Id("Engine").Struct(
		Id("State").Id("*State"),
		Id("Patch").Id("*State"),
		Id("Tree").Id("*Tree"),
		Id("BroadcastingClientID").Id("string"),
		Id("ThisClientID").Id("string"),
		Id("planner").Id("*assemblePlanner"),
		Id("IDgen").Int(),
	)

	s.file.Func().Id("NewEngine").Params().Id("*Engine").Block(
		Return(Id("&Engine").Values(Dict{
			Id("State"):   Id("newState").Call(),
			Id("Patch"):   Id("newState").Call(),
			Id("Tree"):    Id("newTree").Call(),
			Id("planner"): Id("newAssemblePlanner").Call(),
			Id("IDgen"):   Lit(1),
		})),
	)

	return s
}

func (s *Factory) writeGenerateID() *Factory {

	s.file.Func().Params(Id("engine").Id("*Engine")).Id("GenerateID").Params().Int().Block(
		Id("newID").Op(":=").Id("engine").Dot("IDgen"),
		Id("engine").Dot("IDgen").Op("=").Id("engine").Dot("IDgen").Op("+").Lit(1),
		Return(Id("newID")),
	)

	return s
}

func writeImportPatchElement(typeName string) *Statement {
	return For(List(Id("_"), Id(typeName)).Op(":=").Range().Id("patch").Dot(Title(typeName))).Block(
		Id("engine").Dot("Patch").Dot(Title(typeName)).Index(Id(typeName).Dot("ID")).Op("=").Id(typeName),
	)
}

func (s *Factory) writeImportPatch() *Factory {
	s.file.Func().Params(Id("engine").Id("*Engine")).Id("ImportPatch").Params(Id("patch").Id("*State")).Block(
		ForEachBasicType(func(b BasicType) *Statement {
			return writeImportPatchElement(b.Value)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			return writeImportPatchElement(configType.Name)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			return writeImportPatchElement(ValueTypeName(&field))
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			return writeImportPatchElement(AnyValueTypeName(&field))
		}),
	)

	return s
}

func (s *Factory) writeUpdateState() *Factory {

	u := updateStateWriter{}

	s.file.Func().Params(u.receiverParams()).Id("UpdateState").Params().Block(
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			if !configType.IsEvent {
				return Empty()
			}
			u.typeName = func() string {
				return configType.Name
			}
			return writeDeleteEvent(u)
		}),
		ForEachBasicType(func(b BasicType) *Statement {
			u.typeName = func() string {
				return b.Value
			}
			return writeUpdateElement(u)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			u.t = &configType
			defer func() { u.t = nil }()
			return writeUpdateElement(u)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return ValueTypeName(&field)
			}
			return writeUpdateElement(u)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return AnyValueTypeName(&field)
			}
			return writeUpdateElement(u)
		}),
		ForEachBasicType(func(b BasicType) *Statement {
			u.typeName = func() string {
				return b.Value
			}
			return writeClearPatch(u)
		}),
		ForEachTypeInAST(s.config, func(configType ast.ConfigType) *Statement {
			u.typeName = func() string {
				return configType.Name
			}
			return writeClearPatch(u)
		}),
		ForEachRefFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return ValueTypeName(&field)
			}
			return writeClearPatch(u)
		}),
		ForEachAnyFieldInAST(s.config, func(field ast.Field) *Statement {
			u.typeName = func() string {
				return AnyValueTypeName(&field)
			}
			return writeClearPatch(u)
		}),
	)

	return s
}

func writeUpdateElement(u updateStateWriter) *Statement {
	return For(u.loopPatchElementsConditions()).Block(
		If(u.isOperationKindDelete()).Block(
			u.deleteElement(),
		).Else().Block(
			OnlyIf(u.t != nil, u.emptyEvents()),
			u.setOperationKindUnchanged(),
			u.updateElement(),
		),
	)
}

func writeClearPatch(u updateStateWriter) *Statement {
	return For(u.loopPatchKeysConditions()).Block(
		u.clearElementFromPatch(),
	)
}

func writeDeleteEvent(u updateStateWriter) *Statement {
	return For(u.loopPatchElementsConditions()).Block(
		Id("engine").Dot("delete" + Title(u.typeName())).Call(Id(u.typeName()).Dot("ID")),
	)
}
