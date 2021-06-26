package main

import (
	"log"

	"github.com/Java-Jonas/bar-cli/integrationtest/state"
)

func startServer() {
	var playerID state.PlayerID

	err := state.Start(
		func(a state.AddItemToPlayerParams, e *state.Engine) state.AddItemToPlayerResponse {
			log.Println("addItemToPlayer", a)
			if playerID != 0 {
				item := e.Player(playerID).AddItem()
				item.SetName(a.NewName)
				return state.AddItemToPlayerResponse{PlayerPath: item.Name()}
			}
			return state.AddItemToPlayerResponse{}
		},
		func(p state.MovePlayerParams, e *state.Engine) {
			if playerID == 0 {
				player := e.CreatePlayer()
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

	panic(err)
}
