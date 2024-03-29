package utils

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"strings"

	"github.com/jobergner/backent-cli/pkg/ast"

	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

var (
	PackageName   = "main"
	PackageClause = fmt.Sprintf("package %s\n", PackageName)
)

type BasicType struct {
	Name  string
	Value string
}

var (
	BasicTypes = map[string]string{
		"bool":    "boolValue",
		"float64": "floatValue",
		"int64":   "intValue",
		"string":  "stringValue",
	}
)

func RangeBasicTypes(fn func(BasicType)) {
	for _, name := range []string{"bool", "float64", "int64", "string"} {
		fn(BasicType{Name: name, Value: BasicTypes[name]})
	}
}

func ForEachBasicType(fn func(BasicType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	RangeBasicTypes(func(basicType BasicType) {
		statements = append(statements, fn(basicType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

// pluralizeClient.Singular is used to find the singular of field names
// this is necessary for writing coherent method names, eg. in write_adders.go (toSingular)
// with getting the singular form of a plural, this field:
// { pieces []piece }
// can have the coherent adder method of "AddPiece"
var Singular func(string) string = pluralize.NewClient().Singular
var Plural func(string) string = pluralize.NewClient().Plural

func ForEachParamInAction(action ast.Action, fn func(param ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	action.RangeParams(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachResponseValueInAction(action ast.Action, fn func(param ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	action.RangeResponse(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldInType(configType ast.ConfigType, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldInAST(a *ast.AST, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	a.RangeTypes(func(configType ast.ConfigType) {
		configType.RangeFields(func(field ast.Field) {
			statements = append(statements, fn(field))
			statements = append(statements, jen.Line())
		})
	})
	return &statements
}

func ForEachTypeImplementingType(configType ast.ConfigType, fn func(configType *ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeImplementedBy(func(configType *ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachValueOfField(field ast.Field, fn func(configType *ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	field.RangeValueTypes(func(parentType *ast.ConfigType) {
		statements = append(statements, fn(parentType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachReferenceOfType(configType ast.ConfigType, fn func(field *ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	configType.RangeReferencedBy(func(field *ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachTypeInAST(config *ast.AST, fn func(configType ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeTypes(func(configType ast.ConfigType) {
		statements = append(statements, fn(configType))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachRefFieldInAST(config *ast.AST, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeRefFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachAnyFieldInAST(config *ast.AST, fn func(field ast.Field) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeAnyFields(func(field ast.Field) {
		statements = append(statements, fn(field))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachActionInAST(config *ast.AST, fn func(action ast.Action) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	config.RangeActions(func(action ast.Action) {
		statements = append(statements, fn(action))
		statements = append(statements, jen.Line())
	})
	return &statements
}

func ForEachFieldValueComparison(field ast.Field, comparator jen.Statement, fn func(configType *ast.ConfigType) *jen.Statement) *jen.Statement {
	var statements jen.Statement
	first := true
	field.RangeValueTypes(func(valueType *ast.ConfigType) {
		statement := jen.Empty()
		if !first {
			statement.Else()
		}
		_comparator := comparator
		statement.If(_comparator.Op("==").Id("ElementKind" + Title(valueType.Name))).Block(
			fn(valueType),
		)
		statements = append(statements, statement)
		first = false
	})
	return &statements
}

func FieldPathIdentifier(f ast.Field) string {
	return fmt.Sprintf("%s_%sIdentifier", Lower(f.Parent.Name), f.Name)
}

func Title(name string) string {
	return strings.Title(name)
}

func Lower(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func OnlyIf(is bool, statement *jen.Statement) *jen.Statement {
	if is {
		return statement
	}
	return jen.Empty()
}

type DeclSet struct {
	File *jen.File
}

func NewDeclSet() DeclSet {
	return DeclSet{
		File: jen.NewFile("state"),
	}
}

func (d DeclSet) Render(buf *bytes.Buffer) {
	var _buf bytes.Buffer
	err := d.File.Render(&_buf)
	if err != nil {
		panic(err)
	}
	code := TrimPackageClause(_buf.String())
	buf.WriteString("\n" + code + "\n")
}

func TrimPackageClause(sourceCode string) string {
	return strings.TrimPrefix(sourceCode, PackageClause)
}

func Format(buf *bytes.Buffer) error {
	fs := token.NewFileSet()

	f, err := parser.ParseFile(fs, "", buf.Bytes(), parser.AllErrors|parser.ParseComments)
	if err != nil {
		return err
	}

	buf.Reset()
	err = format.Node(buf, fs, f)
	if err != nil {
		return err
	}

	return nil
}

func ValueTypeName(field *ast.Field) string {
	if field.HasPointerValue {
		return field.Parent.Name + Title(Singular(field.Name)) + "Ref"
	}
	if field.HasAnyValue {
		return AnyValueTypeName(field)
	}
	return field.ValueType().Name
}

func AnyValueTypeName(field *ast.Field) string {
	name := "anyOf"
	firstIteration := true
	field.RangeValueTypes(func(configType *ast.ConfigType) {
		if firstIteration {
			name += Title(configType.Name)
		} else {
			name += "_" + Title(configType.Name)
		}
		firstIteration = false
	})
	return name
}
