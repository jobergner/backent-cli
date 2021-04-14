package main

import (
	state "bar-cli/serverfactory/server_example/server"
	"log"
)

func main() {
	state.Start(
		func(a state.PlayerID, x float64, y float64, e *state.Engine) {
			log.Println("moving player..")
			e.Player(a).Position(e).SetX(e, x)
		},
		func(state.Item, state.PlayerID, *state.Engine) {}, func([]state.Item, *state.Engine) {},
		func(*state.Engine) {},
		func(*state.Engine) {},
	)
}
