package main

import (
	"github.com/Java-Jonas/bar-cli/integrationtest/state"
	"log"
)

func startServer() {
	var playerID state.PlayerID

	err := state.Start(
		func(a state.AddItemToPlayerParams, e *state.Engine) {},
		func(p state.MovePlayerParams, e *state.Engine) {
			if playerID == 0 {
				player := e.CreatePlayer()
				playerID = player.ID()
			}
			log.Println("moving player..")
			e.Player(playerID).Position().SetX(p.ChangeX)
		},
		func(a state.SpawnZoneItemsParams, e *state.Engine) {},
		func(*state.Engine) {},
		func(*state.Engine) {},
	)

	panic(err)
}
