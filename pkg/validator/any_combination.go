package validator

import (
	"bytes"
	"fmt"
	"regexp"
)

type anyOfTypeIterator struct {
	parentName       string
	fieldName        string
	hasSliceType     bool
	hasPointerType   bool
	types            []string
	currentTypeIndex int
}

func (a *anyOfTypeIterator) inc() {
	a.currentTypeIndex += 1
}

func (a anyOfTypeIterator) currentValueString() string {
	var buf bytes.Buffer

	if a.hasSliceType {
		buf.WriteString("[]")
	}

	if a.hasPointerType {
		buf.WriteString("*")
	}

	buf.WriteString(a.types[a.currentTypeIndex])

	return buf.String()
}

func newAnyOfTypeIterator(parentName, fieldName, valueString string) anyOfTypeIterator {
	return anyOfTypeIterator{
		parentName:     parentName,
		fieldName:      fieldName,
		hasSliceType:   hasSliceValue(valueString),
		hasPointerType: hasPointerValue(valueString),
		types:          extractTypes(valueString)[1:], // omit first element as `extractTypes` consider the `anyOf` prefix a type
	}
}

func hasSliceValue(valueString string) bool {
	re := regexp.MustCompile(`\[\]`)
	return re.MatchString(valueString)
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

func (a *anyOfTypeCombinator) build() (errs []error) {

	manipulatedData := copyData(a.originalData)

	for k, v := range a.originalData {

		keyName := fmt.Sprintf("%v", k)
		valueObject := v.(map[interface{}]interface{})

		for _k, _v := range valueObject {

			_keyName := fmt.Sprintf("%v", _k)
			valueString := fmt.Sprintf("%v", _v)

			if isAnyOfTypes(valueString) {
				if err := validateAnyOfDefinition(valueString); err != nil {
					errs = append(errs, err)
				}

				a.anyOfTypes = append(a.anyOfTypes, newAnyOfTypeIterator(keyName, _keyName, valueString))

				delete(valueObject, _keyName)

				manipulatedData[keyName] = valueObject
			}

		}

	}

	a.data = manipulatedData

	return
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

		a.anyOfTypes[currentAnyOfTypeIndex].inc()

	}

	a.anyOfTypes[currentAnyOfTypeIndex].currentTypeIndex = 0
}

func (a *anyOfTypeCombinator) generateData() {
	dataCombination := copyData(a.data)

	for _, anyOfType := range a.anyOfTypes {

		value := dataCombination[anyOfType.parentName]
		valueObject := value.(map[interface{}]interface{})
		valueObject[anyOfType.fieldName] = anyOfType.currentValueString()
		dataCombination[anyOfType.parentName] = valueObject

	}

	a.dataCombinations = append(a.dataCombinations, dataCombination)
}

func copyData(data map[interface{}]interface{}) map[interface{}]interface{} {
	newData := make(map[interface{}]interface{})

	for k, v := range data {

		newChildData := make(map[interface{}]interface{})

		if isMap(v) {
			childMapValue := v.(map[interface{}]interface{})

			for _k, _v := range childMapValue {

				newChildData[_k] = _v

			}

			newData[k] = newChildData

		} else {

			newData[k] = v

		}

	}

	return newData
}

func isAnyOfTypes(valueString string) bool {
	re := regexp.MustCompile(`(\[\])?\*?anyOf<\s*(([A-Za-z]+\s*,\s*)*\s*([A-Za-z]+\s*))*>`)
	s := re.FindString(valueString)
	if len(s) == 0 || len(s) != len(valueString) {
		return false
	}
	return true
}

func stateConfigCombinationsFrom(data map[interface{}]interface{}) ([]map[interface{}]interface{}, []error) {
	cmb := newAnyOfTypeCombinator(data)

	invalidAnyOfDefinitionErrs := cmb.build()
	if len(invalidAnyOfDefinitionErrs) != 0 {
		return nil, invalidAnyOfDefinitionErrs
	}

	return cmb.generateCombinations(), nil
}
