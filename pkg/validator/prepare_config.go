package validator

import "fmt"

type metaField struct {
	name  string
	value string
}

var (
	metaFields = []metaField{
		{
			name:  "__event__",
			value: "true",
		},
	}
)

func prepareStateConfig(stateConfigData map[interface{}]interface{}) (map[interface{}]interface{}, []error) {
	// prevalidate structure so we can make our assumptions about how the data looks
	structuralErrs := structuralValidation(stateConfigData)
	if len(structuralErrs) != 0 {
		return nil, structuralErrs
	}

	dataCopy := copyData(stateConfigData)

	for _, value := range dataCopy {
		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			prepareStateConfigObject(mapValue)
		}
	}

	return dataCopy, nil
}

func prepareStateConfigObject(objectData map[interface{}]interface{}) {

	for key, value := range objectData {
		keyName := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)

		for _, mf := range metaFields {
			if keyName == mf.name && valueString == mf.value {
				delete(objectData, key)
			}
		}

	}
}
