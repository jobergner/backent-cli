package validator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func extracpMapDeclExpression(file *ast.File) ast.Expr {
	return file.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type
}

func findMapTypeRecursive(expr ast.Expr) *ast.MapType {
	if mapType, ok := expr.(*ast.MapType); ok {
		return mapType
	}
	if arrayType, ok := expr.(*ast.ArrayType); ok {
		return findMapTypeRecursive(arrayType.Elt)
	}
	if starType, ok := expr.(*ast.StarExpr); ok {
		return findMapTypeRecursive(starType.X)
	}
	return nil
}

func extractMapKeyRecursive(mapKeys []string, mapType *ast.MapType, mockSrc string) []string {

	if mapType == nil || mapType.Key == nil {
		return mapKeys
	}

	keyIdent, ok := mapType.Key.(*ast.Ident)
	// map[string]int -> true; "string" is the Ident
	// map[[]string]int, map[map[int]string]int .. -> false; map keys are not Idents
	if !ok {
		mapKey := mockSrc[mapType.Key.Pos()-1 : mapType.Key.End()-1]
		mapKeys = append(mapKeys, mapKey)
	} else {
		mapKeys = append(mapKeys, keyIdent.Name)
	}

	mapKeys = extractMapKeyRecursive(mapKeys, findMapTypeRecursive(mapType.Value), mockSrc)

	return mapKeys
}

func extractMapKeys(valueString string) []string {
	mockSrc := `
	package main
	type mockType ` + valueString

	file, _ := parser.ParseFile(token.NewFileSet(), "", mockSrc, 0)

	typeExpression := extracpMapDeclExpression(file)
	mapType, ok := typeExpression.(*ast.MapType)
	if !ok {
		return []string{}
	}

	return extractMapKeyRecursive([]string{}, mapType, mockSrc)
}

func findIllegalMapKeys(valueString string, data map[interface{}]interface{}) []string {
	mapKeys := extractMapKeys(valueString)
	if len(mapKeys) == 0 {
		return nil
	}

	var invalidMapKeys []string

	for _, mapKey := range mapKeys {
		var isInvalid bool
		if isReferenceType(mapKey) {
			isInvalid = true
		}
		if containsUncomparableValue(mapKey, data) {
			isInvalid = true
		}
		if isInvalid {
			invalidMapKeys = append(invalidMapKeys, mapKey)
		}
	}

	return invalidMapKeys
}

func validateIllegalMapKeys(data map[interface{}]interface{}) (errs []error) {
	for _, value := range data {

		if isString(value) {
			valueString := fmt.Sprintf("%v", value)
			illegalMapKeys := findIllegalMapKeys(valueString, data)
			for _, illegalMapKey := range illegalMapKeys {
				errs = append(errs, newValidationErrorInvalidMapKey(illegalMapKey, valueString))
			}
			continue
		}

		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			objectValidationErrs := validateIllegalMapKeysObject(mapValue, data)
			errs = append(errs, objectValidationErrs...)
		}
	}

	return
}

func validateIllegalMapKeysObject(objectData, data map[interface{}]interface{}) (errs []error) {
	for _, value := range objectData {
		valueString := fmt.Sprintf("%v", value)
		illegalMapKeys := findIllegalMapKeys(valueString, data)
		for _, illegalMapKey := range illegalMapKeys {
			errs = append(errs, newValidationErrorInvalidMapKey(illegalMapKey, valueString))
		}
	}
	return
}
