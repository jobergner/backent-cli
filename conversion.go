package validator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

// used to process map in determined order based on key names
func rangeInAlphabeticalOrder(data map[interface{}]interface{}, fn func(key string, value interface{})) {
	var keys []string
	for key := range data {
		keyLiteral := fmt.Sprintf("%v", key)
		keys = append(keys, keyLiteral)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fn(key, data[key])
	}
}

func convertToAST(yamlData map[interface{}]interface{}) *ast.File {
	sw := newSourceWriter()

	rangeInAlphabeticalOrder(yamlData, func(keyName string, value interface{}) {
		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			sw.addNamedType(keyName, valueString)
			return
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			sw.startStructType(keyName)
			rangeInAlphabeticalOrder(mapValue, func(_key string, _value interface{}) {
				_valueString := fmt.Sprintf("%v", _value)
				_keyName := fmt.Sprintf("%v", _key)
				sw.addStructField(_keyName, _valueString)
			})
			sw.closeStructType()
		}
	})

	return sw.parse()
}

type sourceWriter struct {
	sourceCode string
}

func newSourceWriter() *sourceWriter {
	return &sourceWriter{"package " + mockPackageName + "\n"}
}

func (s *sourceWriter) parse() *ast.File {
	file, _ := parser.ParseFile(token.NewFileSet(), "", s.sourceCode, 0)
	return file
}

func (s *sourceWriter) addNamedType(name, typeName string) *sourceWriter {
	s.sourceCode = fmt.Sprintf("%s\ntype %s %s", s.sourceCode, name, typeName)
	return s
}

func (s *sourceWriter) startStructType(name string) *sourceWriter {
	s.sourceCode = fmt.Sprintf("%s\ntype %s struct {", s.sourceCode, name)
	return s
}

func (s *sourceWriter) addStructField(name, typeName string) *sourceWriter {
	s.sourceCode = fmt.Sprintf("%s\n%s %s", s.sourceCode, name, typeName)
	return s
}

func (s *sourceWriter) closeStructType() *sourceWriter {
	s.sourceCode = fmt.Sprintf("%s\n}", s.sourceCode)
	return s
}
