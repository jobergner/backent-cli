package statefactory

import (
	"fmt"
	"regexp"
)

type simpleAST struct {
	decls map[string]simpleStructDecl
}

func newSimpleAST() simpleAST {
	return simpleAST{
		decls: make(map[string]simpleStructDecl),
	}
}

type simpleStructDecl struct {
	name   string
	fields map[string]simpleFieldDecl
}

func newSimpleStruct(name string) simpleStructDecl {
	return simpleStructDecl{
		name:   name,
		fields: make(map[string]simpleFieldDecl),
	}
}

type simpleFieldDecl struct {
	name          string
	parent        *simpleStructDecl
	valueType     *simpleStructDecl
	valueString   string
	hasSliceValue bool
}

func buildRudimentarySimpleAST(data map[interface{}]interface{}) simpleAST {
	ast := newSimpleAST()

	for key, value := range data {
		objectValue := value.(map[interface{}]interface{})
		structName := getSring(key)

		structDecl := buildRudimentarySimpleStructDecl(objectValue, structName)

		ast.decls[structName] = structDecl
	}

	return ast
}

func buildRudimentarySimpleStructDecl(objectValue map[interface{}]interface{}, structName string) simpleStructDecl {
	structDecl := newSimpleStruct(structName)

	for key, value := range objectValue {
		fieldName := getSring(key)
		valueString := getSring(value)

		fieldDecl := simpleFieldDecl{
			name:          fieldName,
			hasSliceValue: isSliceValue(valueString),
			valueString:   valueString,
		}

		structDecl.fields[fieldName] = fieldDecl
	}

	return structDecl
}

func (s *simpleAST) fillInReferences() {
	for simpleStructName, simpleStruct := range s.decls {
		ss := simpleStruct
		for simpleFieldName, simpleField := range ss.fields {
			sf := simpleField
			sf.parent = &ss
			referencedStruct, ok := s.decls[extractValueType(sf.valueString)]
			if ok {
				sf.valueType = &referencedStruct
			} else {
				sf.valueType = nil
			}
			ss.fields[simpleFieldName] = sf
		}
		s.decls[simpleStructName] = ss
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
