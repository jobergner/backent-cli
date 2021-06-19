package main

import (
	"log"

	state "github.com/Java-Jonas/bar-cli/examples/application/server"
)

func main() {

	var playerID state.PlayerID

	state.Start(
		func(a state.AddItemToPlayerParams, e *state.Engine) state.AddItemToPlayerResponse {
			return state.AddItemToPlayerResponse{}
		},
		func(p state.MovePlayerParams, e *state.Engine) {
			if playerID == 0 {
				player := e.CreatePlayer()
				log.Println(player.ID())
				playerID = player.ID()
			}
			log.Println("moving player..")
			e.Player(playerID).Position().SetX(p.ChangeX)
		},
		func(a state.SpawnZoneItemsParams, e *state.Engine) state.SpawnZoneItemsResponse {
			return state.SpawnZoneItemsResponse{}
		},
		func(*state.Engine) {},
		func(*state.Engine) {},
	)
}
