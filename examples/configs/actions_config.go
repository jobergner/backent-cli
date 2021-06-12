package configs

var ActionsConfig = map[interface{}]interface{}{
	"movePlayer": map[interface{}]interface{}{
		"player":  "playerID",
		"changeX": "float64",
		"changeY": "float64",
	},
	"addItemToPlayer": map[interface{}]interface{}{
		"item":    "itemID",
		"newName": "string",
	},
	"spawnZoneItems": map[interface{}]interface{}{
		"items": "[]itemID",
	},
}
