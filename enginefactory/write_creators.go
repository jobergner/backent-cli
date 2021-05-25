package enginefactory

import (
	"bar-cli/ast"
	. "bar-cli/factoryutils"

	. "github.com/dave/jennifer/jen"
)

func (s *EngineFactory) writeCreators() *EngineFactory {
	decls := NewDeclSet()
	s.config.RangeTypes(func(configType ast.ConfigType) {

		cw := creatorWrapperWriter{
			t: configType,
		}

		decls.File.Func().Params(cw.receiverParams()).Id(cw.name()).Params().Id(cw.returns()).Block(
			Return(cw.createElement()),
		)

		c := creatorWriter{
			t: configType,
			f: nil,
		}

		decls.File.Func().Params(c.receiverParams()).Id(c.name()).Params(c.params()).Id(c.returns()).Block(
			c.declareElement(),
			c.assignEngine(),
			c.generateID(),
			onlyIf(!configType.IsRootType, c.setHasParent()),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				c.f = &field
				if field.HasSliceValue || field.ValueType().IsBasicType || field.HasPointerValue {
					return Empty()
				}
				return &Statement{
					c.createChildElement(), Line(),
					c.setChildElement(),
				}
			}),
			c.setOperationKind(),
			c.updateElementInPatch(),
			Return(c.returnElement()),
		)
	})

	s.config.RangeRefFields(func(field ast.Field) {
		c := generatedTypeCreatorWriter{
			f:        field,
			typeName: field.ValueTypeName,
		}

		decls.File.Func().Params(c.receiverParams()).Id(c.name()).Params(c.params()).Id(c.returns()).Block(
			c.declareElement(),
			c.assignEngine(),
			c.setReferencedElementID(),
			c.setParentID(),
			c.setID(),
			c.setOperationKind(),
			c.assignElementToPatch(),
			Return(Id("element")),
		)
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		c := generatedTypeCreatorWriter{
			f:        field,
			typeName: anyNameByField(field),
		}

		decls.File.Func().Params(c.receiverParams()).Id(c.name()).Params(Id("setDefaultValue").Bool()).Id(anyNameByField(field)).Block(
			c.declareElement(),
			c.assignEngine(),
			c.setID(),
			If(Id("setDefaultValue")).Block(
				c.createChildElement(),
				c.assignChildElement(),
				c.assignElementKind(),
			),
			c.setOperationKind(),
			c.assignElementToPatch(),
			Return(c.returnElement()),
		)
	})

	decls.Render(s.buf)
	return s
}

type creatorWrapperWriter struct {
	t ast.ConfigType
}

func (cw creatorWrapperWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (cw creatorWrapperWriter) name() string {
	return "Create" + title(cw.t.Name)
}

func (cw creatorWrapperWriter) returns() string {
	return cw.t.Name
}

func (cw creatorWrapperWriter) createElement() *Statement {
	var callParam *Statement
	if cw.t.IsRootType {
		callParam = Empty()
	} else {
		callParam = Lit(false)
	}

	return Id("engine").Dot("create" + title(cw.t.Name)).Call(callParam)
}

type creatorWriter struct {
	t ast.ConfigType
	f *ast.Field
}

func (c creatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c creatorWriter) name() string {
	return "create" + title(c.t.Name)
}

func (c creatorWriter) returns() string {
	return c.t.Name
}

func (c creatorWriter) hasParentParam() string {
	if c.t.IsRootType {
		return ""
	}
	return "hasParent"
}

func (c creatorWriter) params() *Statement {
	if c.t.IsRootType {
		return Empty()
	}
	return Id(c.hasParentParam()).Bool()
}

func (c creatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.t.Name + "Core")
}

func (c creatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c creatorWriter) generateID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(title(c.t.Name) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c creatorWriter) setHasParent() *Statement {
	return Id("element").Dot("HasParent").Op("=").Id(c.hasParentParam())
}

func (c creatorWriter) createChildElement() *Statement {
	return Id("element" + title(c.f.Name)).Op(":=").Id("engine").Dot("create" + title(c.f.ValueTypeName)).Call(Lit(true))
}
func (c creatorWriter) setChildElement() *Statement {
	return Id("element").Dot(title(c.f.Name)).Op("=").Id("element" + title(c.f.Name)).Dot(c.f.ValueTypeName).Dot("ID")
}

func (c creatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c creatorWriter) updateElementInPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(title(c.t.Name)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c creatorWriter) returnElement() *Statement {
	return Id(c.t.Name).Values(Dict{
		Id(c.t.Name): Id("element"),
	})
}

type generatedTypeCreatorWriter struct {
	f        ast.Field
	typeName string
}

func (c generatedTypeCreatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c generatedTypeCreatorWriter) name() string {
	return "create" + title(c.typeName)
}

func (c generatedTypeCreatorWriter) params() *Statement {
	referencedElementType := c.f.ValueType().Name
	if c.f.HasAnyValue {
		referencedElementType = anyNameByField(c.f)
	}
	return List(Id("referencedElementID").Id(title(referencedElementType)+"ID"), Id("parentID").Id(title(c.f.Parent.Name)+"ID"))
}

func (c generatedTypeCreatorWriter) returns() string {
	return c.typeName + "Core"
}

func (c generatedTypeCreatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.typeName + "Core")
}

func (c generatedTypeCreatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c generatedTypeCreatorWriter) setReferencedElementID() *Statement {
	return Id("element").Dot("ReferencedElementID").Op("=").Id("referencedElementID")
}

func (c generatedTypeCreatorWriter) setParentID() *Statement {
	return Id("element").Dot("ParentID").Op("=").Id("parentID")
}

func (c generatedTypeCreatorWriter) setID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(title(c.typeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c generatedTypeCreatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c generatedTypeCreatorWriter) assignElementToPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(title(c.typeName)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}

func (c generatedTypeCreatorWriter) createChildElement() *Statement {
	return Id("element" + title(c.f.ValueType().Name)).Op(":=").Id("engine").Dot("create" + title(c.f.ValueType().Name)).Call(True())
}

func (c generatedTypeCreatorWriter) assignChildElement() *Statement {
	return Id("element").Dot(title(c.f.ValueType().Name)).Op("=").Id("element" + title(c.f.ValueType().Name)).Dot(c.f.ValueType().Name).Dot("ID")
}

func (c generatedTypeCreatorWriter) assignElementKind() *Statement {
	return Id("element").Dot("ElementKind").Op("=").Id("ElementKind" + title(c.f.ValueType().Name))
}

func (c generatedTypeCreatorWriter) returnElement() *Statement {
	return Id(anyNameByField(c.f)).Values(Dict{
		Id(anyNameByField(c.f)): Id("element"),
	})
}
