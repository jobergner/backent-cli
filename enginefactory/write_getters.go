package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeGetters() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {
		t := typeGetter{
			t: configType,
		}

		decls.File.Func().Params(t.receiverParams()).Id(t.name()).Params(t.params()).Id(t.returns()).Block(
			t.definePatchingElement(),
			If(Id("ok")).Block(
				Return(t.earlyReturnPatching()),
			),
			t.defineCurrentElement(),
			If(Id("ok")).Block(
				Return(t.earlyReturnCurrent()),
			),
			Return(t.finalReturn()),
		)

		i := idGetter{
			t: configType,
		}

		decls.File.Func().Params(i.receiverParams()).Id(i.name()).Params(i.params()).Id(i.returns()).Block(
			Return(i.returnID()),
		)

		configType.RangeFields(func(field ast.Field) {
			f := fieldGetter{
				t: configType,
				f: field,
			}

			decls.File.Func().Params(f.receiverParams()).Id(f.name()).Params(f.params()).Id(f.returns()).Block(
				f.reassignElement(),
				// if slice
				onlyIf(field.HasSliceValue, f.declareSliceOfElements()),
				onlyIf(field.HasSliceValue, For(f.loopConditions().Block(
					f.appendElement(),
				))),
				onlyIf(field.HasSliceValue, Return(f.returnSlice())),
				// if not slice
				onlyIf(!field.HasSliceValue, Return(f.returnElement())),
			)
		})

	})

	decls.Render(s.buf)
	return s
}

type typeGetter struct {
	t ast.ConfigType
}

func (t typeGetter) receiverParams() *Statement {
	return Id("se").Id("*Engine")
}

func (t typeGetter) name() string {
	return title(t.t.Name)
}

func (t typeGetter) idParam() string {
	return t.t.Name + "ID"
}

func (t typeGetter) params() *Statement {
	return Id(t.t.Name + "ID").Id(title(t.t.Name) + "ID")
}

func (t typeGetter) returns() string {
	return t.t.Name
}

func (t typeGetter) definePatchingElement() *Statement {
	return List(Id("patching"+title(t.t.Name)), Id("ok")).Op(":=").Id("se").Dot("Patch").Dot(title(t.t.Name)).Index(Id(t.idParam()))
}

func (t typeGetter) earlyReturnPatching() *Statement {
	return Id(t.t.Name).Values(Dict{Id(t.t.Name): Id("patching" + title(t.t.Name))})
}

func (t typeGetter) defineCurrentElement() *Statement {
	return List(Id("current"+title(t.t.Name)), Id("ok")).Op(":=").Id("se").Dot("State").Dot(title(t.t.Name)).Index(Id(t.idParam()))
}

func (t typeGetter) earlyReturnCurrent() *Statement {
	return Id(t.t.Name).Values(Dict{Id(t.t.Name): Id("current" + title(t.t.Name))})
}

func (t typeGetter) finalReturn() *Statement {
	return Id(t.t.Name).Values(Dict{Id(t.t.Name): Id(t.t.Name + "Core").Values(Dict{Id("OperationKind_"): Id("OperationKindDelete")})})
}

type idGetter struct {
	t ast.ConfigType
}

func (i idGetter) receiverName() string {
	return "_" + i.t.Name
}

func (i idGetter) receiverParams() *Statement {
	return Id(i.receiverName()).Id(i.t.Name)
}

func (i idGetter) name() string {
	return "ID"
}

func (i idGetter) params() *Statement {
	return Id("se").Id("*Engine")
}

func (i idGetter) returns() string {
	return title(i.t.Name) + "ID"
}

func (i idGetter) returnID() *Statement {
	return Id(i.receiverName()).Dot(i.t.Name).Dot("ID")
}

type fieldGetter struct {
	t ast.ConfigType
	f ast.Field
}

func (f fieldGetter) receiverName() string {
	return "_" + f.t.Name
}

func (f fieldGetter) receiverParams() *Statement {
	return Id(f.receiverName()).Id(f.t.Name)
}

func (f fieldGetter) name() string {
	return title(f.f.Name)
}

func (f fieldGetter) params() *Statement {
	return Id("se").Id("*Engine")
}

func (f fieldGetter) returns() string {
	var val string

	if f.f.HasSliceValue {
		val = "[]"
	}

	if f.f.ValueType().IsBasicType {
		return val + f.f.ValueType().Name
	}
	return val + f.f.ValueType().Name
}

func (f fieldGetter) reassignElement() *Statement {
	return Id(f.t.Name).Op(":=").Id("se").Dot(title(f.t.Name)).Call(Id(f.receiverName()).Dot(f.t.Name).Dot("ID"))
}

func (f fieldGetter) declareSliceOfElements() *Statement {
	return Var().Id(f.f.Name).Id(f.returns())
}

func (f fieldGetter) loopedElementIdentifier() string {
	if f.f.ValueType().IsBasicType {
		return "element"
	}
	return f.f.ValueType().Name + "ID"
}

func (f fieldGetter) loopConditions() *Statement {
	identifier := f.loopedElementIdentifier()
	return List(Id("_"), Id(identifier)).Op(":=").Range().Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))
}

func (f fieldGetter) appendedItem() *Statement {
	if f.f.ValueType().IsBasicType {
		return Id(f.loopedElementIdentifier())
	}
	return Id("se").Dot(title(f.f.ValueType().Name)).Call(Id(f.f.ValueType().Name + "ID"))
}

func (f fieldGetter) appendElement() *Statement {
	return Id(f.f.Name).Op("=").Append(Id(f.f.Name), f.appendedItem())
}

func (f fieldGetter) returnSlice() *Statement {
	return Id(f.f.Name)
}

func (f fieldGetter) returnBasicType() *Statement {
	return Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))
}

func (f fieldGetter) returnType() *Statement {
	return Id("se").Dot(title(f.f.Name)).Call(Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name)))
}

func (f fieldGetter) returnElement() *Statement {
	if f.f.ValueType().IsBasicType {
		return f.returnBasicType()
	}
	return f.returnType()
}
