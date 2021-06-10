package main

import (
	state "bar-cli/serverfactory/server_example/server"
	"log"
)

func main() {
	state.Start(
		func(a state.AddItemToPlayerParams, e *state.Engine) {},
		func(p state.MovePlayerParams, e *state.Engine) {
			log.Println("moving player..")
			e.Player(p.Player).Position().SetX(p.ChangeX)
		},
		func(a state.SpawnZoneItemsParams, e *state.Engine) {},
		func(*state.Engine) {},
		func(*state.Engine) {},
	)
}
