package examples

var StateConfig = map[interface{}]interface{}{
	"player": map[interface{}]interface{}{
		"items":         "[]item",
		"equipmentSets": "[]*equipmentSet",
		"gearScore":     "gearScore",
		"position":      "position",
		"guildMembers":  "[]*player",
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
		"name":      "string",
		"gearScore": "gearScore",
		"boundTo":   "*player",
	},
	"gearScore": map[interface{}]interface{}{
		"level": "int",
		"score": "int",
	},
	"equipmentSet": map[interface{}]interface{}{
		"name":      "string",
		"equipment": "[]*item",
	},
}
