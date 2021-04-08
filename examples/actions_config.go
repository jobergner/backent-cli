package examples

var ActionsConfig = map[interface{}]interface{}{
	"MovePlayer": map[interface{}]interface{}{
		"playerID": "playerID",
		"changeX":  "float64",
		"changeY":  "float64",
	},
	"addItemToPlayer": map[interface{}]interface{}{
		"item":     "item",
		"playerID": "playerID",
	},
	"spawnZoneItems": map[interface{}]interface{}{
		"items": "[]item",
	},
}
