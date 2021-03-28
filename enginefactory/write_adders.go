package enginefactory

import (
	. "github.com/dave/jennifer/jen"
)

func (s *stateFactory) writeAdders() *stateFactory {
	decls := newDeclSet()
	s.ast.rangeTypes(func(configType stateConfigType) {
		configType.rangeFields(func(field stateConfigField) {

			if !field.HasSliceValue {
				return
			}

			a := adder{
				t: configType,
				f: field,
			}

			decls.file.Func().Params(a.receiverParams()).Id(a.name()).Params(a.params()).Id(a.returns()).Block(
				a.reassignElement(),
				If(a.isOperationKindDelete()).Block(
					Return(a.earlyReturn()),
				),
				onlyIf(!field.ValueType.IsBasicType, a.createNewElement()),
				a.appendElement(),
				a.setOperationKindUpdate(),
				a.updateElementInPatch(),
				onlyIf(!field.ValueType.IsBasicType, Return(Id(field.ValueType.Name))),
			)
		})
	})

	decls.render(s.buf)
	return s
}

type adder struct {
	t stateConfigType
	f stateConfigField
}

func (a adder) receiverParams() *Statement {
	return Id(a.receiverName()).Id(title(a.t.Name))
}

func (a adder) name() string {
	if a.f.ValueType.IsBasicType {
		return "Add" + title(a.f.Name)
	}
	return "Add" + title(pluralizeClient.Singular(a.f.Name))
}

func (a adder) params() *Statement {
	params := Id("se").Id("*Engine")
	if a.f.ValueType.IsBasicType {
		return List(params, Id(a.f.Name).Id("..."+a.f.ValueType.Name))
	}
	return params
}

func (a adder) returns() string {
	if a.f.ValueType.IsBasicType {
		return ""
	}
	return title(a.f.ValueType.Name)
}

func (a adder) reassignElement() *Statement {
	return Id(a.t.Name).Op(":=").Id("se").Dot(title(a.t.Name)).Call(Id(a.receiverName()).Dot(a.t.Name).Dot("ID"))
}

func (a adder) isOperationKindDelete() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind_").Op("==").Id("OperationKindDelete")
}

func (a adder) earlyReturn() *Statement {
	if a.f.ValueType.IsBasicType {
		return Empty()
	}
	return Id(title(a.f.ValueType.Name)).Values(Dict{
		Id(a.f.ValueType.Name): Id(a.f.ValueType.Name + "Core").Values(Dict{
			Id("OperationKind_"): Id("OperationKindDelete"),
		})})
}

func (a adder) createNewElement() *Statement {
	return Id(a.f.ValueType.Name).Op(":=").Id("se").Dot("create" + title(a.f.ValueType.Name)).Params(Id("true"))
}

func (a adder) appendElement() *Statement {

	var toAppend *Statement
	if a.f.ValueType.IsBasicType {
		toAppend = Id(a.f.Name + "...")
	} else {
		toAppend = Id(a.f.ValueType.Name).Dot(a.f.ValueType.Name).Dot("ID")
	}

	appendStatement := Id(a.t.Name).Dot(a.t.Name).Dot(title(a.f.Name)).Op("=").Append(
		Id(a.t.Name).Dot(a.t.Name).Dot(title(a.f.Name)),
		toAppend,
	)

	return appendStatement
}

func (a adder) setOperationKindUpdate() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("OperationKind_").Op("=").Id("OperationKindUpdate")
}

func (a adder) updateElementInPatch() *Statement {
	return Id("se").Dot("Patch").Dot(title(a.t.Name)).Index(a.elementID()).Op("=").Id(a.t.Name).Dot(a.t.Name)
}

func (a adder) elementID() *Statement {
	return Id(a.t.Name).Dot(a.t.Name).Dot("ID")
}

func (a adder) receiverName() string {
	return "_" + a.t.Name
}
