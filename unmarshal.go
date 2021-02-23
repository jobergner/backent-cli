package validator

import (
	"go/ast"

	"gopkg.in/yaml.v2"
)

func convertToDataMap(yamlDataBytes []byte) (map[interface{}]interface{}, error) {
	yamlData := make(map[interface{}]interface{})
	err := yaml.Unmarshal(yamlDataBytes, &yamlData)

	if err != nil {
		return yamlData, err
	}

	return yamlData, err
}

func Unmarshal(yamlDataBytes []byte) ([]ast.Decl, []error) {
	yamlData, err := convertToDataMap(yamlDataBytes)
	if err != nil {
		return nil, []error{err}
	}

	validationErrs := validateYamlData(yamlData)
	if len(validationErrs) > 0 {
		return nil, validationErrs
	}

	file := convertToAST(yamlData)

	return file.Decls, make([]error, 0)
}
