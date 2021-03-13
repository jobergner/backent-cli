package statefactory

import (
	"fmt"
	"regexp"
)

/*
simpleAST is an abstract syntax tree of a state configuration.
I could have used Go's own AST, since the way state is configured leans very heavily onto
Go's structs, but that would have made things more complicated than they need to be.
This way I was also able to add functionality I needed and will be more flexible in the future.
*/
type simpleAST struct {
	Decls map[string]simpleTypeDecl
}

func newSimpleAST() simpleAST {
	return simpleAST{
		Decls: make(map[string]simpleTypeDecl),
	}
}

type simpleTypeDecl struct {
	Name        string
	Fields      map[string]simpleFieldDecl
	IsBasicType bool // is of one of Go's basic types (string, rune, int etc.)
	IsRootType  bool // is not implemented into any other types and thus can not have parent
	IsLeafType  bool // does not implement any other user-defined types in any of its fields
}

func newSimpleTypeDecl(name string) simpleTypeDecl {
	return simpleTypeDecl{
		Name:   name,
		Fields: make(map[string]simpleFieldDecl),
	}
}

type simpleFieldDecl struct {
	Name          string
	Parent        *simpleTypeDecl
	ValueType     *simpleTypeDecl
	ValueString   string // the original value
	HasSliceValue bool   // if the value is a slice value (eg. []string)
}

func buildRudimentarySimpleAST(data map[interface{}]interface{}) simpleAST {
	ast := newSimpleAST()

	for key, value := range data {
		objectValue := value.(map[interface{}]interface{})
		typeName := getSring(key)

		typeDecl := buildRudimentarySimpleTypeDecl(objectValue, typeName)

		ast.Decls[typeName] = typeDecl
	}

	return ast
}

func buildRudimentarySimpleTypeDecl(objectValue map[interface{}]interface{}, typeName string) simpleTypeDecl {
	typeDecl := newSimpleTypeDecl(typeName)

	for key, value := range objectValue {
		fieldName := getSring(key)
		valueString := getSring(value)

		fieldDecl := simpleFieldDecl{
			Name:          fieldName,
			HasSliceValue: isSliceValue(valueString),
			ValueString:   valueString,
		}

		typeDecl.Fields[fieldName] = fieldDecl
	}

	return typeDecl
}

func (s *simpleAST) fillInReferences() *simpleAST {
	for simpleTypeName, simpleType := range s.Decls {
		ss := simpleType
		for simpleFieldName, simpleField := range ss.Fields {
			sf := simpleField
			sf.Parent = &ss
			referencedType, ok := s.Decls[extractValueType(sf.ValueString)]
			if ok {
				sf.ValueType = &referencedType
			} else {
				sf.ValueType = &simpleTypeDecl{Name: extractValueType(sf.ValueString), IsBasicType: true}
			}
			ss.Fields[simpleFieldName] = sf
		}
		s.Decls[simpleTypeName] = ss
	}
	return s
}

func (s *simpleAST) fillInParentalInfo() {
	s.evalRootTypes()
	s.evalLeafTypes()
}

func (s *simpleAST) evalLeafTypes() {
	for simpleTypeName, simpleType := range s.Decls {
		isLeafType := true
		for _, simpleField := range simpleType.Fields {
			if !simpleField.ValueType.IsBasicType {
				isLeafType = false
			}
		}
		if isLeafType {
			simpleType.IsLeafType = true
			s.Decls[simpleTypeName] = simpleType
		}
	}
}

func (s *simpleAST) evalRootTypes() {
	for simpleTypeName, simpleType := range s.Decls {
		isRootType := true
		for _, _simpleType := range s.Decls {
			for _, simpleField := range _simpleType.Fields {
				if simpleField.ValueType.Name == simpleTypeName {
					isRootType = false
				}
			}
		}
		if isRootType {
			simpleType.IsRootType = true
			s.Decls[simpleTypeName] = simpleType
		}
	}
}

// "[]string" -> true
// "string" -> false
func isSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
}

// "[]float64" -> float64
// "float64" -> float64
func extractValueType(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*`)
	return re.FindString(valueString)
}

func getSring(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
