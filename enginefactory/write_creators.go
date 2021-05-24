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
		c := anyCreatorWriter{
			f: field,
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

type anyCreatorWriter struct {
	f ast.Field
}

func (c anyCreatorWriter) receiverParams() *Statement {
	return Id("engine").Id("*Engine")
}

func (c anyCreatorWriter) name() string {
	return "create" + title(c.f.ValueTypeName)
}

func (c anyCreatorWriter) params() *Statement {
	referencedElementType := c.f.ValueType().Name
	if c.f.HasAnyValue {
		referencedElementType = anyNameByField(c.f)
	}
	return List(Id("referencedElementID").Id(title(referencedElementType)+"ID"), Id("parentID").Id(title(c.f.Parent.Name)+"ID"))
}

func (c anyCreatorWriter) returns() string {
	return c.f.ValueTypeName + "Core"
}

func (c anyCreatorWriter) declareElement() *Statement {
	return Var().Id("element").Id(c.f.ValueTypeName + "Core")
}

func (c anyCreatorWriter) assignEngine() *Statement {
	return Id("element").Dot("engine").Op("=").Id("engine")
}

func (c anyCreatorWriter) setReferencedElementID() *Statement {
	return Id("element").Dot("ReferencedElementID").Op("=").Id("referencedElementID")
}

func (c anyCreatorWriter) setParentID() *Statement {
	return Id("element").Dot("ParentID").Op("=").Id("parentID")
}

func (c anyCreatorWriter) setID() *Statement {
	return Id("element").Dot("ID").Op("=").Id(title(c.f.ValueTypeName) + "ID").Call(Id("engine").Dot("GenerateID").Call())
}

func (c anyCreatorWriter) setOperationKind() *Statement {
	return Id("element").Dot("OperationKind").Op("=").Id("OperationKindUpdate")
}

func (c anyCreatorWriter) assignElementToPatch() *Statement {
	return Id("engine").Dot("Patch").Dot(title(c.f.ValueTypeName)).Index(Id("element").Dot("ID")).Op("=").Id("element")
}
