package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeCreators() *Factory {

	RangeBasicTypes(func(b BasicType) {
		c := creatorWriter{
			typeName: b.Value,
			f:        nil,
		}

		s.file.Func().Params(c.receiverParams()).Id(c.name()).Params(Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier"), Id("value").Id(b.Name)).Id(b.Value).Block(
			Var().Id("element").Id(b.Value),
			Id("element").Dot("Value").Op("=").Id("value"),
			c.assignEngine(),
			c.generateID(),
			c.assignExtendedPath(),
			c.assignJsonPath(),
			c.setOperationKind(),
			c.updateElementInPatch(),
			Return(Id("element")),
		)
	})

	s.config.RangeTypes(func(configType ast.ConfigType) {

		cw := creatorWrapperWriter{
			t: configType,
		}

		s.file.Func().Params(cw.receiverParams()).Id(cw.name()).Params().Id(cw.returns()).Block(
			Return(cw.createElement()),
		)

		c := creatorWriter{
			typeName: configType.Name,
			f:        nil,
		}

		s.file.Func().Params(c.receiverParams()).Id(c.name()).Params(c.params()).Id(c.returns()).Block(
			c.declareElement(),
			c.assignEngine(),
			c.generateID(),
			c.assignExtendedPath(),
			c.assignJsonPath(),
			ForEachFieldInType(configType, func(field ast.Field) *Statement {
				c.f = &field
				if field.HasSliceValue || field.HasPointerValue {
					return Empty()
				}
				return &Statement{
					OnlyIf(c.f.HasAnyValue, c.createChildSubElement()), Line(),
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

		s.file.Func().Params(c.receiverParams()).Id(c.name()).Params(c.params()).Id(c.returns()).Block(
			c.declareElement(),
			c.assignEngine(),
			c.setReferencedElementID(),
			c.setParentID(),
			c.setIDRef(),
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

		s.file.Func().Params(c.receiverParams()).Id(c.name()).Params(Id("parentID").Int(), Id("childID").Int(), Id("childKind").Id("ElementKind"), Id("p").Id("path"), Id("fieldIdentifier").Id("treeFieldIdentifier")).Id(Title(anyNameByField(field))).Block(
			c.declareElement(),
			c.assignEngine(),
			c.setIDAny(),
			c.setChildID(),
			c.assignElementKind(),
			c.setOperationKind(),
			c.setChildElementPath(),
			c.setFieldIdentifier(),
			c.assignElementToPatch(),
			Return(c.returnElement()),
		)
	})

	return s
}
