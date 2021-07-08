package main

var exampleConfig = jsonConfig{
	State: map[string]interface{}{
		"player": map[string]interface{}{
			"name":         "string",
			"location":     "location",
			"items":        "[]item",
			"friendsList":  "[]*player",
			"inCombatWith": "*anyOf<npc,player>",
		},
		"item": map[string]interface{}{
			"name":          "string",
			"firstLootedBy": "*player",
			"isRare":        "bool",
		},
		"npc": map[string]interface{}{
			"name":     "string",
			"location": "location",
		},
		"location": map[string]interface{}{
			"x": "float64",
			"y": "float64",
		},
	},
	Actions: map[string]interface{}{
		"createPlayer": map[string]interface{}{
			"name": "string",
		},
		"moveNpc": map[string]interface{}{
			"npc":  "npcID",
			"newX": "float64",
			"newY": "float64",
		},
		"movePlayer": map[string]interface{}{
			"player": "playerID",
			"newX":   "float64",
			"newY":   "float64",
		},
		"addFriend": map[string]interface{}{
			"player":    "playerID",
			"newFriend": "playerID",
		},
		"setPlayerCombat": map[string]interface{}{
			"player":    "playerID",
			"enemyKind": "string",
			"enemyID":   "int",
		},
		"playerLeaveCombat": map[string]interface{}{
			"player": "playerID",
		},
		"addItemToPlayer": map[string]interface{}{
			"player":   "playerID",
			"itemName": "string",
		},
	},
	Responses: map[string]interface{}{
		"createPlayer": map[string]interface{}{
			"playerPath": "string",
		},
		"addFriend": map[string]interface{}{
			"newNumberOfFriends": "int",
		},
		"setPlayerCombat": map[string]interface{}{
			"enemyEntityKind": "string",
			"enemyEntityPath": "string",
		},
		"playerLeaveCombat": map[string]interface{}{
			"combatWon": "bool",
		},
		"addItemToPlayer": map[string]interface{}{
			"itemPath": "string",
		},
	},
}
