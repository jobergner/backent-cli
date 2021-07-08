package main

var exampleConfig = config{
	State: map[interface{}]interface{}{
		"player": map[interface{}]interface{}{
			"name":         "string",
			"location":     "location",
			"items":        "[]item",
			"friendsList":  "[]*player",
			"inCombatWith": "*anyOf<npc,player>",
		},
		"item": map[interface{}]interface{}{
			"name":          "string",
			"firstLootedBy": "*player",
			"isRare":        "bool",
		},
		"npc": map[interface{}]interface{}{
			"name":     "string",
			"location": "location",
		},
		"location": map[interface{}]interface{}{
			"x": "float64",
			"y": "float64",
		},
	},
	Actions: map[interface{}]interface{}{
		"moveNpc": map[interface{}]interface{}{
			"npc":  "npcID",
			"newX": "float64",
			"newY": "float64",
		},
		"movePlayer": map[interface{}]interface{}{
			"player": "playerID",
			"newX":   "float64",
			"newY":   "float64",
		},
		"addFriend": map[interface{}]interface{}{
			"player":    "playerID",
			"newFriend": "playerID",
		},
		"setPlayerCombat": map[interface{}]interface{}{
			"player":    "playerID",
			"enemyKind": "string",
		},
		"playerLeaveCombat": map[interface{}]interface{}{
			"player": "playerID",
		},
		"addItemToPlayer": map[interface{}]interface{}{
			"itemName": "string",
		},
	},
	Responses: map[interface{}]interface{}{
		"addFriend": map[interface{}]interface{}{
			"newNumberOfFriends": "int",
		},
		"setPlayerCombat": map[interface{}]interface{}{
			"enemyEntityKind": "string",
			"enemyEntityPath": "string",
		},
		"playerLeaveCombat": map[interface{}]interface{}{
			"combatWon": "bool",
		},
		"addItemToPlayer": map[interface{}]interface{}{
			"itemPath": "string",
		},
	},
}
