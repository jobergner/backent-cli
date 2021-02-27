package statefactory

import (
	"fmt"
	"regexp"
)

type simpleAST struct {
	decls map[string]simpleTypeDecl
}

func newSimpleAST() simpleAST {
	return simpleAST{
		decls: make(map[string]simpleTypeDecl),
	}
}

type simpleTypeDecl struct {
	name        string
	fields      map[string]simpleFieldDecl
	isBasicType bool
}

func newSimpleType(name string) simpleTypeDecl {
	return simpleTypeDecl{
		name:   name,
		fields: make(map[string]simpleFieldDecl),
	}
}

type simpleFieldDecl struct {
	name          string
	parent        *simpleTypeDecl
	valueType     *simpleTypeDecl
	valueString   string
	hasSliceValue bool
}

func buildRudimentarySimpleAST(data map[interface{}]interface{}) simpleAST {
	ast := newSimpleAST()

	for key, value := range data {
		objectValue := value.(map[interface{}]interface{})
		typeName := getSring(key)

		typeDecl := buildRudimentarySimpleTypeDecl(objectValue, typeName)

		ast.decls[typeName] = typeDecl
	}

	return ast
}

func buildRudimentarySimpleTypeDecl(objectValue map[interface{}]interface{}, typeName string) simpleTypeDecl {
	typeDecl := newSimpleType(typeName)

	for key, value := range objectValue {
		fieldName := getSring(key)
		valueString := getSring(value)

		fieldDecl := simpleFieldDecl{
			name:          fieldName,
			hasSliceValue: isSliceValue(valueString),
			valueString:   valueString,
		}

		typeDecl.fields[fieldName] = fieldDecl
	}

	return typeDecl
}

func (s *simpleAST) fillInReferences() {
	for simpleTypeName, simpleType := range s.decls {
		ss := simpleType
		for simpleFieldName, simpleField := range ss.fields {
			sf := simpleField
			sf.parent = &ss
			referencedType, ok := s.decls[extractValueType(sf.valueString)]
			if ok {
				sf.valueType = &referencedType
			} else {
				sf.valueType = &simpleTypeDecl{name: extractValueType(sf.valueString), isBasicType: true}
			}
			ss.fields[simpleFieldName] = sf
		}
		s.decls[simpleTypeName] = ss
	}
}

func isSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
}

func extractValueType(valueString string) string {
	re := regexp.MustCompile(`[A-Za-z]+[0-9]*`)
	return re.FindString(valueString)
}

func getSring(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
