package main

import (
	"log"

	state "github.com/jobergner/backent-cli/examples/application/server"
)

var playerID state.PlayerID

var actions = state.Actions{
	AddItemToPlayer: func(a state.AddItemToPlayerParams, e *state.Engine) state.AddItemToPlayerResponse {
		return state.AddItemToPlayerResponse{}
	},
	MovePlayer: func(p state.MovePlayerParams, e *state.Engine) {
		if playerID == 0 {
			player := e.CreatePlayer()
			log.Println(player.ID())
			playerID = player.ID()
		}
		log.Println("moving player..")
		e.Player(playerID).Position().SetX(p.ChangeX)
	},
	SpawnZoneItems: func(a state.SpawnZoneItemsParams, e *state.Engine) state.SpawnZoneItemsResponse {
		return state.SpawnZoneItemsResponse{}
	},
}

var sideEffects = state.SideEffects{
	OnDeploy:    func(*state.Engine) {},
	OnFrameTick: func(*state.Engine) {},
}

func main() {
	err := state.Start(actions, sideEffects, 1, 3496)
	if err != nil {
		panic(err)
	}
}
