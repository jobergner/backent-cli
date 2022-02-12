package configs

var StateConfig = map[interface{}]interface{}{
	"player": map[interface{}]interface{}{
		"items":         "[]item",
		"equipmentSets": "[]*equipmentSet",
		"gearScore":     "gearScore",
		"position":      "position",
		"guildMembers":  "[]*player",
		"target":        "*anyOf<player,zoneItem>",
		"targetedBy":    "[]*anyOf<player,zoneItem>",
		"action":        "[]attackEvent",
	},
	"attackEvent": map[interface{}]interface{}{
		"__event__": "true",
		"target":    "*player",
	},
	"zone": map[interface{}]interface{}{
		"items":         "[]zoneItem",
		"players":       "[]player",
		"tags":          "[]string",
		"interactables": "[]anyOf<item,player,zoneItem>",
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
		"origin":    "anyOf<player,position>",
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
