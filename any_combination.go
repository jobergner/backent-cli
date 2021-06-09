package validator

import (
	"fmt"
	"regexp"
)

type anyOfTypeIterator struct {
	parentName       string
	fieldName        string
	types            []string
	currentTypeIndex int
}

func newAnyOfTypeIterator(parentName, fieldName, valueString string) anyOfTypeIterator {
	return anyOfTypeIterator{
		parentName: parentName,
		fieldName:  fieldName,
		types:      extractTypes(valueString)[1:], // omit first element as `extractTypes` consider the `anyOf` prefix a type
	}
}

type anyOfTypeCombinator struct {
	anyOfTypes       []anyOfTypeIterator
	dataCombinations []map[interface{}]interface{}
	data             map[interface{}]interface{} // contains given data excluding `anyOf<>` fields
	originalData     map[interface{}]interface{}
}

func newAnyOfTypeCombinator(data map[interface{}]interface{}) *anyOfTypeCombinator {
	return &anyOfTypeCombinator{
		dataCombinations: make([]map[interface{}]interface{}, 0),
		originalData:     data,
	}
}

func (a *anyOfTypeCombinator) build() []error {
	// pre-validate as we need to have a minimum amount of data entegrity before combining
	structuralErrs := structuralValidation(a.originalData)
	if len(structuralErrs) != 0 {
		return structuralErrs
	}
	nonObjectErrs := validateNonObjectType(a.originalData)
	if len(nonObjectErrs) != 0 {
		return nonObjectErrs
	}

	manipulatedData := copyData(a.originalData)

	for k, v := range manipulatedData {
		keyName := fmt.Sprintf("%v", k)
		valueObject := v.(map[interface{}]interface{})
		for _k, _v := range valueObject {
			_keyName := fmt.Sprintf("%v", _k)
			valueString := fmt.Sprintf("%v", _v)
			if isAnyOfTypes(valueString) {
				a.anyOfTypes = append(a.anyOfTypes, newAnyOfTypeIterator(keyName, _keyName, valueString))
				delete(valueObject, _keyName)
				manipulatedData[keyName] = valueObject
			}
		}
	}

	a.data = manipulatedData

	return nil
}

func (a *anyOfTypeCombinator) generateCombinations() []map[interface{}]interface{} {
	if len(a.anyOfTypes) == 0 {
		a.dataCombinations = append(a.dataCombinations, a.originalData)
		return a.dataCombinations
	}
	a.recursivelyIterateAnyOfTypes(0)
	return a.dataCombinations
}

func (a *anyOfTypeCombinator) recursivelyIterateAnyOfTypes(currentAnyOfTypeIndex int) {
	for range a.anyOfTypes[currentAnyOfTypeIndex].types {
		if currentAnyOfTypeIndex == len(a.anyOfTypes)-1 {
			a.generateData()
		}
		if currentAnyOfTypeIndex < len(a.anyOfTypes)-1 {
			a.recursivelyIterateAnyOfTypes(currentAnyOfTypeIndex + 1)
		}
		a.anyOfTypes[currentAnyOfTypeIndex].currentTypeIndex += 1
	}
	a.anyOfTypes[currentAnyOfTypeIndex].currentTypeIndex = 0
}

func (a *anyOfTypeCombinator) generateData() {
	dataCombination := copyData(a.data)

	for _, anyOfType := range a.anyOfTypes {
		value := dataCombination[anyOfType.parentName]
		valueObject := value.(map[interface{}]interface{})
		valueObject[anyOfType.fieldName] = anyOfType.types[anyOfType.currentTypeIndex]
		dataCombination[anyOfType.parentName] = valueObject
	}

	a.dataCombinations = append(a.dataCombinations, dataCombination)
}

func copyData(data map[interface{}]interface{}) map[interface{}]interface{} {
	newData := make(map[interface{}]interface{})
	for k, v := range data {
		newChildData := make(map[interface{}]interface{})
		childMapValue := v.(map[interface{}]interface{})
		for _k, _v := range childMapValue {
			newChildData[_k] = _v
		}
		newData[k] = newChildData
	}
	return newData
}

func isAnyOfTypes(valueString string) bool {
	re := regexp.MustCompile(`anyOf<\s*([A-Za-z]+\s*,\s*)*\s*([A-Za-z]+\s*)>`)
	return re.MatchString(valueString)
}
