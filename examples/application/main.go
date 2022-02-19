package main

import (
	"log"

	state "github.com/jobergner/backent-cli/examples/application/server"
)

var playerID state.PlayerID

var actions = state.Actions{
	AddItemToPlayer: func(params state.AddItemToPlayerParams, engine *state.Engine, roomName, clientID string) state.AddItemToPlayerResponse {
		return state.AddItemToPlayerResponse{}
	},
	MovePlayer: func(params state.MovePlayerParams, engine *state.Engine, roomName, clientID string) {
		if playerID == 0 {
			player := engine.CreatePlayer()
			log.Println(player.ID())
			playerID = player.ID()
		}
		log.Println("moving player..")
		engine.Player(playerID).Position().SetX(params.ChangeX)
	},
	SpawnZoneItems: func(params state.SpawnZoneItemsParams, engine *state.Engine, roomName, clientID string) state.SpawnZoneItemsResponse {
		return state.SpawnZoneItemsResponse{}
	},
}

var sideEffects = state.SideEffects{
	OnDeploy:    func(*state.Engine) {},
	OnFrameTick: func(*state.Engine) {},
}

var signals = state.LoginSignals{}

func main() {
	err := state.Start(signals, actions, sideEffects, 1, 3496)
	if err != nil {
		panic(err)
	}
}
