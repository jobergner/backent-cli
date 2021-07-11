package integrationtest

import (
	"log"

	"github.com/jobergner/backent-cli/integrationtest/state"
)

var playerID state.PlayerID

var actions = state.Actions{
	AddItemToPlayer: func(a state.AddItemToPlayerParams, e *state.Engine) state.AddItemToPlayerResponse {
		log.Println("addItemToPlayer", a)
		player := e.Player(playerID)
		item := player.AddItem()
		item.SetName(a.NewName)
		return state.AddItemToPlayerResponse{PlayerPath: player.Path()}
	},
	MovePlayer: func(p state.MovePlayerParams, e *state.Engine) {
		log.Println("movePlayer", p)
		playerPosition := e.Player(playerID).Position()
		playerPosition.SetX(playerPosition.X() + p.ChangeX)
	},
	SpawnZoneItems: func(a state.SpawnZoneItemsParams, e *state.Engine) state.SpawnZoneItemsResponse {
		return state.SpawnZoneItemsResponse{}
	},
}

var sideEffects = state.SideEffects{
	OnDeploy: func(e *state.Engine) {
		player := e.CreatePlayer()
		playerID = player.ID()
	},
	OnFrameTick: func(e *state.Engine) {
		playerGearScore := e.Player(playerID).GearScore()
		playerGearScore.SetLevel(playerGearScore.Level() + 1)
	},
}

func startServer() {
	err := state.Start(actions, sideEffects, 10, 3496)
	panic(err)
}
