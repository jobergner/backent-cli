package validator

import "fmt"

type metaField struct {
	name  string
	value string
}

var eventMetaField = metaField{
	name:  "__event__",
	value: "true",
}

var (
	metaFields = []metaField{
		eventMetaField,
	}
)

// prepareStateConfig removes meta fields so we can
// validate state configs without them included.
// this way we only have one point where we have to deal with them
func prepareStateConfig(stateConfigData map[interface{}]interface{}) map[interface{}]interface{} {
	dataCopy := copyData(stateConfigData)

	for _, value := range dataCopy {
		if isMap(value) {
			mapValue := value.(map[interface{}]interface{})
			prepareStateConfigObject(mapValue)
		}
	}

	return dataCopy
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
