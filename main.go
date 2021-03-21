package main

import (
	"bar-cli/enginefactory"
	"fmt"
)

func main() {
	data := map[interface{}]interface{}{
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

	state := enginefactory.WriteEngineFrom(data)
	fmt.Println(string(state))
}
