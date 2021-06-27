package integrationtest

import (
	"log"

	"github.com/Java-Jonas/bar-cli/integrationtest/state"
)

func startServer() {
	var playerID state.PlayerID

	err := state.Start(
		func(a state.AddItemToPlayerParams, e *state.Engine) state.AddItemToPlayerResponse {
			log.Println("addItemToPlayer", a)
			player := e.Player(playerID)
			item := player.AddItem()
			item.SetName(a.NewName)
			return state.AddItemToPlayerResponse{PlayerPath: player.Path()}
		},
		func(p state.MovePlayerParams, e *state.Engine) {
			log.Println("movePlayer", p)
			playerPosition := e.Player(playerID).Position()
			playerPosition.SetX(playerPosition.X() + p.ChangeX)
		},
		func(a state.SpawnZoneItemsParams, e *state.Engine) state.SpawnZoneItemsResponse {
			return state.SpawnZoneItemsResponse{}
		},
		func(e *state.Engine) {
			player := e.CreatePlayer()
			playerID = player.ID()
		},
		func(e *state.Engine) {
			playerGearScore := e.Player(playerID).GearScore()
			playerGearScore.SetLevel(playerGearScore.Level() + 1)
		},
	)

	panic(err)
}
