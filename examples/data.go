package examples

var StateConfig = map[interface{}]interface{}{
	"player": map[interface{}]interface{}{
		"items":     "[]item",
		"gearScore": "gearScore",
		"position":  "position",
	},
	"zone": map[interface{}]interface{}{
		"items":   "[]zoneItem",
		"players": "[]player",
		"tags":    "[]string",
	},
	"zoneItem": map[interface{}]interface{}{
		"position": "position",
		"item":     "item",
	},
	"position": map[interface{}]interface{}{
		"x": "float64",
		"y": "float64",
	},
	"item": map[interface{}]interface{}{
		"gearScore": "gearScore",
	},
	"gearScore": map[interface{}]interface{}{
		"level": "int",
		"score": "int",
	},
}

var ActionsConfig = map[interface{}]interface{}{
	"movePlayer": map[interface{}]interface{}{
		"playerID": "int",
		"changeX":  "int",
		"changeY":  "int",
	},
	"addItemToPlayer": map[interface{}]interface{}{
		"item":     "item",
		"playerID": "int",
	},
	"spawnZoneItems": map[interface{}]interface{}{
		"items": "[]item",
	},
}
