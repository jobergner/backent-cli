package state

import (
	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"

	. "github.com/dave/jennifer/jen"
)

func (s *Factory) writeComplexID() *Factory {

	s.file.Comment("// easyjson:skip ")
	s.file.Id(_ComplexID_type)
	s.file.Id(complexIDStructCache_type)
	s.file.Id(_MarshalJSON_ComplexID_func)
	s.file.Id(_UnmarshalJSON_ComplexID_func)

	s.config.RangeRefFields(func(field ast.Field) {
		writeMarshallingMethods(s.file, field.ValueTypeName)
	})

	s.config.RangeAnyFields(func(field ast.Field) {
		writeMarshallingMethods(s.file, anyNameByField(field))
	})

	return s
}

func writeMarshallingMethods(file *File, typeName string) {
	file.Func().Params(Id("x").Id(Title(typeName)+"ID")).Id("MarshalJSON").Params().Params(Id("[]byte"), Error()).Block(
		Return(Id("ComplexID").Call(Id("x")).Dot("MarshalJSON").Call()),
	)
	file.Func().Params(Id("x").Id("*"+Title(typeName)+"ID")).Id("UnmarshalJSON").Params(Id("s").Id("[]byte")).Error().Block(
		Id("temp").Op(":=").Id("ComplexID").Call(Id("*x")),
		Id("temp").Dot("UnmarshalJSON").Call(Id("s")),
		Id("*x").Op("=").Id(Title(typeName)+"ID").Call(Id("temp")),
		Return(Nil()),
	)
}
