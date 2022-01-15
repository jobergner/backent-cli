package enginefactory

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factoryutils"

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
			c.assignExtendedPath(),
			c.assignJsonPath(),
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
			c.setHasParent(),
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
			c.assignPath(),
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

		decls.File.Func().Params(c.receiverParams()).Id(c.name()).Params(Id("setDefaultValue").Bool(), Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier")).Id(Title(anyNameByField(field))).Block(
			c.declareElement(),
			c.assignEngine(),
			c.setID(),
			If(Id("setDefaultValue")).Block(
				c.createChildElement(),
				c.assignChildElement(),
				c.assignElementKind(),
			),
			c.setOperationKind(),
			c.setChildElementPath(),
			c.setFieldIdentifier(),
			c.assignElementToPatch(),
			Return(c.returnElement()),
		)
	})

	decls.Render(s.buf)
	return s
}
