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
			name: func() string {
				return title(configType.Name)
			},
			typeName: func() string {
				return configType.Name
			},
		}

		i := idGetter{
			typeName: func() string {
				return configType.Name
			},
			returns: func() string {
				return title(configType.Name) + "ID"
			},
			idFieldToReturn: func() string {
				return "ID"
			},
		}

		writeTypeDeleter(&decls, t)
		writeIdDeleter(&decls, i)

		configType.RangeFields(func(field ast.Field) {
			f := fieldGetter{
				t: configType,
				f: field,
			}

			decls.File.Func().Params(f.receiverParams()).Id(f.name()).Params().Id(f.returns()).Block(
				f.reassignElement(),
				// if slice
				onlyIf(field.HasSliceValue, f.declareSliceOfElements()),
				onlyIf(field.HasSliceValue, For(f.loopConditions().Block(
					f.appendElement(),
				))),
				onlyIf(field.HasSliceValue, Return(f.returnSliceOfType())),
				// if not slice
				onlyIf(!field.HasSliceValue, Return(f.returnSingleType())),
			)
		})

	})

	alreadyWrittenCheck := make(map[string]bool)

	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if alreadyWrittenCheck[field.ValueTypeName] {
				return
			}
			if !field.HasPointerValue {
				return
			}

			t := typeGetter{}
			i := idGetter{}

			t.name = func() string {
				return field.ValueTypeName
			}
			i.idFieldToReturn = func() string {
				return "ReferencedElementID"
			}
			if field.HasAnyValue {
				i.returns = func() string {
					return title(anyNameByField(field)) + "ID"
				}
			} else {
				i.returns = func() string {
					return title(field.ValueType().Name) + "ID"
				}
			}

			alreadyWrittenCheck[field.ValueTypeName] = true

			t.typeName = t.name
			i.typeName = t.name

			writeTypeDeleter(&decls, t)
			writeIdDeleter(&decls, i)

		})
	})

	alreadyWrittenCheck = make(map[string]bool)

	s.config.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			if alreadyWrittenCheck[anyNameByField(field)] {
				return
			}
			if !field.HasAnyValue {
				return
			}

			t := typeGetter{}
			i := idGetter{}

			t.name = func() string {
				return anyNameByField(field)
			}
			i.idFieldToReturn = func() string {
				return "ID"
			}
			i.returns = func() string {
				return title(t.name()) + "ID"
			}

			alreadyWrittenCheck[anyNameByField(field)] = true

			t.typeName = t.name
			i.typeName = t.name

			writeTypeDeleter(&decls, t)
			writeIdDeleter(&decls, i)
		})
	})

	decls.Render(s.buf)
	return s
}

func writeTypeDeleter(decls *DeclSet, t typeGetter) {
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
}

func writeIdDeleter(decls *DeclSet, i idGetter) {
	decls.File.Func().Params(i.receiverParams()).Id(i.name()).Params().Id(i.returns()).Block(
		Return(i.returnID()),
	)
}

type typeGetter struct {
	name     func() string
	typeName func() string
}

func (t typeGetter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (t typeGetter) idParam() string {
	return t.typeName() + "ID"
}

func (t typeGetter) params() *Statement {
	return Id(t.idParam()).Id(title(t.typeName()) + "ID")
}

func (t typeGetter) returns() string {
	return t.typeName()
}

func (t typeGetter) definePatchingElement() *Statement {
	return List(Id("patching"+title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("Patch").Dot(title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetter) earlyReturnPatching() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id("patching" + title(t.typeName()))})
}

func (t typeGetter) defineCurrentElement() *Statement {
	return List(Id("current"+title(t.typeName())), Id("ok")).Op(":=").Id("engine").Dot("State").Dot(title(t.typeName())).Index(Id(t.idParam()))
}

func (t typeGetter) earlyReturnCurrent() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id("current" + title(t.typeName()))})
}

func (t typeGetter) finalReturn() *Statement {
	return Id(t.typeName()).Values(Dict{Id(t.typeName()): Id(t.typeName() + "Core").Values(Dict{Id("OperationKind"): Id("OperationKindDelete"), Id("engine"): Id("engine")})})
}

type idGetter struct {
	typeName        func() string
	returns         func() string
	idFieldToReturn func() string
}

func (i idGetter) receiverName() string {
	return "_" + i.typeName()
}

func (i idGetter) receiverParams() *Statement {
	return Id(i.receiverName()).Id(i.typeName())
}

func (i idGetter) name() string {
	return "ID"
}

func (i idGetter) returnID() *Statement {
	return Id(i.receiverName()).Dot(i.typeName()).Dot(i.idFieldToReturn())
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

func (f fieldGetter) returnedType() string {

	if f.f.ValueType().IsBasicType {
		return f.f.ValueType().Name
	}
	if f.f.HasPointerValue {
		return f.f.Parent.Name + title(pluralizeClient.Singular(f.f.Name)) + "Ref"
	}
	if f.f.HasAnyValue {
		return anyNameByField(f.f)
	}
	return f.f.ValueType().Name
}

func (f fieldGetter) returns() string {
	returnedLiteral := f.returnedType()
	if f.f.HasSliceValue {
		return "[]" + returnedLiteral
	} else if f.f.HasPointerValue {
		return "(" + returnedLiteral + ", bool)"
	}
	return returnedLiteral
}

func (f fieldGetter) reassignElement() *Statement {
	return Id(f.t.Name).Op(":=").Id(f.receiverName()).Dot(f.t.Name).Dot("engine").Dot(title(f.t.Name)).Call(Id(f.receiverName()).Dot(f.t.Name).Dot("ID"))
}

func (f fieldGetter) declareSliceOfElements() *Statement {
	returnedType := f.returnedType()
	if f.f.HasSliceValue {
		returnedType = "[]" + returnedType
	}
	return Var().Id(f.f.Name).Id(returnedType)
}

func (f fieldGetter) loopedElementIdentifier() string {
	if f.f.ValueType().IsBasicType {
		return "element"
	}
	if f.f.HasPointerValue {
		return "refID"
	}
	if f.f.HasAnyValue {
		return anyNameByField(f.f) + "ID"
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
	returnedType := f.returnedType()
	if !f.f.HasPointerValue && !f.f.HasAnyValue {
		returnedType = title(returnedType)
	}
	return Id(f.t.Name).Dot(f.t.Name).Dot("engine").Dot(returnedType).Call(Id(f.loopedElementIdentifier()))
}

func (f fieldGetter) appendElement() *Statement {
	return Id(f.f.Name).Op("=").Append(Id(f.f.Name), f.appendedItem())
}

func (f fieldGetter) returnSliceOfType() *Statement {
	return Id(f.f.Name)
}

func (f fieldGetter) returnBasicType() *Statement {
	return Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))
}

func (f fieldGetter) returnNamedType() *Statement {
	engine := Id(f.t.Name).Dot(f.t.Name).Dot("engine")
	if f.f.HasPointerValue {
		return engine.Dot(f.returnedType()).Call(Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name))).Id(",").Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name)).Op("!=").Lit(0)
	}
	returnedType := f.returnedType()
	if !f.f.HasAnyValue {
		returnedType = title(returnedType)
	}
	return engine.Dot(returnedType).Call(Id(f.t.Name).Dot(f.t.Name).Dot(title(f.f.Name)))
}

func (f fieldGetter) returnSingleType() *Statement {
	if f.f.ValueType().IsBasicType {
		return f.returnBasicType()
	}
	return f.returnNamedType()
}
